// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: send_server.proto

package apis

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SendParams struct {
	Contact              string   `protobuf:"bytes,1,opt,name=Contact,proto3" json:"Contact,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	Priority             string   `protobuf:"bytes,5,opt,name=Priority,proto3" json:"Priority,omitempty"`
	RemoteTemplate       string   `protobuf:"bytes,6,opt,name=RemoteTemplate,proto3" json:"RemoteTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendParams) Reset()         { *m = SendParams{} }
func (m *SendParams) String() string { return proto.CompactTextString(m) }
func (*SendParams) ProtoMessage()    {}
func (*SendParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fdd68f7eb311f9, []int{0}
}

func (m *SendParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendParams.Unmarshal(m, b)
}
func (m *SendParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendParams.Marshal(b, m, deterministic)
}
func (m *SendParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendParams.Merge(m, src)
}
func (m *SendParams) XXX_Size() int {
	return xxx_messageInfo_SendParams.Size(m)
}
func (m *SendParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SendParams.DiscardUnknown(m)
}

var xxx_messageInfo_SendParams proto.InternalMessageInfo

func (m *SendParams) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *SendParams) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SendParams) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SendParams) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendParams) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *SendParams) GetRemoteTemplate() string {
	if m != nil {
		return m.RemoteTemplate
	}
	return ""
}

