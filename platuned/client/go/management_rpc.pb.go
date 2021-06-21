// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: management_rpc.proto

package platune

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *Progress) Reset() {
	*x = Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_management_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Progress) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

type FoldersMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folders []string `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
}

func (x *FoldersMessage) Reset() {
	*x = FoldersMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoldersMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoldersMessage) ProtoMessage() {}

func (x *FoldersMessage) ProtoReflect() protoreflect.Message {
	mi := &file_management_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoldersMessage.ProtoReflect.Descriptor instead.
func (*FoldersMessage) Descriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *FoldersMessage) GetFolders() []string {
	if x != nil {
		return x.Folders
	}
	return nil
}

type RegisteredMountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mount string `protobuf:"bytes,1,opt,name=mount,proto3" json:"mount,omitempty"`
}

func (x *RegisteredMountMessage) Reset() {
	*x = RegisteredMountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredMountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredMountMessage) ProtoMessage() {}

func (x *RegisteredMountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_management_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredMountMessage.ProtoReflect.Descriptor instead.
func (*RegisteredMountMessage) Descriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RegisteredMountMessage) GetMount() string {
	if x != nil {
		return x.Mount
	}
	return ""
}

var File_management_rpc_proto protoreflect.FileDescriptor

var file_management_rpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22,
	0x2a, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x16, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x88, 0x03, 0x0a, 0x0a,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x53, 0x79,
	0x6e, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x3e, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x19, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x63, 0x68, 0x65, 0x79, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x75, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_rpc_proto_rawDescOnce sync.Once
	file_management_rpc_proto_rawDescData = file_management_rpc_proto_rawDesc
)

func file_management_rpc_proto_rawDescGZIP() []byte {
	file_management_rpc_proto_rawDescOnce.Do(func() {
		file_management_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_rpc_proto_rawDescData)
	})
	return file_management_rpc_proto_rawDescData
}

var file_management_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_management_rpc_proto_goTypes = []interface{}{
	(*Progress)(nil),               // 0: management_rpc.Progress
	(*FoldersMessage)(nil),         // 1: management_rpc.FoldersMessage
	(*RegisteredMountMessage)(nil), // 2: management_rpc.RegisteredMountMessage
	(*emptypb.Empty)(nil),          // 3: google.protobuf.Empty
}
var file_management_rpc_proto_depIdxs = []int32{
	3, // 0: management_rpc.Management.Sync:input_type -> google.protobuf.Empty
	1, // 1: management_rpc.Management.AddFolders:input_type -> management_rpc.FoldersMessage
	3, // 2: management_rpc.Management.GetAllFolders:input_type -> google.protobuf.Empty
	2, // 3: management_rpc.Management.RegisterMount:input_type -> management_rpc.RegisteredMountMessage
	3, // 4: management_rpc.Management.GetRegisteredMount:input_type -> google.protobuf.Empty
	0, // 5: management_rpc.Management.Sync:output_type -> management_rpc.Progress
	3, // 6: management_rpc.Management.AddFolders:output_type -> google.protobuf.Empty
	1, // 7: management_rpc.Management.GetAllFolders:output_type -> management_rpc.FoldersMessage
	3, // 8: management_rpc.Management.RegisterMount:output_type -> google.protobuf.Empty
	2, // 9: management_rpc.Management.GetRegisteredMount:output_type -> management_rpc.RegisteredMountMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_rpc_proto_init() }
func file_management_rpc_proto_init() {
	if File_management_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_management_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoldersMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredMountMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_rpc_proto_goTypes,
		DependencyIndexes: file_management_rpc_proto_depIdxs,
		MessageInfos:      file_management_rpc_proto_msgTypes,
	}.Build()
	File_management_rpc_proto = out.File
	file_management_rpc_proto_rawDesc = nil
	file_management_rpc_proto_goTypes = nil
	file_management_rpc_proto_depIdxs = nil
}
