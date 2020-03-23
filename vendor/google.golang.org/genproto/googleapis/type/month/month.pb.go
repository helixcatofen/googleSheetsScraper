// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/month.proto

package month

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Represents a month in the Gregorian calendar.
type Month int32

const (
	// The unspecifed month.
	Month_MONTH_UNSPECIFIED Month = 0
	// The month of January.
	Month_JANUARY Month = 1
	// The month of February.
	Month_FEBRUARY Month = 2
	// The month of March.
	Month_MARCH Month = 3
	// The month of April.
	Month_APRIL Month = 4
	// The month of May.
	Month_MAY Month = 5
	// The month of June.
	Month_JUNE Month = 6
	// The month of July.
	Month_JULY Month = 7
	// The month of August.
	Month_AUGUST Month = 8
	// The month of September.
	Month_SEPTEMBER Month = 9
	// The month of October.
	Month_OCTOBER Month = 10
	// The month of November.
	Month_NOVEMBER Month = 11
	// The month of December.
	Month_DECEMBER Month = 12
)

var Month_name = map[int32]string{
	0:  "MONTH_UNSPECIFIED",
	1:  "JANUARY",
	2:  "FEBRUARY",
	3:  "MARCH",
	4:  "APRIL",
	5:  "MAY",
	6:  "JUNE",
	7:  "JULY",
	8:  "AUGUST",
	9:  "SEPTEMBER",
	10: "OCTOBER",
	11: "NOVEMBER",
	12: "DECEMBER",
}

var Month_value = map[string]int32{
	"MONTH_UNSPECIFIED": 0,
	"JANUARY":           1,
	"FEBRUARY":          2,
	"MARCH":             3,
	"APRIL":             4,
	"MAY":               5,
	"JUNE":              6,
	"JULY":              7,
	"AUGUST":            8,
	"SEPTEMBER":         9,
	"OCTOBER":           10,
	"NOVEMBER":          11,
	"DECEMBER":          12,
}

func (x Month) String() string {
	return proto.EnumName(Month_name, int32(x))
}

func (Month) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3bd21fd11ded5e73, []int{0}
}

func init() {
	proto.RegisterEnum("google.type.Month", Month_name, Month_value)
}

func init() {
	proto.RegisterFile("google/type/month.proto", fileDescriptor_3bd21fd11ded5e73)
}

var fileDescriptor_3bd21fd11ded5e73 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0x80, 0xed, 0xba, 0xfe, 0x9d, 0x4e, 0x8c, 0x01, 0xf1, 0x1d, 0xbc, 0x68, 0x2f, 0x04, 0x6f,
	0xbc, 0x4a, 0xbb, 0x6c, 0xeb, 0x58, 0xdb, 0xd0, 0x26, 0x42, 0x05, 0x91, 0x29, 0x25, 0x0a, 0x5b,
	0x53, 0x66, 0x6f, 0x7c, 0x1d, 0x9f, 0xc0, 0x47, 0x94, 0x24, 0x5e, 0xec, 0x26, 0x9c, 0x8f, 0x2f,
	0x70, 0xf8, 0x0e, 0xdc, 0x4a, 0xa5, 0xe4, 0xa1, 0x4f, 0xa7, 0xef, 0xb1, 0x4f, 0x8f, 0x6a, 0x98,
	0x3e, 0x92, 0xf1, 0xa4, 0x26, 0x85, 0x63, 0x2b, 0x12, 0x2d, 0xee, 0x7e, 0x1d, 0xf0, 0x4a, 0x2d,
	0xf1, 0x0d, 0x5c, 0x97, 0x75, 0xc5, 0x37, 0xaf, 0xa2, 0x6a, 0x19, 0xcd, 0x8b, 0x55, 0x41, 0x97,
	0xe8, 0x02, 0xc7, 0x10, 0x6c, 0x49, 0x25, 0x48, 0xd3, 0x21, 0x07, 0x2f, 0x20, 0x5c, 0xd1, 0xac,
	0x31, 0x34, 0xc3, 0x11, 0x78, 0x25, 0x69, 0xf2, 0x0d, 0x72, 0xf5, 0x48, 0x58, 0x53, 0xec, 0xd0,
	0x1c, 0x07, 0xe0, 0x96, 0xa4, 0x43, 0x1e, 0x0e, 0x61, 0xbe, 0x15, 0x15, 0x45, 0xbe, 0x9d, 0x76,
	0x1d, 0x0a, 0x30, 0x80, 0x4f, 0xc4, 0x5a, 0xb4, 0x1c, 0x85, 0xf8, 0x12, 0xa2, 0x96, 0x32, 0x4e,
	0xcb, 0x8c, 0x36, 0x28, 0xd2, 0x8b, 0xea, 0x9c, 0xd7, 0x1a, 0x40, 0x2f, 0xaa, 0xea, 0x27, 0xab,
	0x62, 0x4d, 0x4b, 0x9a, 0x5b, 0x5a, 0x64, 0x2f, 0x70, 0xf5, 0xae, 0x8e, 0xc9, 0x59, 0x45, 0x06,
	0x26, 0x81, 0xe9, 0x3c, 0xe6, 0x3c, 0x3f, 0xfc, 0x2b, 0xa9, 0x0e, 0xfb, 0x41, 0x26, 0xea, 0x24,
	0x53, 0xd9, 0x0f, 0x26, 0x3e, 0xb5, 0x6a, 0x3f, 0x7e, 0x7e, 0x9d, 0x1d, 0xe6, 0xd1, 0xbc, 0x3f,
	0x33, 0x77, 0xcd, 0xd9, 0x9b, 0x6f, 0x3e, 0xde, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x32,
	0x69, 0x57, 0x40, 0x01, 0x00, 0x00,
}
