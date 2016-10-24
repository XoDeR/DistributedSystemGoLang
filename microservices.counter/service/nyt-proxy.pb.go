// Code generated by protoc-gen-go.
// source: nyt-proxy.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	nyt-proxy.proto

It has these top-level messages:
	MostPopularRequest
	MostPopularResponse
	CatsRequest
	CatsResponse
*/
package service

import (
	"fmt"
	"math"

	"github.com/NYTimes/gizmo/examples/nyt"
	proto "github.com/golang/protobuf/proto"
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

type MostPopularRequest struct {
	ResourceType   string `protobuf:"bytes,1,opt,name=resourceType" json:"resourceType,omitempty"`
	Section        string `protobuf:"bytes,2,opt,name=section" json:"section,omitempty"`
	TimePeriodDays uint32 `protobuf:"varint,3,opt,name=timePeriodDays" json:"timePeriodDays,omitempty"`
}

func (m *MostPopularRequest) Reset()                    { *m = MostPopularRequest{} }
func (m *MostPopularRequest) String() string            { return proto.CompactTextString(m) }
func (*MostPopularRequest) ProtoMessage()               {}
func (*MostPopularRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MostPopularResponse struct {
	Result []*nyt.MostPopularResult `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *MostPopularResponse) Reset()                    { *m = MostPopularResponse{} }
func (m *MostPopularResponse) String() string            { return proto.CompactTextString(m) }
func (*MostPopularResponse) ProtoMessage()               {}
func (*MostPopularResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MostPopularResponse) GetResult() []*nyt.MostPopularResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type CatsRequest struct {
}

func (m *CatsRequest) Reset()                    { *m = CatsRequest{} }
func (m *CatsRequest) String() string            { return proto.CompactTextString(m) }
func (*CatsRequest) ProtoMessage()               {}
func (*CatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CatsResponse struct {
	Results []*nyt.SemanticConceptArticle `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *CatsResponse) Reset()                    { *m = CatsResponse{} }
func (m *CatsResponse) String() string            { return proto.CompactTextString(m) }
func (*CatsResponse) ProtoMessage()               {}
func (*CatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CatsResponse) GetResults() []*nyt.SemanticConceptArticle {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*MostPopularRequest)(nil), "service.MostPopularRequest")
	proto.RegisterType((*MostPopularResponse)(nil), "service.MostPopularResponse")
	proto.RegisterType((*CatsRequest)(nil), "service.CatsRequest")
	proto.RegisterType((*CatsResponse)(nil), "service.CatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for NYTProxyService service

type NYTProxyServiceClient interface {
	GetMostPopular(ctx context.Context, in *MostPopularRequest, opts ...grpc.CallOption) (*MostPopularResponse, error)
	GetCats(ctx context.Context, in *CatsRequest, opts ...grpc.CallOption) (*CatsResponse, error)
}

type nYTProxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewNYTProxyServiceClient(cc *grpc.ClientConn) NYTProxyServiceClient {
	return &nYTProxyServiceClient{cc}
}

func (c *nYTProxyServiceClient) GetMostPopular(ctx context.Context, in *MostPopularRequest, opts ...grpc.CallOption) (*MostPopularResponse, error) {
	out := new(MostPopularResponse)
	err := grpc.Invoke(ctx, "/service.NYTProxyService/GetMostPopular", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nYTProxyServiceClient) GetCats(ctx context.Context, in *CatsRequest, opts ...grpc.CallOption) (*CatsResponse, error) {
	out := new(CatsResponse)
	err := grpc.Invoke(ctx, "/service.NYTProxyService/GetCats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NYTProxyService service

type NYTProxyServiceServer interface {
	GetMostPopular(context.Context, *MostPopularRequest) (*MostPopularResponse, error)
	GetCats(context.Context, *CatsRequest) (*CatsResponse, error)
}

func RegisterNYTProxyServiceServer(s *grpc.Server, srv NYTProxyServiceServer) {
	s.RegisterService(&NYTProxyService_serviceDesc, srv)
}

func _NYTProxyService_GetMostPopular_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MostPopularRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NYTProxyServiceServer).GetMostPopular(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NYTProxyService/GetMostPopular",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NYTProxyServiceServer).GetMostPopular(ctx, req.(*MostPopularRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NYTProxyService_GetCats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NYTProxyServiceServer).GetCats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NYTProxyService/GetCats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NYTProxyServiceServer).GetCats(ctx, req.(*CatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var NYTProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.NYTProxyService",
	HandlerType: (*NYTProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMostPopular",
			Handler:    _NYTProxyService_GetMostPopular_Handler,
		},
		{
			MethodName: "GetCats",
			Handler:    _NYTProxyService_GetCats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("nyt-proxy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xe9, 0x57, 0x68, 0xf9, 0x6e, 0xff, 0xc1, 0x68, 0x25, 0xa6, 0x2e, 0x24, 0x0b, 0xe9,
	0xc6, 0x11, 0x2a, 0xba, 0x97, 0x2a, 0x2e, 0x44, 0x29, 0x69, 0x37, 0x2e, 0xe3, 0x78, 0x17, 0x81,
	0x24, 0x33, 0xce, 0x4c, 0xc4, 0xf8, 0x20, 0x3e, 0xaf, 0x37, 0x99, 0xa9, 0x34, 0xd5, 0x4d, 0xc8,
	0xdc, 0xf3, 0xcb, 0xb9, 0x27, 0x67, 0x60, 0x52, 0x54, 0xf6, 0x5c, 0x69, 0xf9, 0x51, 0x71, 0x7a,
	0x5a, 0xc9, 0xfa, 0x06, 0xf5, 0x7b, 0x2a, 0x30, 0x9c, 0x92, 0x72, 0x91, 0x4b, 0x63, 0x95, 0x54,
	0x65, 0x96, 0x68, 0xa7, 0x87, 0xc7, 0xf5, 0xd8, 0x60, 0x9e, 0x14, 0x36, 0x15, 0x42, 0x16, 0x02,
	0x95, 0x75, 0x52, 0xf4, 0x09, 0xec, 0x91, 0xf8, 0x95, 0xe3, 0x63, 0x7c, 0x2b, 0xd1, 0x58, 0x16,
	0xc1, 0x50, 0xa3, 0x91, 0xa5, 0x16, 0xb8, 0xa9, 0x14, 0x06, 0x9d, 0xd3, 0xce, 0xfc, 0x7f, 0xdc,
	0x9a, 0xb1, 0x00, 0x68, 0xad, 0xb0, 0xa9, 0x2c, 0x82, 0x7f, 0x8d, 0xbc, 0x3d, 0xb2, 0x33, 0x18,
	0xdb, 0x34, 0xc7, 0x15, 0xea, 0x54, 0xbe, 0xde, 0x26, 0x95, 0x09, 0xba, 0x04, 0x8c, 0xe2, 0xbd,
	0x69, 0x74, 0x07, 0x07, 0xad, 0xdd, 0x46, 0xc9, 0xc2, 0x20, 0xe3, 0xd0, 0xa3, 0x45, 0x65, 0x66,
	0x69, 0x6d, 0x77, 0x3e, 0x58, 0x1c, 0x71, 0x8a, 0xcf, 0xdb, 0x24, 0xa9, 0xb1, 0xa7, 0xa2, 0x11,
	0x0c, 0x96, 0x89, 0x35, 0x3e, 0x3b, 0xb9, 0x0e, 0xdd, 0xd1, 0xdb, 0x5d, 0x41, 0xdf, 0x81, 0xc6,
	0xfb, 0xcd, 0x1a, 0xbf, 0xb5, 0xaf, 0x63, 0xe9, 0xea, 0xb8, 0xd1, 0xf4, 0x9e, 0x61, 0xbc, 0x65,
	0x17, 0x5f, 0x1d, 0x98, 0x3c, 0x3d, 0x6f, 0x56, 0x75, 0xcd, 0x6b, 0x57, 0x2f, 0x7b, 0x80, 0xf1,
	0x3d, 0xda, 0x9d, 0x24, 0x6c, 0xc6, 0x7d, 0xf5, 0xfc, 0x77, 0x8b, 0xe1, 0xc9, 0xdf, 0xa2, 0xcf,
	0x75, 0x0d, 0x7d, 0x32, 0xab, 0xa3, 0xb2, 0xc3, 0x1f, 0x70, 0xe7, 0x47, 0xc2, 0xe9, 0xde, 0xd4,
	0x7d, 0xf7, 0xd2, 0x6b, 0x2e, 0xee, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x15, 0x72, 0x87, 0xbb,
	0x06, 0x02, 0x00, 0x00,
}
