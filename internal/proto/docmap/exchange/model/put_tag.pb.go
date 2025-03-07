// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: put_tag.proto

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

type PutTagResponse_PutTagResult int32

const (
	PutTagResponse_SUCCESS         PutTagResponse_PutTagResult = 0 // 成功
	PutTagResponse_EXPIRED         PutTagResponse_PutTagResult = 1 // 写入的TAG过期了
	PutTagResponse_TIMEOUT         PutTagResponse_PutTagResult = 2 // 写入超时
	PutTagResponse_UNPACK_FAILED   PutTagResponse_PutTagResult = 3 // 解包载荷失败
	PutTagResponse_READ_TAG_FAILED PutTagResponse_PutTagResult = 4 // 读原始TAG失败
	PutTagResponse_ENCRYPT_FAILED  PutTagResponse_PutTagResult = 5 // 加密失败
)

// Enum value maps for PutTagResponse_PutTagResult.
var (
	PutTagResponse_PutTagResult_name = map[int32]string{
		0: "SUCCESS",
		1: "EXPIRED",
		2: "TIMEOUT",
		3: "UNPACK_FAILED",
		4: "READ_TAG_FAILED",
		5: "ENCRYPT_FAILED",
	}
	PutTagResponse_PutTagResult_value = map[string]int32{
		"SUCCESS":         0,
		"EXPIRED":         1,
		"TIMEOUT":         2,
		"UNPACK_FAILED":   3,
		"READ_TAG_FAILED": 4,
		"ENCRYPT_FAILED":  5,
	}
)

func (x PutTagResponse_PutTagResult) Enum() *PutTagResponse_PutTagResult {
	p := new(PutTagResponse_PutTagResult)
	*p = x
	return p
}

func (x PutTagResponse_PutTagResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PutTagResponse_PutTagResult) Descriptor() protoreflect.EnumDescriptor {
	return file_put_tag_proto_enumTypes[0].Descriptor()
}

func (PutTagResponse_PutTagResult) Type() protoreflect.EnumType {
	return &file_put_tag_proto_enumTypes[0]
}

func (x PutTagResponse_PutTagResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PutTagResponse_PutTagResult.Descriptor instead.
func (PutTagResponse_PutTagResult) EnumDescriptor() ([]byte, []int) {
	return file_put_tag_proto_rawDescGZIP(), []int{1, 0}
}

type PutTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Tags     []*Tag `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Encrypt  bool   `protobuf:"varint,3,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
}

func (x *PutTag) Reset() {
	*x = PutTag{}
	mi := &file_put_tag_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutTag) ProtoMessage() {}

func (x *PutTag) ProtoReflect() protoreflect.Message {
	mi := &file_put_tag_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutTag.ProtoReflect.Descriptor instead.
func (*PutTag) Descriptor() ([]byte, []int) {
	return file_put_tag_proto_rawDescGZIP(), []int{0}
}

func (x *PutTag) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *PutTag) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PutTag) GetEncrypt() bool {
	if x != nil {
		return x.Encrypt
	}
	return false
}

type PutTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string                      `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Tags     []*Tag                      `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Result   PutTagResponse_PutTagResult `protobuf:"varint,3,opt,name=result,proto3,enum=docmap.exchange.model.PutTagResponse_PutTagResult" json:"result,omitempty"`
}

func (x *PutTagResponse) Reset() {
	*x = PutTagResponse{}
	mi := &file_put_tag_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutTagResponse) ProtoMessage() {}

func (x *PutTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_put_tag_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutTagResponse.ProtoReflect.Descriptor instead.
func (*PutTagResponse) Descriptor() ([]byte, []int) {
	return file_put_tag_proto_rawDescGZIP(), []int{1}
}

func (x *PutTagResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *PutTagResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PutTagResponse) GetResult() PutTagResponse_PutTagResult {
	if x != nil {
		return x.Result
	}
	return PutTagResponse_SUCCESS
}

var File_put_tag_proto protoreflect.FileDescriptor

var file_put_tag_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6e, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61,
	0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x22, 0x9b, 0x02, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x4a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x74, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x71, 0x0a, 0x0c,
	0x50, 0x75, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x50, 0x41, 0x43, 0x4b, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x54,
	0x41, 0x47, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x42,
	0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_put_tag_proto_rawDescOnce sync.Once
	file_put_tag_proto_rawDescData = file_put_tag_proto_rawDesc
)

func file_put_tag_proto_rawDescGZIP() []byte {
	file_put_tag_proto_rawDescOnce.Do(func() {
		file_put_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_put_tag_proto_rawDescData)
	})
	return file_put_tag_proto_rawDescData
}

var file_put_tag_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_put_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_put_tag_proto_goTypes = []any{
	(PutTagResponse_PutTagResult)(0), // 0: docmap.exchange.model.PutTagResponse.PutTagResult
	(*PutTag)(nil),                   // 1: docmap.exchange.model.PutTag
	(*PutTagResponse)(nil),           // 2: docmap.exchange.model.PutTagResponse
	(*Tag)(nil),                      // 3: docmap.exchange.model.Tag
}
var file_put_tag_proto_depIdxs = []int32{
	3, // 0: docmap.exchange.model.PutTag.tags:type_name -> docmap.exchange.model.Tag
	3, // 1: docmap.exchange.model.PutTagResponse.tags:type_name -> docmap.exchange.model.Tag
	0, // 2: docmap.exchange.model.PutTagResponse.result:type_name -> docmap.exchange.model.PutTagResponse.PutTagResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_put_tag_proto_init() }
func file_put_tag_proto_init() {
	if File_put_tag_proto != nil {
		return
	}
	file_tag_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_put_tag_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_put_tag_proto_goTypes,
		DependencyIndexes: file_put_tag_proto_depIdxs,
		EnumInfos:         file_put_tag_proto_enumTypes,
		MessageInfos:      file_put_tag_proto_msgTypes,
	}.Build()
	File_put_tag_proto = out.File
	file_put_tag_proto_rawDesc = nil
	file_put_tag_proto_goTypes = nil
	file_put_tag_proto_depIdxs = nil
}
