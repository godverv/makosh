// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: grpc/makosh-be_api.proto

package makosh_be

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{0}
}

type ListEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEndpoints) Reset() {
	*x = ListEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpoints) ProtoMessage() {}

func (x *ListEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpoints.ProtoReflect.Descriptor instead.
func (*ListEndpoints) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{1}
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Addrs       []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{2}
}

func (x *Endpoint) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Endpoint) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

type UpsertEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpsertEndpoints) Reset() {
	*x = UpsertEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertEndpoints) ProtoMessage() {}

func (x *UpsertEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertEndpoints.ProtoReflect.Descriptor instead.
func (*UpsertEndpoints) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{3}
}

type Version_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Version_Request) Reset() {
	*x = Version_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Request) ProtoMessage() {}

func (x *Version_Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Request.ProtoReflect.Descriptor instead.
func (*Version_Request) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{0, 0}
}

type Version_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Version_Response) Reset() {
	*x = Version_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Response) ProtoMessage() {}

func (x *Version_Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Response.ProtoReflect.Descriptor instead.
func (*Version_Response) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Version_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ListEndpoints_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *ListEndpoints_Request) Reset() {
	*x = ListEndpoints_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpoints_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpoints_Request) ProtoMessage() {}

func (x *ListEndpoints_Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpoints_Request.ProtoReflect.Descriptor instead.
func (*ListEndpoints_Request) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListEndpoints_Request) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type ListEndpoints_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ListEndpoints_Response) Reset() {
	*x = ListEndpoints_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpoints_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpoints_Response) ProtoMessage() {}

func (x *ListEndpoints_Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpoints_Response.ProtoReflect.Descriptor instead.
func (*ListEndpoints_Response) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ListEndpoints_Response) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

type UpsertEndpoints_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []*Endpoint `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *UpsertEndpoints_Request) Reset() {
	*x = UpsertEndpoints_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertEndpoints_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertEndpoints_Request) ProtoMessage() {}

func (x *UpsertEndpoints_Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertEndpoints_Request.ProtoReflect.Descriptor instead.
func (*UpsertEndpoints_Request) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{3, 0}
}

func (x *UpsertEndpoints_Request) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type UpsertEndpoints_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpsertEndpoints_Response) Reset() {
	*x = UpsertEndpoints_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_makosh_be_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertEndpoints_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertEndpoints_Response) ProtoMessage() {}

func (x *UpsertEndpoints_Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_makosh_be_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertEndpoints_Response.ProtoReflect.Descriptor instead.
func (*UpsertEndpoints_Response) Descriptor() ([]byte, []int) {
	return file_grpc_makosh_be_api_proto_rawDescGZIP(), []int{3, 1}
}

var File_grpc_makosh_be_api_proto protoreflect.FileDescriptor

var file_grpc_makosh_be_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x6b, 0x6f, 0x73, 0x68, 0x2d, 0x62, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x6b, 0x6f,
	0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x1a, 0x2c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x22, 0x43, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x22, 0x5f, 0x0a, 0x0f, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x40, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x6b, 0x6f, 0x73,
	0x68, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x0a, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf1, 0x02, 0x0a, 0x0b, 0x6d, 0x61, 0x6b,
	0x6f, 0x73, 0x68, 0x42, 0x65, 0x41, 0x50, 0x49, 0x12, 0x5f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6b, 0x6f, 0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x6b, 0x6f, 0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x82, 0x01, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x6d, 0x61,
	0x6b, 0x6f, 0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x6b, 0x6f, 0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7c,
	0x0a, 0x0f, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x26, 0x2e, 0x6d, 0x61, 0x6b, 0x6f, 0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x61, 0x6b, 0x6f,
	0x73, 0x68, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x0c, 0x5a, 0x0a,
	0x2f, 0x6d, 0x61, 0x6b, 0x6f, 0x73, 0x68, 0x5f, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_makosh_be_api_proto_rawDescOnce sync.Once
	file_grpc_makosh_be_api_proto_rawDescData = file_grpc_makosh_be_api_proto_rawDesc
)

func file_grpc_makosh_be_api_proto_rawDescGZIP() []byte {
	file_grpc_makosh_be_api_proto_rawDescOnce.Do(func() {
		file_grpc_makosh_be_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_makosh_be_api_proto_rawDescData)
	})
	return file_grpc_makosh_be_api_proto_rawDescData
}

var file_grpc_makosh_be_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_grpc_makosh_be_api_proto_goTypes = []interface{}{
	(*Version)(nil),                  // 0: makosh_be_api.Version
	(*ListEndpoints)(nil),            // 1: makosh_be_api.ListEndpoints
	(*Endpoint)(nil),                 // 2: makosh_be_api.Endpoint
	(*UpsertEndpoints)(nil),          // 3: makosh_be_api.UpsertEndpoints
	(*Version_Request)(nil),          // 4: makosh_be_api.Version.Request
	(*Version_Response)(nil),         // 5: makosh_be_api.Version.Response
	(*ListEndpoints_Request)(nil),    // 6: makosh_be_api.ListEndpoints.Request
	(*ListEndpoints_Response)(nil),   // 7: makosh_be_api.ListEndpoints.Response
	(*UpsertEndpoints_Request)(nil),  // 8: makosh_be_api.UpsertEndpoints.Request
	(*UpsertEndpoints_Response)(nil), // 9: makosh_be_api.UpsertEndpoints.Response
}
var file_grpc_makosh_be_api_proto_depIdxs = []int32{
	2, // 0: makosh_be_api.UpsertEndpoints.Request.endpoints:type_name -> makosh_be_api.Endpoint
	4, // 1: makosh_be_api.makoshBeAPI.Version:input_type -> makosh_be_api.Version.Request
	6, // 2: makosh_be_api.makoshBeAPI.ListEndpoints:input_type -> makosh_be_api.ListEndpoints.Request
	8, // 3: makosh_be_api.makoshBeAPI.UpsertEndpoints:input_type -> makosh_be_api.UpsertEndpoints.Request
	5, // 4: makosh_be_api.makoshBeAPI.Version:output_type -> makosh_be_api.Version.Response
	7, // 5: makosh_be_api.makoshBeAPI.ListEndpoints:output_type -> makosh_be_api.ListEndpoints.Response
	9, // 6: makosh_be_api.makoshBeAPI.UpsertEndpoints:output_type -> makosh_be_api.UpsertEndpoints.Response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_makosh_be_api_proto_init() }
func file_grpc_makosh_be_api_proto_init() {
	if File_grpc_makosh_be_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_makosh_be_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpoints); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertEndpoints); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version_Request); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version_Response); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpoints_Request); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpoints_Response); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertEndpoints_Request); i {
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
		file_grpc_makosh_be_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertEndpoints_Response); i {
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
			RawDescriptor: file_grpc_makosh_be_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_makosh_be_api_proto_goTypes,
		DependencyIndexes: file_grpc_makosh_be_api_proto_depIdxs,
		MessageInfos:      file_grpc_makosh_be_api_proto_msgTypes,
	}.Build()
	File_grpc_makosh_be_api_proto = out.File
	file_grpc_makosh_be_api_proto_rawDesc = nil
	file_grpc_makosh_be_api_proto_goTypes = nil
	file_grpc_makosh_be_api_proto_depIdxs = nil
}
