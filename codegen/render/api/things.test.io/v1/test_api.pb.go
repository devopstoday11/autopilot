// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test_api.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AcrylicType_Body int32

const (
	AcrylicType_Light  AcrylicType_Body = 0
	AcrylicType_Medium AcrylicType_Body = 1
	AcrylicType_Heavy  AcrylicType_Body = 2
)

var AcrylicType_Body_name = map[int32]string{
	0: "Light",
	1: "Medium",
	2: "Heavy",
}

var AcrylicType_Body_value = map[string]int32{
	"Light":  0,
	"Medium": 1,
	"Heavy":  2,
}

func (x AcrylicType_Body) String() string {
	return proto.EnumName(AcrylicType_Body_name, int32(x))
}

func (AcrylicType_Body) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{2, 0}
}

type PaintSpec struct {
	Color *PaintColor `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	// Types that are valid to be assigned to PaintType:
	//	*PaintSpec_Acrylic
	//	*PaintSpec_Oil
	PaintType            isPaintSpec_PaintType `protobuf_oneof:"paintType"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaintSpec) Reset()         { *m = PaintSpec{} }
func (m *PaintSpec) String() string { return proto.CompactTextString(m) }
func (*PaintSpec) ProtoMessage()    {}
func (*PaintSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{0}
}
func (m *PaintSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintSpec.Unmarshal(m, b)
}
func (m *PaintSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintSpec.Marshal(b, m, deterministic)
}
func (m *PaintSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintSpec.Merge(m, src)
}
func (m *PaintSpec) XXX_Size() int {
	return xxx_messageInfo_PaintSpec.Size(m)
}
func (m *PaintSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PaintSpec proto.InternalMessageInfo

type isPaintSpec_PaintType interface {
	isPaintSpec_PaintType()
}

type PaintSpec_Acrylic struct {
	Acrylic *AcrylicType `protobuf:"bytes,2,opt,name=acrylic,proto3,oneof" json:"acrylic,omitempty"`
}
type PaintSpec_Oil struct {
	Oil *OilType `protobuf:"bytes,3,opt,name=oil,proto3,oneof" json:"oil,omitempty"`
}

func (*PaintSpec_Acrylic) isPaintSpec_PaintType() {}
func (*PaintSpec_Oil) isPaintSpec_PaintType()     {}

func (m *PaintSpec) GetPaintType() isPaintSpec_PaintType {
	if m != nil {
		return m.PaintType
	}
	return nil
}

func (m *PaintSpec) GetColor() *PaintColor {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *PaintSpec) GetAcrylic() *AcrylicType {
	if x, ok := m.GetPaintType().(*PaintSpec_Acrylic); ok {
		return x.Acrylic
	}
	return nil
}

func (m *PaintSpec) GetOil() *OilType {
	if x, ok := m.GetPaintType().(*PaintSpec_Oil); ok {
		return x.Oil
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PaintSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PaintSpec_Acrylic)(nil),
		(*PaintSpec_Oil)(nil),
	}
}

type PaintColor struct {
	Hue                  string   `protobuf:"bytes,1,opt,name=hue,proto3" json:"hue,omitempty"`
	Value                float32  `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaintColor) Reset()         { *m = PaintColor{} }
func (m *PaintColor) String() string { return proto.CompactTextString(m) }
func (*PaintColor) ProtoMessage()    {}
func (*PaintColor) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{1}
}
func (m *PaintColor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintColor.Unmarshal(m, b)
}
func (m *PaintColor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintColor.Marshal(b, m, deterministic)
}
func (m *PaintColor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintColor.Merge(m, src)
}
func (m *PaintColor) XXX_Size() int {
	return xxx_messageInfo_PaintColor.Size(m)
}
func (m *PaintColor) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintColor.DiscardUnknown(m)
}

var xxx_messageInfo_PaintColor proto.InternalMessageInfo

func (m *PaintColor) GetHue() string {
	if m != nil {
		return m.Hue
	}
	return ""
}

func (m *PaintColor) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AcrylicType struct {
	Body                 AcrylicType_Body `protobuf:"varint,3,opt,name=body,proto3,enum=things.test.io.AcrylicType_Body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AcrylicType) Reset()         { *m = AcrylicType{} }
func (m *AcrylicType) String() string { return proto.CompactTextString(m) }
func (*AcrylicType) ProtoMessage()    {}
func (*AcrylicType) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{2}
}
func (m *AcrylicType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcrylicType.Unmarshal(m, b)
}
func (m *AcrylicType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcrylicType.Marshal(b, m, deterministic)
}
func (m *AcrylicType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcrylicType.Merge(m, src)
}
func (m *AcrylicType) XXX_Size() int {
	return xxx_messageInfo_AcrylicType.Size(m)
}
func (m *AcrylicType) XXX_DiscardUnknown() {
	xxx_messageInfo_AcrylicType.DiscardUnknown(m)
}

