// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/click_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
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

// A click view with metrics aggregated at each click level, including both
// valid and invalid clicks. For non-Search campaigns, metrics.clicks
// represents the number of valid and invalid interactions.
// Queries including ClickView must have a filter limiting the results to one
// day and can be requested for dates back to 90 days before the time of the
// request.
type ClickView struct {
	// Output only. The resource name of the click view.
	// Click view resource names have the form:
	//
	// `customers/{customer_id}/clickViews/{date (yyyy-MM-dd)}~{gclid}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The Google Click ID.
	Gclid *wrappers.StringValue `protobuf:"bytes,2,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Output only. The location criteria matching the area of interest associated with the
	// impression.
	AreaOfInterest *common.ClickLocation `protobuf:"bytes,3,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	// Output only. The location criteria matching the location of presence associated with the
	// impression.
	LocationOfPresence *common.ClickLocation `protobuf:"bytes,4,opt,name=location_of_presence,json=locationOfPresence,proto3" json:"location_of_presence,omitempty"`
	// Output only. Page number in search results where the ad was shown.
	PageNumber *wrappers.Int64Value `protobuf:"bytes,5,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// Output only. The associated ad.
	AdGroupAd            *wrappers.StringValue `protobuf:"bytes,7,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClickView) Reset()         { *m = ClickView{} }
func (m *ClickView) String() string { return proto.CompactTextString(m) }
func (*ClickView) ProtoMessage()    {}
func (*ClickView) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b0611b0c8085235, []int{0}
}

func (m *ClickView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickView.Unmarshal(m, b)
}
func (m *ClickView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickView.Marshal(b, m, deterministic)
}
func (m *ClickView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickView.Merge(m, src)
}
func (m *ClickView) XXX_Size() int {
	return xxx_messageInfo_ClickView.Size(m)
}
func (m *ClickView) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickView.DiscardUnknown(m)
}

var xxx_messageInfo_ClickView proto.InternalMessageInfo

func (m *ClickView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ClickView) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickView) GetAreaOfInterest() *common.ClickLocation {
	if m != nil {
		return m.AreaOfInterest
	}
	return nil
}

func (m *ClickView) GetLocationOfPresence() *common.ClickLocation {
	if m != nil {
		return m.LocationOfPresence
	}
	return nil
}

func (m *ClickView) GetPageNumber() *wrappers.Int64Value {
	if m != nil {
		return m.PageNumber
	}
	return nil
}

func (m *ClickView) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func init() {
	proto.RegisterType((*ClickView)(nil), "google.ads.googleads.v2.resources.ClickView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/click_view.proto", fileDescriptor_0b0611b0c8085235)
}

