// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/carrier_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A carrier criterion that can be used in campaign targeting.
type CarrierConstant struct {
	// Output only. The resource name of the carrier criterion.
	// Carrier criterion resource names have the form:
	//
	// `carrierConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the carrier criterion.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The full name of the carrier in English.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The country code of the country where the carrier is located, e.g., "AR",
	// "FR", etc.
	CountryCode          *wrappers.StringValue `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CarrierConstant) Reset()         { *m = CarrierConstant{} }
func (m *CarrierConstant) String() string { return proto.CompactTextString(m) }
func (*CarrierConstant) ProtoMessage()    {}
func (*CarrierConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b8b1da6512447a4, []int{0}
}

func (m *CarrierConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarrierConstant.Unmarshal(m, b)
}
func (m *CarrierConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarrierConstant.Marshal(b, m, deterministic)
}
func (m *CarrierConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarrierConstant.Merge(m, src)
}
func (m *CarrierConstant) XXX_Size() int {
	return xxx_messageInfo_CarrierConstant.Size(m)
}
func (m *CarrierConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_CarrierConstant.DiscardUnknown(m)
}

var xxx_messageInfo_CarrierConstant proto.InternalMessageInfo

func (m *CarrierConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CarrierConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CarrierConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CarrierConstant) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*CarrierConstant)(nil), "google.ads.googleads.v1.resources.CarrierConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/carrier_constant.proto", fileDescriptor_6b8b1da6512447a4)
}

var fileDescriptor_6b8b1da6512447a4 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0xd9, 0x4c, 0x11, 0x4c, 0x2b, 0xc2, 0xe0, 0x61, 0xad, 0x45, 0x5b, 0xa5, 0xb0, 0x78,
	0x48, 0x3a, 0xfe, 0x43, 0xe2, 0x29, 0xbb, 0x42, 0xd1, 0x83, 0x96, 0x15, 0xf7, 0x20, 0x0b, 0x4b,
	0x36, 0x49, 0xc7, 0xc0, 0x4e, 0xde, 0x21, 0xc9, 0xac, 0x88, 0x78, 0xf0, 0xab, 0x78, 0xf4, 0xa3,
	0xf8, 0x29, 0x7a, 0xae, 0xdf, 0xc0, 0x93, 0xec, 0x4c, 0x66, 0xba, 0xae, 0x60, 0x7b, 0x7b, 0x66,
	0xde, 0xe7, 0xf7, 0xe4, 0x79, 0xe1, 0xc5, 0xcf, 0x73, 0x80, 0x7c, 0xa1, 0xa9, 0x50, 0x9e, 0x36,
	0x72, 0xa5, 0x96, 0x19, 0x75, 0xda, 0x43, 0xe5, 0xa4, 0xf6, 0x54, 0x0a, 0xe7, 0x8c, 0x76, 0x33,
	0x09, 0xd6, 0x07, 0x61, 0x03, 0x29, 0x1d, 0x04, 0x48, 0x0f, 0x1a, 0x3b, 0x11, 0xca, 0x93, 0x8e,
	0x24, 0xcb, 0x8c, 0x74, 0xe4, 0xee, 0xbd, 0x36, 0xbc, 0x34, 0xf4, 0xd4, 0xe8, 0x85, 0x9a, 0xcd,
	0xf5, 0x47, 0xb1, 0x34, 0xe0, 0x9a, 0x8c, 0xdd, 0xdb, 0x6b, 0x86, 0x16, 0x8b, 0xa3, 0xbb, 0x71,
	0x54, 0x7f, 0xcd, 0xab, 0x53, 0xfa, 0xc9, 0x89, 0xb2, 0xd4, 0xce, 0xc7, 0xf9, 0xde, 0x1a, 0x2a,
	0xac, 0x85, 0x20, 0x82, 0x01, 0x1b, 0xa7, 0xf7, 0x7f, 0x21, 0x7c, 0x73, 0xd4, 0xf4, 0x1e, 0xc5,
	0xda, 0xe9, 0x7b, 0x7c, 0xa3, 0x7d, 0x63, 0x66, 0x45, 0xa1, 0xfb, 0xbd, 0xfd, 0xde, 0xe0, 0xfa,
	0xf0, 0xe8, 0x8c, 0x27, 0xbf, 0xf9, 0x43, 0x3c, 0xb8, 0x58, 0x22, 0xaa, 0xd2, 0x78, 0x22, 0xa1,
	0xa0, 0x1b, 0x41, 0xe3, 0x9d, 0x36, 0xe6, 0x8d, 0x28, 0x74, 0x7a, 0x84, 0x91, 0x51, 0x7d, 0xb4,
	0xdf, 0x1b, 0x6c, 0x3f, 0xba, 0x13, 0x51, 0xd2, 0xb6, 0x26, 0xaf, 0x6c, 0x78, 0xf6, 0x64, 0x22,
	0x16, 0x95, 0x1e, 0x26, 0x67, 0x3c, 0x19, 0x23, 0xa3, 0xd2, 0xa7, 0x78, 0xab, 0x7e, 0x3f, 0xa9,
	0x99, 0xbd, 0x7f, 0x98, 0x77, 0xc1, 0x19, 0x9b, 0xaf, 0x41, 0xb5, 0x3d, 0x7d, 0x89, 0x77, 0x24,
	0x54, 0x36, 0xb8, 0xcf, 0x33, 0x09, 0x4a, 0xf7, 0xb7, 0xae, 0x8a, 0x6f, 0x47, 0x6c, 0x04, 0x4a,
	0xb3, 0xf1, 0x39, 0x7f, 0x7b, 0xf5, 0x5d, 0xd3, 0x07, 0xf2, 0xef, 0x1f, 0x9e, 0x7e, 0xd9, 0xbc,
	0x87, 0xaf, 0xc3, 0x6f, 0x08, 0x1f, 0x4a, 0x28, 0xc8, 0xa5, 0x17, 0x31, 0xbc, 0xb5, 0x91, 0x7f,
	0xb2, 0x2a, 0x7d, 0xd2, 0xfb, 0xf0, 0x3a, 0xa2, 0x39, 0x2c, 0x84, 0xcd, 0x09, 0xb8, 0x9c, 0xe6,
	0xda, 0xd6, 0x2b, 0xd1, 0x8b, 0x86, 0xff, 0xb9, 0xd2, 0x17, 0x9d, 0xfa, 0x8e, 0x92, 0x63, 0xce,
	0x7f, 0xa0, 0x83, 0xe3, 0x26, 0x92, 0x2b, 0x4f, 0x1a, 0xb9, 0x52, 0x93, 0x8c, 0x8c, 0x5b, 0xe7,
	0xcf, 0xd6, 0x33, 0xe5, 0xca, 0x4f, 0x3b, 0xcf, 0x74, 0x92, 0x4d, 0x3b, 0xcf, 0x39, 0x3a, 0x6c,
	0x06, 0x8c, 0x71, 0xe5, 0x19, 0xeb, 0x5c, 0x8c, 0x4d, 0x32, 0xc6, 0x3a, 0xdf, 0xfc, 0x5a, 0x5d,
	0xf6, 0xf1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xd5, 0xcc, 0xd1, 0x51, 0x03, 0x00, 0x00,
}
