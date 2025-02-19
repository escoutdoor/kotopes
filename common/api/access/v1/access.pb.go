// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: access.proto

package access_v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Method   string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	UserId   string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_access_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *CheckRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CheckRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAllowed bool `protobuf:"varint,1,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_access_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetIsAllowed() bool {
	if x != nil {
		return x.IsAllowed
	}
	return false
}

var File_access_proto protoreflect.FileDescriptor

var file_access_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0c,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d,
	0xba, 0x48, 0x1a, 0x72, 0x18, 0x32, 0x16, 0x5e, 0x2f, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x30, 0x2d, 0x39, 0x2f, 0x7b, 0x7d, 0x2e, 0x5f, 0x2a, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xba, 0x48, 0x21, 0x72, 0x1f, 0x52, 0x04,
	0x50, 0x4f, 0x53, 0x54, 0x52, 0x03, 0x50, 0x55, 0x54, 0x52, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48,
	0x52, 0x03, 0x47, 0x45, 0x54, 0x52, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xba, 0x48, 0x0f, 0x72, 0x0d, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x2e, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x32, 0x3b, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x56, 0x31, 0x12, 0x2f, 0x0a, 0x0e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x0d,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x63, 0x6f,
	0x75, 0x74, 0x64, 0x6f, 0x6f, 0x72, 0x2f, 0x6b, 0x6f, 0x74, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_access_proto_rawDescOnce sync.Once
	file_access_proto_rawDescData = file_access_proto_rawDesc
)

func file_access_proto_rawDescGZIP() []byte {
	file_access_proto_rawDescOnce.Do(func() {
		file_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_access_proto_rawDescData)
	})
	return file_access_proto_rawDescData
}

var file_access_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_access_proto_goTypes = []any{
	(*CheckRequest)(nil),  // 0: CheckRequest
	(*CheckResponse)(nil), // 1: CheckResponse
}
var file_access_proto_depIdxs = []int32{
	0, // 0: AccessV1.CheckIsAllowed:input_type -> CheckRequest
	1, // 1: AccessV1.CheckIsAllowed:output_type -> CheckResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_access_proto_init() }
func file_access_proto_init() {
	if File_access_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_access_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_access_proto_goTypes,
		DependencyIndexes: file_access_proto_depIdxs,
		MessageInfos:      file_access_proto_msgTypes,
	}.Build()
	File_access_proto = out.File
	file_access_proto_rawDesc = nil
	file_access_proto_goTypes = nil
	file_access_proto_depIdxs = nil
}
