package dns

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/etcd/msg"
	"github.com/coredns/coredns/plugin/pkg/dnsutil"
	"github.com/coredns/coredns/plugin/pkg/fall"
	"github.com/coredns/coredns/plugin/pkg/upstream"
	"github.com/coredns/coredns/request"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mholt/caddy"
	"github.com/miekg/dns"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"

	ylog "yunion.io/x/log"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/k8s"
	"yunion.io/x/pkg/utils"
	"yunion.io/x/sqlchemy"
)

const (
	PluginName string = "yunion"

	// defaultTTL to apply to all answers
	defaultTTL           = 10
	defaultDbMaxOpenConn = 32
	defaultDbMaxIdleConn = 32
)

var (
	DNSTypeMap map[uint16]string = map[uint16]string{
		dns.TypeA:     "A",
		dns.TypeAAAA:  "AAAA",
		dns.TypeTXT:   "TXT",
		dns.TypeCNAME: "CNAME",
		dns.TypePTR:   "PTR",
		dns.TypeMX:    "MX",
		dns.TypeSRV:   "SRV",
		dns.TypeSOA:   "SOA",
		dns.TypeNS:    "NS",
	}
)

type SRegionDNS struct {
	Next          plugin.Handler
	Fall          fall.F
	Zones         []string
	PrimaryZone   string
	Upstream      upstream.Upstream
	SqlConnection string
	K8sConfigFile string
	K8sClient     *kubernetes.Clientset
}

func New() *SRegionDNS {
	r := new(SRegionDNS)
	return r
}

func (r *SRegionDNS) initDB(c *caddy.Controller) error {
	dialect, sqlStr, err := utils.TransSQLAchemyURL(r.SqlConnection)
	if err != nil {
		return err
	}
	sqlDb, err := sql.Open(dialect, sqlStr)
	if err != nil {
		return err
	}
	sqlDb.SetMaxOpenConns(defaultDbMaxOpenConn)
	sqlDb.SetMaxIdleConns(defaultDbMaxIdleConn)
	sqlchemy.SetDB(sqlDb)
	db.InitAllManagers()

	c.OnShutdown(func() error {
		sqlchemy.CloseDB()
		return nil
	})
	return nil
}

func (r *SRegionDNS) initK8s(c *caddy.Controller) {
	cli, err := k8s.NewClientByFile(r.K8sConfigFile, nil)
	if err != nil {
		ylog.Errorf("Init kubernetes client error: %v", err)
		return
	}
	r.K8sClient = cli
	pods, err := cli.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		ylog.Errorf("Get all pods in kubernetes cluster error: %v", err)
		return
	}
	ylog.Infof("Init k8s client success, %d pods in the cluster", len(pods.Items))
}

func (r *SRegionDNS) ServeDNS(ctx context.Context, w dns.ResponseWriter, rmsg *dns.Msg) (int, error) {
	var (
		records []dns.RR
		extra   []dns.RR
		err     error
	)

	opt := plugin.Options{}
	state := request.Request{W: w, Req: rmsg, Context: ctx}
	zone := plugin.Zones(r.Zones).Matches(state.Name())
	switch state.QType() {
	case dns.TypeA:
		records, err = plugin.A(r, zone, state, nil, opt)
	case dns.TypeAAAA:
		// TODO fallthrough to next
		records, err = plugin.AAAA(r, zone, state, nil, opt)
	case dns.TypeTXT:
		records, err = plugin.TXT(r, zone, state, opt)
	case dns.TypeCNAME:
		records, err = plugin.CNAME(r, zone, state, opt)
	case dns.TypePTR:
		records, err = plugin.PTR(r, zone, state, opt)
	case dns.TypeMX:
		records, extra, err = plugin.MX(r, zone, state, opt)
	case dns.TypeSRV:
		records, extra, err = plugin.SRV(r, zone, state, opt)
	case dns.TypeSOA:
		records, err = plugin.SOA(r, zone, state, opt)
	case dns.TypeNS:
		if state.Name() == zone {
			records, extra, err = plugin.NS(r, zone, state, opt)
			break
		}
		fallthrough
	default:
		ylog.Warningf("Not processed state: %#v", state)
		// Do a fake A lookup, so we can distinguish between NODATA and NXDOMAIN
		_, err = plugin.A(r, zone, state, nil, opt)
	}

	if err == errCallNext {
		if r.Fall.Through(state.Name()) {
			return plugin.NextOrFailure(r.Name(), r.Next, ctx, w, rmsg)
		}
		return plugin.BackendError(r, zone, dns.RcodeNameError, state, nil /* err */, opt)
	} else if err == errRefused {
		return plugin.BackendError(r, zone, dns.RcodeRefused, state, err, opt)
	} else if err == errNotFound {
		return plugin.BackendError(r, zone, dns.RcodeNameError, state, err, opt)
	}

	if len(records) == 0 {
		return plugin.BackendError(r, zone, dns.RcodeNameError, state, err, opt)
	}

	m := new(dns.Msg)
	m.SetReply(rmsg)
	m.Authoritative, m.RecursionAvailable = true, true
	m.Answer = append(m.Answer, records...)
	m.Extra = append(m.Extra, extra...)

	state.SizeAndDo(m)
	m, _ = state.Scrub(m)
	w.WriteMsg(m)
	return dns.RcodeSuccess, nil
}

