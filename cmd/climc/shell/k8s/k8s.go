package k8s

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/cmd/climc/shell"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/mcclient/modules/k8s"
	"yunion.io/x/onecloud/pkg/util/printutils"
)

func init() {
	// cluster resources
	initCluster()
	initNode()

	// helm resources
	initTiller()
	initRepo()
	initChart()
	initRelease()

	// kubernetes original resources
	initRaw()
	initConfigMap()
	initDeployment()
	initStatefulset()
	initPod()
	initService()
	initIngress()
	initNamespace()
	initK8sNode()
}

type clusterBaseOptions struct {
	Cluster string `default:"$K8S_CLUSTER|default" help:"Kubernetes cluster name"`
}

func (o clusterBaseOptions) ClusterParams() *jsonutils.JSONDict {
	ret := jsonutils.NewDict()
	ret.Add(jsonutils.NewString(o.Cluster), "cluster")
	return ret
}

type baseListOptions struct {
	Limit  int `default:"20" help:"Page limit"`
	Offset int `default:"0" help:"page offset"`
}

type NamespaceResourceListOptions struct {
	namespaceListOptions
	baseListOptions
}

func (o NamespaceResourceListOptions) Params() *jsonutils.JSONDict {
	params := fetchNamespaceParams(o.namespaceListOptions)
	params.Update(fetchPagingParams(o.baseListOptions))
	params.Update(o.ClusterParams())
	return params
}

func fetchPagingParams(opt baseListOptions) *jsonutils.JSONDict {
	params := jsonutils.NewDict()
	if opt.Limit > 0 {
		params.Add(jsonutils.NewInt(int64(opt.Limit)), "limit")
	}
	if opt.Offset > 0 {
		params.Add(jsonutils.NewInt(int64(opt.Offset)), "offset")
	}
	return params
}

type namespaceListOptions struct {
	namespaceOptions
	AllNamespace bool `help:"Show resource in all namespace"`
}

type namespaceOptions struct {
	clusterBaseOptions
	Namespace string `help:"Namespace of this resource"`
}

type resourceGetOptions struct {
	clusterBaseOptions
	Namespace string `help:"Namespace of this resource"`
	NAME      string `help:"Name ident of the resource"`
}

func (o resourceGetOptions) ToJSON() *jsonutils.JSONDict {
	params := o.ClusterParams()
	if o.Namespace != "" {
		params.Add(jsonutils.NewString(o.Namespace), "namespace")
	}
	return params
}

func fetchNamespaceParams(opt namespaceListOptions) *jsonutils.JSONDict {
	params := jsonutils.NewDict()
	if opt.AllNamespace {
		params.Add(jsonutils.JSONTrue, "all_namespace")
		return params
	}
	if opt.Namespace != "" {
		params.Add(jsonutils.NewString(opt.Namespace), "namespace")
	}
	return params
}

var (
	R                 = shell.R
	printList         = printutils.PrintJSONList
	printObject       = printutils.PrintJSONObject
	printBatchResults = printutils.PrintJSONBatchResults
)

func resourceCmdN(prefix, suffix string) string {
	return fmt.Sprintf("k8s-%s-%s", prefix, suffix)
}

func kubeResourceCmdN(prefix, suffix string) string {
	return fmt.Sprintf("kube-%s-%s", prefix, suffix)
}

func clusterContext(clusterId string) modules.ManagerContext {
	return modules.ManagerContext{
		InstanceManager: k8s.Clusters,
		InstanceId:      clusterId,
	}
}

func printObjectYAML(obj jsonutils.JSONObject) {
	fmt.Println(obj.YAMLString())
}
