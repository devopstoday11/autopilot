// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilot-operator.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// MeshProviders provide an interface to monitoring and managing a specific
// mesh.
// Autopilot does not abstract the mesh API - Autopilot developers must
// still reason able about Provider-specific CRDs. Autopilot's job is to
// abstract operational concerns such as discovering control plane configuration
// and monitoring metrics.
type MeshProvider int32

const (
	// the Operator will utilize Istio mesh for metrics and configuration
	MeshProvider_Istio MeshProvider = 0
	// the Operator will utilize a locally deployed Prometheus instance for metrics
	// (Currently unimplemented)
	MeshProvider_Custom MeshProvider = 1
)

var MeshProvider_name = map[int32]string{
	0: "Istio",
	1: "Custom",
}

var MeshProvider_value = map[string]int32{
	"Istio":  0,
	"Custom": 1,
}

func (x MeshProvider) String() string {
	return proto.EnumName(MeshProvider_name, int32(x))
}

func (MeshProvider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f975433f2c607a, []int{0}
}

// The AutopilotOperator file is the bootstrap
// Configuration file for the Operator.
// It is stored and mounted to the operator as a Kubernetes ConfigMap.
// The Operator will hot-reload when the configuration file changes.
// Default name is 'autopilot-operator.yaml' and should be stored in the project root.
type AutopilotOperator struct {
	// version of the operator
	// used for logging and metrics
	// default is "0.0.1"
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// meshProvider determines how the operator will connect to a service mesh
	// Default is "SMI"
	MeshProvider MeshProvider `protobuf:"varint,2,opt,name=meshProvider,proto3,enum=autopilot.MeshProvider" json:"meshProvider,omitempty"`
	// controlPlaneNs is the namespace the control plane lives in
	// Default is "istio-system"
	ControlPlaneNs string `protobuf:"bytes,3,opt,name=controlPlaneNs,proto3" json:"controlPlaneNs,omitempty"`
	// workInterval to sets the interval at which CRD workers resync.
	// Default is 5s
	WorkInterval *duration.Duration `protobuf:"bytes,4,opt,name=workInterval,proto3" json:"workInterval,omitempty"`
	// Serve metrics on this address. Set to empty string to disable metrics
	// defaults to ":9091"
	MetricsAddr string `protobuf:"bytes,5,opt,name=metricsAddr,proto3" json:"metricsAddr,omitempty"`
	// Enable leader election. This will prevent more than one operator from running at a time
	// defaults to true
	EnableLeaderElection bool `protobuf:"varint,6,opt,name=enableLeaderElection,proto3" json:"enableLeaderElection,omitempty"`
	// if non-empty, watchNamespace will restrict the Operator to watching resources in a single namespace
	// if empty (default), the Operator must have Cluster-scope RBAC permissions (ClusterRole/Binding)
	// can also be set via the WATCH_NAMESPACE environment variable
	WatchNamespace string `protobuf:"bytes,7,opt,name=watchNamespace,proto3" json:"watchNamespace,omitempty"`
	// The namespace to use for Leader Election (requires read/write ConfigMap permissions)
	// defaults to the watchNamespace
	LeaderElectionNamespace string `protobuf:"bytes,8,opt,name=leaderElectionNamespace,proto3" json:"leaderElectionNamespace,omitempty"`
	// Log level for the operator's logger
	// values:
	// 0 - Debug
	// 1 - Info
	// 2 - Warn
	// 3 - Error
	// 4 - DPanic
	// 5 - Panic
	// 6 - Fatal
	// Defaults to Info
	LogLevel             *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AutopilotOperator) Reset()         { *m = AutopilotOperator{} }
func (m *AutopilotOperator) String() string { return proto.CompactTextString(m) }
func (*AutopilotOperator) ProtoMessage()    {}
func (*AutopilotOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f975433f2c607a, []int{0}
}

func (m *AutopilotOperator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutopilotOperator.Unmarshal(m, b)
}
func (m *AutopilotOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutopilotOperator.Marshal(b, m, deterministic)
}
func (m *AutopilotOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutopilotOperator.Merge(m, src)
}
func (m *AutopilotOperator) XXX_Size() int {
	return xxx_messageInfo_AutopilotOperator.Size(m)
}
func (m *AutopilotOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_AutopilotOperator.DiscardUnknown(m)
}

var xxx_messageInfo_AutopilotOperator proto.InternalMessageInfo

func (m *AutopilotOperator) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AutopilotOperator) GetMeshProvider() MeshProvider {
	if m != nil {
		return m.MeshProvider
	}
	return MeshProvider_Istio
}

func (m *AutopilotOperator) GetControlPlaneNs() string {
	if m != nil {
		return m.ControlPlaneNs
	}
	return ""
}

func (m *AutopilotOperator) GetWorkInterval() *duration.Duration {
	if m != nil {
		return m.WorkInterval
	}
	return nil
}

func (m *AutopilotOperator) GetMetricsAddr() string {
	if m != nil {
		return m.MetricsAddr
	}
	return ""
}

func (m *AutopilotOperator) GetEnableLeaderElection() bool {
	if m != nil {
		return m.EnableLeaderElection
	}
	return false
}

func (m *AutopilotOperator) GetWatchNamespace() string {
	if m != nil {
		return m.WatchNamespace
	}
	return ""
}

func (m *AutopilotOperator) GetLeaderElectionNamespace() string {
	if m != nil {
		return m.LeaderElectionNamespace
	}
	return ""
}

