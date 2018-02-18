// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shortener.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	shortener.proto

It has these top-level messages:
	ShortenUrlRequest
	ShortenUrlResponse
	ShortenUrlsResponse
	RemoveShortenedUrlRequest
	RemoveShortenedUrlResponse
	InfoRequest
	InfoResponse
	ShortenedUrl
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ShortenUrlRequest struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	OwnerId uint32 `protobuf:"varint,2,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
}

func (m *ShortenUrlRequest) Reset()                    { *m = ShortenUrlRequest{} }
func (m *ShortenUrlRequest) String() string            { return proto1.CompactTextString(m) }
func (*ShortenUrlRequest) ProtoMessage()               {}
func (*ShortenUrlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ShortenUrlRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ShortenUrlRequest) GetOwnerId() uint32 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *ShortenUrlRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

type ShortenUrlResponse struct {
	ShortenedUrl string `protobuf:"bytes,1,opt,name=shortened_url,json=shortenedUrl" json:"shortened_url,omitempty"`
}

func (m *ShortenUrlResponse) Reset()                    { *m = ShortenUrlResponse{} }
func (m *ShortenUrlResponse) String() string            { return proto1.CompactTextString(m) }
func (*ShortenUrlResponse) ProtoMessage()               {}
func (*ShortenUrlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ShortenUrlResponse) GetShortenedUrl() string {
	if m != nil {
		return m.ShortenedUrl
	}
	return ""
}

type ShortenUrlsResponse struct {
	Url          string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	ShortenedUrl string `protobuf:"bytes,2,opt,name=shortened_url,json=shortenedUrl" json:"shortened_url,omitempty"`
}

func (m *ShortenUrlsResponse) Reset()                    { *m = ShortenUrlsResponse{} }
func (m *ShortenUrlsResponse) String() string            { return proto1.CompactTextString(m) }
func (*ShortenUrlsResponse) ProtoMessage()               {}
func (*ShortenUrlsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ShortenUrlsResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ShortenUrlsResponse) GetShortenedUrl() string {
	if m != nil {
		return m.ShortenedUrl
	}
	return ""
}

type RemoveShortenedUrlRequest struct {
	ShortenedUrl string `protobuf:"bytes,1,opt,name=shortened_url,json=shortenedUrl" json:"shortened_url,omitempty"`
}

func (m *RemoveShortenedUrlRequest) Reset()                    { *m = RemoveShortenedUrlRequest{} }
func (m *RemoveShortenedUrlRequest) String() string            { return proto1.CompactTextString(m) }
func (*RemoveShortenedUrlRequest) ProtoMessage()               {}
func (*RemoveShortenedUrlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoveShortenedUrlRequest) GetShortenedUrl() string {
	if m != nil {
		return m.ShortenedUrl
	}
	return ""
}

type RemoveShortenedUrlResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RemoveShortenedUrlResponse) Reset()                    { *m = RemoveShortenedUrlResponse{} }
func (m *RemoveShortenedUrlResponse) String() string            { return proto1.CompactTextString(m) }
func (*RemoveShortenedUrlResponse) ProtoMessage()               {}
func (*RemoveShortenedUrlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveShortenedUrlResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type InfoRequest struct {
	ShortenedUrl string `protobuf:"bytes,1,opt,name=shortened_url,json=shortenedUrl" json:"shortened_url,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto1.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InfoRequest) GetShortenedUrl() string {
	if m != nil {
		return m.ShortenedUrl
	}
	return ""
}

type InfoResponse struct {
	ShortenedUrl *ShortenedUrl `protobuf:"bytes,1,opt,name=shortened_url,json=shortenedUrl" json:"shortened_url,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto1.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InfoResponse) GetShortenedUrl() *ShortenedUrl {
	if m != nil {
		return m.ShortenedUrl
	}
	return nil
}

type ShortenedUrl struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	OwnerId uint32 `protobuf:"varint,2,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	Clicks  uint32 `protobuf:"varint,4,opt,name=clicks" json:"clicks,omitempty"`
}

func (m *ShortenedUrl) Reset()                    { *m = ShortenedUrl{} }
func (m *ShortenedUrl) String() string            { return proto1.CompactTextString(m) }
func (*ShortenedUrl) ProtoMessage()               {}
func (*ShortenedUrl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ShortenedUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ShortenedUrl) GetOwnerId() uint32 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *ShortenedUrl) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *ShortenedUrl) GetClicks() uint32 {
	if m != nil {
		return m.Clicks
	}
	return 0
}

