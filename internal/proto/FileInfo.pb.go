// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: FileInfo.proto

package __

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

// 文件属性
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityIdentity     *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=entityIdentity,proto3" json:"entityIdentity,omitempty"`         // 文件识别ID，1个文件只有唯一值
	ContentIdentities  []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=contentIdentities,proto3" json:"contentIdentities,omitempty"`   // 用于跟踪内容，一个文件可能有多值
	Path               *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`                             // 带路径的完整文件名
	Filename           *wrapperspb.StringValue   `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`                     // 文件名
	DriveType          *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=driveType,proto3" json:"driveType,omitempty"`                   // 驱动器类型
	SecurityLevel      *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=securityLevel,proto3" json:"securityLevel,omitempty"`           // 安全级别
	ClassificationType []*wrapperspb.StringValue `protobuf:"bytes,7,rep,name=classificationType,proto3" json:"classificationType,omitempty"` // 分类
	CreateTime         *wrapperspb.UInt64Value   `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`                 // 文件创建时间，Unix时间戳，毫秒
	ModifyTime         *wrapperspb.UInt64Value   `protobuf:"bytes,9,opt,name=modifyTime,proto3" json:"modifyTime,omitempty"`                 // 文件修改时间，Unix时间戳，毫秒
	Sha1               *wrapperspb.StringValue   `protobuf:"bytes,10,opt,name=sha1,proto3" json:"sha1,omitempty"`                            // 文件 sha1 值
	Size               *wrapperspb.UInt64Value   `protobuf:"bytes,11,opt,name=size,proto3" json:"size,omitempty"`                            // 文件大小
	Ext                *wrapperspb.StringValue   `protobuf:"bytes,12,opt,name=ext,proto3" json:"ext,omitempty"`                              // 文件扩展名
	EncryptChannel     *wrapperspb.StringValue   `protobuf:"bytes,13,opt,name=encryptChannel,proto3" json:"encryptChannel,omitempty"`        // 文件被哪个应用加密过
	Encrypted          *wrapperspb.Int32Value    `protobuf:"bytes,14,opt,name=encrypted,proto3" json:"encrypted,omitempty"`                  // 是否被客户端加密过(0:未加密;1:已加密)
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_FileInfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FileInfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_FileInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetEntityIdentity() *wrapperspb.StringValue {
	if x != nil {
		return x.EntityIdentity
	}
	return nil
}

func (x *FileInfo) GetContentIdentities() []*wrapperspb.StringValue {
	if x != nil {
		return x.ContentIdentities
	}
	return nil
}

func (x *FileInfo) GetPath() *wrapperspb.StringValue {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *FileInfo) GetFilename() *wrapperspb.StringValue {
	if x != nil {
		return x.Filename
	}
	return nil
}

func (x *FileInfo) GetDriveType() *wrapperspb.StringValue {
	if x != nil {
		return x.DriveType
	}
	return nil
}

func (x *FileInfo) GetSecurityLevel() *wrapperspb.StringValue {
	if x != nil {
		return x.SecurityLevel
	}
	return nil
}

func (x *FileInfo) GetClassificationType() []*wrapperspb.StringValue {
	if x != nil {
		return x.ClassificationType
	}
	return nil
}

func (x *FileInfo) GetCreateTime() *wrapperspb.UInt64Value {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FileInfo) GetModifyTime() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ModifyTime
	}
	return nil
}

func (x *FileInfo) GetSha1() *wrapperspb.StringValue {
	if x != nil {
		return x.Sha1
	}
	return nil
}

func (x *FileInfo) GetSize() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *FileInfo) GetExt() *wrapperspb.StringValue {
	if x != nil {
		return x.Ext
	}
	return nil
}

func (x *FileInfo) GetEncryptChannel() *wrapperspb.StringValue {
	if x != nil {
		return x.EncryptChannel
	}
	return nil
}

func (x *FileInfo) GetEncrypted() *wrapperspb.Int32Value {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

var File_FileInfo_proto protoreflect.FileDescriptor

var file_FileInfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe7, 0x06, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a,
	0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x4c, 0x0a, 0x12, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x31, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x73, 0x68, 0x61, 0x31, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x65, 0x78, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x44, 0x0a, 0x0e, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x39, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FileInfo_proto_rawDescOnce sync.Once
	file_FileInfo_proto_rawDescData = file_FileInfo_proto_rawDesc
)

func file_FileInfo_proto_rawDescGZIP() []byte {
	file_FileInfo_proto_rawDescOnce.Do(func() {
		file_FileInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FileInfo_proto_rawDescData)
	})
	return file_FileInfo_proto_rawDescData
}

var file_FileInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FileInfo_proto_goTypes = []any{
	(*FileInfo)(nil),               // 0: FileInfo
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.UInt64Value)(nil), // 2: google.protobuf.UInt64Value
	(*wrapperspb.Int32Value)(nil),  // 3: google.protobuf.Int32Value
}
var file_FileInfo_proto_depIdxs = []int32{
	1,  // 0: FileInfo.entityIdentity:type_name -> google.protobuf.StringValue
	1,  // 1: FileInfo.contentIdentities:type_name -> google.protobuf.StringValue
	1,  // 2: FileInfo.path:type_name -> google.protobuf.StringValue
	1,  // 3: FileInfo.filename:type_name -> google.protobuf.StringValue
	1,  // 4: FileInfo.driveType:type_name -> google.protobuf.StringValue
	1,  // 5: FileInfo.securityLevel:type_name -> google.protobuf.StringValue
	1,  // 6: FileInfo.classificationType:type_name -> google.protobuf.StringValue
	2,  // 7: FileInfo.createTime:type_name -> google.protobuf.UInt64Value
	2,  // 8: FileInfo.modifyTime:type_name -> google.protobuf.UInt64Value
	1,  // 9: FileInfo.sha1:type_name -> google.protobuf.StringValue
	2,  // 10: FileInfo.size:type_name -> google.protobuf.UInt64Value
	1,  // 11: FileInfo.ext:type_name -> google.protobuf.StringValue
	1,  // 12: FileInfo.encryptChannel:type_name -> google.protobuf.StringValue
	3,  // 13: FileInfo.encrypted:type_name -> google.protobuf.Int32Value
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_FileInfo_proto_init() }
func file_FileInfo_proto_init() {
	if File_FileInfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FileInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FileInfo_proto_goTypes,
		DependencyIndexes: file_FileInfo_proto_depIdxs,
		MessageInfos:      file_FileInfo_proto_msgTypes,
	}.Build()
	File_FileInfo_proto = out.File
	file_FileInfo_proto_rawDesc = nil
	file_FileInfo_proto_goTypes = nil
	file_FileInfo_proto_depIdxs = nil
}
