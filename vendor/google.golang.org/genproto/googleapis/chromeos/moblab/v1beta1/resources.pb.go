// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/chromeos/moblab/v1beta1/resources.proto

package moblab

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

// Resource that represents a build target.
type BuildTarget struct {
	// The resource name of the build target.
	// Format: buildTargets/{build_target}
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildTarget) Reset()         { *m = BuildTarget{} }
func (m *BuildTarget) String() string { return proto.CompactTextString(m) }
func (*BuildTarget) ProtoMessage()    {}
func (*BuildTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17c8fac142feee0, []int{0}
}

func (m *BuildTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildTarget.Unmarshal(m, b)
}
func (m *BuildTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildTarget.Marshal(b, m, deterministic)
}
func (m *BuildTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildTarget.Merge(m, src)
}
func (m *BuildTarget) XXX_Size() int {
	return xxx_messageInfo_BuildTarget.Size(m)
}
func (m *BuildTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildTarget.DiscardUnknown(m)
}

var xxx_messageInfo_BuildTarget proto.InternalMessageInfo

func (m *BuildTarget) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Resource that represents a model. Each model belongs to a build target. For
// non-unified build, the model name is the same as its build target name.
type Model struct {
	// The resource name of the model.
	// Format: buildTargets/{build_target}/models/{model}
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17c8fac142feee0, []int{1}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Resource that represents a chrome OS milestone.
type Milestone struct {
	// The resource name of the milestone.
	// Format: milestones/{milestone}
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Milestone) Reset()         { *m = Milestone{} }
func (m *Milestone) String() string { return proto.CompactTextString(m) }
func (*Milestone) ProtoMessage()    {}
func (*Milestone) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17c8fac142feee0, []int{2}
}

func (m *Milestone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Milestone.Unmarshal(m, b)
}
func (m *Milestone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Milestone.Marshal(b, m, deterministic)
}
func (m *Milestone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Milestone.Merge(m, src)
}
func (m *Milestone) XXX_Size() int {
	return xxx_messageInfo_Milestone.Size(m)
}
func (m *Milestone) XXX_DiscardUnknown() {
	xxx_messageInfo_Milestone.DiscardUnknown(m)
}

var xxx_messageInfo_Milestone proto.InternalMessageInfo

