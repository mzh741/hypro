// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hypro.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	hypro.proto

It has these top-level messages:
	CheckVersionRequest
	CheckVersionResponse
	RegisterRequest
	RegisterResponse
	Packet
*/
package protos

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

type CheckVersionRequest struct {
	ClientVersion string `protobuf:"bytes,10,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
}

func (m *CheckVersionRequest) Reset()                    { *m = CheckVersionRequest{} }
func (m *CheckVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckVersionRequest) ProtoMessage()               {}
func (*CheckVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckVersionRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type CheckVersionResponse struct {
	Compatible    bool   `protobuf:"varint,10,opt,name=compatible" json:"compatible,omitempty"`
	ServerVersion string `protobuf:"bytes,20,opt,name=server_version,json=serverVersion" json:"server_version,omitempty"`
	MinVersion    string `protobuf:"bytes,30,opt,name=min_version,json=minVersion" json:"min_version,omitempty"`
}

func (m *CheckVersionResponse) Reset()                    { *m = CheckVersionResponse{} }
func (m *CheckVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckVersionResponse) ProtoMessage()               {}
func (*CheckVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckVersionResponse) GetCompatible() bool {
	if m != nil {
		return m.Compatible
	}
	return false
}

func (m *CheckVersionResponse) GetServerVersion() string {
	if m != nil {
		return m.ServerVersion
	}
	return ""
}

func (m *CheckVersionResponse) GetMinVersion() string {
	if m != nil {
		return m.MinVersion
	}
	return ""
}

type RegisterRequest struct {
	Domain string `protobuf:"bytes,10,opt,name=domain" json:"domain,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type RegisterResponse struct {
	Token      string `protobuf:"bytes,10,opt,name=token" json:"token,omitempty"`
	FullDomain string `protobuf:"bytes,20,opt,name=full_domain,json=fullDomain" json:"full_domain,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterResponse) GetFullDomain() string {
	if m != nil {
		return m.FullDomain
	}
	return ""
}

type Packet struct {
	Data []byte `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Packet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckVersionRequest)(nil), "protos.CheckVersionRequest")
	proto.RegisterType((*CheckVersionResponse)(nil), "protos.CheckVersionResponse")
	proto.RegisterType((*RegisterRequest)(nil), "protos.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "protos.RegisterResponse")
	proto.RegisterType((*Packet)(nil), "protos.Packet")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tunnel service

type TunnelClient interface {
	CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	CreateTunnel(ctx context.Context, opts ...grpc.CallOption) (Tunnel_CreateTunnelClient, error)
}

type tunnelClient struct {
	cc *grpc.ClientConn
}

func NewTunnelClient(cc *grpc.ClientConn) TunnelClient {
	return &tunnelClient{cc}
}

func (c *tunnelClient) CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionResponse, error) {
	out := new(CheckVersionResponse)
	err := grpc.Invoke(ctx, "/protos.Tunnel/CheckVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/protos.Tunnel/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelClient) CreateTunnel(ctx context.Context, opts ...grpc.CallOption) (Tunnel_CreateTunnelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tunnel_serviceDesc.Streams[0], c.cc, "/protos.Tunnel/CreateTunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelCreateTunnelClient{stream}
	return x, nil
}

type Tunnel_CreateTunnelClient interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ClientStream
}

type tunnelCreateTunnelClient struct {
	grpc.ClientStream
}

func (x *tunnelCreateTunnelClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelCreateTunnelClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Tunnel service

type TunnelServer interface {
	CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	CreateTunnel(Tunnel_CreateTunnelServer) error
}

func RegisterTunnelServer(s *grpc.Server, srv TunnelServer) {
	s.RegisterService(&_Tunnel_serviceDesc, srv)
}

func _Tunnel_CheckVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServer).CheckVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Tunnel/CheckVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServer).CheckVersion(ctx, req.(*CheckVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tunnel_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Tunnel/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tunnel_CreateTunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).CreateTunnel(&tunnelCreateTunnelServer{stream})
}

type Tunnel_CreateTunnelServer interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type tunnelCreateTunnelServer struct {
	grpc.ServerStream
}

