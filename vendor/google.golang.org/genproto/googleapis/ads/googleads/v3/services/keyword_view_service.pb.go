// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/keyword_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [KeywordViewService.GetKeywordView][google.ads.googleads.v3.services.KeywordViewService.GetKeywordView].
type GetKeywordViewRequest struct {
	// Required. The resource name of the keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordViewRequest) Reset()         { *m = GetKeywordViewRequest{} }
func (m *GetKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordViewRequest) ProtoMessage()    {}
func (*GetKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9ec08eaed78b3, []int{0}
}

func (m *GetKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordViewRequest.Merge(m, src)
}
func (m *GetKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordViewRequest.Size(m)
}
func (m *GetKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordViewRequest proto.InternalMessageInfo

func (m *GetKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordViewRequest)(nil), "google.ads.googleads.v3.services.GetKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/keyword_view_service.proto", fileDescriptor_38e9ec08eaed78b3)
}

var fileDescriptor_38e9ec08eaed78b3 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0xeb, 0xd3, 0x40,
	0x18, 0x26, 0x29, 0x08, 0x06, 0x75, 0x08, 0x48, 0x4b, 0x14, 0x2c, 0xa5, 0x43, 0xe9, 0x70, 0x27,
	0x46, 0x10, 0xae, 0x38, 0x5c, 0x07, 0x2b, 0x08, 0x52, 0x2a, 0x64, 0x90, 0x40, 0xb8, 0x26, 0xaf,
	0xf1, 0x30, 0xc9, 0xd5, 0xbb, 0x34, 0x45, 0xc4, 0xc5, 0xaf, 0xe0, 0x37, 0x70, 0xf4, 0x5b, 0xb8,
	0x76, 0x75, 0x73, 0x72, 0x70, 0x72, 0x77, 0x15, 0x49, 0x2f, 0x97, 0xa6, 0xfa, 0x2b, 0xdd, 0x1e,
	0xee, 0xf9, 0xf3, 0xbe, 0xef, 0x93, 0x38, 0xb3, 0x54, 0x88, 0x34, 0x03, 0xcc, 0x12, 0x85, 0x35,
	0xac, 0x51, 0xe5, 0x63, 0x05, 0xb2, 0xe2, 0x31, 0x28, 0xfc, 0x06, 0xde, 0xed, 0x84, 0x4c, 0xa2,
	0x8a, 0xc3, 0x2e, 0x6a, 0x5e, 0xd1, 0x46, 0x8a, 0x52, 0xb8, 0x43, 0xed, 0x40, 0x2c, 0x51, 0xa8,
	0x35, 0xa3, 0xca, 0x47, 0xc6, 0xec, 0x3d, 0x3c, 0x17, 0x2f, 0x41, 0x89, 0xad, 0xfc, 0x37, 0x5f,
	0xe7, 0x7a, 0x77, 0x8d, 0x6b, 0xc3, 0x31, 0x2b, 0x0a, 0x51, 0xb2, 0x92, 0x8b, 0x42, 0x35, 0x6c,
	0xbf, 0xc3, 0xc6, 0x19, 0x87, 0xa2, 0x6c, 0x88, 0x7b, 0x1d, 0xe2, 0x15, 0x87, 0x2c, 0x89, 0xd6,
	0xf0, 0x9a, 0x55, 0x5c, 0x48, 0x2d, 0x18, 0x51, 0xe7, 0xf6, 0x02, 0xca, 0x67, 0x7a, 0x60, 0xc0,
	0x61, 0xb7, 0x82, 0xb7, 0x5b, 0x50, 0xa5, 0x3b, 0x71, 0x6e, 0x9a, 0x85, 0xa2, 0x82, 0xe5, 0x30,
	0xb0, 0x86, 0xd6, 0xe4, 0xfa, 0xbc, 0xf7, 0x83, 0xda, 0xab, 0x1b, 0x86, 0x79, 0xce, 0x72, 0x78,
	0xf0, 0xdb, 0x72, 0xdc, 0x4e, 0xc0, 0x0b, 0x7d, 0xa8, 0xfb, 0xd5, 0x72, 0x6e, 0x9d, 0x46, 0xbb,
	0x8f, 0xd0, 0xa5, 0x76, 0xd0, 0x95, 0xcb, 0x78, 0xe8, 0xac, 0xb1, 0x2d, 0x0d, 0x75, 0x6c, 0xa3,
	0x27, 0xdf, 0xe9, 0xe9, 0xf6, 0x1f, 0xbf, 0xfd, 0xfc, 0x64, 0xdf, 0x77, 0x51, 0xdd, 0xf3, 0xfb,
	0x13, 0xe6, 0x71, 0xbc, 0x55, 0xa5, 0xc8, 0x41, 0x2a, 0x3c, 0x35, 0xc5, 0xd7, 0x19, 0x0a, 0x4f,
	0x3f, 0x78, 0x77, 0xf6, 0x74, 0x70, 0x1c, 0xd7, 0xa0, 0x0d, 0x57, 0x28, 0x16, 0xf9, 0xfc, 0x8f,
	0xe5, 0x8c, 0x63, 0x91, 0x5f, 0xbc, 0x69, 0xde, 0xff, 0xbf, 0x9d, 0x65, 0x5d, 0xfe, 0xd2, 0x7a,
	0xf9, 0xb4, 0x31, 0xa7, 0x22, 0x63, 0x45, 0x8a, 0x84, 0x4c, 0x71, 0x0a, 0xc5, 0xe1, 0xd3, 0xe0,
	0xe3, 0xb8, 0xf3, 0xbf, 0xe2, 0xcc, 0x80, 0xcf, 0x76, 0x6f, 0x41, 0xe9, 0x17, 0x7b, 0xb8, 0xd0,
	0x81, 0x34, 0x51, 0x48, 0xc3, 0x1a, 0x05, 0x3e, 0x6a, 0x06, 0xab, 0xbd, 0x91, 0x84, 0x34, 0x51,
	0x61, 0x2b, 0x09, 0x03, 0x3f, 0x34, 0x92, 0x5f, 0xf6, 0x58, 0xbf, 0x13, 0x42, 0x13, 0x45, 0x48,
	0x2b, 0x22, 0x24, 0xf0, 0x09, 0x31, 0xb2, 0xf5, 0xb5, 0xc3, 0x9e, 0xfe, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0xbe, 0x83, 0xd8, 0x31, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordViewServiceClient is the client API for KeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordViewServiceClient interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error)
}

type keywordViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordViewServiceClient(cc grpc.ClientConnInterface) KeywordViewServiceClient {
	return &keywordViewServiceClient{cc}
}

func (c *keywordViewServiceClient) GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error) {
	out := new(resources.KeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordViewService/GetKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordViewServiceServer is the server API for KeywordViewService service.
type KeywordViewServiceServer interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(context.Context, *GetKeywordViewRequest) (*resources.KeywordView, error)
}

// UnimplementedKeywordViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordViewServiceServer struct {
}

func (*UnimplementedKeywordViewServiceServer) GetKeywordView(ctx context.Context, req *GetKeywordViewRequest) (*resources.KeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordView not implemented")
}

func RegisterKeywordViewServiceServer(s *grpc.Server, srv KeywordViewServiceServer) {
	s.RegisterService(&_KeywordViewService_serviceDesc, srv)
}

func _KeywordViewService_GetKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordViewService/GetKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, req.(*GetKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.KeywordViewService",
	HandlerType: (*KeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordView",
			Handler:    _KeywordViewService_GetKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/keyword_view_service.proto",
}