var fileDescriptor_0b0611b0c8085235 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x49, 0xd2, 0x2a, 0x9d, 0x68, 0x91, 0xc5, 0xc3, 0x5a, 0x8b, 0xb6, 0xc5, 0x42, 0x11,
	0x9d, 0x85, 0xad, 0x28, 0xac, 0xa7, 0x89, 0x87, 0x50, 0x91, 0x26, 0x44, 0xcc, 0x41, 0x02, 0xcb,
	0x64, 0xe7, 0x65, 0x3b, 0xb8, 0x3b, 0xb3, 0xcc, 0xec, 0x26, 0x07, 0xe9, 0x97, 0xf1, 0xe8, 0x47,
	0xf1, 0x3b, 0x08, 0x3d, 0xf7, 0x23, 0xe8, 0x45, 0x76, 0x77, 0x66, 0x1a, 0x90, 0xda, 0xe2, 0xed,
	0x9f, 0xbc, 0xff, 0xff, 0xf7, 0xde, 0xbc, 0xd9, 0x41, 0x61, 0x2a, 0x65, 0x9a, 0x41, 0x40, 0x99,
	0x0e, 0x5a, 0x59, 0xab, 0x65, 0x18, 0x28, 0xd0, 0xb2, 0x52, 0x09, 0xe8, 0x20, 0xc9, 0x78, 0xf2,
	0x25, 0x5e, 0x72, 0x58, 0xe1, 0x42, 0xc9, 0x52, 0x7a, 0xfb, 0xad, 0x11, 0x53, 0xa6, 0xb1, 0xcb,
	0xe0, 0x65, 0x88, 0x5d, 0x66, 0xe7, 0xf8, 0x3a, 0x6c, 0x22, 0xf3, 0x5c, 0x0a, 0xc3, 0xcc, 0x64,
	0x42, 0x4b, 0x2e, 0x45, 0xcb, 0xdd, 0x79, 0x6a, 0x43, 0x05, 0x0f, 0x16, 0x1c, 0x32, 0x16, 0xcf,
	0xe1, 0x8c, 0x2e, 0xb9, 0x54, 0xc6, 0xf0, 0x68, 0xcd, 0x60, 0x7b, 0x99, 0xd2, 0x13, 0x53, 0x6a,
	0x7e, 0xcd, 0xab, 0x45, 0xb0, 0x52, 0xb4, 0x28, 0x40, 0x69, 0x53, 0xdf, 0x5d, 0x8b, 0x52, 0x21,
	0x64, 0xd9, 0x34, 0x36, 0xd5, 0x83, 0x9f, 0x1b, 0x68, 0xeb, 0x5d, 0x3d, 0xd2, 0x94, 0xc3, 0xca,
	0x1b, 0xa1, 0xfb, 0x96, 0x1e, 0x0b, 0x9a, 0x83, 0xdf, 0xd9, 0xeb, 0x1c, 0x6d, 0x0d, 0x9e, 0x5f,
	0x90, 0xde, 0x2f, 0xf2, 0x0c, 0x1d, 0x5c, 0x9d, 0xd9, 0xa8, 0x82, 0x6b, 0x9c, 0xc8, 0x3c, 0x70,
	0x88, 0xc9, 0x3d, 0x0b, 0x38, 0xa5, 0x39, 0x78, 0x6f, 0xd0, 0x66, 0x9a, 0x64, 0x9c, 0xf9, 0xdd,
	0xbd, 0xce, 0x51, 0x3f, 0xdc, 0x35, 0x39, 0x6c, 0x87, 0xc5, 0x1f, 0x4b, 0xc5, 0x45, 0x3a, 0xa5,
	0x59, 0x05, 0x83, 0xde, 0x05, 0xe9, 0x4d, 0x5a, 0xbf, 0x37, 0x43, 0x0f, 0xa8, 0x02, 0x1a, 0xcb,
	0x45, 0xcc, 0x45, 0x09, 0x0a, 0x74, 0xe9, 0xf7, 0x1a, 0xc6, 0x4b, 0x7c, 0xdd, 0x25, 0xb4, 0x1b,
	0xc6, 0xcd, 0x2c, 0x1f, 0xcc, 0x82, 0x5b, 0xe8, 0x76, 0xcd, 0x1a, 0x2d, 0x4e, 0x0c, 0xc9, 0x63,
	0xe8, 0xa1, 0xbd, 0x81, 0xba, 0x43, 0xa1, 0x40, 0x83, 0x48, 0xc0, 0xdf, 0xf8, 0xef, 0x0e, 0x9e,
	0xe5, 0x8d, 0x16, 0x63, 0x43, 0xf3, 0x08, 0xea, 0x17, 0x34, 0x85, 0x58, 0x54, 0xf9, 0x1c, 0x94,
	0xbf, 0xd9, 0xc0, 0x1f, 0xff, 0xb5, 0x82, 0x13, 0x51, 0xbe, 0x7e, 0xb5, 0xb6, 0x01, 0x54, 0x87,
	0x4e, 0x9b, 0x8c, 0x77, 0x86, 0xfa, 0x94, 0xc5, 0xa9, 0x92, 0x55, 0x11, 0x53, 0xe6, 0xdf, 0xbd,
	0xc5, 0x16, 0x6f, 0xbe, 0x2c, 0xc2, 0x86, 0x35, 0x90, 0xb0, 0xc9, 0x16, 0xb5, 0x32, 0xfa, 0x74,
	0x49, 0x26, 0xb7, 0xb9, 0x60, 0xef, 0x45, 0x52, 0xe9, 0x52, 0xe6, 0xa0, 0x74, 0xf0, 0xd5, 0xca,
	0xf3, 0xf6, 0xb3, 0xae, 0xeb, 0xf5, 0xbf, 0xee, 0xd9, 0x9c, 0x0f, 0x7e, 0x77, 0xd0, 0x61, 0x22,
	0x73, 0x7c, 0xe3, 0xc3, 0x19, 0x6c, 0xbb, 0x16, 0xe3, 0xfa, 0x58, 0xe3, 0xce, 0xe7, 0xf7, 0x26,
	0x94, 0xca, 0x8c, 0x8a, 0x14, 0x4b, 0x95, 0x06, 0x29, 0x88, 0xe6, 0xd0, 0xc1, 0xd5, 0x78, 0xff,
	0x78, 0xc0, 0x6f, 0x9d, 0xfa, 0xd6, 0xed, 0x0d, 0x09, 0xf9, 0xde, 0xdd, 0x1f, 0xb6, 0x48, 0xc2,
	0x34, 0x6e, 0x65, 0xad, 0xa6, 0x21, 0x9e, 0x58, 0xe7, 0x0f, 0xeb, 0x99, 0x11, 0xa6, 0x67, 0xce,
	0x33, 0x9b, 0x86, 0x33, 0xe7, 0xb9, 0xec, 0x1e, 0xb6, 0x85, 0x28, 0x22, 0x4c, 0x47, 0x91, 0x73,
	0x45, 0xd1, 0x34, 0x8c, 0x22, 0xe7, 0x9b, 0xdf, 0x69, 0x86, 0x3d, 0xfe, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x11, 0x08, 0x61, 0x6c, 0x04, 0x00, 0x00,
}
