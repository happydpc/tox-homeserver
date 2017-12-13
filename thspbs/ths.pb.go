// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ths.proto

/*
Package thspbs is a generated protocol buffer package.

It is generated from these files:
	ths.proto

It has these top-level messages:
	Event
	BaseInfo
	FriendInfo
	GroupInfo
	MemberInfo
	EmptyReq
	HelloReq
	HelloResp
*/
package thspbs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id    int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Args  []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	Margs []string `protobuf:"bytes,4,rep,name=margs" json:"margs,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Event) GetMargs() []string {
	if m != nil {
		return m.Margs
	}
	return nil
}

type BaseInfo struct {
	Id      string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Stmsg   string                 `protobuf:"bytes,3,opt,name=stmsg" json:"stmsg,omitempty"`
	Status  uint32                 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	Friends map[uint32]*FriendInfo `protobuf:"bytes,5,rep,name=friends" json:"friends,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Groups  map[uint32]*GroupInfo  `protobuf:"bytes,6,rep,name=groups" json:"groups,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BaseInfo) Reset()                    { *m = BaseInfo{} }
func (m *BaseInfo) String() string            { return proto.CompactTextString(m) }
func (*BaseInfo) ProtoMessage()               {}
func (*BaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BaseInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BaseInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *BaseInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BaseInfo) GetFriends() map[uint32]*FriendInfo {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *BaseInfo) GetGroups() map[uint32]*GroupInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

type FriendInfo struct {
	Fnum   uint32 `protobuf:"varint,1,opt,name=fnum" json:"fnum,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Pubkey string `protobuf:"bytes,3,opt,name=pubkey" json:"pubkey,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Stmsg  string `protobuf:"bytes,5,opt,name=stmsg" json:"stmsg,omitempty"`
	Avatar string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty"`
	Seen   uint64 `protobuf:"varint,7,opt,name=seen" json:"seen,omitempty"`
}

func (m *FriendInfo) Reset()                    { *m = FriendInfo{} }
func (m *FriendInfo) String() string            { return proto.CompactTextString(m) }
func (*FriendInfo) ProtoMessage()               {}
func (*FriendInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FriendInfo) GetFnum() uint32 {
	if m != nil {
		return m.Fnum
	}
	return 0
}

func (m *FriendInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FriendInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *FriendInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FriendInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *FriendInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *FriendInfo) GetSeen() uint64 {
	if m != nil {
		return m.Seen
	}
	return 0
}