var (
	errRefused  = errors.New("refused the query")
	errNotFound = errors.New("not found")
	errCallNext = errors.New("continue to next")
)

// Services implements the ServiceBackend interface
func (r *SRegionDNS) Services(state request.Request, exact bool, opt plugin.Options) (services []msg.Service, err error) {
	switch state.QType() {
	case dns.TypeTXT:
		t, _ := dnsutil.TrimZone(state.Name(), state.Zone)

		segs := dns.SplitDomainName(t)
		if len(segs) != 1 {
			return nil, fmt.Errorf("yunion region: TXT query can onlyu be for dns-version: %s", state.QName())
		}
		if segs[0] != "dns-version" {
			return nil, nil
		}
		svc := msg.Service{Text: "0.0.1", TTL: 28800, Key: msg.Path(state.QName(), "coredns")}
		return []msg.Service{svc}, nil
	case dns.TypeNS:
		ns := r.nsAddr()
		svc := msg.Service{Host: ns.A.String(), Key: msg.Path(state.QName(), "coredns")}
		return []msg.Service{svc}, nil
	}

	if state.QType() == dns.TypeA && isDefaultNS(state.Name(), state.Zone) {
		// If this is an A request for "ns.dns", respond with a "fake" record for coredns.
		// SOA records always use this hardcoded name
		ns := r.nsAddr()
		svc := msg.Service{Host: ns.A.String(), Key: msg.Path(state.QName(), "coredns")}
		return []msg.Service{svc}, nil
	}

	services, err = r.Records(state, false)
	return
}

// Lookup implements the ServiceBackend interface
func (r *SRegionDNS) Lookup(state request.Request, name string, typ uint16) (*dns.Msg, error) {
	return r.Upstream.Lookup(state, name, typ)
}

// IsNameError implements the ServiceBackend interface
func (r *SRegionDNS) IsNameError(err error) bool {
	return err == errCallNext
}

// Records looks up records in region mysql
func (r *SRegionDNS) Records(state request.Request, exact bool) ([]msg.Service, error) {
	req, e := parseRequest(state)
	if e != nil {
		return nil, e
	}
	return r.findRecords(req)
}

func (r *SRegionDNS) getHostIpWithName(req *recordRequest) string {
	name := req.QueryName()
	host, _ := models.HostManager.FetchByName("", name)
	if host == nil {
		return ""
	}
	ip := host.(*models.SHost).AccessIp
	return ip
}

func (r *SRegionDNS) getGuestIpWithName(req *recordRequest) []string {
	ips := []string{}
	name := req.QueryName()
	projectId := req.ProjectId()
	wantOnlyExit := false
	ips = models.GuestManager.GetIpInProjectWithName(projectId, name, wantOnlyExit)
	return ips
}