var xxx_messageInfo_AcrylicType proto.InternalMessageInfo

func (m *AcrylicType) GetBody() AcrylicType_Body {
	if m != nil {
		return m.Body
	}
	return AcrylicType_Light
}

type OilType struct {
	WaterMixable         bool     `protobuf:"varint,1,opt,name=waterMixable,proto3" json:"waterMixable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OilType) Reset()         { *m = OilType{} }
func (m *OilType) String() string { return proto.CompactTextString(m) }
func (*OilType) ProtoMessage()    {}
func (*OilType) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{3}
}
func (m *OilType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OilType.Unmarshal(m, b)
}
func (m *OilType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OilType.Marshal(b, m, deterministic)
}
func (m *OilType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OilType.Merge(m, src)
}
func (m *OilType) XXX_Size() int {
	return xxx_messageInfo_OilType.Size(m)
}
func (m *OilType) XXX_DiscardUnknown() {
	xxx_messageInfo_OilType.DiscardUnknown(m)
}

var xxx_messageInfo_OilType proto.InternalMessageInfo

func (m *OilType) GetWaterMixable() bool {
	if m != nil {
		return m.WaterMixable
	}
	return false
}

type PaintStatus struct {
	ObservedGeneration   int64    `protobuf:"varint,1,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	PercentRemaining     int64    `protobuf:"varint,2,opt,name=percentRemaining,proto3" json:"percentRemaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaintStatus) Reset()         { *m = PaintStatus{} }
