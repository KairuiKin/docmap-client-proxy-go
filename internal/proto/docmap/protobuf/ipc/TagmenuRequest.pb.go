// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: TagmenuRequest.proto

package ipc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TagmenuRequest_RequestType int32

const (
	TagmenuRequest_GetLevelAndCategory      TagmenuRequest_RequestType = 0 // 获取文件级别和类别范围的请求(发送给Proxy)
	TagmenuRequest_ReadFileLevelAndCategory TagmenuRequest_RequestType = 1 // 读取指定文件的级别和类别
	TagmenuRequest_WriteFileLevel           TagmenuRequest_RequestType = 2 // 写文件密级
	TagmenuRequest_WriteFileCategory        TagmenuRequest_RequestType = 3 // 写文件类别
	TagmenuRequest_RemoveFolderUserDatas    TagmenuRequest_RequestType = 4 // 移除目录标记
	TagmenuRequest_ApplyForOutdoor          TagmenuRequest_RequestType = 5 // 申请出边界
	TagmenuRequest_AllowOutdoor             TagmenuRequest_RequestType = 6 // 允许或禁止文件出边界
)

// Enum value maps for TagmenuRequest_RequestType.
var (
	TagmenuRequest_RequestType_name = map[int32]string{
		0: "GetLevelAndCategory",
		1: "ReadFileLevelAndCategory",
		2: "WriteFileLevel",
		3: "WriteFileCategory",
		4: "RemoveFolderUserDatas",
		5: "ApplyForOutdoor",
		6: "AllowOutdoor",
	}
	TagmenuRequest_RequestType_value = map[string]int32{
		"GetLevelAndCategory":      0,
		"ReadFileLevelAndCategory": 1,
		"WriteFileLevel":           2,
		"WriteFileCategory":        3,
		"RemoveFolderUserDatas":    4,
		"ApplyForOutdoor":          5,
		"AllowOutdoor":             6,
	}
)

func (x TagmenuRequest_RequestType) Enum() *TagmenuRequest_RequestType {
	p := new(TagmenuRequest_RequestType)
	*p = x
	return p
}

func (x TagmenuRequest_RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TagmenuRequest_RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_TagmenuRequest_proto_enumTypes[0].Descriptor()
}

func (TagmenuRequest_RequestType) Type() protoreflect.EnumType {
	return &file_TagmenuRequest_proto_enumTypes[0]
}

func (x TagmenuRequest_RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TagmenuRequest_RequestType.Descriptor instead.
func (TagmenuRequest_RequestType) EnumDescriptor() ([]byte, []int) {
	return file_TagmenuRequest_proto_rawDescGZIP(), []int{0, 0}
}

// TagMenu发出的请求
type TagmenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    TagmenuRequest_RequestType `protobuf:"varint,1,opt,name=type,proto3,enum=docmap.protobuf.ipc.TagmenuRequest_RequestType" json:"type,omitempty"` // 请求类型
	Payload *anypb.Any                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                                // 请求载荷
}

func (x *TagmenuRequest) Reset() {
	*x = TagmenuRequest{}
	mi := &file_TagmenuRequest_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagmenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagmenuRequest) ProtoMessage() {}

func (x *TagmenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_TagmenuRequest_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagmenuRequest.ProtoReflect.Descriptor instead.
func (*TagmenuRequest) Descriptor() ([]byte, []int) {
	return file_TagmenuRequest_proto_rawDescGZIP(), []int{0}
}

func (x *TagmenuRequest) GetType() TagmenuRequest_RequestType {
	if x != nil {
		return x.Type
	}
	return TagmenuRequest_GetLevelAndCategory
}

func (x *TagmenuRequest) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_TagmenuRequest_proto protoreflect.FileDescriptor

var file_TagmenuRequest_proto_rawDesc = []byte{
	0x0a, 0x14, 0x54, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x69, 0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x0e, 0x54, 0x61, 0x67, 0x6d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x54, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb1,
	0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x41, 0x6e, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x41, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x10, 0x03,
	0x12, 0x19, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x73, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x46, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x64, 0x6f, 0x6f, 0x72, 0x10, 0x05,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x64, 0x6f, 0x6f, 0x72,
	0x10, 0x06, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x6d, 0x61, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TagmenuRequest_proto_rawDescOnce sync.Once
	file_TagmenuRequest_proto_rawDescData = file_TagmenuRequest_proto_rawDesc
)

func file_TagmenuRequest_proto_rawDescGZIP() []byte {
	file_TagmenuRequest_proto_rawDescOnce.Do(func() {
		file_TagmenuRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_TagmenuRequest_proto_rawDescData)
	})
	return file_TagmenuRequest_proto_rawDescData
}

var file_TagmenuRequest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TagmenuRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TagmenuRequest_proto_goTypes = []any{
	(TagmenuRequest_RequestType)(0), // 0: docmap.protobuf.ipc.TagmenuRequest.RequestType
	(*TagmenuRequest)(nil),          // 1: docmap.protobuf.ipc.TagmenuRequest
	(*anypb.Any)(nil),               // 2: google.protobuf.Any
}
var file_TagmenuRequest_proto_depIdxs = []int32{
	0, // 0: docmap.protobuf.ipc.TagmenuRequest.type:type_name -> docmap.protobuf.ipc.TagmenuRequest.RequestType
	2, // 1: docmap.protobuf.ipc.TagmenuRequest.payload:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TagmenuRequest_proto_init() }
func file_TagmenuRequest_proto_init() {
	if File_TagmenuRequest_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TagmenuRequest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TagmenuRequest_proto_goTypes,
		DependencyIndexes: file_TagmenuRequest_proto_depIdxs,
		EnumInfos:         file_TagmenuRequest_proto_enumTypes,
		MessageInfos:      file_TagmenuRequest_proto_msgTypes,
	}.Build()
	File_TagmenuRequest_proto = out.File
	file_TagmenuRequest_proto_rawDesc = nil
	file_TagmenuRequest_proto_goTypes = nil
	file_TagmenuRequest_proto_depIdxs = nil
}