func init() {
	proto1.RegisterType((*ShortenUrlRequest)(nil), "proto.ShortenUrlRequest")
	proto1.RegisterType((*ShortenUrlResponse)(nil), "proto.ShortenUrlResponse")
	proto1.RegisterType((*ShortenUrlsResponse)(nil), "proto.ShortenUrlsResponse")
	proto1.RegisterType((*RemoveShortenedUrlRequest)(nil), "proto.RemoveShortenedUrlRequest")
	proto1.RegisterType((*RemoveShortenedUrlResponse)(nil), "proto.RemoveShortenedUrlResponse")
	proto1.RegisterType((*InfoRequest)(nil), "proto.InfoRequest")
	proto1.RegisterType((*InfoResponse)(nil), "proto.InfoResponse")
	proto1.RegisterType((*ShortenedUrl)(nil), "proto.ShortenedUrl")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Shortener service

type ShortenerClient interface {
	ShortenUrl(ctx context.Context, in *ShortenUrlRequest, opts ...grpc.CallOption) (*ShortenUrlResponse, error)
	ShortenUrls(ctx context.Context, opts ...grpc.CallOption) (Shortener_ShortenUrlsClient, error)
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	RemoveShortenedUrl(ctx context.Context, in *RemoveShortenedUrlRequest, opts ...grpc.CallOption) (*RemoveShortenedUrlResponse, error)
}

type shortenerClient struct {
	cc *grpc.ClientConn
}

func NewShortenerClient(cc *grpc.ClientConn) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) ShortenUrl(ctx context.Context, in *ShortenUrlRequest, opts ...grpc.CallOption) (*ShortenUrlResponse, error) {
	out := new(ShortenUrlResponse)
	err := grpc.Invoke(ctx, "/proto.Shortener/ShortenUrl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) ShortenUrls(ctx context.Context, opts ...grpc.CallOption) (Shortener_ShortenUrlsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Shortener_serviceDesc.Streams[0], c.cc, "/proto.Shortener/ShortenUrls", opts...)
	if err != nil {
		return nil, err
	}
	x := &shortenerShortenUrlsClient{stream}
	return x, nil
}

type Shortener_ShortenUrlsClient interface {
	Send(*ShortenUrlRequest) error
	Recv() (*ShortenUrlsResponse, error)
	grpc.ClientStream
}

type shortenerShortenUrlsClient struct {
	grpc.ClientStream
}

func (x *shortenerShortenUrlsClient) Send(m *ShortenUrlRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shortenerShortenUrlsClient) Recv() (*ShortenUrlsResponse, error) {
	m := new(ShortenUrlsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shortenerClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/proto.Shortener/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) RemoveShortenedUrl(ctx context.Context, in *RemoveShortenedUrlRequest, opts ...grpc.CallOption) (*RemoveShortenedUrlResponse, error) {
	out := new(RemoveShortenedUrlResponse)
	err := grpc.Invoke(ctx, "/proto.Shortener/RemoveShortenedUrl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shortener service

type ShortenerServer interface {
	ShortenUrl(context.Context, *ShortenUrlRequest) (*ShortenUrlResponse, error)
	ShortenUrls(Shortener_ShortenUrlsServer) error
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	RemoveShortenedUrl(context.Context, *RemoveShortenedUrlRequest) (*RemoveShortenedUrlResponse, error)
}

func RegisterShortenerServer(s *grpc.Server, srv ShortenerServer) {
	s.RegisterService(&_Shortener_serviceDesc, srv)
}

func _Shortener_ShortenUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).ShortenUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Shortener/ShortenUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).ShortenUrl(ctx, req.(*ShortenUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_ShortenUrls_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShortenerServer).ShortenUrls(&shortenerShortenUrlsServer{stream})
}

type Shortener_ShortenUrlsServer interface {
	Send(*ShortenUrlsResponse) error
	Recv() (*ShortenUrlRequest, error)
	grpc.ServerStream
}

type shortenerShortenUrlsServer struct {
	grpc.ServerStream
}

func (x *shortenerShortenUrlsServer) Send(m *ShortenUrlsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shortenerShortenUrlsServer) Recv() (*ShortenUrlRequest, error) {
	m := new(ShortenUrlRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Shortener_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Shortener/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_RemoveShortenedUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveShortenedUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).RemoveShortenedUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Shortener/RemoveShortenedUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).RemoveShortenedUrl(ctx, req.(*RemoveShortenedUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shortener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShortenUrl",
			Handler:    _Shortener_ShortenUrl_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Shortener_Info_Handler,
		},
		{
			MethodName: "RemoveShortenedUrl",
			Handler:    _Shortener_RemoveShortenedUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShortenUrls",
			Handler:       _Shortener_ShortenUrls_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "shortener.proto",
}

func init() { proto1.RegisterFile("shortener.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4b, 0x4f, 0xb3, 0x40,
	0x14, 0x86, 0x3f, 0xda, 0x7e, 0xbd, 0x9c, 0xb6, 0x51, 0x4f, 0x13, 0x33, 0xb0, 0x42, 0xdc, 0xb0,
	0x6a, 0x14, 0x13, 0xa3, 0x3b, 0x13, 0x37, 0x92, 0xb8, 0xa2, 0xe9, 0x46, 0x17, 0x4d, 0x84, 0x31,
	0x12, 0x29, 0x83, 0x33, 0xa0, 0x3f, 0xca, 0x3f, 0x69, 0x3a, 0x30, 0x5c, 0x02, 0x36, 0x9a, 0xb8,
	0x82, 0x73, 0x79, 0x9f, 0xf3, 0xce, 0x9c, 0x81, 0x03, 0xf1, 0xc2, 0x78, 0x4a, 0x63, 0xca, 0x97,
	0x09, 0x67, 0x29, 0xc3, 0xff, 0xf2, 0x63, 0x3d, 0xc0, 0xd1, 0x2a, 0xaf, 0xac, 0x79, 0xe4, 0xd1,
	0xb7, 0x8c, 0x8a, 0x14, 0x0f, 0xa1, 0x9f, 0xf1, 0x88, 0x68, 0xa6, 0x66, 0x4f, 0xbc, 0xdd, 0x2f,
	0xea, 0x30, 0x66, 0x1f, 0x31, 0xe5, 0x9b, 0x30, 0x20, 0x3d, 0x53, 0xb3, 0xe7, 0xde, 0x48, 0xc6,
	0x6e, 0x80, 0x04, 0x46, 0x2c, 0x49, 0x43, 0x16, 0x0b, 0xd2, 0x97, 0x02, 0x15, 0x5a, 0xd7, 0x80,
	0x75, 0xb6, 0x48, 0x58, 0x2c, 0x28, 0x9e, 0xc2, 0x5c, 0x79, 0x09, 0x36, 0xd5, 0x98, 0x59, 0x99,
	0x5c, 0xf3, 0xc8, 0xba, 0x87, 0x45, 0x25, 0x15, 0xa5, 0xb6, 0x6d, 0xac, 0x45, 0xeb, 0x75, 0xd0,
	0x6e, 0x40, 0xf7, 0xe8, 0x96, 0xbd, 0xd3, 0x55, 0x2d, 0xab, 0x0e, 0xfb, 0x23, 0x3f, 0x97, 0x60,
	0x74, 0x11, 0x0a, 0x5b, 0x04, 0x46, 0x22, 0xf3, 0x7d, 0x2a, 0x84, 0x14, 0x8f, 0x3d, 0x15, 0x5a,
	0x0e, 0x4c, 0xdd, 0xf8, 0x99, 0xfd, 0x6a, 0xd6, 0x1d, 0xcc, 0x72, 0x4d, 0x41, 0xbf, 0xea, 0x12,
	0x4d, 0x9d, 0x45, 0xbe, 0xc8, 0x65, 0xc3, 0x51, 0x93, 0xb4, 0x85, 0x59, 0xbd, 0xfa, 0x47, 0x7b,
	0xc5, 0x63, 0x18, 0xfa, 0x51, 0xe8, 0xbf, 0x0a, 0x32, 0x90, 0x92, 0x22, 0x72, 0x3e, 0x7b, 0x30,
	0x51, 0xf3, 0x38, 0xde, 0x02, 0x54, 0x2b, 0x44, 0xd2, 0x74, 0x5b, 0xdd, 0xbf, 0xa1, 0x77, 0x54,
	0xf2, 0x93, 0x5b, 0xff, 0xd0, 0x85, 0x69, 0xed, 0x1d, 0xec, 0xa1, 0x18, 0xad, 0x8a, 0xa8, 0x30,
	0xb6, 0x76, 0xa6, 0xe1, 0x39, 0x0c, 0x76, 0xd7, 0x8a, 0x58, 0x74, 0xd6, 0xf6, 0x62, 0x2c, 0x1a,
	0xb9, 0x72, 0xfa, 0x23, 0x60, 0x7b, 0xeb, 0x68, 0x16, 0xcd, 0xdf, 0x3e, 0x29, 0xe3, 0x64, 0x4f,
	0x87, 0x82, 0x3f, 0x0d, 0x65, 0xcf, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x40, 0x8b,
	0xb5, 0x9a, 0x03, 0x00, 0x00,
}