func (x *tunnelCreateTunnelServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelCreateTunnelServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Tunnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Tunnel",
	HandlerType: (*TunnelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckVersion",
			Handler:    _Tunnel_CheckVersion_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Tunnel_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateTunnel",
			Handler:       _Tunnel_CreateTunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hypro.proto",
}

func init() { proto.RegisterFile("hypro.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0xa0, 0xa1, 0x4e, 0x63, 0x95, 0x35, 0x68, 0x89, 0xa5, 0x4a, 0x40, 0xa8, 0x97, 0x22,
	0xf5, 0xaa, 0xa7, 0x7a, 0xe9, 0x4d, 0x82, 0x78, 0x2d, 0xdb, 0x74, 0xb4, 0x4b, 0x92, 0xdd, 0xb8,
	0xbb, 0x29, 0x78, 0xf1, 0xe7, 0xf9, 0xbb, 0x24, 0xfb, 0x51, 0x6b, 0xed, 0x29, 0x99, 0x37, 0xef,
	0xbd, 0x79, 0x33, 0x2c, 0x74, 0x57, 0x9f, 0xb5, 0x14, 0xe3, 0x5a, 0x0a, 0x2d, 0x48, 0x68, 0x3e,
	0x2a, 0x7d, 0x80, 0xb3, 0xe9, 0x0a, 0xf3, 0xe2, 0x15, 0xa5, 0x62, 0x82, 0x67, 0xf8, 0xd1, 0xa0,
	0xd2, 0xe4, 0x06, 0x7a, 0x79, 0xc9, 0x90, 0xeb, 0xf9, 0xda, 0x36, 0xfa, 0x70, 0x1d, 0x8c, 0x8e,
	0xb2, 0x63, 0x8b, 0x3a, 0x76, 0xfa, 0x05, 0xf1, 0x5f, 0xb5, 0xaa, 0x05, 0x57, 0x48, 0x86, 0x00,
	0xb9, 0xa8, 0x6a, 0xaa, 0xd9, 0xa2, 0x44, 0x23, 0xed, 0x64, 0x5b, 0x48, 0x6b, 0xaf, 0x50, 0xae,
	0x51, 0x6e, 0xec, 0x63, 0x6b, 0x6f, 0x51, 0x67, 0x47, 0xae, 0xa0, 0x5b, 0x31, 0xbe, 0xe1, 0x0c,
	0x0d, 0x07, 0x2a, 0xc6, 0xfd, 0xfc, 0x5b, 0x38, 0xc9, 0xf0, 0x9d, 0x29, 0x8d, 0xd2, 0x27, 0x3f,
	0x87, 0x70, 0x29, 0x2a, 0xca, 0x7c, 0x62, 0x57, 0xa5, 0x33, 0x38, 0xfd, 0xa5, 0xba, 0x98, 0x31,
	0x1c, 0x6a, 0x51, 0xa0, 0xa7, 0xda, 0xa2, 0x9d, 0xfa, 0xd6, 0x94, 0xe5, 0xdc, 0xd9, 0xd8, 0x64,
	0xd0, 0x42, 0x4f, 0xd6, 0x6a, 0x00, 0xe1, 0x33, 0xcd, 0x0b, 0xd4, 0x84, 0xc0, 0xc1, 0x92, 0x6a,
	0x6a, 0xf4, 0x51, 0x66, 0xfe, 0x27, 0xdf, 0x01, 0x84, 0x2f, 0x0d, 0xe7, 0x58, 0x92, 0x19, 0x44,
	0xdb, 0xe7, 0x21, 0x97, 0xf6, 0xf8, 0x6a, 0xbc, 0xe7, 0xe4, 0xc9, 0x60, 0x7f, 0xd3, 0x45, 0x7d,
	0x84, 0x8e, 0x8f, 0x4f, 0x2e, 0x3c, 0x73, 0x67, 0xf7, 0xa4, 0xff, 0xbf, 0xe1, 0xe4, 0x13, 0x88,
	0xa6, 0x12, 0xa9, 0x46, 0x97, 0xac, 0xe7, 0x99, 0x76, 0x91, 0x64, 0xa7, 0x1e, 0x05, 0x77, 0xc1,
	0xc2, 0x3e, 0x91, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x9c, 0xf2, 0x98, 0x38, 0x02,
	0x00, 0x00,
}