func (r *SRegionDNS) getK8sServiceBackends(req *recordRequest) ([]string, error) {
	queryInfo := req.GetK8sQueryInfo()
	pods, err := r.getK8sServicePods(queryInfo.Namespace, queryInfo.ServiceName)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			err = nil
		}
		return nil, err
	}
	ips := make([]string, 0)
	for _, pod := range pods {
		ip := pod.Status.PodIP
		if len(ip) != 0 {
			ips = append(ips, ip)
		}
	}
	return ips, nil
}

func (r *SRegionDNS) getK8sServicePods(namespace, name string) ([]v1.Pod, error) {
	cli := r.K8sClient
	svc, err := cli.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	labelSelector := labels.SelectorFromSet(svc.Spec.Selector)
	pods, err := cli.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector.String(),
		FieldSelector: fields.Everything().String(),
	})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

func (r *SRegionDNS) Name() string {
	return PluginName
}

func (r *SRegionDNS) queryLocalDnsRecords(req *recordRequest) (recs []msg.Service) {
	ips := models.DnsRecordManager.QueryDnsIps(req.ProjectId(), req.Name(), req.Type())
	if len(ips) == 0 {
		return
	}

	for _, ip := range ips {
		var s = msg.Service{}
		var ttl uint32 = uint32(ip.Ttl)
		if ttl == 0 {
			ttl = defaultTTL
		}
		if req.IsSRV() {
			parts := strings.SplitN(ip.Addr, ":", 2)
			if len(parts) != 2 {
				ylog.Errorf("Invalid SRV records: %q", ip.Addr)
				return
			}
			port, e := strconv.Atoi(parts[1])
			if e != nil {
				ylog.Errorf("Invalid SRV records: %q", ip.Addr)
				return
			}
			s = msg.Service{Host: parts[0], Port: port, TTL: ttl}
		} else {
			s = msg.Service{Host: ip.Addr, TTL: ttl}
		}
		recs = append(recs, s)
	}
	return
}

func (r *SRegionDNS) isMyDomain(req *recordRequest) bool {
	zones := []string{fmt.Sprintf("%s.", r.PrimaryZone)}
	zone := plugin.Zones(zones).Matches(req.state.Name())
	if zone != "" {
		return true
	}
	return false
}

func (r *SRegionDNS) findRecords(req *recordRequest) ([]msg.Service, error) {
	// 1. try local dns records table
	rrs := r.queryLocalDnsRecords(req)
	if len(rrs) > 0 {
		return rrs, nil
	}

	isPlainName := req.IsPlainName()
	isMyDomain := r.isMyDomain(req)
	if isPlainName {
		isCloudIp := req.SrcInCloud()
		if isCloudIp {
			ips := r.findInternalRecordIps(req)
			if len(ips) > 0 {
				return ips2DnsRecords(ips), nil
			} else {
				return nil, errNotFound
			}
		} else {
			return nil, errRefused
		}
	} else if isMyDomain {
		ips := r.findInternalRecordIps(req)
		if len(ips) > 0 {
			return ips2DnsRecords(ips), nil
		} else {
			return nil, errNotFound
		}
	} else {
		return nil, errCallNext
	}
}

func (r *SRegionDNS) findInternalRecordIps(req *recordRequest) []string {
	{
		// 1. try host table
		ip := r.getHostIpWithName(req)
		if len(ip) > 0 {
			return []string{ip}
		}
	}
	{
		// 2. try guest table
		ips := r.getGuestIpWithName(req)
		if len(ips) > 0 {
			return ips
		}
	}

	if r.K8sClient == nil {
		ylog.Warningf("K8s client not ready, skip it.")
		return nil
	}
	// 3. try k8s service backends
	ips, err := r.getK8sServiceBackends(req)
	if err != nil {
		ylog.Errorf("Get k8s service backends error: %v", err)
	}
	return ips
}

func ips2DnsRecords(ips []string) []msg.Service {
	recs := make([]msg.Service, 0)
	for _, ip := range ips {
		s := msg.Service{Host: ip, TTL: defaultTTL}
		recs = append(recs, s)
	}
	return recs
}
