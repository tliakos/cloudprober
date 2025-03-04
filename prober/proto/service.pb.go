// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: github.com/google/cloudprober/prober/proto/service.proto

package proto

import (
	context "context"
	proto "github.com/google/cloudprober/probes/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProbeConfig *proto.ProbeDef `protobuf:"bytes,1,opt,name=probe_config,json=probeConfig" json:"probe_config,omitempty"`
}

func (x *AddProbeRequest) Reset() {
	*x = AddProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProbeRequest) ProtoMessage() {}

func (x *AddProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProbeRequest.ProtoReflect.Descriptor instead.
func (*AddProbeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddProbeRequest) GetProbeConfig() *proto.ProbeDef {
	if x != nil {
		return x.ProbeConfig
	}
	return nil
}

type AddProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddProbeResponse) Reset() {
	*x = AddProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProbeResponse) ProtoMessage() {}

func (x *AddProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProbeResponse.ProtoReflect.Descriptor instead.
func (*AddProbeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{1}
}

type RemoveProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProbeName *string `protobuf:"bytes,1,opt,name=probe_name,json=probeName" json:"probe_name,omitempty"`
}

func (x *RemoveProbeRequest) Reset() {
	*x = RemoveProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProbeRequest) ProtoMessage() {}

func (x *RemoveProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProbeRequest.ProtoReflect.Descriptor instead.
func (*RemoveProbeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveProbeRequest) GetProbeName() string {
	if x != nil && x.ProbeName != nil {
		return *x.ProbeName
	}
	return ""
}

type RemoveProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveProbeResponse) Reset() {
	*x = RemoveProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProbeResponse) ProtoMessage() {}

func (x *RemoveProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProbeResponse.ProtoReflect.Descriptor instead.
func (*RemoveProbeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{3}
}

type ListProbesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProbesRequest) Reset() {
	*x = ListProbesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProbesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProbesRequest) ProtoMessage() {}

func (x *ListProbesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProbesRequest.ProtoReflect.Descriptor instead.
func (*ListProbesRequest) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{4}
}

