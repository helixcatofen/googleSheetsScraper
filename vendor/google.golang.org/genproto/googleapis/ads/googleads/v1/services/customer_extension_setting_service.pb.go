// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/customer_extension_setting_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for
// [CustomerExtensionSettingService.GetCustomerExtensionSetting][google.ads.googleads.v1.services.CustomerExtensionSettingService.GetCustomerExtensionSetting].
type GetCustomerExtensionSettingRequest struct {
	// Required. The resource name of the customer extension setting to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerExtensionSettingRequest) Reset()         { *m = GetCustomerExtensionSettingRequest{} }
func (m *GetCustomerExtensionSettingRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerExtensionSettingRequest) ProtoMessage()    {}
func (*GetCustomerExtensionSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f039497e3ad343eb, []int{0}
}

func (m *GetCustomerExtensionSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerExtensionSettingRequest.Unmarshal(m, b)
}
func (m *GetCustomerExtensionSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerExtensionSettingRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerExtensionSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerExtensionSettingRequest.Merge(m, src)
}
func (m *GetCustomerExtensionSettingRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerExtensionSettingRequest.Size(m)
}
func (m *GetCustomerExtensionSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerExtensionSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerExtensionSettingRequest proto.InternalMessageInfo

func (m *GetCustomerExtensionSettingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerExtensionSettingService.MutateCustomerExtensionSettings][google.ads.googleads.v1.services.CustomerExtensionSettingService.MutateCustomerExtensionSettings].
type MutateCustomerExtensionSettingsRequest struct {
	// Required. The ID of the customer whose customer extension settings are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual customer extension
	// settings.
	Operations []*CustomerExtensionSettingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCustomerExtensionSettingsRequest) Reset() {
	*m = MutateCustomerExtensionSettingsRequest{}
}
func (m *MutateCustomerExtensionSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerExtensionSettingsRequest) ProtoMessage()    {}
func (*MutateCustomerExtensionSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f039497e3ad343eb, []int{1}
}

func (m *MutateCustomerExtensionSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Unmarshal(m, b)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Merge(m, src)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerExtensionSettingsRequest.Size(m)
}
func (m *MutateCustomerExtensionSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerExtensionSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerExtensionSettingsRequest proto.InternalMessageInfo

func (m *MutateCustomerExtensionSettingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerExtensionSettingsRequest) GetOperations() []*CustomerExtensionSettingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCustomerExtensionSettingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCustomerExtensionSettingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a customer extension setting.
type CustomerExtensionSettingOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerExtensionSettingOperation_Create
	//	*CustomerExtensionSettingOperation_Update
	//	*CustomerExtensionSettingOperation_Remove
	Operation            isCustomerExtensionSettingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CustomerExtensionSettingOperation) Reset()         { *m = CustomerExtensionSettingOperation{} }
func (m *CustomerExtensionSettingOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerExtensionSettingOperation) ProtoMessage()    {}
func (*CustomerExtensionSettingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f039497e3ad343eb, []int{2}
}

func (m *CustomerExtensionSettingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerExtensionSettingOperation.Unmarshal(m, b)
}
func (m *CustomerExtensionSettingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerExtensionSettingOperation.Marshal(b, m, deterministic)
}
func (m *CustomerExtensionSettingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerExtensionSettingOperation.Merge(m, src)
}
func (m *CustomerExtensionSettingOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerExtensionSettingOperation.Size(m)
}
func (m *CustomerExtensionSettingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerExtensionSettingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerExtensionSettingOperation proto.InternalMessageInfo

func (m *CustomerExtensionSettingOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCustomerExtensionSettingOperation_Operation interface {
	isCustomerExtensionSettingOperation_Operation()
}

type CustomerExtensionSettingOperation_Create struct {
	Create *resources.CustomerExtensionSetting `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerExtensionSettingOperation_Update struct {
	Update *resources.CustomerExtensionSetting `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CustomerExtensionSettingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CustomerExtensionSettingOperation_Create) isCustomerExtensionSettingOperation_Operation() {}

func (*CustomerExtensionSettingOperation_Update) isCustomerExtensionSettingOperation_Operation() {}

func (*CustomerExtensionSettingOperation_Remove) isCustomerExtensionSettingOperation_Operation() {}

func (m *CustomerExtensionSettingOperation) GetOperation() isCustomerExtensionSettingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerExtensionSettingOperation) GetCreate() *resources.CustomerExtensionSetting {
	if x, ok := m.GetOperation().(*CustomerExtensionSettingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerExtensionSettingOperation) GetUpdate() *resources.CustomerExtensionSetting {
	if x, ok := m.GetOperation().(*CustomerExtensionSettingOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CustomerExtensionSettingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CustomerExtensionSettingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerExtensionSettingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerExtensionSettingOperation_Create)(nil),
		(*CustomerExtensionSettingOperation_Update)(nil),
		(*CustomerExtensionSettingOperation_Remove)(nil),
	}
}

// Response message for a customer extension setting mutate.
type MutateCustomerExtensionSettingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCustomerExtensionSettingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MutateCustomerExtensionSettingsResponse) Reset() {
	*m = MutateCustomerExtensionSettingsResponse{}
}
func (m *MutateCustomerExtensionSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerExtensionSettingsResponse) ProtoMessage()    {}
func (*MutateCustomerExtensionSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f039497e3ad343eb, []int{3}
}

func (m *MutateCustomerExtensionSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Unmarshal(m, b)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Merge(m, src)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerExtensionSettingsResponse.Size(m)
}
func (m *MutateCustomerExtensionSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerExtensionSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerExtensionSettingsResponse proto.InternalMessageInfo

func (m *MutateCustomerExtensionSettingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCustomerExtensionSettingsResponse) GetResults() []*MutateCustomerExtensionSettingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the customer extension setting mutate.
type MutateCustomerExtensionSettingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerExtensionSettingResult) Reset()         { *m = MutateCustomerExtensionSettingResult{} }
func (m *MutateCustomerExtensionSettingResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerExtensionSettingResult) ProtoMessage()    {}
func (*MutateCustomerExtensionSettingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f039497e3ad343eb, []int{4}
}

func (m *MutateCustomerExtensionSettingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerExtensionSettingResult.Unmarshal(m, b)
}
func (m *MutateCustomerExtensionSettingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerExtensionSettingResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerExtensionSettingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerExtensionSettingResult.Merge(m, src)
}
func (m *MutateCustomerExtensionSettingResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerExtensionSettingResult.Size(m)
}
func (m *MutateCustomerExtensionSettingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerExtensionSettingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerExtensionSettingResult proto.InternalMessageInfo

func (m *MutateCustomerExtensionSettingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerExtensionSettingRequest)(nil), "google.ads.googleads.v1.services.GetCustomerExtensionSettingRequest")
	proto.RegisterType((*MutateCustomerExtensionSettingsRequest)(nil), "google.ads.googleads.v1.services.MutateCustomerExtensionSettingsRequest")
	proto.RegisterType((*CustomerExtensionSettingOperation)(nil), "google.ads.googleads.v1.services.CustomerExtensionSettingOperation")
	proto.RegisterType((*MutateCustomerExtensionSettingsResponse)(nil), "google.ads.googleads.v1.services.MutateCustomerExtensionSettingsResponse")
	proto.RegisterType((*MutateCustomerExtensionSettingResult)(nil), "google.ads.googleads.v1.services.MutateCustomerExtensionSettingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/customer_extension_setting_service.proto", fileDescriptor_f039497e3ad343eb)
}

var fileDescriptor_f039497e3ad343eb = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xbf, 0x6f, 0xdb, 0x46,
	0x14, 0xae, 0x28, 0xc3, 0xad, 0x4f, 0x76, 0x0b, 0x5c, 0xd1, 0x56, 0x90, 0x0b, 0x58, 0x65, 0x85,
	0x5a, 0x10, 0x0a, 0x12, 0x52, 0x37, 0x0a, 0x2e, 0x40, 0xb9, 0x96, 0x6d, 0x14, 0xfe, 0x01, 0x0a,
	0xf5, 0x50, 0x08, 0x60, 0x4f, 0xe4, 0x59, 0x26, 0x4c, 0xf2, 0xd8, 0xbb, 0xa3, 0x10, 0xc3, 0xf0,
	0x92, 0x21, 0x4b, 0xc6, 0x6c, 0x19, 0x33, 0x66, 0xcf, 0x3f, 0xe1, 0xd5, 0x9b, 0xa7, 0x0c, 0x99,
	0xb2, 0x67, 0x0f, 0xc8, 0xe3, 0xe9, 0x87, 0x13, 0x9a, 0x01, 0xec, 0xed, 0xe9, 0xde, 0xc7, 0xef,
	0x7d, 0xf7, 0xde, 0x77, 0x4f, 0x60, 0x7f, 0x4c, 0xc8, 0xd8, 0xc7, 0x3a, 0x72, 0x99, 0x2e, 0xc2,
	0x24, 0x9a, 0xb4, 0x75, 0x86, 0xe9, 0xc4, 0x73, 0x30, 0xd3, 0x9d, 0x98, 0x71, 0x12, 0x60, 0x6a,
	0xe3, 0x27, 0x1c, 0x87, 0xcc, 0x23, 0xa1, 0xcd, 0x30, 0xe7, 0x5e, 0x38, 0xb6, 0x33, 0x8c, 0x16,
	0x51, 0xc2, 0x09, 0xac, 0x8b, 0xef, 0x35, 0xe4, 0x32, 0x6d, 0x4a, 0xa5, 0x4d, 0xda, 0x9a, 0xa4,
	0xaa, 0xf5, 0xf2, 0x8a, 0x51, 0xcc, 0x48, 0x4c, 0xef, 0xaf, 0x26, 0xaa, 0xd4, 0x7e, 0x96, 0x1c,
	0x91, 0xa7, 0xa3, 0x30, 0x24, 0x1c, 0x71, 0x8f, 0x84, 0x2c, 0xcb, 0xfe, 0x34, 0x97, 0x75, 0x7c,
	0x0f, 0x87, 0x3c, 0x4b, 0x6c, 0xcc, 0x25, 0x4e, 0x3d, 0xec, 0xbb, 0xf6, 0x08, 0x9f, 0xa1, 0x89,
	0x47, 0x68, 0x06, 0xc8, 0xd4, 0xeb, 0xe9, 0xaf, 0x51, 0x7c, 0x9a, 0xa1, 0x02, 0xc4, 0xce, 0xef,
	0x70, 0xd3, 0xc8, 0xd1, 0x19, 0x47, 0x3c, 0xce, 0x8a, 0xaa, 0x87, 0x40, 0xdd, 0xc5, 0x7c, 0x3b,
	0x53, 0xbe, 0x23, 0x85, 0x0f, 0x84, 0x6e, 0x0b, 0xff, 0x1f, 0x63, 0xc6, 0x61, 0x13, 0xac, 0xc9,
	0x6b, 0xda, 0x21, 0x0a, 0x70, 0xb5, 0x54, 0x2f, 0x35, 0x57, 0x7a, 0xe5, 0xb7, 0xa6, 0x62, 0xad,
	0xca, 0xcc, 0x21, 0x0a, 0xb0, 0xfa, 0x4c, 0x01, 0xbf, 0x1d, 0xc4, 0x1c, 0x71, 0x9c, 0xc7, 0xc9,
	0x24, 0x69, 0x03, 0x54, 0xa6, 0x1d, 0xf3, 0xdc, 0x79, 0x4a, 0x20, 0xcf, 0xf7, 0x5d, 0x78, 0x06,
	0x00, 0x89, 0x30, 0x15, 0x9d, 0xaa, 0x2a, 0xf5, 0x72, 0xb3, 0xd2, 0xd9, 0xd6, 0x8a, 0xc6, 0xa5,
	0xe5, 0x55, 0x3f, 0x92, 0x5c, 0x59, 0xa5, 0x19, 0x37, 0xdc, 0x04, 0xdf, 0x45, 0x88, 0x72, 0x0f,
	0xf9, 0xf6, 0x29, 0xf2, 0xfc, 0x98, 0xe2, 0x6a, 0xb9, 0x5e, 0x6a, 0x7e, 0x63, 0x7d, 0x9b, 0x1d,
	0xf7, 0xc5, 0x29, 0xfc, 0x15, 0xac, 0x4d, 0x90, 0xef, 0xb9, 0x88, 0x63, 0x9b, 0x84, 0xfe, 0x45,
	0x75, 0x29, 0x85, 0xad, 0xca, 0xc3, 0xa3, 0xd0, 0xbf, 0x50, 0xdf, 0x28, 0xe0, 0x97, 0x42, 0x11,
	0xb0, 0x0b, 0x2a, 0x71, 0x94, 0x12, 0x25, 0xc3, 0x4a, 0x89, 0x2a, 0x9d, 0x9a, 0xbc, 0x9e, 0x9c,
	0xa7, 0xd6, 0x4f, 0xe6, 0x79, 0x80, 0xd8, 0xb9, 0x05, 0x04, 0x3c, 0x89, 0xe1, 0x3f, 0x60, 0xd9,
	0xa1, 0x18, 0x71, 0x31, 0x8e, 0x4a, 0xa7, 0x9b, 0xdb, 0x96, 0xa9, 0x47, 0x73, 0xfb, 0xb2, 0xf7,
	0x95, 0x95, 0x91, 0x25, 0xb4, 0xa2, 0x48, 0x55, 0x79, 0x14, 0x5a, 0x41, 0x06, 0xab, 0x60, 0x99,
	0xe2, 0x80, 0x4c, 0x44, 0x57, 0x57, 0x92, 0x8c, 0xf8, 0xdd, 0xab, 0x80, 0x95, 0xe9, 0x18, 0xd4,
	0x9b, 0x12, 0xd8, 0x2c, 0x34, 0x10, 0x8b, 0x48, 0xc8, 0x30, 0xec, 0x83, 0x1f, 0xee, 0x4c, 0xcc,
	0xc6, 0x94, 0x12, 0x9a, 0x56, 0xa8, 0x74, 0xa0, 0x14, 0x4e, 0x23, 0x47, 0x1b, 0xa4, 0xae, 0xb7,
	0xbe, 0x5f, 0x9c, 0xe5, 0x4e, 0x02, 0x87, 0xff, 0x81, 0xaf, 0x29, 0x66, 0xb1, 0xcf, 0xa5, 0xc1,
	0xfa, 0xc5, 0x06, 0xbb, 0x5f, 0xa3, 0x95, 0xd2, 0x59, 0x92, 0x56, 0xfd, 0x1b, 0x34, 0xbe, 0xe4,
	0x83, 0xc4, 0x5a, 0x9f, 0x79, 0x68, 0x8b, 0x6f, 0xac, 0x73, 0xb3, 0x04, 0x36, 0xf2, 0x78, 0x06,
	0x42, 0x1f, 0xfc, 0x50, 0x02, 0xeb, 0xf7, 0x3c, 0x6c, 0xf8, 0x57, 0xf1, 0x0d, 0x8b, 0xf7, 0x42,
	0xed, 0x21, 0xd6, 0x50, 0x07, 0xb7, 0xe6, 0xe2, 0x65, 0x9f, 0xde, 0xbc, 0x7b, 0xa1, 0x6c, 0xc1,
	0x6e, 0xb2, 0x55, 0x2f, 0x17, 0x32, 0x5b, 0x72, 0x21, 0x30, 0xbd, 0x35, 0x5d, 0xb3, 0x9f, 0xf8,
	0x42, 0x6f, 0x5d, 0xc1, 0x97, 0x0a, 0xd8, 0x28, 0xb0, 0x0f, 0xdc, 0x7b, 0xe8, 0x74, 0xe5, 0x0a,
	0xab, 0xed, 0x3f, 0x02, 0x93, 0xf0, 0xb2, 0x3a, 0xba, 0x35, 0x7f, 0x9c, 0x5b, 0x87, 0xbf, 0xcf,
	0x16, 0x53, 0xda, 0x96, 0x6d, 0xf5, 0xcf, 0xa4, 0x2d, 0xb3, 0x3e, 0x5c, 0xce, 0x81, 0xb7, 0x5a,
	0x57, 0xf9, 0x5d, 0x31, 0x82, 0x54, 0x81, 0x51, 0x6a, 0xd5, 0xd6, 0xaf, 0xcd, 0xea, 0x4c, 0x65,
	0x16, 0x45, 0x1e, 0xd3, 0x1c, 0x12, 0xf4, 0x9e, 0x2b, 0xa0, 0xe1, 0x90, 0xa0, 0xf0, 0x46, 0xbd,
	0x46, 0x81, 0xf7, 0x8e, 0x93, 0xad, 0x75, 0x5c, 0xfa, 0x77, 0x2f, 0x63, 0x1a, 0x13, 0x1f, 0x85,
	0x63, 0x8d, 0xd0, 0xb1, 0x3e, 0xc6, 0x61, 0xba, 0xd3, 0xf4, 0x59, 0xed, 0xfc, 0x7f, 0xef, 0xae,
	0x0c, 0x5e, 0x29, 0xe5, 0x5d, 0xd3, 0x7c, 0xad, 0xd4, 0x77, 0x05, 0xa1, 0xe9, 0x32, 0x4d, 0x84,
	0x49, 0x74, 0xd2, 0xd6, 0xb2, 0xc2, 0xec, 0x5a, 0x42, 0x86, 0xa6, 0xcb, 0x86, 0x53, 0xc8, 0xf0,
	0xa4, 0x3d, 0x94, 0x90, 0xf7, 0x4a, 0x43, 0x9c, 0x1b, 0x86, 0xe9, 0x32, 0xc3, 0x98, 0x82, 0x0c,
	0xe3, 0xa4, 0x6d, 0x18, 0x12, 0x36, 0x5a, 0x4e, 0x75, 0xfe, 0xf1, 0x31, 0x00, 0x00, 0xff, 0xff,
	0x47, 0xfa, 0xcf, 0x92, 0x64, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerExtensionSettingServiceClient is the client API for CustomerExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerExtensionSettingServiceClient interface {
	// Returns the requested customer extension setting in full detail.
	GetCustomerExtensionSetting(ctx context.Context, in *GetCustomerExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CustomerExtensionSetting, error)
	// Creates, updates, or removes customer extension settings. Operation
	// statuses are returned.
	MutateCustomerExtensionSettings(ctx context.Context, in *MutateCustomerExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCustomerExtensionSettingsResponse, error)
}

type customerExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerExtensionSettingServiceClient(cc grpc.ClientConnInterface) CustomerExtensionSettingServiceClient {
	return &customerExtensionSettingServiceClient{cc}
}

func (c *customerExtensionSettingServiceClient) GetCustomerExtensionSetting(ctx context.Context, in *GetCustomerExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CustomerExtensionSetting, error) {
	out := new(resources.CustomerExtensionSetting)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerExtensionSettingService/GetCustomerExtensionSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerExtensionSettingServiceClient) MutateCustomerExtensionSettings(ctx context.Context, in *MutateCustomerExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCustomerExtensionSettingsResponse, error) {
	out := new(MutateCustomerExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerExtensionSettingService/MutateCustomerExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerExtensionSettingServiceServer is the server API for CustomerExtensionSettingService service.
type CustomerExtensionSettingServiceServer interface {
	// Returns the requested customer extension setting in full detail.
	GetCustomerExtensionSetting(context.Context, *GetCustomerExtensionSettingRequest) (*resources.CustomerExtensionSetting, error)
	// Creates, updates, or removes customer extension settings. Operation
	// statuses are returned.
	MutateCustomerExtensionSettings(context.Context, *MutateCustomerExtensionSettingsRequest) (*MutateCustomerExtensionSettingsResponse, error)
}

// UnimplementedCustomerExtensionSettingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerExtensionSettingServiceServer struct {
}

func (*UnimplementedCustomerExtensionSettingServiceServer) GetCustomerExtensionSetting(ctx context.Context, req *GetCustomerExtensionSettingRequest) (*resources.CustomerExtensionSetting, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerExtensionSetting not implemented")
}
func (*UnimplementedCustomerExtensionSettingServiceServer) MutateCustomerExtensionSettings(ctx context.Context, req *MutateCustomerExtensionSettingsRequest) (*MutateCustomerExtensionSettingsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCustomerExtensionSettings not implemented")
}

func RegisterCustomerExtensionSettingServiceServer(s *grpc.Server, srv CustomerExtensionSettingServiceServer) {
	s.RegisterService(&_CustomerExtensionSettingService_serviceDesc, srv)
}

func _CustomerExtensionSettingService_GetCustomerExtensionSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerExtensionSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerExtensionSettingServiceServer).GetCustomerExtensionSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerExtensionSettingService/GetCustomerExtensionSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerExtensionSettingServiceServer).GetCustomerExtensionSetting(ctx, req.(*GetCustomerExtensionSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerExtensionSettingService_MutateCustomerExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerExtensionSettingServiceServer).MutateCustomerExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerExtensionSettingService/MutateCustomerExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerExtensionSettingServiceServer).MutateCustomerExtensionSettings(ctx, req.(*MutateCustomerExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerExtensionSettingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CustomerExtensionSettingService",
	HandlerType: (*CustomerExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerExtensionSetting",
			Handler:    _CustomerExtensionSettingService_GetCustomerExtensionSetting_Handler,
		},
		{
			MethodName: "MutateCustomerExtensionSettings",
			Handler:    _CustomerExtensionSettingService_MutateCustomerExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/customer_extension_setting_service.proto",
}