type GroupInfo struct {
	Gnum    uint32                 `protobuf:"varint,1,opt,name=gnum" json:"gnum,omitempty"`
	Cookie  string                 `protobuf:"bytes,2,opt,name=cookie" json:"cookie,omitempty"`
	Title   string                 `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Mtype   uint32                 `protobuf:"varint,4,opt,name=mtype" json:"mtype,omitempty"`
	Ours    bool                   `protobuf:"varint,5,opt,name=ours" json:"ours,omitempty"`
	Members map[uint32]*MemberInfo `protobuf:"bytes,6,rep,name=members" json:"members,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GroupInfo) Reset()                    { *m = GroupInfo{} }
func (m *GroupInfo) String() string            { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()               {}
func (*GroupInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupInfo) GetGnum() uint32 {
	if m != nil {
		return m.Gnum
	}
	return 0
}

func (m *GroupInfo) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

func (m *GroupInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GroupInfo) GetMtype() uint32 {
	if m != nil {
		return m.Mtype
	}
	return 0
}

func (m *GroupInfo) GetOurs() bool {
	if m != nil {
		return m.Ours
	}
	return false
}

func (m *GroupInfo) GetMembers() map[uint32]*MemberInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberInfo struct {
	Pnum   uint32 `protobuf:"varint,1,opt,name=pnum" json:"pnum,omitempty"`
	Pubkey string `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *MemberInfo) Reset()                    { *m = MemberInfo{} }
func (m *MemberInfo) String() string            { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()               {}
func (*MemberInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MemberInfo) GetPnum() uint32 {
	if m != nil {
		return m.Pnum
	}
	return 0
}

func (m *MemberInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *MemberInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EmptyReq struct {
}

func (m *EmptyReq) Reset()                    { *m = EmptyReq{} }
func (m *EmptyReq) String() string            { return proto.CompactTextString(m) }
func (*EmptyReq) ProtoMessage()               {}
func (*EmptyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type HelloReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *HelloReq) Reset()                    { *m = HelloReq{} }
func (m *HelloReq) String() string            { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()               {}
func (*HelloReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type HelloResp struct {
	Code   int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Status int64 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
}

func (m *HelloResp) Reset()                    { *m = HelloResp{} }
func (m *HelloResp) String() string            { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()               {}
func (*HelloResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *HelloResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HelloResp) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "thspbs.Event")
	proto.RegisterType((*BaseInfo)(nil), "thspbs.BaseInfo")
	proto.RegisterType((*FriendInfo)(nil), "thspbs.FriendInfo")
	proto.RegisterType((*GroupInfo)(nil), "thspbs.GroupInfo")
	proto.RegisterType((*MemberInfo)(nil), "thspbs.MemberInfo")
	proto.RegisterType((*EmptyReq)(nil), "thspbs.EmptyReq")
	proto.RegisterType((*HelloReq)(nil), "thspbs.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "thspbs.HelloResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Toxhs service

type ToxhsClient interface {
	PollCallback(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (Toxhs_PollCallbackClient, error)
	GetBaseInfo(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BaseInfo, error)
}

type toxhsClient struct {
	cc *grpc.ClientConn
}

func NewToxhsClient(cc *grpc.ClientConn) ToxhsClient {
	return &toxhsClient{cc}
}

func (c *toxhsClient) PollCallback(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (Toxhs_PollCallbackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Toxhs_serviceDesc.Streams[0], c.cc, "/thspbs.Toxhs/PollCallback", opts...)
	if err != nil {
		return nil, err
	}
	x := &toxhsPollCallbackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Toxhs_PollCallbackClient interface {
	Recv() (*EmptyReq, error)
	grpc.ClientStream
}

type toxhsPollCallbackClient struct {
	grpc.ClientStream
}

func (x *toxhsPollCallbackClient) Recv() (*EmptyReq, error) {
	m := new(EmptyReq)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *toxhsClient) GetBaseInfo(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BaseInfo, error) {
	out := new(BaseInfo)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/GetBaseInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Toxhs service

type ToxhsServer interface {
	PollCallback(*EmptyReq, Toxhs_PollCallbackServer) error
	GetBaseInfo(context.Context, *EmptyReq) (*BaseInfo, error)
}

func RegisterToxhsServer(s *grpc.Server, srv ToxhsServer) {
	s.RegisterService(&_Toxhs_serviceDesc, srv)
}

func _Toxhs_PollCallback_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ToxhsServer).PollCallback(m, &toxhsPollCallbackServer{stream})
}

type Toxhs_PollCallbackServer interface {
	Send(*EmptyReq) error
	grpc.ServerStream
}

type toxhsPollCallbackServer struct {
	grpc.ServerStream
}

func (x *toxhsPollCallbackServer) Send(m *EmptyReq) error {
	return x.ServerStream.SendMsg(m)
}

func _Toxhs_GetBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).GetBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/GetBaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).GetBaseInfo(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Toxhs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thspbs.Toxhs",
	HandlerType: (*ToxhsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBaseInfo",
			Handler:    _Toxhs_GetBaseInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PollCallback",
			Handler:       _Toxhs_PollCallback_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ths.proto",
}

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*HelloReq, error)
	// 测试带参数的hello
	SayHellox(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := grpc.Invoke(ctx, "/thspbs.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHellox(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := grpc.Invoke(ctx, "/thspbs.Greeter/SayHellox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *EmptyReq) (*HelloReq, error)
	// 测试带参数的hello
	SayHellox(context.Context, *HelloReq) (*HelloReq, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHellox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHellox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Greeter/SayHellox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHellox(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thspbs.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHellox",
			Handler:    _Greeter_SayHellox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ths.proto",
}

func init() { proto.RegisterFile("ths.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xed, 0xd8, 0x89, 0x27, 0xed, 0xaf, 0xfc, 0xab, 0xaa, 0xb2, 0x22, 0x40, 0x91, 0x6f,
	0xf0, 0x95, 0x29, 0x01, 0xa9, 0x15, 0x97, 0xa0, 0x10, 0x90, 0x0a, 0x42, 0x06, 0x1e, 0x60, 0x93,
	0x6c, 0x0e, 0xc4, 0xf6, 0xba, 0xbb, 0xeb, 0xa8, 0xb9, 0xe4, 0x65, 0x78, 0x37, 0xde, 0x02, 0xed,
	0x29, 0x71, 0x93, 0x50, 0x89, 0xbb, 0x99, 0xd9, 0xf9, 0xbe, 0x39, 0x7c, 0x63, 0x43, 0x28, 0x96,
	0x3c, 0xad, 0x18, 0x15, 0x14, 0x05, 0x62, 0xc9, 0xab, 0x09, 0x8f, 0xbf, 0x83, 0x3f, 0xda, 0x90,
	0x52, 0xa0, 0xff, 0xc0, 0x5d, 0xcd, 0x22, 0x67, 0xe0, 0x24, 0x7e, 0xe6, 0xae, 0x66, 0x08, 0x41,
	0xab, 0xc4, 0x05, 0x89, 0xdc, 0x81, 0x93, 0x84, 0x99, 0xb2, 0x65, 0x0c, 0xb3, 0x05, 0x8f, 0xbc,
	0x81, 0x27, 0x63, 0xd2, 0x46, 0x17, 0xe0, 0x17, 0x2a, 0xd8, 0x52, 0x41, 0xed, 0xc4, 0xbf, 0x5d,
	0xe8, 0xbc, 0xc5, 0x9c, 0x7c, 0x2c, 0xe7, 0xb4, 0x41, 0x1d, 0xfe, 0x95, 0xfa, 0x02, 0x7c, 0x2e,
	0x0a, 0xbe, 0x88, 0x3c, 0x15, 0xd4, 0x0e, 0xba, 0x84, 0x80, 0x0b, 0x2c, 0x6a, 0xc9, 0xee, 0x24,
	0xe7, 0x99, 0xf1, 0xd0, 0x35, 0xb4, 0xe7, 0x6c, 0x45, 0xca, 0x19, 0x8f, 0xfc, 0x81, 0x97, 0x74,
	0x87, 0x4f, 0x53, 0x3d, 0x4f, 0x6a, 0x8b, 0xa6, 0xef, 0xf5, 0xfb, 0xa8, 0x14, 0x6c, 0x9b, 0xd9,
	0x6c, 0xf4, 0x1a, 0x82, 0x05, 0xa3, 0x75, 0xc5, 0xa3, 0x40, 0xe1, 0x9e, 0x1c, 0xe1, 0xc6, 0xea,
	0x59, 0xc3, 0x4c, 0x6e, 0xff, 0x33, 0x9c, 0x35, 0xe9, 0x50, 0x0f, 0xbc, 0x35, 0xd9, 0xaa, 0x89,
	0xce, 0x33, 0x69, 0xa2, 0x04, 0xfc, 0x0d, 0xce, 0x6b, 0x3d, 0x53, 0x77, 0x88, 0x2c, 0xad, 0x86,
	0x49, 0xe2, 0x4c, 0x27, 0xbc, 0x71, 0x6f, 0x9c, 0xfe, 0x2d, 0x74, 0x1b, 0x65, 0x4e, 0xd0, 0x3d,
	0x7f, 0x48, 0xf7, 0xbf, 0xa5, 0x53, 0xa8, 0x03, 0xb6, 0xf8, 0x97, 0x03, 0xb0, 0xaf, 0x23, 0xb7,
	0x3b, 0x2f, 0xeb, 0xc2, 0xd0, 0x29, 0xbb, 0xb1, 0x47, 0xf7, 0xc1, 0x1e, 0x2f, 0x21, 0xa8, 0xea,
	0x89, 0x2c, 0xae, 0xd7, 0x6e, 0xbc, 0x9d, 0x42, 0xad, 0x53, 0x0a, 0xf9, 0x07, 0x0a, 0xe1, 0x0d,
	0x16, 0x98, 0x45, 0x81, 0x66, 0xd0, 0x9e, 0x64, 0xe0, 0x84, 0x94, 0x51, 0x7b, 0xe0, 0x24, 0xad,
	0x4c, 0xd9, 0xf1, 0x4f, 0x17, 0xc2, 0xdd, 0x04, 0x32, 0x63, 0xd1, 0xe8, 0x73, 0x61, 0xfa, 0x9c,
	0x52, 0xba, 0x5e, 0xd9, 0xdb, 0x30, 0x9e, 0xac, 0x2d, 0x56, 0x22, 0x27, 0xf6, 0x3a, 0x94, 0xa3,
	0x4e, 0x4f, 0x6c, 0x2b, 0x62, 0x8e, 0x43, 0x3b, 0x92, 0x97, 0xd6, 0x8c, 0xab, 0x36, 0x3b, 0x99,
	0xb2, 0xd1, 0x0d, 0xb4, 0x0b, 0x52, 0x4c, 0x08, 0xb3, 0xba, 0x3f, 0x3b, 0xda, 0x68, 0xfa, 0x49,
	0x27, 0x98, 0x83, 0x31, 0xe9, 0x52, 0xfa, 0xe6, 0xc3, 0x3f, 0x48, 0xaf, 0x61, 0x87, 0x62, 0xdd,
	0x02, 0xec, 0x1f, 0x64, 0xaf, 0x55, 0x63, 0x07, 0x95, 0xd9, 0x81, 0xd1, 0xc4, 0x3d, 0xa9, 0x89,
	0xb7, 0xd7, 0x24, 0x06, 0xe8, 0x8c, 0x8a, 0x4a, 0x6c, 0x33, 0x72, 0x17, 0x5f, 0x41, 0xe7, 0x03,
	0xc9, 0x73, 0x9a, 0x91, 0xbb, 0x5d, 0xae, 0xd3, 0xd0, 0xaf, 0x07, 0x9e, 0x54, 0x4f, 0x93, 0x4a,
	0x33, 0xbe, 0x86, 0xd0, 0x20, 0x78, 0x25, 0x21, 0x53, 0x3a, 0x23, 0xe6, 0x0f, 0xa0, 0xec, 0x83,
	0xb3, 0xf1, 0xec, 0xd9, 0x0c, 0x2b, 0xf0, 0xbf, 0xd1, 0xfb, 0xa5, 0xfc, 0x9c, 0xce, 0xbe, 0xd0,
	0x3c, 0x7f, 0x87, 0xf3, 0x7c, 0x82, 0xa7, 0x6b, 0xd4, 0xb3, 0xc3, 0xdb, 0xae, 0xfa, 0x47, 0x91,
	0x2b, 0x07, 0xbd, 0x84, 0xee, 0x98, 0x88, 0xdd, 0xef, 0xe1, 0x11, 0x90, 0xcd, 0x19, 0xfe, 0x80,
	0xf6, 0x98, 0x11, 0x22, 0x08, 0x43, 0x29, 0x74, 0xbe, 0xe2, 0xad, 0x6a, 0xfc, 0x31, 0xe8, 0x6e,
	0x17, 0x2f, 0x20, 0xb4, 0xf9, 0xf7, 0xe8, 0xe8, 0xf9, 0x18, 0x30, 0x09, 0xd4, 0x1f, 0xf2, 0xd5,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x3b, 0x2d, 0xea, 0x2e, 0x05, 0x00, 0x00,
}