func (m *AutopilotOperator) GetLogLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func init() {
	proto.RegisterEnum("autopilot.MeshProvider", MeshProvider_name, MeshProvider_value)
	proto.RegisterType((*AutopilotOperator)(nil), "autopilot.AutopilotOperator")
}

func init() { proto.RegisterFile("autopilot-operator.proto", fileDescriptor_56f975433f2c607a) }

var fileDescriptor_56f975433f2c607a = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0xcb, 0xd3, 0x40,
	0x14, 0x85, 0x8d, 0xaf, 0xfd, 0xc8, 0xb4, 0x94, 0x3a, 0x08, 0x1d, 0x45, 0x24, 0x28, 0x95, 0x20,
	0x34, 0xc1, 0x74, 0x53, 0x10, 0x17, 0xf5, 0x63, 0x51, 0xa8, 0xb5, 0x04, 0x74, 0xe1, 0x6e, 0x92,
	0x5c, 0xd3, 0xc1, 0x49, 0x6e, 0x98, 0x99, 0xa4, 0x3f, 0xd0, 0x3f, 0x26, 0x4d, 0xda, 0xd8, 0x58,
	0xdf, 0x65, 0xee, 0x73, 0xee, 0x3d, 0xe7, 0x64, 0x08, 0xe3, 0xa5, 0xc1, 0x42, 0x48, 0x34, 0x0b,
	0x2c, 0x40, 0x71, 0x83, 0xca, 0x2b, 0x14, 0x1a, 0xa4, 0x76, 0x4b, 0x9e, 0xbd, 0x48, 0x11, 0x53,
	0x09, 0x7e, 0x0d, 0xa2, 0xf2, 0xa7, 0x9f, 0x94, 0x8a, 0x1b, 0x81, 0x79, 0x23, 0xbd, 0xe5, 0x47,
	0xc5, 0x8b, 0x02, 0x94, 0x6e, 0xf8, 0xcb, 0xdf, 0x77, 0xe4, 0xf1, 0xfa, 0x72, 0xed, 0xeb, 0xd9,
	0x86, 0x32, 0x32, 0xa8, 0x40, 0x69, 0x81, 0x39, 0xb3, 0x1c, 0xcb, 0xb5, 0xc3, 0xcb, 0x27, 0x7d,
	0x47, 0xc6, 0x19, 0xe8, 0xc3, 0x5e, 0x61, 0x25, 0x12, 0x50, 0xec, 0xa1, 0x63, 0xb9, 0x93, 0x60,
	0xe6, 0xb5, 0x89, 0xbc, 0x2f, 0x57, 0x38, 0xec, 0x88, 0xe9, 0x6b, 0x32, 0x89, 0x31, 0x37, 0x0a,
	0xe5, 0x5e, 0xf2, 0x1c, 0x76, 0x9a, 0xdd, 0xd5, 0xd7, 0xff, 0x99, 0xd2, 0xf7, 0x64, 0x7c, 0x44,
	0xf5, 0x6b, 0x93, 0x1b, 0x50, 0x15, 0x97, 0xec, 0x91, 0x63, 0xb9, 0xa3, 0xe0, 0xa9, 0xd7, 0x74,
	0xf1, 0x2e, 0x5d, 0xbc, 0x4f, 0xe7, 0xae, 0x61, 0x47, 0x4e, 0x1d, 0x32, 0xca, 0xc0, 0x28, 0x11,
	0xeb, 0x75, 0x92, 0x28, 0xd6, 0xab, 0x3d, 0xae, 0x47, 0x34, 0x20, 0x4f, 0x20, 0xe7, 0x91, 0x84,
	0x2d, 0xf0, 0x04, 0xd4, 0x67, 0x09, 0xf1, 0xe9, 0x0e, 0xeb, 0x3b, 0x96, 0x3b, 0x0c, 0xff, 0xcb,
	0x4e, 0xe1, 0x8f, 0xdc, 0xc4, 0x87, 0x1d, 0xcf, 0x40, 0x17, 0x3c, 0x06, 0x36, 0x68, 0xc2, 0x77,
	0xa7, 0x74, 0x45, 0x66, 0xb2, 0xb3, 0xf9, 0x77, 0x61, 0x58, 0x2f, 0xdc, 0x87, 0xe9, 0x8a, 0x0c,
	0x25, 0xa6, 0x5b, 0xa8, 0x40, 0x32, 0xbb, 0xae, 0xfc, 0xfc, 0xa6, 0xf2, 0xb7, 0x4d, 0x6e, 0x96,
	0xc1, 0x77, 0x2e, 0x4b, 0x08, 0x5b, 0xf5, 0x9b, 0x39, 0x19, 0x5f, 0xff, 0x76, 0x6a, 0x93, 0xde,
	0x46, 0x1b, 0x81, 0xd3, 0x07, 0x94, 0x90, 0xfe, 0xc7, 0x52, 0x1b, 0xcc, 0xa6, 0xd6, 0x87, 0xf9,
	0x8f, 0x57, 0xa9, 0x30, 0x87, 0x32, 0xf2, 0x62, 0xcc, 0x7c, 0x8d, 0x12, 0x17, 0x02, 0xfd, 0xf6,
	0xe9, 0x7c, 0x5e, 0x08, 0xbf, 0x7a, 0x1b, 0xf5, 0x6b, 0xb7, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0xd0, 0xeb, 0x13, 0x81, 0x02, 0x00, 0x00,
}