func (m *Milestone) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Resource that represents a build for the given build target, model, milestone
// and build version.
// NEXT_TAG: 4
type Build struct {
	// The resource name of the build.
	// Format: buildTargets/{build_target}/models/{model}/builds/{build}
	// Example: buildTargets/octopus/models/bobba/builds/1234.0.0
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The milestone that owns the build.
	// Format: milestones/{milestone}
	Milestone string `protobuf:"bytes,2,opt,name=milestone,proto3" json:"milestone,omitempty"`
	// The build version of the build, e.g. 1234.0.0.
	BuildVersion         string   `protobuf:"bytes,3,opt,name=build_version,json=buildVersion,proto3" json:"build_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17c8fac142feee0, []int{3}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Build) GetMilestone() string {
	if m != nil {
		return m.Milestone
	}
	return ""
}

func (m *Build) GetBuildVersion() string {
	if m != nil {
		return m.BuildVersion
	}
	return ""
}

// Resource that represents a build artifact stored in Google Cloud Storage for
// the given build target, model, build version and bucket. NEXT_TAG: 6
type BuildArtifact struct {
	// The resource name of the build artifact.
	// Format:
	// buildTargets/{build_target}/models/{model}/builds/{build}/artifacts/{artifact}
	// Example:
	// buildTargets/octopus/models/bobba/builds/1234.0.0/artifacts/chromeos-moblab-peng-staging
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The build metadata of the build artifact.
	Build string `protobuf:"bytes,2,opt,name=build,proto3" json:"build,omitempty"`
	// The bucket that stores the build artifact.
	Bucket string `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// The path of the build artifact in the bucket.
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// The number of objects in the build artifact folder. The object number can
	// be used to calculated the stage progress by comparing the source build
	// artifact with the destination build artifact.
	ObjectCount          uint32   `protobuf:"varint,5,opt,name=object_count,json=objectCount,proto3" json:"object_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildArtifact) Reset()         { *m = BuildArtifact{} }
func (m *BuildArtifact) String() string { return proto.CompactTextString(m) }
func (*BuildArtifact) ProtoMessage()    {}
func (*BuildArtifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17c8fac142feee0, []int{4}
}

func (m *BuildArtifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildArtifact.Unmarshal(m, b)
}
func (m *BuildArtifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildArtifact.Marshal(b, m, deterministic)
}
func (m *BuildArtifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildArtifact.Merge(m, src)
}
func (m *BuildArtifact) XXX_Size() int {
	return xxx_messageInfo_BuildArtifact.Size(m)
}
func (m *BuildArtifact) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildArtifact.DiscardUnknown(m)
}

var xxx_messageInfo_BuildArtifact proto.InternalMessageInfo

func (m *BuildArtifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildArtifact) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *BuildArtifact) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *BuildArtifact) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *BuildArtifact) GetObjectCount() uint32 {
	if m != nil {
		return m.ObjectCount
	}
	return 0
}

func init() {
	proto.RegisterType((*BuildTarget)(nil), "google.chromeos.moblab.v1beta1.BuildTarget")
	proto.RegisterType((*Model)(nil), "google.chromeos.moblab.v1beta1.Model")
	proto.RegisterType((*Milestone)(nil), "google.chromeos.moblab.v1beta1.Milestone")
	proto.RegisterType((*Build)(nil), "google.chromeos.moblab.v1beta1.Build")
	proto.RegisterType((*BuildArtifact)(nil), "google.chromeos.moblab.v1beta1.BuildArtifact")
}

func init() {
	proto.RegisterFile("google/chromeos/moblab/v1beta1/resources.proto", fileDescriptor_d17c8fac142feee0)
}

var fileDescriptor_d17c8fac142feee0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x4b, 0x53, 0x29, 0xd3, 0x86, 0x83, 0x0f, 0x95, 0x01, 0x09, 0x15, 0x83, 0x50, 0x88,
	0xa8, 0x57, 0x15, 0x27, 0xc2, 0x01, 0xad, 0x01, 0x09, 0x04, 0x45, 0x95, 0x05, 0x1c, 0xb8, 0x54,
	0xeb, 0xcd, 0xd6, 0x35, 0xd8, 0x9e, 0x68, 0x77, 0xd3, 0x4b, 0xe4, 0x4a, 0x3c, 0x15, 0xef, 0xe4,
	0x2b, 0xb7, 0x9e, 0x90, 0x77, 0xd7, 0x09, 0x87, 0x90, 0x44, 0x3d, 0x65, 0x66, 0xf6, 0xfb, 0x99,
	0x19, 0x65, 0x0c, 0x51, 0x86, 0x98, 0x15, 0x82, 0xf0, 0x4b, 0x89, 0xa5, 0x40, 0x45, 0x4a, 0x4c,
	0x0b, 0x96, 0x92, 0xab, 0x93, 0x54, 0x68, 0x76, 0x42, 0xa4, 0x50, 0x38, 0x93, 0x5c, 0xa8, 0x68,
	0x2a, 0x51, 0xa3, 0xff, 0xd0, 0xe2, 0xa3, 0x0e, 0x1f, 0x59, 0x7c, 0xe4, 0xf0, 0xf7, 0xef, 0x39,
	0x3d, 0x36, 0xcd, 0x17, 0x5c, 0x4b, 0x0d, 0x11, 0xf6, 0xe3, 0x59, 0x5e, 0x4c, 0xbe, 0x30, 0x99,
	0x09, 0xed, 0xfb, 0xb0, 0x5b, 0xb1, 0x52, 0x04, 0xde, 0x91, 0x37, 0xec, 0x27, 0x26, 0x1e, 0x7f,
	0x6a, 0xe8, 0x07, 0x18, 0x75, 0xda, 0xc7, 0x4e, 0xdb, 0x4a, 0xb2, 0x69, 0xae, 0x22, 0x8e, 0x25,
	0xf9, 0x57, 0xe4, 0x41, 0xba, 0x4c, 0x14, 0x99, 0x9b, 0xec, 0x5c, 0x9b, 0xb4, 0x0e, 0x25, 0xf4,
	0x4e, 0x71, 0x22, 0x8a, 0x95, 0x56, 0x5f, 0x1b, 0x9a, 0xc0, 0x93, 0x0d, 0x56, 0x96, 0x3e, 0x5a,
	0x63, 0x42, 0xca, 0x16, 0xa2, 0xc8, 0xdc, 0xfc, 0xd6, 0xe1, 0x05, 0xf4, 0x4f, 0xf3, 0x42, 0x28,
	0x8d, 0x95, 0x58, 0xe9, 0xfb, 0xae, 0xa1, 0x31, 0x0c, 0x37, 0xf9, 0x2e, 0x24, 0x0e, 0xcb, 0x2e,
	0x6c, 0x3d, 0xba, 0xb8, 0x0e, 0xff, 0x78, 0xd0, 0x33, 0x8b, 0x58, 0x65, 0xe2, 0x7f, 0x84, 0xfe,
	0x02, 0x1b, 0xec, 0xb4, 0x0f, 0xf1, 0xf1, 0x0d, 0x1d, 0x6d, 0xef, 0x9b, 0x2c, 0xf9, 0xfe, 0x63,
	0x18, 0xd8, 0x91, 0xaf, 0x84, 0x54, 0x39, 0x56, 0xc1, 0x1d, 0xe3, 0x74, 0x60, 0x8a, 0xdf, 0x6c,
	0x6d, 0x3c, 0x69, 0x28, 0xdb, 0xb8, 0x4e, 0xdb, 0xf0, 0xcb, 0xed, 0xd7, 0x49, 0xcc, 0x63, 0x07,
	0xaa, 0xc3, 0xdf, 0x3b, 0x30, 0x30, 0x22, 0x54, 0xea, 0xfc, 0x82, 0xf1, 0x95, 0xff, 0x22, 0xff,
	0x35, 0xf4, 0x0c, 0xde, 0x4d, 0xfe, 0xec, 0x86, 0x3e, 0xdd, 0xae, 0xb5, 0xc4, 0xf2, 0xfc, 0x43,
	0xd8, 0x4b, 0x67, 0xfc, 0xa7, 0xd0, 0x6e, 0x54, 0x97, 0xb5, 0x66, 0x53, 0xa6, 0x2f, 0x83, 0x5d,
	0x6b, 0xd6, 0xc6, 0xfe, 0x23, 0x38, 0xc0, 0xf4, 0x87, 0xe0, 0xfa, 0x9c, 0xe3, 0xac, 0xd2, 0x41,
	0xef, 0xc8, 0x1b, 0x0e, 0x92, 0x7d, 0x5b, 0x7b, 0xd3, 0x96, 0xc6, 0xbf, 0xbc, 0x86, 0x5e, 0xc3,
	0xf3, 0x6d, 0x3a, 0x58, 0xcc, 0xf5, 0xf9, 0xd6, 0x4b, 0x22, 0xcc, 0x69, 0x28, 0x32, 0xef, 0xc2,
	0x3a, 0xbe, 0x86, 0x90, 0x63, 0x19, 0xad, 0xbf, 0xde, 0xf8, 0x6e, 0xd2, 0x9d, 0xfb, 0x59, 0x7b,
	0xb2, 0xef, 0xbd, 0x33, 0xef, 0xfb, 0x5b, 0xc7, 0xc9, 0xb0, 0x60, 0x55, 0x16, 0xa1, 0xcc, 0x48,
	0x26, 0x2a, 0x73, 0xd2, 0x64, 0x39, 0xc2, 0xff, 0x3e, 0x20, 0xaf, 0x6c, 0x9a, 0xee, 0x19, 0xc2,
	0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x88, 0x2b, 0x60, 0x70, 0x04, 0x00, 0x00,
}
