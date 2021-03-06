// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/user_list.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A user list. This is a list of users a customer may target.
type UserList struct {
	// Immutable. The resource name of the user list.
	// User list resource names have the form:
	//
	// `customers/{customer_id}/userLists/{user_list_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Id of the user list.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. A flag that indicates if a user may edit a list. Depends on the list
	// ownership and list type. For example, external remarketing user lists are
	// not editable.
	//
	// This field is read-only.
	ReadOnly *wrappers.BoolValue `protobuf:"bytes,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// Name of this user list. Depending on its access_reason, the user list name
	// may not be unique (e.g. if access_reason=SHARED)
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of this user list.
	Description *wrappers.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Membership status of this user list. Indicates whether a user list is open
	// or active. Only open user lists can accumulate more users and can be
	// targeted to.
	MembershipStatus enums.UserListMembershipStatusEnum_UserListMembershipStatus `protobuf:"varint,6,opt,name=membership_status,json=membershipStatus,proto3,enum=google.ads.googleads.v3.enums.UserListMembershipStatusEnum_UserListMembershipStatus" json:"membership_status,omitempty"`
	// An ID from external system. It is used by user list sellers to correlate
	// IDs on their systems.
	IntegrationCode *wrappers.StringValue `protobuf:"bytes,7,opt,name=integration_code,json=integrationCode,proto3" json:"integration_code,omitempty"`
	// Number of days a user's cookie stays on your list since its most recent
	// addition to the list. This field must be between 0 and 540 inclusive.
	// However, for CRM based userlists, this field can be set to 10000 which
	// means no expiration.
	//
	// It'll be ignored for logical_user_list.
	MembershipLifeSpan *wrappers.Int64Value `protobuf:"bytes,8,opt,name=membership_life_span,json=membershipLifeSpan,proto3" json:"membership_life_span,omitempty"`
	// Output only. Estimated number of users in this user list, on the Google Display Network.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForDisplay *wrappers.Int64Value `protobuf:"bytes,9,opt,name=size_for_display,json=sizeForDisplay,proto3" json:"size_for_display,omitempty"`
	// Output only. Size range in terms of number of users of the UserList, on the Google
	// Display Network.
	//
	// This field is read-only.
	SizeRangeForDisplay enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,10,opt,name=size_range_for_display,json=sizeRangeForDisplay,proto3,enum=google.ads.googleads.v3.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_display,omitempty"`
	// Output only. Estimated number of users in this user list in the google.com domain.
	// These are the users available for targeting in Search campaigns.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForSearch *wrappers.Int64Value `protobuf:"bytes,11,opt,name=size_for_search,json=sizeForSearch,proto3" json:"size_for_search,omitempty"`
	// Output only. Size range in terms of number of users of the UserList, for Search ads.
	//
	// This field is read-only.
	SizeRangeForSearch enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,12,opt,name=size_range_for_search,json=sizeRangeForSearch,proto3,enum=google.ads.googleads.v3.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_search,omitempty"`
	// Output only. Type of this list.
	//
	// This field is read-only.
	Type enums.UserListTypeEnum_UserListType `protobuf:"varint,13,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.UserListTypeEnum_UserListType" json:"type,omitempty"`
	// Indicating the reason why this user list membership status is closed. It is
	// only populated on lists that were automatically closed due to inactivity,
	// and will be cleared once the list membership status becomes open.
	ClosingReason enums.UserListClosingReasonEnum_UserListClosingReason `protobuf:"varint,14,opt,name=closing_reason,json=closingReason,proto3,enum=google.ads.googleads.v3.enums.UserListClosingReasonEnum_UserListClosingReason" json:"closing_reason,omitempty"`
	// Output only. Indicates the reason this account has been granted access to the list.
	// The reason can be SHARED, OWNED, LICENSED or SUBSCRIBED.
	//
	// This field is read-only.
	AccessReason enums.AccessReasonEnum_AccessReason `protobuf:"varint,15,opt,name=access_reason,json=accessReason,proto3,enum=google.ads.googleads.v3.enums.AccessReasonEnum_AccessReason" json:"access_reason,omitempty"`
	// Indicates if this share is still enabled. When a UserList is shared with
	// the user this field is set to ENABLED. Later the userList owner can decide
	// to revoke the share and make it DISABLED.
	// The default value of this field is set to ENABLED.
	AccountUserListStatus enums.UserListAccessStatusEnum_UserListAccessStatus `protobuf:"varint,16,opt,name=account_user_list_status,json=accountUserListStatus,proto3,enum=google.ads.googleads.v3.enums.UserListAccessStatusEnum_UserListAccessStatus" json:"account_user_list_status,omitempty"`
	// Indicates if this user list is eligible for Google Search Network.
	EligibleForSearch *wrappers.BoolValue `protobuf:"bytes,17,opt,name=eligible_for_search,json=eligibleForSearch,proto3" json:"eligible_for_search,omitempty"`
	// Output only. Indicates this user list is eligible for Google Display Network.
	//
	// This field is read-only.
	EligibleForDisplay *wrappers.BoolValue `protobuf:"bytes,18,opt,name=eligible_for_display,json=eligibleForDisplay,proto3" json:"eligible_for_display,omitempty"`
	// The user list.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to UserList:
	//	*UserList_CrmBasedUserList
	//	*UserList_SimilarUserList
	//	*UserList_RuleBasedUserList
	//	*UserList_LogicalUserList
	//	*UserList_BasicUserList
	UserList             isUserList_UserList `protobuf_oneof:"user_list"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efbaf923f128250, []int{0}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserList) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserList) GetReadOnly() *wrappers.BoolValue {
	if m != nil {
		return m.ReadOnly
	}
	return nil
}

func (m *UserList) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserList) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *UserList) GetMembershipStatus() enums.UserListMembershipStatusEnum_UserListMembershipStatus {
	if m != nil {
		return m.MembershipStatus
	}
	return enums.UserListMembershipStatusEnum_UNSPECIFIED
}

func (m *UserList) GetIntegrationCode() *wrappers.StringValue {
	if m != nil {
		return m.IntegrationCode
	}
	return nil
}

func (m *UserList) GetMembershipLifeSpan() *wrappers.Int64Value {
	if m != nil {
		return m.MembershipLifeSpan
	}
	return nil
}

func (m *UserList) GetSizeForDisplay() *wrappers.Int64Value {
	if m != nil {
		return m.SizeForDisplay
	}
	return nil
}

func (m *UserList) GetSizeRangeForDisplay() enums.UserListSizeRangeEnum_UserListSizeRange {
	if m != nil {
		return m.SizeRangeForDisplay
	}
	return enums.UserListSizeRangeEnum_UNSPECIFIED
}

func (m *UserList) GetSizeForSearch() *wrappers.Int64Value {
	if m != nil {
		return m.SizeForSearch
	}
	return nil
}

func (m *UserList) GetSizeRangeForSearch() enums.UserListSizeRangeEnum_UserListSizeRange {
	if m != nil {
		return m.SizeRangeForSearch
	}
	return enums.UserListSizeRangeEnum_UNSPECIFIED
}

func (m *UserList) GetType() enums.UserListTypeEnum_UserListType {
	if m != nil {
		return m.Type
	}
	return enums.UserListTypeEnum_UNSPECIFIED
}

func (m *UserList) GetClosingReason() enums.UserListClosingReasonEnum_UserListClosingReason {
	if m != nil {
		return m.ClosingReason
	}
	return enums.UserListClosingReasonEnum_UNSPECIFIED
}

func (m *UserList) GetAccessReason() enums.AccessReasonEnum_AccessReason {
	if m != nil {
		return m.AccessReason
	}
	return enums.AccessReasonEnum_UNSPECIFIED
}

func (m *UserList) GetAccountUserListStatus() enums.UserListAccessStatusEnum_UserListAccessStatus {
	if m != nil {
		return m.AccountUserListStatus
	}
	return enums.UserListAccessStatusEnum_UNSPECIFIED
}

func (m *UserList) GetEligibleForSearch() *wrappers.BoolValue {
	if m != nil {
		return m.EligibleForSearch
	}
	return nil
}

func (m *UserList) GetEligibleForDisplay() *wrappers.BoolValue {
	if m != nil {
		return m.EligibleForDisplay
	}
	return nil
}

type isUserList_UserList interface {
	isUserList_UserList()
}

type UserList_CrmBasedUserList struct {
	CrmBasedUserList *common.CrmBasedUserListInfo `protobuf:"bytes,19,opt,name=crm_based_user_list,json=crmBasedUserList,proto3,oneof"`
}

type UserList_SimilarUserList struct {
	SimilarUserList *common.SimilarUserListInfo `protobuf:"bytes,20,opt,name=similar_user_list,json=similarUserList,proto3,oneof"`
}

type UserList_RuleBasedUserList struct {
	RuleBasedUserList *common.RuleBasedUserListInfo `protobuf:"bytes,21,opt,name=rule_based_user_list,json=ruleBasedUserList,proto3,oneof"`
}

type UserList_LogicalUserList struct {
	LogicalUserList *common.LogicalUserListInfo `protobuf:"bytes,22,opt,name=logical_user_list,json=logicalUserList,proto3,oneof"`
}

type UserList_BasicUserList struct {
	BasicUserList *common.BasicUserListInfo `protobuf:"bytes,23,opt,name=basic_user_list,json=basicUserList,proto3,oneof"`
}

func (*UserList_CrmBasedUserList) isUserList_UserList() {}

func (*UserList_SimilarUserList) isUserList_UserList() {}

func (*UserList_RuleBasedUserList) isUserList_UserList() {}

func (*UserList_LogicalUserList) isUserList_UserList() {}

func (*UserList_BasicUserList) isUserList_UserList() {}

func (m *UserList) GetUserList() isUserList_UserList {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *UserList) GetCrmBasedUserList() *common.CrmBasedUserListInfo {
	if x, ok := m.GetUserList().(*UserList_CrmBasedUserList); ok {
		return x.CrmBasedUserList
	}
	return nil
}

func (m *UserList) GetSimilarUserList() *common.SimilarUserListInfo {
	if x, ok := m.GetUserList().(*UserList_SimilarUserList); ok {
		return x.SimilarUserList
	}
	return nil
}

func (m *UserList) GetRuleBasedUserList() *common.RuleBasedUserListInfo {
	if x, ok := m.GetUserList().(*UserList_RuleBasedUserList); ok {
		return x.RuleBasedUserList
	}
	return nil
}

func (m *UserList) GetLogicalUserList() *common.LogicalUserListInfo {
	if x, ok := m.GetUserList().(*UserList_LogicalUserList); ok {
		return x.LogicalUserList
	}
	return nil
}

func (m *UserList) GetBasicUserList() *common.BasicUserListInfo {
	if x, ok := m.GetUserList().(*UserList_BasicUserList); ok {
		return x.BasicUserList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserList) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserList_CrmBasedUserList)(nil),
		(*UserList_SimilarUserList)(nil),
		(*UserList_RuleBasedUserList)(nil),
		(*UserList_LogicalUserList)(nil),
		(*UserList_BasicUserList)(nil),
	}
}

func init() {
	proto.RegisterType((*UserList)(nil), "google.ads.googleads.v3.resources.UserList")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/user_list.proto", fileDescriptor_2efbaf923f128250)
}

var fileDescriptor_2efbaf923f128250 = []byte{
	// 1022 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4f, 0x6f, 0xdb, 0xb6,
	0x1b, 0xc7, 0x6b, 0x27, 0xed, 0x2f, 0x61, 0xe2, 0x38, 0x66, 0x92, 0xfe, 0xb4, 0xac, 0xd8, 0x92,
	0x0d, 0x05, 0xb2, 0x1d, 0xa4, 0x36, 0xee, 0x86, 0xc1, 0x0d, 0x36, 0xc8, 0xd9, 0x9a, 0x34, 0x48,
	0xd3, 0x4e, 0x6e, 0x73, 0xd8, 0x02, 0x08, 0xb4, 0x44, 0x2b, 0x04, 0x24, 0x52, 0x20, 0xa5, 0x0c,
	0x6e, 0xd1, 0x61, 0x3b, 0xec, 0x8d, 0xec, 0x34, 0xec, 0xa5, 0xec, 0x55, 0xf4, 0xdc, 0x97, 0xd0,
	0xd3, 0x20, 0x4a, 0x94, 0x29, 0xa7, 0xae, 0x75, 0xd8, 0x8d, 0xe2, 0xf3, 0x7c, 0x9e, 0x2f, 0xbf,
	0xfc, 0x27, 0x82, 0xfb, 0x01, 0x63, 0x41, 0x88, 0x2d, 0xe4, 0x0b, 0x2b, 0x6f, 0x66, 0xad, 0xab,
	0xae, 0xc5, 0xb1, 0x60, 0x29, 0xf7, 0xb0, 0xb0, 0x52, 0x81, 0xb9, 0x1b, 0x12, 0x91, 0x98, 0x31,
	0x67, 0x09, 0x83, 0xbb, 0x79, 0x9e, 0x89, 0x7c, 0x61, 0x96, 0x88, 0x79, 0xd5, 0x35, 0x4b, 0x64,
	0xdb, 0x9a, 0x55, 0xd5, 0x63, 0x51, 0xc4, 0xe8, 0xa4, 0xa4, 0xc8, 0x6b, 0x6e, 0xcf, 0x1c, 0x06,
	0xa6, 0x69, 0x24, 0x2c, 0xe4, 0x79, 0x58, 0x08, 0x97, 0x63, 0x24, 0x18, 0x2d, 0x90, 0x87, 0x1f,
	0x46, 0x4a, 0x09, 0xb7, 0x80, 0x45, 0x82, 0x92, 0x54, 0xe9, 0x1d, 0xd4, 0x85, 0xbd, 0x90, 0x09,
	0x42, 0x83, 0xaa, 0xf4, 0x77, 0x75, 0xe9, 0x08, 0x47, 0x43, 0xcc, 0xc5, 0x25, 0x89, 0xab, 0xf2,
	0xdf, 0xd4, 0x2d, 0x20, 0xc8, 0x4b, 0xec, 0x72, 0x44, 0x03, 0x5c, 0x90, 0xfb, 0x75, 0xc9, 0x64,
	0x1c, 0x2b, 0xe6, 0x53, 0xc5, 0xc4, 0xc4, 0x1a, 0x11, 0x1c, 0xfa, 0xee, 0x10, 0x5f, 0xa2, 0x2b,
	0xc2, 0x78, 0x91, 0xf0, 0x91, 0x96, 0xa0, 0x16, 0xb1, 0x08, 0x7d, 0x52, 0x84, 0xe4, 0xd7, 0x30,
	0x1d, 0x59, 0xbf, 0x70, 0x14, 0xc7, 0x98, 0x2b, 0x27, 0x77, 0x34, 0x14, 0x51, 0xca, 0x12, 0x94,
	0x10, 0x46, 0x8b, 0xe8, 0x67, 0x7f, 0x75, 0xc0, 0xd2, 0x0b, 0x81, 0xf9, 0x29, 0x11, 0x09, 0x3c,
	0x03, 0x2d, 0x55, 0xdc, 0xa5, 0x28, 0xc2, 0x46, 0x63, 0xa7, 0xb1, 0xb7, 0xdc, 0xff, 0xe2, 0x8d,
	0x7d, 0xf3, 0x9d, 0xfd, 0x39, 0xd8, 0x9d, 0xec, 0xa5, 0xa2, 0x15, 0x13, 0x61, 0x7a, 0x2c, 0xb2,
	0x54, 0x05, 0x67, 0x55, 0xf1, 0x67, 0x28, 0xc2, 0xf0, 0x1e, 0x68, 0x12, 0xdf, 0x68, 0xee, 0x34,
	0xf6, 0x56, 0xf6, 0x3f, 0x2e, 0x18, 0x53, 0x8d, 0xd3, 0x7c, 0x4c, 0x93, 0xaf, 0x1f, 0x9c, 0xa3,
	0x30, 0xc5, 0xfd, 0x85, 0x37, 0xf6, 0x82, 0xd3, 0x24, 0x3e, 0x3c, 0x00, 0xcb, 0x1c, 0x23, 0xdf,
	0x65, 0x34, 0x1c, 0x1b, 0x0b, 0x12, 0xdc, 0xbe, 0x06, 0xf6, 0x19, 0x0b, 0x35, 0x6e, 0x29, 0x23,
	0x9e, 0xd2, 0x70, 0x0c, 0xef, 0x81, 0x45, 0x39, 0xec, 0x45, 0x09, 0xde, 0xb9, 0x06, 0x0e, 0x12,
	0x4e, 0x68, 0x20, 0x51, 0x47, 0x66, 0xc2, 0x6f, 0xc1, 0x8a, 0x8f, 0x85, 0xc7, 0x49, 0x9c, 0x4d,
	0x8a, 0x71, 0xb3, 0x06, 0xa8, 0x03, 0xf0, 0xf7, 0x06, 0xe8, 0x5c, 0xdb, 0x42, 0xc6, 0xad, 0x9d,
	0xc6, 0xde, 0xda, 0xfe, 0x73, 0x73, 0xd6, 0x31, 0x94, 0x3b, 0xc1, 0x54, 0x93, 0xf6, 0xa4, 0xe4,
	0x07, 0x12, 0xff, 0x81, 0xa6, 0xd1, 0xcc, 0xa0, 0xb3, 0x1e, 0x4d, 0xf5, 0xc0, 0x23, 0xb0, 0x4e,
	0x68, 0x82, 0x03, 0x2e, 0x17, 0xd6, 0xf5, 0x98, 0x8f, 0x8d, 0xff, 0xd5, 0x30, 0xd2, 0xd6, 0xa8,
	0x43, 0xe6, 0x63, 0xf8, 0x04, 0x6c, 0x6a, 0x5e, 0x42, 0x32, 0xc2, 0xae, 0x88, 0x11, 0x35, 0x96,
	0xe6, 0x2e, 0xa0, 0x03, 0x27, 0xe0, 0x29, 0x19, 0xe1, 0x41, 0x8c, 0x28, 0x3c, 0x01, 0xeb, 0xf2,
	0x70, 0x8c, 0x18, 0x77, 0x7d, 0x22, 0xe2, 0x10, 0x8d, 0x8d, 0xe5, 0x9a, 0x7b, 0x61, 0x2d, 0x23,
	0x1f, 0x31, 0xfe, 0x7d, 0xce, 0xc1, 0xdf, 0x1a, 0xe0, 0xf6, 0xe4, 0xa4, 0x55, 0x4a, 0x02, 0x39,
	0xd9, 0x8f, 0x6a, 0x4e, 0xf6, 0x80, 0xbc, 0xc4, 0x4e, 0x56, 0xa3, 0x32, 0xcb, 0x65, 0x6f, 0xae,
	0xbe, 0x21, 0xd4, 0xb7, 0x36, 0x84, 0x63, 0xd0, 0x2e, 0xed, 0x08, 0x8c, 0xb8, 0x77, 0x69, 0xac,
	0xd4, 0x74, 0xd3, 0x2a, 0xdc, 0x0c, 0x24, 0x06, 0x7f, 0x05, 0x5b, 0x53, 0x5e, 0x8a, 0x7a, 0xab,
	0xff, 0xbd, 0x15, 0xa8, 0x5b, 0x29, 0xf4, 0x5f, 0x80, 0xc5, 0xec, 0xee, 0x31, 0x5a, 0x52, 0xee,
	0xa0, 0xa6, 0xdc, 0xf3, 0x71, 0x5c, 0x55, 0xca, 0x3a, 0x72, 0x11, 0x59, 0x0e, 0xa6, 0x60, 0xad,
	0x7a, 0x17, 0x1b, 0x6b, 0x52, 0xe0, 0xac, 0xa6, 0xc0, 0x61, 0x0e, 0x3b, 0x92, 0xad, 0x28, 0x55,
	0x22, 0x4e, 0xcb, 0xd3, 0x3f, 0xe1, 0x08, 0xb4, 0x2a, 0x3f, 0x1f, 0xa3, 0x5d, 0xcb, 0x96, 0x2d,
	0x19, 0x4d, 0x4c, 0xef, 0xc8, 0x6d, 0xad, 0x22, 0xad, 0x0b, 0xfe, 0xd1, 0x00, 0x06, 0xf2, 0x3c,
	0x96, 0xd2, 0xc4, 0xd5, 0xae, 0xff, 0xfc, 0xc4, 0xaf, 0x4b, 0xcd, 0xd3, 0x9a, 0x4e, 0x73, 0xa9,
	0xf7, 0x9c, 0x76, 0x3d, 0xe0, 0x6c, 0x15, 0x6a, 0xe5, 0xca, 0xe6, 0xc7, 0xfd, 0x04, 0x6c, 0xe0,
	0x90, 0x04, 0x64, 0x18, 0x56, 0xf6, 0x4e, 0x67, 0xde, 0x65, 0xe9, 0x74, 0x14, 0x36, 0xd9, 0x09,
	0x3f, 0x82, 0xcd, 0x4a, 0x2d, 0x75, 0xa6, 0x60, 0xbd, 0x9b, 0x17, 0x6a, 0x15, 0xd5, 0x31, 0xc1,
	0x60, 0xc3, 0xe3, 0x91, 0x3b, 0x44, 0x02, 0xfb, 0x93, 0x79, 0x32, 0x36, 0x64, 0xc5, 0x07, 0x33,
	0x27, 0x28, 0x7f, 0x76, 0x98, 0x87, 0x3c, 0xea, 0x67, 0xa4, 0xf2, 0xfc, 0x98, 0x8e, 0xd8, 0xf1,
	0x0d, 0x67, 0xdd, 0x9b, 0xea, 0x87, 0x23, 0xd0, 0x11, 0x24, 0x22, 0x21, 0xe2, 0x9a, 0xc8, 0xa6,
	0x14, 0xe9, 0xce, 0x13, 0x19, 0xe4, 0xa0, 0xae, 0x21, 0xfd, 0x1c, 0xdf, 0x70, 0xda, 0xa2, 0x1a,
	0x83, 0x97, 0x60, 0x93, 0xa7, 0x21, 0xbe, 0xe6, 0x67, 0x4b, 0x4a, 0x7d, 0x35, 0x4f, 0xca, 0x49,
	0x43, 0xfc, 0x3e, 0x43, 0x1d, 0x3e, 0x1d, 0x80, 0x08, 0x74, 0x42, 0x16, 0x10, 0x0f, 0x85, 0x9a,
	0xcc, 0xed, 0x7a, 0x8e, 0x4e, 0x73, 0x70, 0x4a, 0xa4, 0x1d, 0x56, 0xbb, 0xe1, 0xcf, 0xa0, 0x3d,
	0x44, 0x82, 0x78, 0x9a, 0xc0, 0xff, 0xa5, 0xc0, 0xfd, 0x79, 0x02, 0xfd, 0x0c, 0x9b, 0x2a, 0xdf,
	0x1a, 0xea, 0x9d, 0x3d, 0xe7, 0xad, 0xfd, 0xb4, 0xc6, 0x13, 0x01, 0x7e, 0xe9, 0xa5, 0x22, 0x61,
	0x11, 0xe6, 0xc2, 0x7a, 0xa5, 0x9a, 0xaf, 0xe5, 0xb3, 0x28, 0x0b, 0x0b, 0xeb, 0x55, 0x39, 0xb8,
	0xd7, 0xfd, 0x15, 0xb0, 0x5c, 0x7e, 0xf5, 0xdf, 0x35, 0xc0, 0x5d, 0x8f, 0x45, 0xe6, 0xdc, 0xc7,
	0x6d, 0xbf, 0xa5, 0xc4, 0x9e, 0x65, 0x1b, 0xf7, 0x59, 0xe3, 0xa7, 0x93, 0x82, 0x09, 0x58, 0x88,
	0x68, 0x60, 0x32, 0x1e, 0x58, 0x01, 0xa6, 0x72, 0x5b, 0x5b, 0x93, 0x71, 0x7e, 0xe0, 0x89, 0xfd,
	0xb0, 0x6c, 0xfd, 0xd9, 0x5c, 0x38, 0xb2, 0xed, 0xbf, 0x9b, 0xbb, 0x47, 0x79, 0x49, 0xdb, 0x17,
	0x66, 0xde, 0xcc, 0x5a, 0xe7, 0x5d, 0xd3, 0x51, 0x99, 0xff, 0xa8, 0x9c, 0x0b, 0xdb, 0x17, 0x17,
	0x65, 0xce, 0xc5, 0x79, 0xf7, 0xa2, 0xcc, 0x79, 0xdb, 0xbc, 0x9b, 0x07, 0x7a, 0x3d, 0xdb, 0x17,
	0xbd, 0x5e, 0x99, 0xd5, 0xeb, 0x9d, 0x77, 0x7b, 0xbd, 0x32, 0x6f, 0x78, 0x4b, 0x0e, 0xb6, 0xfb,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x2b, 0x95, 0xee, 0x0e, 0x0c, 0x00, 0x00,
}
