// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/media_file_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [MediaFileService.GetMediaFile][google.ads.googleads.v1.services.MediaFileService.GetMediaFile]
type GetMediaFileRequest struct {
	// Required. The resource name of the media file to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMediaFileRequest) Reset()         { *m = GetMediaFileRequest{} }
func (m *GetMediaFileRequest) String() string { return proto.CompactTextString(m) }
func (*GetMediaFileRequest) ProtoMessage()    {}
func (*GetMediaFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{0}
}

func (m *GetMediaFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMediaFileRequest.Unmarshal(m, b)
}
func (m *GetMediaFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMediaFileRequest.Marshal(b, m, deterministic)
}
func (m *GetMediaFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMediaFileRequest.Merge(m, src)
}
func (m *GetMediaFileRequest) XXX_Size() int {
	return xxx_messageInfo_GetMediaFileRequest.Size(m)
}
func (m *GetMediaFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMediaFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMediaFileRequest proto.InternalMessageInfo

func (m *GetMediaFileRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MediaFileService.MutateMediaFiles][google.ads.googleads.v1.services.MediaFileService.MutateMediaFiles]
type MutateMediaFilesRequest struct {
	// Required. The ID of the customer whose media files are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual media file.
	Operations []*MediaFileOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMediaFilesRequest) Reset()         { *m = MutateMediaFilesRequest{} }
func (m *MutateMediaFilesRequest) String() string { return proto.CompactTextString(m) }
func (*MutateMediaFilesRequest) ProtoMessage()    {}
func (*MutateMediaFilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{1}
}

func (m *MutateMediaFilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMediaFilesRequest.Unmarshal(m, b)
}
func (m *MutateMediaFilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMediaFilesRequest.Marshal(b, m, deterministic)
}
func (m *MutateMediaFilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMediaFilesRequest.Merge(m, src)
}
func (m *MutateMediaFilesRequest) XXX_Size() int {
	return xxx_messageInfo_MutateMediaFilesRequest.Size(m)
}
func (m *MutateMediaFilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMediaFilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMediaFilesRequest proto.InternalMessageInfo

func (m *MutateMediaFilesRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateMediaFilesRequest) GetOperations() []*MediaFileOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateMediaFilesRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateMediaFilesRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation to create media file.
type MediaFileOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*MediaFileOperation_Create
	Operation            isMediaFileOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MediaFileOperation) Reset()         { *m = MediaFileOperation{} }
func (m *MediaFileOperation) String() string { return proto.CompactTextString(m) }
func (*MediaFileOperation) ProtoMessage()    {}
func (*MediaFileOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{2}
}

func (m *MediaFileOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaFileOperation.Unmarshal(m, b)
}
func (m *MediaFileOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaFileOperation.Marshal(b, m, deterministic)
}
func (m *MediaFileOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaFileOperation.Merge(m, src)
}
func (m *MediaFileOperation) XXX_Size() int {
	return xxx_messageInfo_MediaFileOperation.Size(m)
}
func (m *MediaFileOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaFileOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MediaFileOperation proto.InternalMessageInfo

type isMediaFileOperation_Operation interface {
	isMediaFileOperation_Operation()
}

type MediaFileOperation_Create struct {
	Create *resources.MediaFile `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*MediaFileOperation_Create) isMediaFileOperation_Operation() {}

func (m *MediaFileOperation) GetOperation() isMediaFileOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MediaFileOperation) GetCreate() *resources.MediaFile {
	if x, ok := m.GetOperation().(*MediaFileOperation_Create); ok {
		return x.Create
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MediaFileOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MediaFileOperation_Create)(nil),
	}
}

// Response message for a media file mutate.
type MutateMediaFilesResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateMediaFileResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateMediaFilesResponse) Reset()         { *m = MutateMediaFilesResponse{} }
func (m *MutateMediaFilesResponse) String() string { return proto.CompactTextString(m) }
func (*MutateMediaFilesResponse) ProtoMessage()    {}
func (*MutateMediaFilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{3}
}

func (m *MutateMediaFilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMediaFilesResponse.Unmarshal(m, b)
}
func (m *MutateMediaFilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMediaFilesResponse.Marshal(b, m, deterministic)
}
func (m *MutateMediaFilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMediaFilesResponse.Merge(m, src)
}
func (m *MutateMediaFilesResponse) XXX_Size() int {
	return xxx_messageInfo_MutateMediaFilesResponse.Size(m)
}
func (m *MutateMediaFilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMediaFilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMediaFilesResponse proto.InternalMessageInfo

func (m *MutateMediaFilesResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateMediaFilesResponse) GetResults() []*MutateMediaFileResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the media file mutate.
type MutateMediaFileResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMediaFileResult) Reset()         { *m = MutateMediaFileResult{} }
func (m *MutateMediaFileResult) String() string { return proto.CompactTextString(m) }
func (*MutateMediaFileResult) ProtoMessage()    {}
func (*MutateMediaFileResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_793e9e1492d555b0, []int{4}
}

func (m *MutateMediaFileResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMediaFileResult.Unmarshal(m, b)
}
func (m *MutateMediaFileResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMediaFileResult.Marshal(b, m, deterministic)
}
func (m *MutateMediaFileResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMediaFileResult.Merge(m, src)
}
func (m *MutateMediaFileResult) XXX_Size() int {
	return xxx_messageInfo_MutateMediaFileResult.Size(m)
}
func (m *MutateMediaFileResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMediaFileResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMediaFileResult proto.InternalMessageInfo

func (m *MutateMediaFileResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMediaFileRequest)(nil), "google.ads.googleads.v1.services.GetMediaFileRequest")
	proto.RegisterType((*MutateMediaFilesRequest)(nil), "google.ads.googleads.v1.services.MutateMediaFilesRequest")
	proto.RegisterType((*MediaFileOperation)(nil), "google.ads.googleads.v1.services.MediaFileOperation")
	proto.RegisterType((*MutateMediaFilesResponse)(nil), "google.ads.googleads.v1.services.MutateMediaFilesResponse")
	proto.RegisterType((*MutateMediaFileResult)(nil), "google.ads.googleads.v1.services.MutateMediaFileResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/media_file_service.proto", fileDescriptor_793e9e1492d555b0)
}

var fileDescriptor_793e9e1492d555b0 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x77, 0x13, 0xa9, 0x76, 0xd2, 0x6a, 0x99, 0x52, 0x1b, 0xa2, 0x60, 0x58, 0x0b, 0x86, 0x50,
	0x66, 0x49, 0xac, 0x48, 0x57, 0x8b, 0x6c, 0xd0, 0xb4, 0x1e, 0x6a, 0x6b, 0x0a, 0x05, 0x25, 0xb0,
	0x4c, 0x77, 0xa7, 0x71, 0x60, 0x77, 0x67, 0x9d, 0x99, 0x0d, 0x94, 0xd2, 0x8b, 0x5f, 0xc1, 0x6f,
	0xe0, 0xd1, 0xbb, 0x07, 0xbf, 0x42, 0xf1, 0xe6, 0xad, 0x07, 0xf1, 0xe0, 0x41, 0xfc, 0x0c, 0x1e,
	0x64, 0xff, 0xcc, 0x66, 0x93, 0xb6, 0xd4, 0x7a, 0x7b, 0xbc, 0xf7, 0x7e, 0xbf, 0xf7, 0x7b, 0x7f,
	0x66, 0xc0, 0xea, 0x80, 0xb1, 0x81, 0x47, 0x0c, 0xec, 0x0a, 0x23, 0x35, 0x63, 0x6b, 0xd8, 0x32,
	0x04, 0xe1, 0x43, 0xea, 0x10, 0x61, 0xf8, 0xc4, 0xa5, 0xd8, 0xde, 0xa7, 0x1e, 0xb1, 0x33, 0x1f,
	0x0a, 0x39, 0x93, 0x0c, 0xd6, 0xd3, 0x7c, 0x84, 0x5d, 0x81, 0x72, 0x28, 0x1a, 0xb6, 0x90, 0x82,
	0xd6, 0xda, 0xe7, 0x91, 0x73, 0x22, 0x58, 0xc4, 0xc7, 0xd9, 0x53, 0xd6, 0xda, 0x1d, 0x85, 0x09,
	0xa9, 0x81, 0x83, 0x80, 0x49, 0x2c, 0x29, 0x0b, 0x44, 0x16, 0x5d, 0x2c, 0x44, 0x1d, 0x8f, 0x92,
	0x40, 0x66, 0x81, 0xbb, 0x85, 0xc0, 0x3e, 0x25, 0x9e, 0x6b, 0xef, 0x91, 0xb7, 0x78, 0x48, 0x19,
	0x9f, 0x40, 0xf2, 0xd0, 0x31, 0x84, 0xc4, 0x32, 0xca, 0x28, 0xf5, 0xa7, 0x60, 0x7e, 0x9d, 0xc8,
	0xcd, 0x58, 0x47, 0x97, 0x7a, 0xa4, 0x47, 0xde, 0x45, 0x44, 0x48, 0xd8, 0x00, 0xb3, 0x4a, 0xa5,
	0x1d, 0x60, 0x9f, 0x54, 0xb5, 0xba, 0xd6, 0x98, 0xee, 0x94, 0x7f, 0x58, 0xa5, 0xde, 0x8c, 0x8a,
	0xbc, 0xc4, 0x3e, 0xd1, 0x7f, 0x69, 0x60, 0x71, 0x33, 0x92, 0x58, 0x92, 0x9c, 0x44, 0x28, 0x96,
	0x25, 0x50, 0x71, 0x22, 0x21, 0x99, 0x4f, 0xb8, 0x4d, 0xdd, 0x22, 0x07, 0x50, 0xfe, 0x17, 0x2e,
	0x7c, 0x0d, 0x00, 0x0b, 0x09, 0x4f, 0x3b, 0xad, 0x96, 0xea, 0xe5, 0x46, 0xa5, 0xbd, 0x82, 0x2e,
	0x1a, 0x2f, 0xca, 0xcb, 0x6d, 0x29, 0x70, 0x46, 0x3d, 0x22, 0x83, 0xf7, 0xc1, 0xcd, 0x10, 0x73,
	0x49, 0xb1, 0x67, 0xef, 0x63, 0xea, 0x45, 0x9c, 0x54, 0xcb, 0x75, 0xad, 0x71, 0xbd, 0x77, 0x23,
	0x73, 0x77, 0x53, 0x2f, 0xbc, 0x07, 0x66, 0x87, 0xd8, 0xa3, 0x2e, 0x96, 0xc4, 0x66, 0x81, 0x77,
	0x50, 0xbd, 0x9a, 0xa4, 0xcd, 0x28, 0xe7, 0x56, 0xe0, 0x1d, 0xe8, 0x14, 0xc0, 0xd3, 0x45, 0x61,
	0x17, 0x4c, 0x39, 0x9c, 0x60, 0x99, 0xce, 0xa8, 0xd2, 0x5e, 0x3e, 0x57, 0x7a, 0xbe, 0xf7, 0x91,
	0xf6, 0x8d, 0x2b, 0xbd, 0x0c, 0xdd, 0xa9, 0x80, 0xe9, 0x5c, 0xb9, 0xfe, 0x59, 0x03, 0xd5, 0xd3,
	0x53, 0x15, 0x21, 0x0b, 0x04, 0x81, 0x5d, 0xb0, 0x30, 0xd1, 0x95, 0x4d, 0x38, 0x67, 0x3c, 0xe9,
	0xad, 0xd2, 0x86, 0x4a, 0x00, 0x0f, 0x1d, 0xb4, 0x93, 0x2c, 0xbb, 0x37, 0x3f, 0xde, 0xef, 0xf3,
	0x38, 0x1d, 0xbe, 0x02, 0xd7, 0x38, 0x11, 0x91, 0x27, 0xd5, 0xd4, 0x1f, 0xfd, 0xc3, 0xd4, 0xc7,
	0x45, 0xf5, 0x12, 0x7c, 0x4f, 0xf1, 0xe8, 0x4f, 0xc0, 0xc2, 0x99, 0x19, 0xf1, 0x80, 0xcf, 0x38,
	0xa8, 0xf1, 0x5b, 0x6a, 0x7f, 0x2d, 0x83, 0xb9, 0x1c, 0xb8, 0x93, 0x96, 0x84, 0x5f, 0x34, 0x30,
	0x53, 0x3c, 0x51, 0xf8, 0xf0, 0x62, 0x95, 0x67, 0x9c, 0x74, 0xed, 0x52, 0x7b, 0xd1, 0x9f, 0x9d,
	0x58, 0xe3, 0x82, 0xdf, 0x7f, 0xfb, 0xf9, 0xa1, 0x84, 0xe0, 0x72, 0xfc, 0x80, 0x0f, 0xc7, 0x22,
	0x6b, 0xea, 0x96, 0x85, 0xd1, 0x4c, 0x5f, 0x74, 0xb2, 0x2e, 0xa3, 0x79, 0x04, 0xbf, 0x6b, 0x60,
	0x6e, 0x72, 0x8d, 0x70, 0xf5, 0xd2, 0x53, 0x56, 0x0f, 0xaa, 0x66, 0xfe, 0x0f, 0x34, 0xbd, 0x1a,
	0x7d, 0xe7, 0xc4, 0xba, 0x55, 0x78, 0x8d, 0xcb, 0xa3, 0x67, 0x92, 0xb4, 0xb6, 0xa2, 0x1b, 0x71,
	0x6b, 0xa3, 0x5e, 0x0e, 0x0b, 0xc9, 0x6b, 0xcd, 0xa3, 0x42, 0x67, 0xa6, 0x9f, 0xd4, 0x30, 0xb5,
	0x66, 0xed, 0xf6, 0xb1, 0x55, 0x1d, 0xe9, 0xc8, 0xac, 0x90, 0x0a, 0xe4, 0x30, 0xbf, 0xf3, 0x47,
	0x03, 0x4b, 0x0e, 0xf3, 0x2f, 0xd4, 0xdc, 0x59, 0x98, 0x5c, 0xfa, 0x76, 0xfc, 0x37, 0x6d, 0x6b,
	0x6f, 0x36, 0x32, 0xe8, 0x80, 0x79, 0x38, 0x18, 0x20, 0xc6, 0x07, 0xc6, 0x80, 0x04, 0xc9, 0xcf,
	0x65, 0x8c, 0x8a, 0x9d, 0xff, 0x7d, 0x3f, 0x56, 0xc6, 0xc7, 0x52, 0x79, 0xdd, 0xb2, 0x3e, 0x95,
	0xea, 0xeb, 0x29, 0xa1, 0xe5, 0x0a, 0x94, 0x9a, 0xb1, 0xb5, 0xdb, 0x42, 0x59, 0x61, 0x71, 0xac,
	0x52, 0xfa, 0x96, 0x2b, 0xfa, 0x79, 0x4a, 0x7f, 0xb7, 0xd5, 0x57, 0x29, 0xbf, 0x4b, 0x4b, 0xa9,
	0xdf, 0x34, 0x2d, 0x57, 0x98, 0x66, 0x9e, 0x64, 0x9a, 0xbb, 0x2d, 0xd3, 0x54, 0x69, 0x7b, 0x53,
	0x89, 0xce, 0x07, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xd4, 0xd5, 0x0a, 0x65, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MediaFileServiceClient is the client API for MediaFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MediaFileServiceClient interface {
	// Returns the requested media file in full detail.
	GetMediaFile(ctx context.Context, in *GetMediaFileRequest, opts ...grpc.CallOption) (*resources.MediaFile, error)
	// Creates media files. Operation statuses are returned.
	MutateMediaFiles(ctx context.Context, in *MutateMediaFilesRequest, opts ...grpc.CallOption) (*MutateMediaFilesResponse, error)
}

type mediaFileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaFileServiceClient(cc grpc.ClientConnInterface) MediaFileServiceClient {
	return &mediaFileServiceClient{cc}
}

func (c *mediaFileServiceClient) GetMediaFile(ctx context.Context, in *GetMediaFileRequest, opts ...grpc.CallOption) (*resources.MediaFile, error) {
	out := new(resources.MediaFile)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MediaFileService/GetMediaFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaFileServiceClient) MutateMediaFiles(ctx context.Context, in *MutateMediaFilesRequest, opts ...grpc.CallOption) (*MutateMediaFilesResponse, error) {
	out := new(MutateMediaFilesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MediaFileService/MutateMediaFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaFileServiceServer is the server API for MediaFileService service.
type MediaFileServiceServer interface {
	// Returns the requested media file in full detail.
	GetMediaFile(context.Context, *GetMediaFileRequest) (*resources.MediaFile, error)
	// Creates media files. Operation statuses are returned.
	MutateMediaFiles(context.Context, *MutateMediaFilesRequest) (*MutateMediaFilesResponse, error)
}

// UnimplementedMediaFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMediaFileServiceServer struct {
}

func (*UnimplementedMediaFileServiceServer) GetMediaFile(ctx context.Context, req *GetMediaFileRequest) (*resources.MediaFile, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetMediaFile not implemented")
}
func (*UnimplementedMediaFileServiceServer) MutateMediaFiles(ctx context.Context, req *MutateMediaFilesRequest) (*MutateMediaFilesResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateMediaFiles not implemented")
}

func RegisterMediaFileServiceServer(s *grpc.Server, srv MediaFileServiceServer) {
	s.RegisterService(&_MediaFileService_serviceDesc, srv)
}

func _MediaFileService_GetMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).GetMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MediaFileService/GetMediaFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).GetMediaFile(ctx, req.(*GetMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaFileService_MutateMediaFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMediaFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).MutateMediaFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MediaFileService/MutateMediaFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).MutateMediaFiles(ctx, req.(*MutateMediaFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaFileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.MediaFileService",
	HandlerType: (*MediaFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMediaFile",
			Handler:    _MediaFileService_GetMediaFile_Handler,
		},
		{
			MethodName: "MutateMediaFiles",
			Handler:    _MediaFileService_MutateMediaFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/media_file_service.proto",
}
