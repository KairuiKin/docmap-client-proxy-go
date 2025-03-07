// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: decrypt.proto

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

type DecryptResponse_DecryptResult int32

const (
	DecryptResponse_SUCCESS             DecryptResponse_DecryptResult = 0 // 解密成功
	DecryptResponse_NOTENCRYPTED        DecryptResponse_DecryptResult = 1 // 源文件未加密
	DecryptResponse_SOURCEFILE_NOTFOUND DecryptResponse_DecryptResult = 2 // 未找到源文件
	DecryptResponse_UNPACK_FAILED       DecryptResponse_DecryptResult = 3 // 解包载荷失败
	DecryptResponse_FAILED              DecryptResponse_DecryptResult = 4 // 解密失败
)

// Enum value maps for DecryptResponse_DecryptResult.
var (
	DecryptResponse_DecryptResult_name = map[int32]string{
		0: "SUCCESS",
		1: "NOTENCRYPTED",
		2: "SOURCEFILE_NOTFOUND",
		3: "UNPACK_FAILED",
		4: "FAILED",
	}
	DecryptResponse_DecryptResult_value = map[string]int32{
		"SUCCESS":             0,
		"NOTENCRYPTED":        1,
		"SOURCEFILE_NOTFOUND": 2,
		"UNPACK_FAILED":       3,
		"FAILED":              4,
	}
)

func (x DecryptResponse_DecryptResult) Enum() *DecryptResponse_DecryptResult {
	p := new(DecryptResponse_DecryptResult)
	*p = x
	return p
}

func (x DecryptResponse_DecryptResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DecryptResponse_DecryptResult) Descriptor() protoreflect.EnumDescriptor {
	return file_decrypt_proto_enumTypes[0].Descriptor()
}

func (DecryptResponse_DecryptResult) Type() protoreflect.EnumType {
	return &file_decrypt_proto_enumTypes[0]
}

func (x DecryptResponse_DecryptResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DecryptResponse_DecryptResult.Descriptor instead.
func (DecryptResponse_DecryptResult) EnumDescriptor() ([]byte, []int) {
	return file_decrypt_proto_rawDescGZIP(), []int{1, 0}
}

type Decrypt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceFilePath      string  `protobuf:"bytes,1,opt,name=sourceFilePath,proto3" json:"sourceFilePath,omitempty"`                 // 源文件路径
	DestinationFilePath *string `protobuf:"bytes,2,opt,name=destinationFilePath,proto3,oneof" json:"destinationFilePath,omitempty"` // 目标文件路径（若不给出，则不解密，仅测试源文件是否被加密了）
	DecryptRemove       *bool   `protobuf:"varint,3,opt,name=decryptRemove,proto3,oneof" json:"decryptRemove,omitempty"`            // 解密成功后自动移除来源文件（测试加密状态时忽略此字段，且仅在源文件和目标文件路径不是同一个时生效）
}

func (x *Decrypt) Reset() {
	*x = Decrypt{}
	mi := &file_decrypt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Decrypt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decrypt) ProtoMessage() {}

func (x *Decrypt) ProtoReflect() protoreflect.Message {
	mi := &file_decrypt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decrypt.ProtoReflect.Descriptor instead.
func (*Decrypt) Descriptor() ([]byte, []int) {
	return file_decrypt_proto_rawDescGZIP(), []int{0}
}

func (x *Decrypt) GetSourceFilePath() string {
	if x != nil {
		return x.SourceFilePath
	}
	return ""
}

func (x *Decrypt) GetDestinationFilePath() string {
	if x != nil && x.DestinationFilePath != nil {
		return *x.DestinationFilePath
	}
	return ""
}

func (x *Decrypt) GetDecryptRemove() bool {
	if x != nil && x.DecryptRemove != nil {
		return *x.DecryptRemove
	}
	return false
}

type DecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceFilePath      string                        `protobuf:"bytes,1,opt,name=sourceFilePath,proto3" json:"sourceFilePath,omitempty"`                 // 源文件路径
	DestinationFilePath *string                       `protobuf:"bytes,2,opt,name=destinationFilePath,proto3,oneof" json:"destinationFilePath,omitempty"` // 目标文件路径
	Result              DecryptResponse_DecryptResult `protobuf:"varint,3,opt,name=result,proto3,enum=docmap.exchange.model.DecryptResponse_DecryptResult" json:"result,omitempty"`
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	mi := &file_decrypt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_decrypt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_decrypt_proto_rawDescGZIP(), []int{1}
}

func (x *DecryptResponse) GetSourceFilePath() string {
	if x != nil {
		return x.SourceFilePath
	}
	return ""
}

func (x *DecryptResponse) GetDestinationFilePath() string {
	if x != nil && x.DestinationFilePath != nil {
		return *x.DestinationFilePath
	}
	return ""
}

func (x *DecryptResponse) GetResult() DecryptResponse_DecryptResult {
	if x != nil {
		return x.Result
	}
	return DecryptResponse_SUCCESS
}

var File_decrypt_proto protoreflect.FileDescriptor

var file_decrypt_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xbd, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x13, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x0d, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x35, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x64, 0x6f, 0x63, 0x6d,
	0x61, 0x70, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x66, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x45, 0x4e, 0x43, 0x52,
	0x59, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x50, 0x41, 0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42,
	0x16, 0x0a, 0x14, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x64, 0x6f, 0x63,
	0x6d, 0x61, 0x70, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_decrypt_proto_rawDescOnce sync.Once
	file_decrypt_proto_rawDescData = file_decrypt_proto_rawDesc
)

func file_decrypt_proto_rawDescGZIP() []byte {
	file_decrypt_proto_rawDescOnce.Do(func() {
		file_decrypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_decrypt_proto_rawDescData)
	})
	return file_decrypt_proto_rawDescData
}

var file_decrypt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_decrypt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_decrypt_proto_goTypes = []any{
	(DecryptResponse_DecryptResult)(0), // 0: docmap.exchange.model.DecryptResponse.DecryptResult
	(*Decrypt)(nil),                    // 1: docmap.exchange.model.Decrypt
	(*DecryptResponse)(nil),            // 2: docmap.exchange.model.DecryptResponse
}
var file_decrypt_proto_depIdxs = []int32{
	0, // 0: docmap.exchange.model.DecryptResponse.result:type_name -> docmap.exchange.model.DecryptResponse.DecryptResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_decrypt_proto_init() }
func file_decrypt_proto_init() {
	if File_decrypt_proto != nil {
		return
	}
	file_decrypt_proto_msgTypes[0].OneofWrappers = []any{}
	file_decrypt_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_decrypt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_decrypt_proto_goTypes,
		DependencyIndexes: file_decrypt_proto_depIdxs,
		EnumInfos:         file_decrypt_proto_enumTypes,
		MessageInfos:      file_decrypt_proto_msgTypes,
	}.Build()
	File_decrypt_proto = out.File
	file_decrypt_proto_rawDesc = nil
	file_decrypt_proto_goTypes = nil
	file_decrypt_proto_depIdxs = nil
}
