// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: whoami.proto

package model

import (
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

type WhoAmI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientInstanceId string `protobuf:"bytes,1,opt,name=clientInstanceId,proto3" json:"clientInstanceId,omitempty"`
	App              string `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	ProcessId        string `protobuf:"bytes,3,opt,name=processId,proto3" json:"processId,omitempty"`
}

func (x *WhoAmI) Reset() {
	*x = WhoAmI{}
	mi := &file_whoami_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhoAmI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmI) ProtoMessage() {}

func (x *WhoAmI) ProtoReflect() protoreflect.Message {
	mi := &file_whoami_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmI.ProtoReflect.Descriptor instead.
func (*WhoAmI) Descriptor() ([]byte, []int) {
	return file_whoami_proto_rawDescGZIP(), []int{0}
}

func (x *WhoAmI) GetClientInstanceId() string {
	if x != nil {
		return x.ClientInstanceId
	}
	return ""
}

func (x *WhoAmI) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *WhoAmI) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

type WhoAmIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoAmIResponse) Reset() {
	*x = WhoAmIResponse{}
	mi := &file_whoami_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIResponse) ProtoMessage() {}

func (x *WhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whoami_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIResponse.ProtoReflect.Descriptor instead.
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_whoami_proto_rawDescGZIP(), []int{1}
}

var File_whoami_proto protoreflect.FileDescriptor

var file_whoami_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x64, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12,
	0x2a, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x57,
	0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a,
	0x17, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_whoami_proto_rawDescOnce sync.Once
	file_whoami_proto_rawDescData = file_whoami_proto_rawDesc
)

func file_whoami_proto_rawDescGZIP() []byte {
	file_whoami_proto_rawDescOnce.Do(func() {
		file_whoami_proto_rawDescData = protoimpl.X.CompressGZIP(file_whoami_proto_rawDescData)
	})
	return file_whoami_proto_rawDescData
}

var file_whoami_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_whoami_proto_goTypes = []any{
	(*WhoAmI)(nil),         // 0: docmap.exchange.model.WhoAmI
	(*WhoAmIResponse)(nil), // 1: docmap.exchange.model.WhoAmIResponse
}
var file_whoami_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_whoami_proto_init() }
func file_whoami_proto_init() {
	if File_whoami_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_whoami_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_whoami_proto_goTypes,
		DependencyIndexes: file_whoami_proto_depIdxs,
		MessageInfos:      file_whoami_proto_msgTypes,
	}.Build()
	File_whoami_proto = out.File
	file_whoami_proto_rawDesc = nil
	file_whoami_proto_goTypes = nil
	file_whoami_proto_depIdxs = nil
}