func (m *PaintStatus) String() string { return proto.CompactTextString(m) }
func (*PaintStatus) ProtoMessage()    {}
func (*PaintStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{4}
}
func (m *PaintStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintStatus.Unmarshal(m, b)
}
func (m *PaintStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintStatus.Marshal(b, m, deterministic)
}
func (m *PaintStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintStatus.Merge(m, src)
}
func (m *PaintStatus) XXX_Size() int {
	return xxx_messageInfo_PaintStatus.Size(m)
}
func (m *PaintStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PaintStatus proto.InternalMessageInfo

func (m *PaintStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *PaintStatus) GetPercentRemaining() int64 {
	if m != nil {
		return m.PercentRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("things.test.io.AcrylicType_Body", AcrylicType_Body_name, AcrylicType_Body_value)
	proto.RegisterType((*PaintSpec)(nil), "things.test.io.PaintSpec")
	proto.RegisterType((*PaintColor)(nil), "things.test.io.PaintColor")
	proto.RegisterType((*AcrylicType)(nil), "things.test.io.AcrylicType")
	proto.RegisterType((*OilType)(nil), "things.test.io.OilType")
	proto.RegisterType((*PaintStatus)(nil), "things.test.io.PaintStatus")
}

func init() { proto.RegisterFile("test_api.proto", fileDescriptor_77683351be7bc655) }

var fileDescriptor_77683351be7bc655 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd4, 0x30,
	0x10, 0x85, 0xbb, 0x9b, 0x6e, 0xcb, 0xce, 0xa2, 0x55, 0x64, 0x21, 0xb1, 0x82, 0x4b, 0x95, 0x53,
	0x05, 0xaa, 0x03, 0xa5, 0x12, 0xe7, 0x6e, 0x0f, 0xf4, 0x40, 0x05, 0x32, 0x9c, 0xb8, 0x20, 0x27,
	0x19, 0x65, 0x47, 0xf2, 0x7a, 0x2c, 0xc7, 0x09, 0xe4, 0x5f, 0xf1, 0x13, 0x91, 0x9d, 0x22, 0xd8,
	0x82, 0xb8, 0x79, 0xde, 0x7c, 0x4f, 0xf3, 0xc6, 0x36, 0xac, 0x03, 0x76, 0xe1, 0xab, 0x76, 0x24,
	0x9d, 0xe7, 0xc0, 0x62, 0x1d, 0x76, 0x64, 0xdb, 0x4e, 0x46, 0x59, 0x12, 0x17, 0x3f, 0x66, 0xb0,
	0xfc, 0xa8, 0xc9, 0x86, 0x4f, 0x0e, 0x6b, 0xf1, 0x0a, 0x16, 0x35, 0x1b, 0xf6, 0x9b, 0xd9, 0xd9,
	0xec, 0x7c, 0x75, 0xf9, 0x4c, 0x1e, 0xd2, 0x32, 0x91, 0x37, 0x91, 0x50, 0x13, 0x28, 0xde, 0xc2,
	0xa9, 0xae, 0xfd, 0x68, 0xa8, 0xde, 0xcc, 0x93, 0xe7, 0xf9, 0x43, 0xcf, 0xf5, 0xd4, 0xfe, 0x3c,
	0x3a, 0xbc, 0x3d, 0x52, 0xbf, 0x68, 0xf1, 0x12, 0x32, 0x26, 0xb3, 0xc9, 0x92, 0xe9, 0xe9, 0x43,
	0xd3, 0x07, 0x32, 0xf7, 0x86, 0x48, 0x6d, 0x57, 0xb0, 0x74, 0x71, 0x74, 0xd4, 0x8a, 0x2b, 0x80,
	0xdf, 0x39, 0x44, 0x0e, 0xd9, 0xae, 0xc7, 0x14, 0x78, 0xa9, 0xe2, 0x51, 0x3c, 0x81, 0xc5, 0xa0,
	0x4d, 0x8f, 0x29, 0xd0, 0x5c, 0x4d, 0x45, 0xb1, 0x87, 0xd5, 0x1f, 0x49, 0xc4, 0x15, 0x1c, 0x57,
	0xdc, 0x8c, 0x69, 0xfe, 0xfa, 0xf2, 0xec, 0x3f, 0xa1, 0xe5, 0x96, 0x9b, 0x51, 0x25, 0xba, 0x38,
	0x87, 0xe3, 0x58, 0x89, 0x25, 0x2c, 0xde, 0x53, 0xbb, 0x0b, 0xf9, 0x91, 0x00, 0x38, 0xb9, 0xc3,
	0x86, 0xfa, 0x7d, 0x3e, 0x8b, 0xf2, 0x2d, 0xea, 0x61, 0xcc, 0xe7, 0xc5, 0x05, 0x9c, 0xde, 0xef,
	0x20, 0x0a, 0x78, 0xfc, 0x4d, 0x07, 0xf4, 0x77, 0xf4, 0x5d, 0x57, 0x66, 0x8a, 0xfa, 0x48, 0x1d,
	0x68, 0x05, 0xc1, 0x6a, 0x7a, 0x85, 0xa0, 0x43, 0xdf, 0x09, 0x09, 0x82, 0xab, 0x0e, 0xfd, 0x80,
	0xcd, 0x3b, 0xb4, 0xe8, 0x75, 0x20, 0xb6, 0xc9, 0x98, 0xa9, 0x7f, 0x74, 0xc4, 0x0b, 0xc8, 0x1d,
	0xfa, 0x1a, 0x6d, 0x50, 0xb8, 0xd7, 0x64, 0xc9, 0xb6, 0x69, 0xfb, 0x4c, 0xfd, 0xa5, 0x6f, 0x6f,
	0xbe, 0x5c, 0xb7, 0x14, 0x76, 0x7d, 0x25, 0x6b, 0xde, 0x97, 0x1d, 0x1b, 0xbe, 0x20, 0x2e, 0x75,
	0x1f, 0xd8, 0x91, 0xe1, 0x50, 0xd6, 0xdc, 0x60, 0x8b, 0xb6, 0xf4, 0x68, 0x1b, 0xf4, 0xa5, 0x76,
	0x54, 0x1e, 0x5e, 0x4e, 0x39, 0xbc, 0xae, 0x4e, 0xd2, 0x6f, 0x7a, 0xf3, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x52, 0xf7, 0x98, 0xea, 0x5f, 0x02, 0x00, 0x00,
}