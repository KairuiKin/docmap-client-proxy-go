// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: TagmenuResponseWriteFileLevel.proto

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

type WriteFileLevelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`          // 文件路径
	ResultCode *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`    // 写入结果（零：成功；其它：失败）
	ResultMsg  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=result_msg,json=resultMsg,proto3,oneof" json:"result_msg,omitempty"` // 失败信息
}

func (x *WriteFileLevelResult) Reset() {
	*x = WriteFileLevelResult{}
	mi := &file_TagmenuResponseWriteFileLevel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteFileLevelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileLevelResult) ProtoMessage() {}

func (x *WriteFileLevelResult) ProtoReflect() protoreflect.Message {
	mi := &file_TagmenuResponseWriteFileLevel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileLevelResult.ProtoReflect.Descriptor instead.
func (*WriteFileLevelResult) Descriptor() ([]byte, []int) {
	return file_TagmenuResponseWriteFileLevel_proto_rawDescGZIP(), []int{0}
}

func (x *WriteFileLevelResult) GetFilePath() *wrapperspb.StringValue {
	if x != nil {
		return x.FilePath
	}
	return nil
}

func (x *WriteFileLevelResult) GetResultCode() *wrapperspb.Int32Value {
	if x != nil {
		return x.ResultCode
	}
	return nil
}

func (x *WriteFileLevelResult) GetResultMsg() *wrapperspb.StringValue {
	if x != nil {
		return x.ResultMsg
	}
	return nil
}

// 写指定文件的密级的回应
type TagmenuResponseWriteFileLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*WriteFileLevelResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"` // 写密级的结果清单
}

func (x *TagmenuResponseWriteFileLevel) Reset() {
	*x = TagmenuResponseWriteFileLevel{}
	mi := &file_TagmenuResponseWriteFileLevel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagmenuResponseWriteFileLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagmenuResponseWriteFileLevel) ProtoMessage() {}

func (x *TagmenuResponseWriteFileLevel) ProtoReflect() protoreflect.Message {
	mi := &file_TagmenuResponseWriteFileLevel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagmenuResponseWriteFileLevel.ProtoReflect.Descriptor instead.
func (*TagmenuResponseWriteFileLevel) Descriptor() ([]byte, []int) {
	return file_TagmenuResponseWriteFileLevel_proto_rawDescGZIP(), []int{1}
}

func (x *TagmenuResponseWriteFileLevel) GetResults() []*WriteFileLevelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_TagmenuResponseWriteFileLevel_proto protoreflect.FileDescriptor

var file_TagmenuResponseWriteFileLevel_proto_rawDesc = []byte{
	0x0a, 0x23, 0x54, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x69, 0x70, 0x63, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x14, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3c,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x64, 0x0a,
	0x1d, 0x54, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x43,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TagmenuResponseWriteFileLevel_proto_rawDescOnce sync.Once
	file_TagmenuResponseWriteFileLevel_proto_rawDescData = file_TagmenuResponseWriteFileLevel_proto_rawDesc
)

func file_TagmenuResponseWriteFileLevel_proto_rawDescGZIP() []byte {
	file_TagmenuResponseWriteFileLevel_proto_rawDescOnce.Do(func() {
		file_TagmenuResponseWriteFileLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_TagmenuResponseWriteFileLevel_proto_rawDescData)
	})
	return file_TagmenuResponseWriteFileLevel_proto_rawDescData
}

var file_TagmenuResponseWriteFileLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TagmenuResponseWriteFileLevel_proto_goTypes = []any{
	(*WriteFileLevelResult)(nil),          // 0: docmap.protobuf.ipc.WriteFileLevelResult
	(*TagmenuResponseWriteFileLevel)(nil), // 1: docmap.protobuf.ipc.TagmenuResponseWriteFileLevel
	(*wrapperspb.StringValue)(nil),        // 2: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),         // 3: google.protobuf.Int32Value
}
var file_TagmenuResponseWriteFileLevel_proto_depIdxs = []int32{
	2, // 0: docmap.protobuf.ipc.WriteFileLevelResult.file_path:type_name -> google.protobuf.StringValue
	3, // 1: docmap.protobuf.ipc.WriteFileLevelResult.result_code:type_name -> google.protobuf.Int32Value
	2, // 2: docmap.protobuf.ipc.WriteFileLevelResult.result_msg:type_name -> google.protobuf.StringValue
	0, // 3: docmap.protobuf.ipc.TagmenuResponseWriteFileLevel.results:type_name -> docmap.protobuf.ipc.WriteFileLevelResult
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_TagmenuResponseWriteFileLevel_proto_init() }
func file_TagmenuResponseWriteFileLevel_proto_init() {
	if File_TagmenuResponseWriteFileLevel_proto != nil {
		return
	}
	file_TagmenuResponseWriteFileLevel_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TagmenuResponseWriteFileLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TagmenuResponseWriteFileLevel_proto_goTypes,
		DependencyIndexes: file_TagmenuResponseWriteFileLevel_proto_depIdxs,
		MessageInfos:      file_TagmenuResponseWriteFileLevel_proto_msgTypes,
	}.Build()
	File_TagmenuResponseWriteFileLevel_proto = out.File
	file_TagmenuResponseWriteFileLevel_proto_rawDesc = nil
	file_TagmenuResponseWriteFileLevel_proto_goTypes = nil
	file_TagmenuResponseWriteFileLevel_proto_depIdxs = nil
}