type UpdateConfigParams struct {
	Configs              map[string]string `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateConfigParams) Reset()         { *m = UpdateConfigParams{} }
func (m *UpdateConfigParams) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigParams) ProtoMessage()    {}
func (*UpdateConfigParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fdd68f7eb311f9, []int{1}
}

func (m *UpdateConfigParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigParams.Unmarshal(m, b)
}
func (m *UpdateConfigParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigParams.Marshal(b, m, deterministic)
}
func (m *UpdateConfigParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigParams.Merge(m, src)
}
func (m *UpdateConfigParams) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigParams.Size(m)
}
func (m *UpdateConfigParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigParams proto.InternalMessageInfo

func (m *UpdateConfigParams) GetConfigs() map[string]string {
	if m != nil {
		return m.Configs
	}
	return nil
}

type UseridByMobileParams struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseridByMobileParams) Reset()         { *m = UseridByMobileParams{} }
func (m *UseridByMobileParams) String() string { return proto.CompactTextString(m) }
func (*UseridByMobileParams) ProtoMessage()    {}
func (*UseridByMobileParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fdd68f7eb311f9, []int{2}
}

func (m *UseridByMobileParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseridByMobileParams.Unmarshal(m, b)
}
func (m *UseridByMobileParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseridByMobileParams.Marshal(b, m, deterministic)
}
func (m *UseridByMobileParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseridByMobileParams.Merge(m, src)
}
func (m *UseridByMobileParams) XXX_Size() int {
	return xxx_messageInfo_UseridByMobileParams.Size(m)
}
func (m *UseridByMobileParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UseridByMobileParams.DiscardUnknown(m)
}

var xxx_messageInfo_UseridByMobileParams proto.InternalMessageInfo

func (m *UseridByMobileParams) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fdd68f7eb311f9, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type UseridByMobileReply struct {
	Userid               string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseridByMobileReply) Reset()         { *m = UseridByMobileReply{} }
func (m *UseridByMobileReply) String() string { return proto.CompactTextString(m) }
func (*UseridByMobileReply) ProtoMessage()    {}
func (*UseridByMobileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fdd68f7eb311f9, []int{4}
}

func (m *UseridByMobileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseridByMobileReply.Unmarshal(m, b)
}
func (m *UseridByMobileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseridByMobileReply.Marshal(b, m, deterministic)
}
func (m *UseridByMobileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseridByMobileReply.Merge(m, src)
}
func (m *UseridByMobileReply) XXX_Size() int {
	return xxx_messageInfo_UseridByMobileReply.Size(m)
}
func (m *UseridByMobileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UseridByMobileReply.DiscardUnknown(m)
}

var xxx_messageInfo_UseridByMobileReply proto.InternalMessageInfo

func (m *UseridByMobileReply) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func init() {
	proto.RegisterType((*SendParams)(nil), "apis.SendParams")
	proto.RegisterType((*UpdateConfigParams)(nil), "apis.UpdateConfigParams")
	proto.RegisterMapType((map[string]string)(nil), "apis.UpdateConfigParams.ConfigsEntry")
	proto.RegisterType((*UseridByMobileParams)(nil), "apis.UseridByMobileParams")
	proto.RegisterType((*Empty)(nil), "apis.Empty")
	proto.RegisterType((*UseridByMobileReply)(nil), "apis.UseridByMobileReply")
}

func init() { proto.RegisterFile("send_server.proto", fileDescriptor_63fdd68f7eb311f9) }

var fileDescriptor_63fdd68f7eb311f9 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0x71, 0xfe, 0x6e, 0x26, 0x21, 0x64, 0xb5, 0x61, 0xd1, 0xfa, 0x14, 0x0c, 0x59, 0x72,
	0x59, 0x1f, 0xb2, 0x2c, 0x2c, 0xb9, 0x94, 0x36, 0x84, 0x9e, 0x02, 0xc1, 0x4d, 0xce, 0x45, 0x89,
	0xa7, 0x41, 0xd4, 0xb6, 0x8c, 0xa4, 0x04, 0xfc, 0x18, 0x7d, 0x93, 0xd2, 0x27, 0x2c, 0x92, 0xe5,
	0x36, 0x69, 0xd3, 0x9b, 0x7f, 0xdf, 0xe8, 0x1b, 0xcf, 0x37, 0x12, 0x7c, 0x57, 0x98, 0xc5, 0xf7,
	0x0a, 0xe5, 0x11, 0x65, 0x98, 0x4b, 0xa1, 0x05, 0x69, 0xb0, 0x9c, 0xab, 0xe0, 0xd9, 0x03, 0xb8,
	0xc3, 0x2c, 0x5e, 0x31, 0xc9, 0x52, 0x45, 0x28, 0xb4, 0xe7, 0x22, 0xd3, 0x6c, 0xa7, 0xa9, 0x37,
	0xf2, 0x26, 0x9d, 0xa8, 0x42, 0x32, 0x84, 0xe6, 0x5a, 0xe4, 0x7c, 0x47, 0x6b, 0x56, 0x2f, 0xc1,
	0xaa, 0x5c, 0x27, 0x48, 0xeb, 0x4e, 0x35, 0x60, 0xba, 0x2c, 0x51, 0x29, 0xb6, 0x47, 0xda, 0x28,
	0xbb, 0x38, 0x24, 0x3e, 0x7c, 0x5b, 0x49, 0x2e, 0x24, 0xd7, 0x05, 0x6d, 0xda, 0xd2, 0x1b, 0x93,
	0xdf, 0xd0, 0x8f, 0x30, 0x15, 0x1a, 0xd7, 0x98, 0xe6, 0x09, 0xd3, 0x48, 0x5b, 0xf6, 0xc4, 0x07,
	0x35, 0x78, 0xf2, 0x80, 0x6c, 0xf2, 0x98, 0x69, 0x9c, 0x8b, 0xec, 0x81, 0xef, 0xdd, 0xe8, 0x57,
	0xd0, 0xde, 0x59, 0x56, 0xd4, 0x1b, 0xd5, 0x27, 0xdd, 0xe9, 0x38, 0x34, 0x09, 0xc3, 0xcf, 0x47,
	0xc3, 0x12, 0xd4, 0x22, 0xd3, 0xb2, 0x88, 0x2a, 0x97, 0x3f, 0x83, 0xde, 0x69, 0x81, 0x0c, 0xa0,
	0xfe, 0x88, 0x85, 0xdb, 0x83, 0xf9, 0x34, 0x69, 0x8f, 0x2c, 0x39, 0x60, 0xb5, 0x03, 0x0b, 0xb3,
	0xda, 0x7f, 0x2f, 0x08, 0x61, 0xb8, 0x51, 0x28, 0x79, 0x7c, 0x53, 0x2c, 0xc5, 0x96, 0x27, 0xe8,
	0x86, 0xfa, 0x09, 0xad, 0xd4, 0xb2, 0x6b, 0xe3, 0x28, 0x68, 0x43, 0x73, 0x91, 0xe6, 0xba, 0x08,
	0xfe, 0xc0, 0x8f, 0x73, 0x63, 0x84, 0x79, 0x52, 0x18, 0xdf, 0xc1, 0xca, 0x95, 0xaf, 0xa4, 0xe9,
	0x8b, 0x07, 0x1d, 0x73, 0x5d, 0xd7, 0x7b, 0xcc, 0x34, 0x19, 0x43, 0xc3, 0x00, 0x19, 0x94, 0x49,
	0xdf, 0xef, 0xd1, 0xef, 0x96, 0x8a, 0xfd, 0x07, 0xf9, 0x07, 0xbd, 0xd3, 0x25, 0x10, 0xfa, 0xd5,
	0x62, 0xce, 0x6d, 0xb7, 0xd0, 0x3f, 0x1f, 0x8d, 0xf8, 0xce, 0x78, 0x21, 0xa9, 0xff, 0xeb, 0x52,
	0xcd, 0x86, 0xd9, 0xb6, 0xec, 0x83, 0xfb, 0xfb, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x6c, 0xcb,
	0x1b, 0x85, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SendAgentClient is the client API for SendAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendAgentClient interface {
	Send(ctx context.Context, in *SendParams, opts ...grpc.CallOption) (*Empty, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigParams, opts ...grpc.CallOption) (*Empty, error)
	UseridByMobile(ctx context.Context, in *UseridByMobileParams, opts ...grpc.CallOption) (*UseridByMobileReply, error)
}

type sendAgentClient struct {
	cc *grpc.ClientConn
}

func NewSendAgentClient(cc *grpc.ClientConn) SendAgentClient {
	return &sendAgentClient{cc}
}

func (c *sendAgentClient) Send(ctx context.Context, in *SendParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/apis.SendAgent/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendAgentClient) UpdateConfig(ctx context.Context, in *UpdateConfigParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/apis.SendAgent/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendAgentClient) UseridByMobile(ctx context.Context, in *UseridByMobileParams, opts ...grpc.CallOption) (*UseridByMobileReply, error) {
	out := new(UseridByMobileReply)
	err := c.cc.Invoke(ctx, "/apis.SendAgent/UseridByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendAgentServer is the server API for SendAgent service.
type SendAgentServer interface {
	Send(context.Context, *SendParams) (*Empty, error)
	UpdateConfig(context.Context, *UpdateConfigParams) (*Empty, error)
	UseridByMobile(context.Context, *UseridByMobileParams) (*UseridByMobileReply, error)
}

// UnimplementedSendAgentServer can be embedded to have forward compatible implementations.
type UnimplementedSendAgentServer struct {
}

func (*UnimplementedSendAgentServer) Send(ctx context.Context, req *SendParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedSendAgentServer) UpdateConfig(ctx context.Context, req *UpdateConfigParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (*UnimplementedSendAgentServer) UseridByMobile(ctx context.Context, req *UseridByMobileParams) (*UseridByMobileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseridByMobile not implemented")
}

func RegisterSendAgentServer(s *grpc.Server, srv SendAgentServer) {
	s.RegisterService(&_SendAgent_serviceDesc, srv)
}

func _SendAgent_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendAgentServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.SendAgent/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendAgentServer).Send(ctx, req.(*SendParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendAgent_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendAgentServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.SendAgent/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendAgentServer).UpdateConfig(ctx, req.(*UpdateConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendAgent_UseridByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseridByMobileParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendAgentServer).UseridByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.SendAgent/UseridByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendAgentServer).UseridByMobile(ctx, req.(*UseridByMobileParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _SendAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.SendAgent",
	HandlerType: (*SendAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SendAgent_Send_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _SendAgent_UpdateConfig_Handler,
		},
		{
			MethodName: "UseridByMobile",
			Handler:    _SendAgent_UseridByMobile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "send_server.proto",
}