type Probe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Config *proto.ProbeDef `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (x *Probe) Reset() {
	*x = Probe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Probe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Probe) ProtoMessage() {}

func (x *Probe) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Probe.ProtoReflect.Descriptor instead.
func (*Probe) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *Probe) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Probe) GetConfig() *proto.ProbeDef {
	if x != nil {
		return x.Config
	}
	return nil
}

type ListProbesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Probe []*Probe `protobuf:"bytes,1,rep,name=probe" json:"probe,omitempty"`
}

func (x *ListProbesResponse) Reset() {
	*x = ListProbesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProbesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProbesResponse) ProtoMessage() {}

func (x *ListProbesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProbesResponse.ProtoReflect.Descriptor instead.
func (*ListProbesResponse) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListProbesResponse) GetProbe() []*Probe {
	if x != nil {
		return x.Probe
	}
	return nil
}

var File_github_com_google_cloudprober_prober_proto_service_proto protoreflect.FileDescriptor

var file_github_com_google_cloudprober_prober_proto_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x1a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x44, 0x65, 0x66, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a,
	0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x44, 0x65, 0x66, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3e, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x32, 0xfd, 0x01, 0x0a,
	0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_google_cloudprober_prober_proto_service_proto_rawDescOnce sync.Once
	file_github_com_google_cloudprober_prober_proto_service_proto_rawDescData = file_github_com_google_cloudprober_prober_proto_service_proto_rawDesc
)

func file_github_com_google_cloudprober_prober_proto_service_proto_rawDescGZIP() []byte {
	file_github_com_google_cloudprober_prober_proto_service_proto_rawDescOnce.Do(func() {
		file_github_com_google_cloudprober_prober_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_google_cloudprober_prober_proto_service_proto_rawDescData)
	})
	return file_github_com_google_cloudprober_prober_proto_service_proto_rawDescData
}

var file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_google_cloudprober_prober_proto_service_proto_goTypes = []interface{}{
	(*AddProbeRequest)(nil),     // 0: cloudprober.AddProbeRequest
	(*AddProbeResponse)(nil),    // 1: cloudprober.AddProbeResponse
	(*RemoveProbeRequest)(nil),  // 2: cloudprober.RemoveProbeRequest
	(*RemoveProbeResponse)(nil), // 3: cloudprober.RemoveProbeResponse
	(*ListProbesRequest)(nil),   // 4: cloudprober.ListProbesRequest
	(*Probe)(nil),               // 5: cloudprober.Probe
	(*ListProbesResponse)(nil),  // 6: cloudprober.ListProbesResponse
	(*proto.ProbeDef)(nil),      // 7: cloudprober.probes.ProbeDef
}
var file_github_com_google_cloudprober_prober_proto_service_proto_depIdxs = []int32{
	7, // 0: cloudprober.AddProbeRequest.probe_config:type_name -> cloudprober.probes.ProbeDef
	7, // 1: cloudprober.Probe.config:type_name -> cloudprober.probes.ProbeDef
	5, // 2: cloudprober.ListProbesResponse.probe:type_name -> cloudprober.Probe
	0, // 3: cloudprober.Cloudprober.AddProbe:input_type -> cloudprober.AddProbeRequest
	2, // 4: cloudprober.Cloudprober.RemoveProbe:input_type -> cloudprober.RemoveProbeRequest
	4, // 5: cloudprober.Cloudprober.ListProbes:input_type -> cloudprober.ListProbesRequest
	1, // 6: cloudprober.Cloudprober.AddProbe:output_type -> cloudprober.AddProbeResponse
	3, // 7: cloudprober.Cloudprober.RemoveProbe:output_type -> cloudprober.RemoveProbeResponse
	6, // 8: cloudprober.Cloudprober.ListProbes:output_type -> cloudprober.ListProbesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_google_cloudprober_prober_proto_service_proto_init() }
func file_github_com_google_cloudprober_prober_proto_service_proto_init() {
	if File_github_com_google_cloudprober_prober_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProbeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProbeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProbeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProbeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProbesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Probe); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProbesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_google_cloudprober_prober_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_google_cloudprober_prober_proto_service_proto_goTypes,
		DependencyIndexes: file_github_com_google_cloudprober_prober_proto_service_proto_depIdxs,
		MessageInfos:      file_github_com_google_cloudprober_prober_proto_service_proto_msgTypes,
	}.Build()
	File_github_com_google_cloudprober_prober_proto_service_proto = out.File
	file_github_com_google_cloudprober_prober_proto_service_proto_rawDesc = nil
	file_github_com_google_cloudprober_prober_proto_service_proto_goTypes = nil
	file_github_com_google_cloudprober_prober_proto_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CloudproberClient is the client API for Cloudprober service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudproberClient interface {
	// AddProbe adds a probe to cloudprober. Error is returned if probe is already
	// defined or there is an error during initialization of the probe.
	AddProbe(ctx context.Context, in *AddProbeRequest, opts ...grpc.CallOption) (*AddProbeResponse, error)
	// RemoveProbe stops the probe and removes it from the in-memory database.
	RemoveProbe(ctx context.Context, in *RemoveProbeRequest, opts ...grpc.CallOption) (*RemoveProbeResponse, error)
	// ListProbes lists active probes.
	ListProbes(ctx context.Context, in *ListProbesRequest, opts ...grpc.CallOption) (*ListProbesResponse, error)
}

type cloudproberClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudproberClient(cc grpc.ClientConnInterface) CloudproberClient {
	return &cloudproberClient{cc}
}

func (c *cloudproberClient) AddProbe(ctx context.Context, in *AddProbeRequest, opts ...grpc.CallOption) (*AddProbeResponse, error) {
	out := new(AddProbeResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.Cloudprober/AddProbe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudproberClient) RemoveProbe(ctx context.Context, in *RemoveProbeRequest, opts ...grpc.CallOption) (*RemoveProbeResponse, error) {
	out := new(RemoveProbeResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.Cloudprober/RemoveProbe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudproberClient) ListProbes(ctx context.Context, in *ListProbesRequest, opts ...grpc.CallOption) (*ListProbesResponse, error) {
	out := new(ListProbesResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.Cloudprober/ListProbes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudproberServer is the server API for Cloudprober service.
type CloudproberServer interface {
	// AddProbe adds a probe to cloudprober. Error is returned if probe is already
	// defined or there is an error during initialization of the probe.
	AddProbe(context.Context, *AddProbeRequest) (*AddProbeResponse, error)
	// RemoveProbe stops the probe and removes it from the in-memory database.
	RemoveProbe(context.Context, *RemoveProbeRequest) (*RemoveProbeResponse, error)
	// ListProbes lists active probes.
	ListProbes(context.Context, *ListProbesRequest) (*ListProbesResponse, error)
}

// UnimplementedCloudproberServer can be embedded to have forward compatible implementations.
type UnimplementedCloudproberServer struct {
}

func (*UnimplementedCloudproberServer) AddProbe(context.Context, *AddProbeRequest) (*AddProbeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProbe not implemented")
}
func (*UnimplementedCloudproberServer) RemoveProbe(context.Context, *RemoveProbeRequest) (*RemoveProbeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProbe not implemented")
}
func (*UnimplementedCloudproberServer) ListProbes(context.Context, *ListProbesRequest) (*ListProbesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProbes not implemented")
}

func RegisterCloudproberServer(s *grpc.Server, srv CloudproberServer) {
	s.RegisterService(&_Cloudprober_serviceDesc, srv)
}

func _Cloudprober_AddProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudproberServer).AddProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.Cloudprober/AddProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudproberServer).AddProbe(ctx, req.(*AddProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cloudprober_RemoveProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudproberServer).RemoveProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.Cloudprober/RemoveProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudproberServer).RemoveProbe(ctx, req.(*RemoveProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cloudprober_ListProbes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProbesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudproberServer).ListProbes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.Cloudprober/ListProbes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudproberServer).ListProbes(ctx, req.(*ListProbesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cloudprober_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprober.Cloudprober",
	HandlerType: (*CloudproberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProbe",
			Handler:    _Cloudprober_AddProbe_Handler,
		},
		{
			MethodName: "RemoveProbe",
			Handler:    _Cloudprober_RemoveProbe_Handler,
		},
		{
			MethodName: "ListProbes",
			Handler:    _Cloudprober_ListProbes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/google/cloudprober/prober/proto/service.proto",
}
