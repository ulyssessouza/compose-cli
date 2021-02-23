//
//  Copyright 2021 Docker Compose CLI authors

//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: cli/server/protos/volumes/v1/volumes.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{0}
}

func (x *Volume) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Volume) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AciVolumeCreateOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageAccount string `protobuf:"bytes,1,opt,name=storage_account,json=storageAccount,proto3" json:"storage_account,omitempty"`
}

func (x *AciVolumeCreateOptions) Reset() {
	*x = AciVolumeCreateOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AciVolumeCreateOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AciVolumeCreateOptions) ProtoMessage() {}

func (x *AciVolumeCreateOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AciVolumeCreateOptions.ProtoReflect.Descriptor instead.
func (*AciVolumeCreateOptions) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{1}
}

func (x *AciVolumeCreateOptions) GetStorageAccount() string {
	if x != nil {
		return x.StorageAccount
	}
	return ""
}

type VolumesCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Options:
	//	*VolumesCreateRequest_AciOption
	Options isVolumesCreateRequest_Options `protobuf_oneof:"options"`
}

func (x *VolumesCreateRequest) Reset() {
	*x = VolumesCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesCreateRequest) ProtoMessage() {}

func (x *VolumesCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesCreateRequest.ProtoReflect.Descriptor instead.
func (*VolumesCreateRequest) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{2}
}

func (x *VolumesCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *VolumesCreateRequest) GetOptions() isVolumesCreateRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *VolumesCreateRequest) GetAciOption() *AciVolumeCreateOptions {
	if x, ok := x.GetOptions().(*VolumesCreateRequest_AciOption); ok {
		return x.AciOption
	}
	return nil
}

type isVolumesCreateRequest_Options interface {
	isVolumesCreateRequest_Options()
}

type VolumesCreateRequest_AciOption struct {
	AciOption *AciVolumeCreateOptions `protobuf:"bytes,2,opt,name=aci_option,json=aciOption,proto3,oneof"`
}

func (*VolumesCreateRequest_AciOption) isVolumesCreateRequest_Options() {}

type VolumesCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume *Volume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *VolumesCreateResponse) Reset() {
	*x = VolumesCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesCreateResponse) ProtoMessage() {}

func (x *VolumesCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesCreateResponse.ProtoReflect.Descriptor instead.
func (*VolumesCreateResponse) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{3}
}

func (x *VolumesCreateResponse) GetVolume() *Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

type VolumesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VolumesListRequest) Reset() {
	*x = VolumesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesListRequest) ProtoMessage() {}

func (x *VolumesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesListRequest.ProtoReflect.Descriptor instead.
func (*VolumesListRequest) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{4}
}

type VolumesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volumes []*Volume `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *VolumesListResponse) Reset() {
	*x = VolumesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesListResponse) ProtoMessage() {}

func (x *VolumesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesListResponse.ProtoReflect.Descriptor instead.
func (*VolumesListResponse) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{5}
}

func (x *VolumesListResponse) GetVolumes() []*Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

type VolumesDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VolumesDeleteRequest) Reset() {
	*x = VolumesDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesDeleteRequest) ProtoMessage() {}

func (x *VolumesDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesDeleteRequest.ProtoReflect.Descriptor instead.
func (*VolumesDeleteRequest) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{6}
}

func (x *VolumesDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VolumesDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VolumesDeleteResponse) Reset() {
	*x = VolumesDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesDeleteResponse) ProtoMessage() {}

func (x *VolumesDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesDeleteResponse.ProtoReflect.Descriptor instead.
func (*VolumesDeleteResponse) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{7}
}

type VolumesInspectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VolumesInspectRequest) Reset() {
	*x = VolumesInspectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesInspectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesInspectRequest) ProtoMessage() {}

func (x *VolumesInspectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesInspectRequest.ProtoReflect.Descriptor instead.
func (*VolumesInspectRequest) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{8}
}

func (x *VolumesInspectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VolumesInspectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume *Volume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *VolumesInspectResponse) Reset() {
	*x = VolumesInspectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumesInspectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumesInspectResponse) ProtoMessage() {}

func (x *VolumesInspectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumesInspectResponse.ProtoReflect.Descriptor instead.
func (*VolumesInspectResponse) Descriptor() ([]byte, []int) {
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP(), []int{9}
}

func (x *VolumesInspectResponse) GetVolume() *Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

var File_cli_server_protos_volumes_v1_volumes_proto protoreflect.FileDescriptor

var file_cli_server_protos_volumes_v1_volumes_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6c, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x06, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x16, 0x41, 0x63, 0x69, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x61, 0x63, 0x69, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x69,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x69, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x59, 0x0a, 0x15, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x13,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x07,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x17, 0x0a, 0x15, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5a, 0x0a, 0x16, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x91, 0x04,
	0x0a, 0x07, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x0d, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a, 0x0b,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0d, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x12, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2d, 0x63,
	0x6c, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cli_server_protos_volumes_v1_volumes_proto_rawDescOnce sync.Once
	file_cli_server_protos_volumes_v1_volumes_proto_rawDescData = file_cli_server_protos_volumes_v1_volumes_proto_rawDesc
)

func file_cli_server_protos_volumes_v1_volumes_proto_rawDescGZIP() []byte {
	file_cli_server_protos_volumes_v1_volumes_proto_rawDescOnce.Do(func() {
		file_cli_server_protos_volumes_v1_volumes_proto_rawDescData = protoimpl.X.CompressGZIP(file_cli_server_protos_volumes_v1_volumes_proto_rawDescData)
	})
	return file_cli_server_protos_volumes_v1_volumes_proto_rawDescData
}

var file_cli_server_protos_volumes_v1_volumes_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cli_server_protos_volumes_v1_volumes_proto_goTypes = []interface{}{
	(*Volume)(nil),                 // 0: com.docker.api.protos.volumes.v1.Volume
	(*AciVolumeCreateOptions)(nil), // 1: com.docker.api.protos.volumes.v1.AciVolumeCreateOptions
	(*VolumesCreateRequest)(nil),   // 2: com.docker.api.protos.volumes.v1.VolumesCreateRequest
	(*VolumesCreateResponse)(nil),  // 3: com.docker.api.protos.volumes.v1.VolumesCreateResponse
	(*VolumesListRequest)(nil),     // 4: com.docker.api.protos.volumes.v1.VolumesListRequest
	(*VolumesListResponse)(nil),    // 5: com.docker.api.protos.volumes.v1.VolumesListResponse
	(*VolumesDeleteRequest)(nil),   // 6: com.docker.api.protos.volumes.v1.VolumesDeleteRequest
	(*VolumesDeleteResponse)(nil),  // 7: com.docker.api.protos.volumes.v1.VolumesDeleteResponse
	(*VolumesInspectRequest)(nil),  // 8: com.docker.api.protos.volumes.v1.VolumesInspectRequest
	(*VolumesInspectResponse)(nil), // 9: com.docker.api.protos.volumes.v1.VolumesInspectResponse
}
var file_cli_server_protos_volumes_v1_volumes_proto_depIdxs = []int32{
	1, // 0: com.docker.api.protos.volumes.v1.VolumesCreateRequest.aci_option:type_name -> com.docker.api.protos.volumes.v1.AciVolumeCreateOptions
	0, // 1: com.docker.api.protos.volumes.v1.VolumesCreateResponse.volume:type_name -> com.docker.api.protos.volumes.v1.Volume
	0, // 2: com.docker.api.protos.volumes.v1.VolumesListResponse.volumes:type_name -> com.docker.api.protos.volumes.v1.Volume
	0, // 3: com.docker.api.protos.volumes.v1.VolumesInspectResponse.volume:type_name -> com.docker.api.protos.volumes.v1.Volume
	2, // 4: com.docker.api.protos.volumes.v1.Volumes.VolumesCreate:input_type -> com.docker.api.protos.volumes.v1.VolumesCreateRequest
	4, // 5: com.docker.api.protos.volumes.v1.Volumes.VolumesList:input_type -> com.docker.api.protos.volumes.v1.VolumesListRequest
	6, // 6: com.docker.api.protos.volumes.v1.Volumes.VolumesDelete:input_type -> com.docker.api.protos.volumes.v1.VolumesDeleteRequest
	8, // 7: com.docker.api.protos.volumes.v1.Volumes.VolumesInspect:input_type -> com.docker.api.protos.volumes.v1.VolumesInspectRequest
	3, // 8: com.docker.api.protos.volumes.v1.Volumes.VolumesCreate:output_type -> com.docker.api.protos.volumes.v1.VolumesCreateResponse
	5, // 9: com.docker.api.protos.volumes.v1.Volumes.VolumesList:output_type -> com.docker.api.protos.volumes.v1.VolumesListResponse
	7, // 10: com.docker.api.protos.volumes.v1.Volumes.VolumesDelete:output_type -> com.docker.api.protos.volumes.v1.VolumesDeleteResponse
	9, // 11: com.docker.api.protos.volumes.v1.Volumes.VolumesInspect:output_type -> com.docker.api.protos.volumes.v1.VolumesInspectResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cli_server_protos_volumes_v1_volumes_proto_init() }
func file_cli_server_protos_volumes_v1_volumes_proto_init() {
	if File_cli_server_protos_volumes_v1_volumes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AciVolumeCreateOptions); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesCreateRequest); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesCreateResponse); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesListRequest); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesListResponse); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesDeleteRequest); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesDeleteResponse); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesInspectRequest); i {
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
		file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumesInspectResponse); i {
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
	file_cli_server_protos_volumes_v1_volumes_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*VolumesCreateRequest_AciOption)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cli_server_protos_volumes_v1_volumes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cli_server_protos_volumes_v1_volumes_proto_goTypes,
		DependencyIndexes: file_cli_server_protos_volumes_v1_volumes_proto_depIdxs,
		MessageInfos:      file_cli_server_protos_volumes_v1_volumes_proto_msgTypes,
	}.Build()
	File_cli_server_protos_volumes_v1_volumes_proto = out.File
	file_cli_server_protos_volumes_v1_volumes_proto_rawDesc = nil
	file_cli_server_protos_volumes_v1_volumes_proto_goTypes = nil
	file_cli_server_protos_volumes_v1_volumes_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VolumesClient is the client API for Volumes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VolumesClient interface {
	VolumesCreate(ctx context.Context, in *VolumesCreateRequest, opts ...grpc.CallOption) (*VolumesCreateResponse, error)
	VolumesList(ctx context.Context, in *VolumesListRequest, opts ...grpc.CallOption) (*VolumesListResponse, error)
	VolumesDelete(ctx context.Context, in *VolumesDeleteRequest, opts ...grpc.CallOption) (*VolumesDeleteResponse, error)
	VolumesInspect(ctx context.Context, in *VolumesInspectRequest, opts ...grpc.CallOption) (*VolumesInspectResponse, error)
}

type volumesClient struct {
	cc grpc.ClientConnInterface
}

func NewVolumesClient(cc grpc.ClientConnInterface) VolumesClient {
	return &volumesClient{cc}
}

func (c *volumesClient) VolumesCreate(ctx context.Context, in *VolumesCreateRequest, opts ...grpc.CallOption) (*VolumesCreateResponse, error) {
	out := new(VolumesCreateResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.volumes.v1.Volumes/VolumesCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) VolumesList(ctx context.Context, in *VolumesListRequest, opts ...grpc.CallOption) (*VolumesListResponse, error) {
	out := new(VolumesListResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.volumes.v1.Volumes/VolumesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) VolumesDelete(ctx context.Context, in *VolumesDeleteRequest, opts ...grpc.CallOption) (*VolumesDeleteResponse, error) {
	out := new(VolumesDeleteResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.volumes.v1.Volumes/VolumesDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) VolumesInspect(ctx context.Context, in *VolumesInspectRequest, opts ...grpc.CallOption) (*VolumesInspectResponse, error) {
	out := new(VolumesInspectResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.volumes.v1.Volumes/VolumesInspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumesServer is the server API for Volumes service.
type VolumesServer interface {
	VolumesCreate(context.Context, *VolumesCreateRequest) (*VolumesCreateResponse, error)
	VolumesList(context.Context, *VolumesListRequest) (*VolumesListResponse, error)
	VolumesDelete(context.Context, *VolumesDeleteRequest) (*VolumesDeleteResponse, error)
	VolumesInspect(context.Context, *VolumesInspectRequest) (*VolumesInspectResponse, error)
}

// UnimplementedVolumesServer can be embedded to have forward compatible implementations.
type UnimplementedVolumesServer struct {
}

func (*UnimplementedVolumesServer) VolumesCreate(context.Context, *VolumesCreateRequest) (*VolumesCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumesCreate not implemented")
}
func (*UnimplementedVolumesServer) VolumesList(context.Context, *VolumesListRequest) (*VolumesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumesList not implemented")
}
func (*UnimplementedVolumesServer) VolumesDelete(context.Context, *VolumesDeleteRequest) (*VolumesDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumesDelete not implemented")
}
func (*UnimplementedVolumesServer) VolumesInspect(context.Context, *VolumesInspectRequest) (*VolumesInspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumesInspect not implemented")
}

func RegisterVolumesServer(s *grpc.Server, srv VolumesServer) {
	s.RegisterService(&_Volumes_serviceDesc, srv)
}

func _Volumes_VolumesCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumesCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).VolumesCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.volumes.v1.Volumes/VolumesCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).VolumesCreate(ctx, req.(*VolumesCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_VolumesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).VolumesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.volumes.v1.Volumes/VolumesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).VolumesList(ctx, req.(*VolumesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_VolumesDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumesDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).VolumesDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.volumes.v1.Volumes/VolumesDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).VolumesDelete(ctx, req.(*VolumesDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_VolumesInspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumesInspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).VolumesInspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.volumes.v1.Volumes/VolumesInspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).VolumesInspect(ctx, req.(*VolumesInspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.docker.api.protos.volumes.v1.Volumes",
	HandlerType: (*VolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VolumesCreate",
			Handler:    _Volumes_VolumesCreate_Handler,
		},
		{
			MethodName: "VolumesList",
			Handler:    _Volumes_VolumesList_Handler,
		},
		{
			MethodName: "VolumesDelete",
			Handler:    _Volumes_VolumesDelete_Handler,
		},
		{
			MethodName: "VolumesInspect",
			Handler:    _Volumes_VolumesInspect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli/server/protos/volumes/v1/volumes.proto",
}
