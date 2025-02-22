// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: pet.proto

package pet_v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Age         int32                `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	OwnerId     string               `protobuf:"bytes,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	mi := &file_pet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Pet) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Pet) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Pet) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Age         int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	OwnerId     string `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_pet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_pet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_pet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_pet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type ListPetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	PetIds []string `protobuf:"bytes,3,rep,name=pet_ids,json=petIds,proto3" json:"pet_ids,omitempty"`
}

func (x *ListPetsRequest) Reset() {
	*x = ListPetsRequest{}
	mi := &file_pet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetsRequest) ProtoMessage() {}

func (x *ListPetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetsRequest.ProtoReflect.Descriptor instead.
func (*ListPetsRequest) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{5}
}

func (x *ListPetsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPetsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPetsRequest) GetPetIds() []string {
	if x != nil {
		return x.PetIds
	}
	return nil
}

type ListPetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *ListPetsResponse) Reset() {
	*x = ListPetsResponse{}
	mi := &file_pet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetsResponse) ProtoMessage() {}

func (x *ListPetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetsResponse.ProtoReflect.Descriptor instead.
func (*ListPetsResponse) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{6}
}

func (x *ListPetsResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId     string                `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Name        *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description *wrappers.StringValue `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Age         *wrappers.Int32Value  `protobuf:"bytes,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_pet_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *UpdateRequest) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateRequest) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdateRequest) GetAge() *wrappers.Int32Value {
	if x != nil {
		return x.Age
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_pet_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

var File_pet_proto protoreflect.FileDescriptor

var file_pet_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x74,
	0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8,
	0x01, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b,
	0xba, 0x48, 0x18, 0x72, 0x16, 0x10, 0x01, 0x18, 0x32, 0x32, 0x10, 0x5e, 0x5b, 0x5c, 0x70, 0x7b,
	0x4c, 0x7d, 0x30, 0x2d, 0x39, 0x20, 0x5c, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x03,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04,
	0x18, 0x64, 0x28, 0x00, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xba, 0x48, 0x18, 0x72, 0x16,
	0x10, 0x01, 0x18, 0x32, 0x32, 0x10, 0x5e, 0x5b, 0x5c, 0x70, 0x7b, 0x4c, 0x7d, 0x30, 0x2d, 0x39,
	0x20, 0x5c, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04, 0x18, 0x64, 0x28, 0x00, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03,
	0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x65, 0x74, 0x5f,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x7d, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba,
	0x48, 0x06, 0x1a, 0x04, 0x18, 0x64, 0x28, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1f, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x28, 0x0a, 0x07, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09, 0x18, 0x01, 0x22, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x06, 0x70, 0x65, 0x74, 0x49, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x22,
	0xa0, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x4d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1b, 0xba, 0x48,
	0x18, 0x72, 0x16, 0x10, 0x01, 0x18, 0x32, 0x32, 0x10, 0x5e, 0x5b, 0x5c, 0x70, 0x7b, 0x4c, 0x7d,
	0x30, 0x2d, 0x39, 0x20, 0x5c, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x47, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04, 0x18, 0x64, 0x28, 0x00, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x22, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x32, 0xa1, 0x02, 0x0a, 0x05, 0x50, 0x65, 0x74, 0x56, 0x31, 0x12, 0x37, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70,
	0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74,
	0x73, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x65, 0x74,
	0x5f, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x63, 0x6f, 0x75, 0x74, 0x64, 0x6f, 0x6f, 0x72, 0x2f,
	0x6b, 0x6f, 0x74, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pet_proto_rawDescOnce sync.Once
	file_pet_proto_rawDescData = file_pet_proto_rawDesc
)

func file_pet_proto_rawDescGZIP() []byte {
	file_pet_proto_rawDescOnce.Do(func() {
		file_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_pet_proto_rawDescData)
	})
	return file_pet_proto_rawDescData
}

var file_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pet_proto_goTypes = []any{
	(*Pet)(nil),                  // 0: pet_v1.Pet
	(*CreateRequest)(nil),        // 1: pet_v1.CreateRequest
	(*CreateResponse)(nil),       // 2: pet_v1.CreateResponse
	(*GetRequest)(nil),           // 3: pet_v1.GetRequest
	(*GetResponse)(nil),          // 4: pet_v1.GetResponse
	(*ListPetsRequest)(nil),      // 5: pet_v1.ListPetsRequest
	(*ListPetsResponse)(nil),     // 6: pet_v1.ListPetsResponse
	(*UpdateRequest)(nil),        // 7: pet_v1.UpdateRequest
	(*DeleteRequest)(nil),        // 8: pet_v1.DeleteRequest
	(*timestamp.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 10: google.protobuf.StringValue
	(*wrappers.Int32Value)(nil),  // 11: google.protobuf.Int32Value
	(*empty.Empty)(nil),          // 12: google.protobuf.Empty
}
var file_pet_proto_depIdxs = []int32{
	9,  // 0: pet_v1.Pet.created_at:type_name -> google.protobuf.Timestamp
	0,  // 1: pet_v1.GetResponse.pet:type_name -> pet_v1.Pet
	0,  // 2: pet_v1.ListPetsResponse.pets:type_name -> pet_v1.Pet
	10, // 3: pet_v1.UpdateRequest.name:type_name -> google.protobuf.StringValue
	10, // 4: pet_v1.UpdateRequest.description:type_name -> google.protobuf.StringValue
	11, // 5: pet_v1.UpdateRequest.age:type_name -> google.protobuf.Int32Value
	1,  // 6: pet_v1.PetV1.Create:input_type -> pet_v1.CreateRequest
	3,  // 7: pet_v1.PetV1.Get:input_type -> pet_v1.GetRequest
	5,  // 8: pet_v1.PetV1.ListPets:input_type -> pet_v1.ListPetsRequest
	7,  // 9: pet_v1.PetV1.Update:input_type -> pet_v1.UpdateRequest
	8,  // 10: pet_v1.PetV1.Delete:input_type -> pet_v1.DeleteRequest
	2,  // 11: pet_v1.PetV1.Create:output_type -> pet_v1.CreateResponse
	4,  // 12: pet_v1.PetV1.Get:output_type -> pet_v1.GetResponse
	6,  // 13: pet_v1.PetV1.ListPets:output_type -> pet_v1.ListPetsResponse
	12, // 14: pet_v1.PetV1.Update:output_type -> google.protobuf.Empty
	12, // 15: pet_v1.PetV1.Delete:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pet_proto_init() }
func file_pet_proto_init() {
	if File_pet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pet_proto_goTypes,
		DependencyIndexes: file_pet_proto_depIdxs,
		MessageInfos:      file_pet_proto_msgTypes,
	}.Build()
	File_pet_proto = out.File
	file_pet_proto_rawDesc = nil
	file_pet_proto_goTypes = nil
	file_pet_proto_depIdxs = nil
}
