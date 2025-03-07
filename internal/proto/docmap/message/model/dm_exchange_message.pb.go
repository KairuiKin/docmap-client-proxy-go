// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: dm_exchange_message.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DMExchangeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid     string     `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Type     uint32     `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Response bool       `protobuf:"varint,4,opt,name=response,proto3" json:"response,omitempty"`
	Payload  *anypb.Any `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Devid    *string    `protobuf:"bytes,6,opt,name=devid,proto3,oneof" json:"devid,omitempty"`
}

func (x *DMExchangeMessage) Reset() {
	*x = DMExchangeMessage{}
	mi := &file_dm_exchange_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DMExchangeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DMExchangeMessage) ProtoMessage() {}

func (x *DMExchangeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_dm_exchange_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DMExchangeMessage.ProtoReflect.Descriptor instead.
func (*DMExchangeMessage) Descriptor() ([]byte, []int) {
	return file_dm_exchange_message_proto_rawDescGZIP(), []int{0}
}

func (x *DMExchangeMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DMExchangeMessage) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DMExchangeMessage) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *DMExchangeMessage) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

func (x *DMExchangeMessage) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *DMExchangeMessage) GetDevid() string {
	if x != nil && x.Devid != nil {
		return *x.Devid
	}
	return ""
}

var File_dm_exchange_message_proto protoreflect.FileDescriptor

var file_dm_exchange_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x6d, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x6f, 0x63,
	0x6d, 0x61, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a,
	0x11, 0x44, 0x4d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x64, 0x65, 0x76, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x64, 0x65, 0x76, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x2e,
	0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dm_exchange_message_proto_rawDescOnce sync.Once
	file_dm_exchange_message_proto_rawDescData = file_dm_exchange_message_proto_rawDesc
)

func file_dm_exchange_message_proto_rawDescGZIP() []byte {
	file_dm_exchange_message_proto_rawDescOnce.Do(func() {
		file_dm_exchange_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_dm_exchange_message_proto_rawDescData)
	})
	return file_dm_exchange_message_proto_rawDescData
}

var file_dm_exchange_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dm_exchange_message_proto_goTypes = []any{
	(*DMExchangeMessage)(nil), // 0: docmap.message.model.DMExchangeMessage
	(*anypb.Any)(nil),         // 1: google.protobuf.Any
}
var file_dm_exchange_message_proto_depIdxs = []int32{
	1, // 0: docmap.message.model.DMExchangeMessage.payload:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dm_exchange_message_proto_init() }
func file_dm_exchange_message_proto_init() {
	if File_dm_exchange_message_proto != nil {
		return
	}
	file_dm_exchange_message_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dm_exchange_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dm_exchange_message_proto_goTypes,
		DependencyIndexes: file_dm_exchange_message_proto_depIdxs,
		MessageInfos:      file_dm_exchange_message_proto_msgTypes,
	}.Build()
	File_dm_exchange_message_proto = out.File
	file_dm_exchange_message_proto_rawDesc = nil
	file_dm_exchange_message_proto_goTypes = nil
	file_dm_exchange_message_proto_depIdxs = nil
}
