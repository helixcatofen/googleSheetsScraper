// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/display_keyword_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [DisplayKeywordViewService.GetDisplayKeywordView][google.ads.googleads.v1.services.DisplayKeywordViewService.GetDisplayKeywordView].
type GetDisplayKeywordViewRequest struct {
	// Required. The resource name of the display keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDisplayKeywordViewRequest) Reset()         { *m = GetDisplayKeywordViewRequest{} }
func (m *GetDisplayKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDisplayKeywordViewRequest) ProtoMessage()    {}
func (*GetDisplayKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b29fb51aa49493b, []int{0}
}

func (m *GetDisplayKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDisplayKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetDisplayKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDisplayKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetDisplayKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDisplayKeywordViewRequest.Merge(m, src)
}
func (m *GetDisplayKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDisplayKeywordViewRequest.Size(m)
}
func (m *GetDisplayKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDisplayKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDisplayKeywordViewRequest proto.InternalMessageInfo

func (m *GetDisplayKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDisplayKeywordViewRequest)(nil), "google.ads.googleads.v1.services.GetDisplayKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/display_keyword_view_service.proto", fileDescriptor_5b29fb51aa49493b)
}

var fileDescriptor_5b29fb51aa49493b = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6b, 0xd5, 0x40,
	0x14, 0xe5, 0xa5, 0x20, 0x18, 0x74, 0x13, 0x10, 0x6b, 0x2c, 0xfa, 0x28, 0x5d, 0x94, 0x2e, 0x66,
	0x88, 0x52, 0x84, 0xf1, 0x03, 0xe6, 0x29, 0xbc, 0x82, 0x28, 0xa5, 0x42, 0x16, 0x12, 0x08, 0xd3,
	0xcc, 0x35, 0x0e, 0x26, 0x99, 0x38, 0x37, 0x2f, 0x8f, 0x22, 0x6e, 0x04, 0x7f, 0x81, 0x1b, 0xd7,
	0x2e, 0xfd, 0x29, 0xdd, 0xba, 0x13, 0x04, 0x17, 0xae, 0xfc, 0x15, 0x92, 0x4e, 0x26, 0x7d, 0xa5,
	0xa6, 0xdd, 0x1d, 0xe6, 0x9c, 0x7b, 0xce, 0xfd, 0x18, 0xff, 0x69, 0xae, 0x75, 0x5e, 0x00, 0x15,
	0x12, 0xa9, 0x85, 0x1d, 0x6a, 0x23, 0x8a, 0x60, 0x5a, 0x95, 0x01, 0x52, 0xa9, 0xb0, 0x2e, 0xc4,
	0x51, 0xfa, 0x0e, 0x8e, 0x96, 0xda, 0xc8, 0xb4, 0x55, 0xb0, 0x4c, 0x7b, 0x96, 0xd4, 0x46, 0x37,
	0x3a, 0x98, 0xda, 0x4a, 0x22, 0x24, 0x92, 0xc1, 0x84, 0xb4, 0x11, 0x71, 0x26, 0xe1, 0xa3, 0xb1,
	0x18, 0x03, 0xa8, 0x17, 0x66, 0x2c, 0xc7, 0xfa, 0x87, 0x1b, 0xae, 0xba, 0x56, 0x54, 0x54, 0x95,
	0x6e, 0x44, 0xa3, 0x74, 0x85, 0x3d, 0x7b, 0x73, 0x85, 0xcd, 0x0a, 0x05, 0x55, 0xd3, 0x13, 0x77,
	0x57, 0x88, 0x37, 0x0a, 0x0a, 0x99, 0x1e, 0xc2, 0x5b, 0xd1, 0x2a, 0x6d, 0xac, 0x60, 0x73, 0xcf,
	0xdf, 0x98, 0x43, 0xf3, 0xcc, 0x06, 0x3f, 0xb7, 0xb9, 0xb1, 0x82, 0xe5, 0x01, 0xbc, 0x5f, 0x00,
	0x36, 0xc1, 0xb6, 0x7f, 0xdd, 0xf5, 0x97, 0x56, 0xa2, 0x84, 0xf5, 0xc9, 0x74, 0xb2, 0x7d, 0x75,
	0xb6, 0xf6, 0x9b, 0x7b, 0x07, 0xd7, 0x1c, 0xf3, 0x52, 0x94, 0x70, 0xef, 0xab, 0xe7, 0xdf, 0x3a,
	0xef, 0xf3, 0xca, 0x8e, 0x1f, 0xfc, 0x9a, 0xf8, 0x37, 0xfe, 0x1b, 0x14, 0x3c, 0x21, 0x97, 0xad,
	0x8e, 0x5c, 0xd4, 0x61, 0xb8, 0x3b, 0x5a, 0x3f, 0x2c, 0x96, 0x9c, 0xaf, 0xde, 0x7c, 0xf1, 0x93,
	0x9f, 0x9d, 0xec, 0xd3, 0x8f, 0x3f, 0x5f, 0xbc, 0x07, 0xc1, 0x6e, 0x77, 0x92, 0x0f, 0x67, 0x98,
	0xc7, 0xd9, 0x02, 0x1b, 0x5d, 0x82, 0x41, 0xba, 0xe3, 0x6e, 0xb4, 0x62, 0x85, 0x74, 0xe7, 0x63,
	0x78, 0xfb, 0x98, 0xaf, 0x9f, 0x86, 0xf7, 0xa8, 0x56, 0x48, 0x32, 0x5d, 0xce, 0x3e, 0x7b, 0xfe,
	0x56, 0xa6, 0xcb, 0x4b, 0x07, 0x9d, 0xdd, 0x19, 0x5d, 0xe0, 0x7e, 0x77, 0xad, 0xfd, 0xc9, 0xeb,
	0xbd, 0xde, 0x23, 0xd7, 0x85, 0xa8, 0x72, 0xa2, 0x4d, 0x4e, 0x73, 0xa8, 0x4e, 0x6e, 0x49, 0x4f,
	0x53, 0xc7, 0xff, 0xf2, 0x43, 0x07, 0xbe, 0x79, 0x6b, 0x73, 0xce, 0xbf, 0x7b, 0xd3, 0xb9, 0x35,
	0xe4, 0x12, 0x89, 0x85, 0x1d, 0x8a, 0x23, 0xd2, 0x07, 0xe3, 0xb1, 0x93, 0x24, 0x5c, 0x62, 0x32,
	0x48, 0x92, 0x38, 0x4a, 0x9c, 0xe4, 0xaf, 0xb7, 0x65, 0xdf, 0x19, 0xe3, 0x12, 0x19, 0x1b, 0x44,
	0x8c, 0xc5, 0x11, 0x63, 0x4e, 0x76, 0x78, 0xe5, 0xa4, 0xcf, 0xfb, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x93, 0x66, 0xd2, 0x7b, 0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DisplayKeywordViewServiceClient is the client API for DisplayKeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DisplayKeywordViewServiceClient interface {
	// Returns the requested display keyword view in full detail.
	GetDisplayKeywordView(ctx context.Context, in *GetDisplayKeywordViewRequest, opts ...grpc.CallOption) (*resources.DisplayKeywordView, error)
}

type displayKeywordViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDisplayKeywordViewServiceClient(cc grpc.ClientConnInterface) DisplayKeywordViewServiceClient {
	return &displayKeywordViewServiceClient{cc}
}

func (c *displayKeywordViewServiceClient) GetDisplayKeywordView(ctx context.Context, in *GetDisplayKeywordViewRequest, opts ...grpc.CallOption) (*resources.DisplayKeywordView, error) {
	out := new(resources.DisplayKeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.DisplayKeywordViewService/GetDisplayKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisplayKeywordViewServiceServer is the server API for DisplayKeywordViewService service.
type DisplayKeywordViewServiceServer interface {
	// Returns the requested display keyword view in full detail.
	GetDisplayKeywordView(context.Context, *GetDisplayKeywordViewRequest) (*resources.DisplayKeywordView, error)
}

// UnimplementedDisplayKeywordViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDisplayKeywordViewServiceServer struct {
}

func (*UnimplementedDisplayKeywordViewServiceServer) GetDisplayKeywordView(ctx context.Context, req *GetDisplayKeywordViewRequest) (*resources.DisplayKeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisplayKeywordView not implemented")
}

func RegisterDisplayKeywordViewServiceServer(s *grpc.Server, srv DisplayKeywordViewServiceServer) {
	s.RegisterService(&_DisplayKeywordViewService_serviceDesc, srv)
}

func _DisplayKeywordViewService_GetDisplayKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDisplayKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayKeywordViewServiceServer).GetDisplayKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.DisplayKeywordViewService/GetDisplayKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayKeywordViewServiceServer).GetDisplayKeywordView(ctx, req.(*GetDisplayKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DisplayKeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.DisplayKeywordViewService",
	HandlerType: (*DisplayKeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDisplayKeywordView",
			Handler:    _DisplayKeywordViewService_GetDisplayKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/display_keyword_view_service.proto",
}