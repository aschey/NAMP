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

type EntryType int32

const (
	EntryType_ALBUM        EntryType = 0
	EntryType_SONG         EntryType = 1
	EntryType_ARTIST       EntryType = 2
	EntryType_ALBUM_ARTIST EntryType = 3
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "ALBUM",
		1: "SONG",
		2: "ARTIST",
		3: "ALBUM_ARTIST",
	}
	EntryType_value = map[string]int32{
		"ALBUM":        0,
		"SONG":         1,
		"ARTIST":       2,
		"ALBUM_ARTIST": 3,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_management_rpc_proto_enumTypes[0].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_management_rpc_proto_enumTypes[0]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{0}
}

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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query          string  `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	StartSeparator *string `protobuf:"bytes,2,opt,name=start_separator,json=startSeparator,proto3,oneof" json:"start_separator,omitempty"`
	EndSeparator   *string `protobuf:"bytes,3,opt,name=end_separator,json=endSeparator,proto3,oneof" json:"end_separator,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetStartSeparator() string {
	if x != nil && x.StartSeparator != nil {
		return *x.StartSeparator
	}
	return ""
}

func (x *SearchRequest) GetEndSeparator() string {
	if x != nil && x.EndSeparator != nil {
		return *x.EndSeparator
	}
	return ""
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry         string    `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	EntryType     EntryType `protobuf:"varint,2,opt,name=entry_type,json=entryType,proto3,enum=management_rpc.EntryType" json:"entry_type,omitempty"`
	Artist        *string   `protobuf:"bytes,3,opt,name=artist,proto3,oneof" json:"artist,omitempty"`
	CorrelationId int32     `protobuf:"varint,4,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Description   string    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_management_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResult) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

func (x *SearchResult) GetEntryType() EntryType {
	if x != nil {
		return x.EntryType
	}
	return EntryType_ALBUM
}

func (x *SearchResult) GetArtist() string {
	if x != nil && x.Artist != nil {
		return *x.Artist
	}
	return ""
}

func (x *SearchResult) GetCorrelationId() int32 {
	if x != nil {
		return x.CorrelationId
	}
	return 0
}

func (x *SearchResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SearchResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_management_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResponse) GetResults() []*SearchResult {
	if x != nil {
		return x.Results
	}
	return nil
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
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x70,
	0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x53,
	0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0xcf, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2a, 0x3e, 0x0a,
	0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c,
	0x42, 0x55, 0x4d, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x52, 0x54, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x4c, 0x42, 0x55, 0x4d, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x53, 0x54, 0x10, 0x03, 0x32, 0xd7, 0x03,
	0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x04,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3e, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0f, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x19, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x63, 0x68, 0x65, 0x79, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x75, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_management_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_management_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_management_rpc_proto_goTypes = []interface{}{
	(EntryType)(0),                 // 0: management_rpc.EntryType
	(*Progress)(nil),               // 1: management_rpc.Progress
	(*FoldersMessage)(nil),         // 2: management_rpc.FoldersMessage
	(*RegisteredMountMessage)(nil), // 3: management_rpc.RegisteredMountMessage
	(*SearchRequest)(nil),          // 4: management_rpc.SearchRequest
	(*SearchResult)(nil),           // 5: management_rpc.SearchResult
	(*SearchResponse)(nil),         // 6: management_rpc.SearchResponse
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_management_rpc_proto_depIdxs = []int32{
	0, // 0: management_rpc.SearchResult.entry_type:type_name -> management_rpc.EntryType
	5, // 1: management_rpc.SearchResponse.results:type_name -> management_rpc.SearchResult
	7, // 2: management_rpc.Management.Sync:input_type -> google.protobuf.Empty
	2, // 3: management_rpc.Management.AddFolders:input_type -> management_rpc.FoldersMessage
	7, // 4: management_rpc.Management.GetAllFolders:input_type -> google.protobuf.Empty
	3, // 5: management_rpc.Management.RegisterMount:input_type -> management_rpc.RegisteredMountMessage
	7, // 6: management_rpc.Management.GetRegisteredMount:input_type -> google.protobuf.Empty
	4, // 7: management_rpc.Management.Search:input_type -> management_rpc.SearchRequest
	1, // 8: management_rpc.Management.Sync:output_type -> management_rpc.Progress
	7, // 9: management_rpc.Management.AddFolders:output_type -> google.protobuf.Empty
	2, // 10: management_rpc.Management.GetAllFolders:output_type -> management_rpc.FoldersMessage
	7, // 11: management_rpc.Management.RegisterMount:output_type -> google.protobuf.Empty
	3, // 12: management_rpc.Management.GetRegisteredMount:output_type -> management_rpc.RegisteredMountMessage
	6, // 13: management_rpc.Management.Search:output_type -> management_rpc.SearchResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_management_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_management_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
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
		file_management_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
	file_management_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_management_rpc_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_rpc_proto_goTypes,
		DependencyIndexes: file_management_rpc_proto_depIdxs,
		EnumInfos:         file_management_rpc_proto_enumTypes,
		MessageInfos:      file_management_rpc_proto_msgTypes,
	}.Build()
	File_management_rpc_proto = out.File
	file_management_rpc_proto_rawDesc = nil
	file_management_rpc_proto_goTypes = nil
	file_management_rpc_proto_depIdxs = nil
}
