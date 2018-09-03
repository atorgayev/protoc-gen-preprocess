// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package demo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Demo struct {
	PreprocessedField         string   `protobuf:"bytes,1,opt,name=preprocessedField,proto3" json:"preprocessedField,omitempty"`
	PreprocessedRepeatedField []string `protobuf:"bytes,2,rep,name=preprocessedRepeatedField,proto3" json:"preprocessedRepeatedField,omitempty"`
	Untouched                 string   `protobuf:"bytes,3,opt,name=untouched,proto3" json:"untouched,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Demo) Reset()         { *m = Demo{} }
func (m *Demo) String() string { return proto.CompactTextString(m) }
func (*Demo) ProtoMessage()    {}
func (*Demo) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_23bf4c656b62c6d3, []int{0}
}
func (m *Demo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Demo.Unmarshal(m, b)
}
func (m *Demo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Demo.Marshal(b, m, deterministic)
}
func (dst *Demo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Demo.Merge(dst, src)
}
func (m *Demo) XXX_Size() int {
	return xxx_messageInfo_Demo.Size(m)
}
func (m *Demo) XXX_DiscardUnknown() {
	xxx_messageInfo_Demo.DiscardUnknown(m)
}

var xxx_messageInfo_Demo proto.InternalMessageInfo

func (m *Demo) GetPreprocessedField() string {
	if m != nil {
		return m.PreprocessedField
	}
	return ""
}

func (m *Demo) GetPreprocessedRepeatedField() []string {
	if m != nil {
		return m.PreprocessedRepeatedField
	}
	return nil
}

func (m *Demo) GetUntouched() string {
	if m != nil {
		return m.Untouched
	}
	return ""
}

func init() {
	proto.RegisterType((*Demo)(nil), "Demo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	Echo(ctx context.Context, in *Demo, opts ...grpc.CallOption) (*Demo, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Echo(ctx context.Context, in *Demo, opts ...grpc.CallOption) (*Demo, error) {
	out := new(Demo)
	err := c.cc.Invoke(ctx, "/DemoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	Echo(context.Context, *Demo) (*Demo, error)
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Demo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DemoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Echo(ctx, req.(*Demo))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _DemoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_demo_23bf4c656b62c6d3) }

var fileDescriptor_demo_23bf4c656b62c6d3 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x72, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0xcf, 0xcc, 0x4b, 0xcb, 0x4f, 0xca, 0xc9, 0xaf, 0xc8, 0x2f, 0x48, 0xcd, 0xd3,
	0x07, 0x4b, 0x27, 0xeb, 0xa6, 0xa7, 0xe6, 0xe9, 0x16, 0x14, 0xa5, 0x16, 0x14, 0xe5, 0x27, 0xa7,
	0x16, 0x17, 0xeb, 0xe7, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x15, 0xeb, 0x23, 0x84, 0xa0, 0xe6, 0xc8,
	0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97,
	0x24, 0x82, 0x15, 0x42, 0x64, 0x95, 0x96, 0x30, 0x72, 0xb1, 0xb8, 0xa4, 0xe6, 0xe6, 0x0b, 0x99,
	0x71, 0x09, 0x22, 0xb4, 0xa6, 0xa6, 0xb8, 0x65, 0xa6, 0xe6, 0xa4, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x3a, 0x71, 0xec, 0xda, 0xc6, 0xca, 0x22, 0xc4, 0xc4, 0xc1, 0x18, 0x84, 0xa9, 0x44, 0xc8,
	0x8d, 0x4b, 0x12, 0x59, 0x30, 0x28, 0xb5, 0x20, 0x35, 0xb1, 0x04, 0xa6, 0x9f, 0x49, 0x81, 0x19,
	0x45, 0x3f, 0x6e, 0xa5, 0x42, 0x32, 0x5c, 0x9c, 0xa5, 0x79, 0x25, 0xf9, 0xa5, 0xc9, 0x19, 0xa9,
	0x29, 0x12, 0xcc, 0x20, 0x7b, 0x83, 0x10, 0x02, 0x46, 0xa6, 0x5c, 0xdc, 0x20, 0x57, 0x06, 0xa7,
	0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xa9, 0x71, 0xb1, 0xb8, 0x26, 0x67, 0xe4, 0x0b, 0xb1, 0xea,
	0x81, 0x44, 0xa5, 0x20, 0x94, 0x92, 0x40, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xb8, 0x94, 0x58, 0xf5,
	0x53, 0x93, 0x33, 0xf2, 0xad, 0x18, 0xb5, 0x92, 0xd8, 0xc0, 0x9e, 0x34, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x96, 0xa2, 0xfe, 0x58, 0x01, 0x00, 0x00,
}
