// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: TagmenuRequestWriteFileLevel.proto

package ipc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 文件密级
type FileLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`    // 文件路径
	LevelCode *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=level_code,json=levelCode,proto3" json:"level_code,omitempty"` // 密级代码
}

func (x *FileLevel) Reset() {
	*x = FileLevel{}
	mi := &file_TagmenuRequestWriteFileLevel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLevel) ProtoMessage() {}

func (x *FileLevel) ProtoReflect() protoreflect.Message {
	mi := &file_TagmenuRequestWriteFileLevel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileLevel.ProtoReflect.Descriptor instead.
func (*FileLevel) Descriptor() ([]byte, []int) {
	return file_TagmenuRequestWriteFileLevel_proto_rawDescGZIP(), []int{0}
}

func (x *FileLevel) GetFilePath() *wrapperspb.StringValue {
	if x != nil {
		return x.FilePath
	}
	return nil
}

func (x *FileLevel) GetLevelCode() *wrapperspb.UInt32Value {
	if x != nil {
		return x.LevelCode
	}
	return nil
}

// 写指定文件的密级
type TagmenuRequestWriteFileLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileLevels []*FileLevel `protobuf:"bytes,1,rep,name=file_levels,json=fileLevels,proto3" json:"file_levels,omitempty"` // 要写入的文件密级
}

func (x *TagmenuRequestWriteFileLevel) Reset() {
	*x = TagmenuRequestWriteFileLevel{}
	mi := &file_TagmenuRequestWriteFileLevel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagmenuRequestWriteFileLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagmenuRequestWriteFileLevel) ProtoMessage() {}

func (x *TagmenuRequestWriteFileLevel) ProtoReflect() protoreflect.Message {
	mi := &file_TagmenuRequestWriteFileLevel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagmenuRequestWriteFileLevel.ProtoReflect.Descriptor instead.
func (*TagmenuRequestWriteFileLevel) Descriptor() ([]byte, []int) {
	return file_TagmenuRequestWriteFileLevel_proto_rawDescGZIP(), []int{1}
}

func (x *TagmenuRequestWriteFileLevel) GetFileLevels() []*FileLevel {
	if x != nil {
		return x.FileLevels
	}
	return nil
}

var File_TagmenuRequestWriteFileLevel_proto protoreflect.FileDescriptor

var file_TagmenuRequestWriteFileLevel_proto_rawDesc = []byte{
	0x0a, 0x22, 0x54, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x69, 0x70, 0x63, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x3b, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x5f, 0x0a, 0x1c, 0x54, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x3f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73,
	0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TagmenuRequestWriteFileLevel_proto_rawDescOnce sync.Once
	file_TagmenuRequestWriteFileLevel_proto_rawDescData = file_TagmenuRequestWriteFileLevel_proto_rawDesc
)

func file_TagmenuRequestWriteFileLevel_proto_rawDescGZIP() []byte {
	file_TagmenuRequestWriteFileLevel_proto_rawDescOnce.Do(func() {
		file_TagmenuRequestWriteFileLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_TagmenuRequestWriteFileLevel_proto_rawDescData)
	})
	return file_TagmenuRequestWriteFileLevel_proto_rawDescData
}

var file_TagmenuRequestWriteFileLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TagmenuRequestWriteFileLevel_proto_goTypes = []any{
	(*FileLevel)(nil),                    // 0: docmap.protobuf.ipc.FileLevel
	(*TagmenuRequestWriteFileLevel)(nil), // 1: docmap.protobuf.ipc.TagmenuRequestWriteFileLevel
	(*wrapperspb.StringValue)(nil),       // 2: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil),       // 3: google.protobuf.UInt32Value
}
var file_TagmenuRequestWriteFileLevel_proto_depIdxs = []int32{
	2, // 0: docmap.protobuf.ipc.FileLevel.file_path:type_name -> google.protobuf.StringValue
	3, // 1: docmap.protobuf.ipc.FileLevel.level_code:type_name -> google.protobuf.UInt32Value
	0, // 2: docmap.protobuf.ipc.TagmenuRequestWriteFileLevel.file_levels:type_name -> docmap.protobuf.ipc.FileLevel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_TagmenuRequestWriteFileLevel_proto_init() }
func file_TagmenuRequestWriteFileLevel_proto_init() {
	if File_TagmenuRequestWriteFileLevel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TagmenuRequestWriteFileLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TagmenuRequestWriteFileLevel_proto_goTypes,
		DependencyIndexes: file_TagmenuRequestWriteFileLevel_proto_depIdxs,
		MessageInfos:      file_TagmenuRequestWriteFileLevel_proto_msgTypes,
	}.Build()
	File_TagmenuRequestWriteFileLevel_proto = out.File
	file_TagmenuRequestWriteFileLevel_proto_rawDesc = nil
	file_TagmenuRequestWriteFileLevel_proto_goTypes = nil
	file_TagmenuRequestWriteFileLevel_proto_depIdxs = nil
}
