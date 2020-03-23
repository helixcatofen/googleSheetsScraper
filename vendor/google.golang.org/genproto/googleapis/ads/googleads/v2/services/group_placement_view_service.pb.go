// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/group_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [GroupPlacementViewService.GetGroupPlacementView][google.ads.googleads.v2.services.GroupPlacementViewService.GetGroupPlacementView].
type GetGroupPlacementViewRequest struct {
	// Required. The resource name of the Group Placement view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupPlacementViewRequest) Reset()         { *m = GetGroupPlacementViewRequest{} }
func (m *GetGroupPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupPlacementViewRequest) ProtoMessage()    {}
func (*GetGroupPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fd8b1c75e923f6a, []int{0}
}

func (m *GetGroupPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetGroupPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupPlacementViewRequest.Merge(m, src)
}
func (m *GetGroupPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Size(m)
}
func (m *GetGroupPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupPlacementViewRequest proto.InternalMessageInfo

func (m *GetGroupPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGroupPlacementViewRequest)(nil), "google.ads.googleads.v2.services.GetGroupPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/group_placement_view_service.proto", fileDescriptor_3fd8b1c75e923f6a)
}

var fileDescriptor_3fd8b1c75e923f6a = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6a, 0xd5, 0x40,
	0x18, 0x25, 0x29, 0x08, 0x06, 0xdd, 0x04, 0xc4, 0x1a, 0x8b, 0x5e, 0x4a, 0x17, 0xa5, 0x8b, 0x19,
	0x88, 0x14, 0x61, 0xfc, 0x81, 0xb9, 0x2e, 0xd2, 0x8d, 0x72, 0xa9, 0x90, 0x85, 0x04, 0xc2, 0x34,
	0xf9, 0x8c, 0x03, 0xc9, 0x4c, 0x9c, 0x99, 0xa4, 0x0b, 0x71, 0x23, 0xf8, 0x04, 0x6e, 0x5c, 0xbb,
	0xf4, 0x51, 0xba, 0x75, 0x27, 0x08, 0x2e, 0x5c, 0xf9, 0x14, 0x92, 0x4e, 0x26, 0xed, 0xa5, 0xa6,
	0x77, 0x77, 0x98, 0x73, 0xbe, 0x73, 0xbe, 0x9f, 0x09, 0x5e, 0x54, 0x52, 0x56, 0x35, 0x60, 0x56,
	0x6a, 0x6c, 0xe1, 0x80, 0xfa, 0x18, 0x6b, 0x50, 0x3d, 0x2f, 0x40, 0xe3, 0x4a, 0xc9, 0xae, 0xcd,
	0xdb, 0x9a, 0x15, 0xd0, 0x80, 0x30, 0x79, 0xcf, 0xe1, 0x34, 0x1f, 0x59, 0xd4, 0x2a, 0x69, 0x64,
	0xb8, 0xb0, 0x95, 0x88, 0x95, 0x1a, 0x4d, 0x26, 0xa8, 0x8f, 0x91, 0x33, 0x89, 0x9e, 0xce, 0xc5,
	0x28, 0xd0, 0xb2, 0x53, 0x73, 0x39, 0xd6, 0x3f, 0xda, 0x71, 0xd5, 0x2d, 0xc7, 0x4c, 0x08, 0x69,
	0x98, 0xe1, 0x52, 0xe8, 0x91, 0xbd, 0x7b, 0x89, 0x2d, 0x6a, 0x0e, 0xc2, 0x8c, 0xc4, 0xc3, 0x4b,
	0xc4, 0x5b, 0x0e, 0x75, 0x99, 0x9f, 0xc0, 0x3b, 0xd6, 0x73, 0xa9, 0xac, 0x60, 0xf7, 0x28, 0xd8,
	0x49, 0xc0, 0x24, 0x43, 0xf0, 0xca, 0xe5, 0xa6, 0x1c, 0x4e, 0x8f, 0xe1, 0x7d, 0x07, 0xda, 0x84,
	0xfb, 0xc1, 0x6d, 0xd7, 0x5f, 0x2e, 0x58, 0x03, 0xdb, 0xde, 0xc2, 0xdb, 0xbf, 0xb9, 0xdc, 0xfa,
	0x4d, 0xfd, 0xe3, 0x5b, 0x8e, 0x79, 0xc5, 0x1a, 0x88, 0xbf, 0xfa, 0xc1, 0xbd, 0xab, 0x3e, 0xaf,
	0xed, 0xf8, 0xe1, 0x2f, 0x2f, 0xb8, 0xf3, 0xdf, 0xa0, 0xf0, 0x39, 0xda, 0xb4, 0x3a, 0x74, 0x5d,
	0x87, 0xd1, 0xe1, 0x6c, 0xfd, 0xb4, 0x58, 0x74, 0xb5, 0x7a, 0xf7, 0xe5, 0x4f, 0xba, 0x3e, 0xd9,
	0xa7, 0x1f, 0x7f, 0xbe, 0xf8, 0x8f, 0xc3, 0xc3, 0xe1, 0x24, 0x1f, 0xd6, 0x98, 0x67, 0x45, 0xa7,
	0x8d, 0x6c, 0x40, 0x69, 0x7c, 0x60, 0x6f, 0xb4, 0x66, 0xa5, 0xf1, 0xc1, 0xc7, 0xe8, 0xfe, 0x19,
	0xdd, 0xbe, 0x08, 0x1f, 0x51, 0xcb, 0x35, 0x2a, 0x64, 0xb3, 0xfc, 0xec, 0x07, 0x7b, 0x85, 0x6c,
	0x36, 0x0e, 0xba, 0x7c, 0x30, 0xbb, 0xc0, 0xd5, 0x70, 0xad, 0x95, 0xf7, 0xe6, 0x68, 0xf4, 0xa8,
	0x64, 0xcd, 0x44, 0x85, 0xa4, 0xaa, 0x70, 0x05, 0xe2, 0xfc, 0x96, 0xf8, 0x22, 0x75, 0xfe, 0x2f,
	0x3f, 0x71, 0xe0, 0x9b, 0xbf, 0x95, 0x50, 0xfa, 0xdd, 0x5f, 0x24, 0xd6, 0x90, 0x96, 0x1a, 0x59,
	0x38, 0xa0, 0x34, 0x46, 0x63, 0xb0, 0x3e, 0x73, 0x92, 0x8c, 0x96, 0x3a, 0x9b, 0x24, 0x59, 0x1a,
	0x67, 0x4e, 0xf2, 0xd7, 0xdf, 0xb3, 0xef, 0x84, 0xd0, 0x52, 0x13, 0x32, 0x89, 0x08, 0x49, 0x63,
	0x42, 0x9c, 0xec, 0xe4, 0xc6, 0x79, 0x9f, 0x8f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xdb,
	0xd6, 0xe9, 0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupPlacementViewServiceClient is the client API for GroupPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupPlacementViewServiceClient interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error)
}

type groupPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupPlacementViewServiceClient(cc grpc.ClientConnInterface) GroupPlacementViewServiceClient {
	return &groupPlacementViewServiceClient{cc}
}

func (c *groupPlacementViewServiceClient) GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error) {
	out := new(resources.GroupPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GroupPlacementViewService/GetGroupPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupPlacementViewServiceServer is the server API for GroupPlacementViewService service.
type GroupPlacementViewServiceServer interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(context.Context, *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error)
}

// UnimplementedGroupPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroupPlacementViewServiceServer struct {
}

func (*UnimplementedGroupPlacementViewServiceServer) GetGroupPlacementView(ctx context.Context, req *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupPlacementView not implemented")
}

func RegisterGroupPlacementViewServiceServer(s *grpc.Server, srv GroupPlacementViewServiceServer) {
	s.RegisterService(&_GroupPlacementViewService_serviceDesc, srv)
}

func _GroupPlacementViewService_GetGroupPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GroupPlacementViewService/GetGroupPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, req.(*GetGroupPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.GroupPlacementViewService",
	HandlerType: (*GroupPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupPlacementView",
			Handler:    _GroupPlacementViewService_GetGroupPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/group_placement_view_service.proto",
}