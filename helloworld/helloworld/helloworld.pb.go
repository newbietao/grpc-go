// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloRes struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRes) Reset()         { *m = HelloRes{} }
func (m *HelloRes) String() string { return proto.CompactTextString(m) }
func (*HelloRes) ProtoMessage()    {}
func (*HelloRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRes.Unmarshal(m, b)
}
func (m *HelloRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRes.Marshal(b, m, deterministic)
}
func (m *HelloRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRes.Merge(m, src)
}
func (m *HelloRes) XXX_Size() int {
	return xxx_messageInfo_HelloRes.Size(m)
}
func (m *HelloRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRes.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRes proto.InternalMessageInfo

func (m *HelloRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "helloworld.HelloReq")
	proto.RegisterType((*HelloRes)(nil), "helloworld.HelloRes")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xc9, 0x71, 0x71, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89,
	0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x0c, 0x5c, 0xbe, 0x58,
	0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x1d, 0x2a, 0x0d, 0x62, 0x1a, 0xb9, 0x71, 0x71, 0x81, 0x65,
	0xc3, 0x41, 0x66, 0x09, 0x59, 0x70, 0x71, 0x04, 0x27, 0x56, 0x82, 0x05, 0x84, 0x44, 0xf4, 0x90,
	0xac, 0x85, 0xd9, 0x20, 0x85, 0x4d, 0xb4, 0x58, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x30, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x86, 0x07, 0x11, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldClient interface {
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error)
}

type helloWorldClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldClient(cc *grpc.ClientConn) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error) {
	out := new(HelloRes)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorld/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
type HelloWorldServer interface {
	SayHello(context.Context, *HelloReq) (*HelloRes, error)
}

// UnimplementedHelloWorldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (*UnimplementedHelloWorldServer) SayHello(ctx context.Context, req *HelloReq) (*HelloRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).SayHello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloWorld_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
