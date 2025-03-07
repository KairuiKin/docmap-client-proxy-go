// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: encrypt_method.proto

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

type EncryptMethod int32

const (
	EncryptMethod_AUTO      EncryptMethod = 0 // 自动判断
	EncryptMethod_ENCRYPT   EncryptMethod = 1 // 加密
	EncryptMethod_UNENCRYPT EncryptMethod = 2 // 不加密
)

// Enum value maps for EncryptMethod.
var (
	EncryptMethod_name = map[int32]string{
		0: "AUTO",
		1: "ENCRYPT",
		2: "UNENCRYPT",
	}
	EncryptMethod_value = map[string]int32{
		"AUTO":      0,
		"ENCRYPT":   1,
		"UNENCRYPT": 2,
	}
)

func (x EncryptMethod) Enum() *EncryptMethod {
	p := new(EncryptMethod)
	*p = x
	return p
}

func (x EncryptMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_encrypt_method_proto_enumTypes[0].Descriptor()
}

func (EncryptMethod) Type() protoreflect.EnumType {
	return &file_encrypt_method_proto_enumTypes[0]
}

func (x EncryptMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptMethod.Descriptor instead.
func (EncryptMethod) EnumDescriptor() ([]byte, []int) {
	return file_encrypt_method_proto_rawDescGZIP(), []int{0}
}

var File_encrypt_method_proto protoreflect.FileDescriptor

var file_encrypt_method_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2a, 0x35, 0x0a,
	0x0d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x43, 0x52,
	0x59, 0x50, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x45, 0x4e, 0x43, 0x52, 0x59,
	0x50, 0x54, 0x10, 0x02, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70,
	0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encrypt_method_proto_rawDescOnce sync.Once
	file_encrypt_method_proto_rawDescData = file_encrypt_method_proto_rawDesc
)

func file_encrypt_method_proto_rawDescGZIP() []byte {
	file_encrypt_method_proto_rawDescOnce.Do(func() {
		file_encrypt_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_encrypt_method_proto_rawDescData)
	})
	return file_encrypt_method_proto_rawDescData
}

var file_encrypt_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_encrypt_method_proto_goTypes = []any{
	(EncryptMethod)(0), // 0: docmap.exchange.model.EncryptMethod
}
var file_encrypt_method_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_encrypt_method_proto_init() }
func file_encrypt_method_proto_init() {
	if File_encrypt_method_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_encrypt_method_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encrypt_method_proto_goTypes,
		DependencyIndexes: file_encrypt_method_proto_depIdxs,
		EnumInfos:         file_encrypt_method_proto_enumTypes,
	}.Build()
	File_encrypt_method_proto = out.File
	file_encrypt_method_proto_rawDesc = nil
	file_encrypt_method_proto_goTypes = nil
	file_encrypt_method_proto_depIdxs = nil
}
