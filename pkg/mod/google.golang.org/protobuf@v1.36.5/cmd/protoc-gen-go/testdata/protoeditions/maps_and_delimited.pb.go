// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/protoeditions/maps_and_delimited.proto

package protoeditions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type MessageWithMaps struct {
	state              protoimpl.MessageState                   `protogen:"open.v1"`
	MapWithoutMessage  map[string]string                        `protobuf:"bytes,1,rep,name=map_without_message,json=mapWithoutMessage" json:"map_without_message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapWithoutMessageB map[uint32][]byte                        `protobuf:"bytes,2,rep,name=map_without_message_b,json=mapWithoutMessageB" json:"map_without_message_b,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapWithMessage     map[int64]*MessageWithMaps_NestedMessage `protobuf:"bytes,3,rep,name=map_with_message,json=mapWithMessage" json:"map_with_message,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NestedMessage      *MessageWithMaps_NestedMessage           `protobuf:"group,4,opt,name=NestedMessage,json=nestedMessage" json:"nested_message,omitempty"`
	RepeatedMessage    []*MessageWithMaps_NestedMessage         `protobuf:"group,5,rep,name=NestedMessage,json=repeatedMessage" json:"repeated_message,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *MessageWithMaps) Reset() {
	*x = MessageWithMaps{}
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageWithMaps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithMaps) ProtoMessage() {}

func (x *MessageWithMaps) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithMaps.ProtoReflect.Descriptor instead.
func (*MessageWithMaps) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithMaps) GetMapWithoutMessage() map[string]string {
	if x != nil {
		return x.MapWithoutMessage
	}
	return nil
}

func (x *MessageWithMaps) GetMapWithoutMessageB() map[uint32][]byte {
	if x != nil {
		return x.MapWithoutMessageB
	}
	return nil
}

func (x *MessageWithMaps) GetMapWithMessage() map[int64]*MessageWithMaps_NestedMessage {
	if x != nil {
		return x.MapWithMessage
	}
	return nil
}

func (x *MessageWithMaps) GetNestedMessage() *MessageWithMaps_NestedMessage {
	if x != nil {
		return x.NestedMessage
	}
	return nil
}

func (x *MessageWithMaps) GetRepeatedMessage() []*MessageWithMaps_NestedMessage {
	if x != nil {
		return x.RepeatedMessage
	}
	return nil
}

type MessageWithMaps_NestedMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint64                `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageWithMaps_NestedMessage) Reset() {
	*x = MessageWithMaps_NestedMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageWithMaps_NestedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithMaps_NestedMessage) ProtoMessage() {}

func (x *MessageWithMaps_NestedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithMaps_NestedMessage.ProtoReflect.Descriptor instead.
func (*MessageWithMaps_NestedMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescGZIP(), []int{0, 3}
}

func (x *MessageWithMaps_NestedMessage) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *MessageWithMaps_NestedMessage) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDesc = string([]byte{
	0x0a, 0x41, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x5f,
	0x61, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xfc, 0x06, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x4d, 0x61, 0x70, 0x73, 0x12, 0x74, 0x0a, 0x13, 0x6d, 0x61, 0x70, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x6d, 0x61, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x78, 0x0a, 0x15, 0x6d,
	0x61, 0x70, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x12, 0x6d, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x12, 0x6b, 0x0a, 0x10, 0x6d, 0x61, 0x70, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x4d,
	0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0e, 0x6d, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x62, 0x0a, 0x0e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x66, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x73, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x44,
	0x0a, 0x16, 0x4d, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x45, 0x0a, 0x17, 0x4d, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x7e, 0x0a, 0x13, 0x4d,
	0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61,
	0x70, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x33, 0x0a, 0x0d, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x4a, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x92, 0x03, 0x02, 0x28, 0x02, 0x62, 0x08, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescData []byte
)

func file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDesc)))
	})
	return file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_goTypes = []any{
	(*MessageWithMaps)(nil),               // 0: goproto.protoc.protoeditions.MessageWithMaps
	nil,                                   // 1: goproto.protoc.protoeditions.MessageWithMaps.MapWithoutMessageEntry
	nil,                                   // 2: goproto.protoc.protoeditions.MessageWithMaps.MapWithoutMessageBEntry
	nil,                                   // 3: goproto.protoc.protoeditions.MessageWithMaps.MapWithMessageEntry
	(*MessageWithMaps_NestedMessage)(nil), // 4: goproto.protoc.protoeditions.MessageWithMaps.NestedMessage
}
var file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_depIdxs = []int32{
	1, // 0: goproto.protoc.protoeditions.MessageWithMaps.map_without_message:type_name -> goproto.protoc.protoeditions.MessageWithMaps.MapWithoutMessageEntry
	2, // 1: goproto.protoc.protoeditions.MessageWithMaps.map_without_message_b:type_name -> goproto.protoc.protoeditions.MessageWithMaps.MapWithoutMessageBEntry
	3, // 2: goproto.protoc.protoeditions.MessageWithMaps.map_with_message:type_name -> goproto.protoc.protoeditions.MessageWithMaps.MapWithMessageEntry
	4, // 3: goproto.protoc.protoeditions.MessageWithMaps.nested_message:type_name -> goproto.protoc.protoeditions.MessageWithMaps.NestedMessage
	4, // 4: goproto.protoc.protoeditions.MessageWithMaps.repeated_message:type_name -> goproto.protoc.protoeditions.MessageWithMaps.NestedMessage
	4, // 5: goproto.protoc.protoeditions.MessageWithMaps.MapWithMessageEntry.value:type_name -> goproto.protoc.protoeditions.MessageWithMaps.NestedMessage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_init() }
func file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_init() {
	if File_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto = out.File
	file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_protoeditions_maps_and_delimited_proto_depIdxs = nil
}
