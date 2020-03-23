// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/hotel_group_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A hotel group view.
type HotelGroupView struct {
	// Output only. The resource name of the hotel group view.
	// Hotel Group view resource names have the form:
	//
	// `customers/{customer_id}/hotelGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelGroupView) Reset()         { *m = HotelGroupView{} }
func (m *HotelGroupView) String() string { return proto.CompactTextString(m) }
func (*HotelGroupView) ProtoMessage()    {}
func (*HotelGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cc5584a36eaf63e, []int{0}
}

func (m *HotelGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelGroupView.Unmarshal(m, b)
}
func (m *HotelGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelGroupView.Marshal(b, m, deterministic)
}
func (m *HotelGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelGroupView.Merge(m, src)
}
func (m *HotelGroupView) XXX_Size() int {
	return xxx_messageInfo_HotelGroupView.Size(m)
}
func (m *HotelGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_HotelGroupView proto.InternalMessageInfo

func (m *HotelGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HotelGroupView)(nil), "google.ads.googleads.v2.resources.HotelGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/hotel_group_view.proto", fileDescriptor_8cc5584a36eaf63e)
}

var fileDescriptor_8cc5584a36eaf63e = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x0a, 0x17, 0x6e, 0xb8, 0xf7, 0x2e, 0x7a, 0x37, 0x5a, 0x04, 0xad, 0x50, 0xd4,
	0xcd, 0x0c, 0xc4, 0x85, 0x32, 0xae, 0xa6, 0x9b, 0x8a, 0x0b, 0x29, 0x45, 0xb2, 0x90, 0x40, 0x98,
	0x26, 0x63, 0x3a, 0x90, 0xe4, 0x0f, 0x33, 0x69, 0xba, 0x28, 0x05, 0x9f, 0xc5, 0xa5, 0x8f, 0x22,
	0xf8, 0x0e, 0xae, 0xfb, 0x08, 0xae, 0x24, 0x9d, 0xce, 0xb4, 0x75, 0xa1, 0xee, 0x0e, 0xfc, 0xdf,
	0x39, 0x73, 0x38, 0xe3, 0x5d, 0xa6, 0x00, 0x69, 0xc6, 0x31, 0x4b, 0x14, 0xd6, 0xb2, 0x51, 0xb5,
	0x8f, 0x25, 0x57, 0x30, 0x95, 0x31, 0x57, 0x78, 0x02, 0x15, 0xcf, 0xa2, 0x54, 0xc2, 0xb4, 0x8c,
	0x6a, 0xc1, 0x67, 0xa8, 0x94, 0x50, 0x41, 0xbb, 0xab, 0x71, 0xc4, 0x12, 0x85, 0xac, 0x13, 0xd5,
	0x3e, 0xb2, 0xce, 0xce, 0xa1, 0x09, 0x2f, 0x05, 0x7e, 0x10, 0x3c, 0x4b, 0xa2, 0x31, 0x9f, 0xb0,
	0x5a, 0x80, 0xd4, 0x19, 0x9d, 0xfd, 0x2d, 0xc0, 0xd8, 0xd6, 0xa7, 0x83, 0xad, 0x13, 0x2b, 0x0a,
	0xa8, 0x58, 0x25, 0xa0, 0x50, 0xfa, 0x7a, 0xfc, 0xea, 0x78, 0xff, 0xae, 0x9b, 0x5e, 0x83, 0xa6,
	0x56, 0x20, 0xf8, 0xac, 0x7d, 0xe7, 0xfd, 0x35, 0x11, 0x51, 0xc1, 0x72, 0xbe, 0xe7, 0x1c, 0x39,
	0xa7, 0xbf, 0xfb, 0xf8, 0x8d, 0xb6, 0xde, 0xe9, 0x99, 0x77, 0xb2, 0xe9, 0xb8, 0x56, 0xa5, 0x50,
	0x28, 0x86, 0x1c, 0xef, 0xe6, 0x8c, 0xfe, 0x98, 0x94, 0x5b, 0x96, 0x73, 0xc2, 0x97, 0x74, 0xfc,
	0x63, 0x6f, 0xfb, 0x22, 0x9e, 0xaa, 0x0a, 0x72, 0x2e, 0x15, 0x9e, 0x1b, 0xb9, 0xd0, 0x03, 0x5a,
	0x48, 0xe1, 0xf9, 0xe7, 0x45, 0x17, 0xfd, 0x47, 0xd7, 0xeb, 0xc5, 0x90, 0xa3, 0x6f, 0x37, 0xed,
	0xff, 0xdf, 0x7d, 0x72, 0xd8, 0xcc, 0x31, 0x74, 0xee, 0x6f, 0xd6, 0xce, 0x14, 0x32, 0x56, 0xa4,
	0x08, 0x64, 0x8a, 0x53, 0x5e, 0xac, 0xc6, 0xc2, 0x9b, 0xce, 0x5f, 0x7c, 0xf3, 0x95, 0x55, 0x4f,
	0x6e, 0x6b, 0x40, 0xe9, 0xb3, 0xdb, 0x1d, 0xe8, 0x48, 0x9a, 0x28, 0xa4, 0x65, 0xa3, 0x02, 0x1f,
	0x8d, 0x0c, 0xf9, 0x62, 0x98, 0x90, 0x26, 0x2a, 0xb4, 0x4c, 0x18, 0xf8, 0xa1, 0x65, 0x96, 0x6e,
	0x4f, 0x1f, 0x08, 0xa1, 0x89, 0x22, 0xc4, 0x52, 0x84, 0x04, 0x3e, 0x21, 0x96, 0x1b, 0xff, 0x5a,
	0x95, 0x3d, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x03, 0xf8, 0x8c, 0x92, 0x02, 0x00, 0x00,
}
