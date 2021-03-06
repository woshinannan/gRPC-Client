// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type ReqMsg struct {
	Age                  int64    `protobuf:"varint,1,opt,name=Age,proto3" json:"Age,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMsg) Reset()         { *m = ReqMsg{} }
func (m *ReqMsg) String() string { return proto.CompactTextString(m) }
func (*ReqMsg) ProtoMessage()    {}
func (*ReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *ReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMsg.Unmarshal(m, b)
}
func (m *ReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMsg.Marshal(b, m, deterministic)
}
func (m *ReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMsg.Merge(m, src)
}
func (m *ReqMsg) XXX_Size() int {
	return xxx_messageInfo_ReqMsg.Size(m)
}
func (m *ReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMsg proto.InternalMessageInfo

func (m *ReqMsg) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *ReqMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RspMsg struct {
	Age_Name             string   `protobuf:"bytes,1,opt,name=Age_Name,json=AgeName,proto3" json:"Age_Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspMsg) Reset()         { *m = RspMsg{} }
func (m *RspMsg) String() string { return proto.CompactTextString(m) }
func (*RspMsg) ProtoMessage()    {}
func (*RspMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *RspMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspMsg.Unmarshal(m, b)
}
func (m *RspMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspMsg.Marshal(b, m, deterministic)
}
func (m *RspMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspMsg.Merge(m, src)
}
func (m *RspMsg) XXX_Size() int {
	return xxx_messageInfo_RspMsg.Size(m)
}
func (m *RspMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RspMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RspMsg proto.InternalMessageInfo

func (m *RspMsg) GetAge_Name() string {
	if m != nil {
		return m.Age_Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqMsg)(nil), "pb.ReqMsg")
	proto.RegisterType((*RspMsg)(nil), "pb.RspMsg")
}

func init() {
	proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d)
}

var fileDescriptor_ca53982754088a9d = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0xa8, 0x50, 0xd2, 0xe3, 0x62, 0x0b, 0x4a, 0x2d, 0xf4, 0x2d, 0x4e, 0x17, 0x12,
	0xe0, 0x62, 0x76, 0x4c, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x85, 0x84,
	0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25,
	0x65, 0x2e, 0xb6, 0xa0, 0xe2, 0x02, 0x90, 0x7a, 0x49, 0x2e, 0x0e, 0xc7, 0xf4, 0xd4, 0x78, 0xb0,
	0x0a, 0x46, 0xb0, 0x0a, 0x76, 0xc7, 0xf4, 0x54, 0x10, 0xd7, 0xc8, 0x91, 0x8b, 0xd7, 0xb7, 0x32,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x19, 0x2c, 0x20, 0x64, 0xc0, 0xc5, 0xe2, 0x9a, 0x9c, 0x91, 0x2f,
	0xc4, 0xa5, 0x57, 0x90, 0xa4, 0x07, 0xb1, 0x4f, 0x0a, 0xc2, 0x06, 0x9b, 0xa5, 0x24, 0xd0, 0x74,
	0xf9, 0xc9, 0x64, 0x26, 0x2e, 0x25, 0x56, 0xfd, 0xd4, 0xe4, 0x8c, 0x7c, 0x2b, 0x46, 0xad, 0x24,
	0x36, 0xb0, 0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xef, 0x17, 0x32, 0xce, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MyServiceNameClient is the client API for MyServiceName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyServiceNameClient interface {
	Echo(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (*RspMsg, error)
}

type myServiceNameClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceNameClient(cc grpc.ClientConnInterface) MyServiceNameClient {
	return &myServiceNameClient{cc}
}

func (c *myServiceNameClient) Echo(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (*RspMsg, error) {
	out := new(RspMsg)
	err := c.cc.Invoke(ctx, "/pb.MyServiceName/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceNameServer is the server API for MyServiceName service.
type MyServiceNameServer interface {
	Echo(context.Context, *ReqMsg) (*RspMsg, error)
}

// UnimplementedMyServiceNameServer can be embedded to have forward compatible implementations.
type UnimplementedMyServiceNameServer struct {
}

func (*UnimplementedMyServiceNameServer) Echo(ctx context.Context, req *ReqMsg) (*RspMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterMyServiceNameServer(s *grpc.Server, srv MyServiceNameServer) {
	s.RegisterService(&_MyServiceName_serviceDesc, srv)
}

func _MyServiceName_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceNameServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MyServiceName/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceNameServer).Echo(ctx, req.(*ReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyServiceName_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MyServiceName",
	HandlerType: (*MyServiceNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _MyServiceName_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
