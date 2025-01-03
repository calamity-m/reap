// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: proto/sow/v1/sow.proto

package sow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Records represent an individual record of some food
//
// Each record must have at least a user_id and description.
// The remaining options are all optional to maintain
// ease of use by users.
type Record struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique Id of this record. Should be a UUID in string encoding.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User that owns this record. Should be a UUID in string
	// encoding.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Friendly description of this food record, or what was eaten,
	// i.e. "chicken parma with some veggies"
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// A specific mapping name of some meal or object that can be
	// referenced for nutritional information later, i.e. "kellog's nutrigrain".
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Kilojules.
	//
	// kj will always take priority over the imperial "calories"
	Kj float32 `protobuf:"fixed32,5,opt,name=kj,proto3" json:"kj,omitempty"`
	// Milliliters
	//
	// ml will always take priority over the imperial "fl_oz"
	Ml float32 `protobuf:"fixed32,6,opt,name=ml,proto3" json:"ml,omitempty"`
	// Grams, 1/1000 of a kg
	//
	// grams will always take priority over the imperial "oz"
	Grams float32 `protobuf:"fixed32,7,opt,name=grams,proto3" json:"grams,omitempty"`
	// Known as calories but effectively kilocalorie.
	// (I hate imperial)
	Calories float32 `protobuf:"fixed32,8,opt,name=calories,proto3" json:"calories,omitempty"`
	// Fluid Ounce
	FlOz float32 `protobuf:"fixed32,9,opt,name=fl_oz,json=flOz,proto3" json:"fl_oz,omitempty"`
	// Ounce
	Oz float32 `protobuf:"fixed32,10,opt,name=oz,proto3" json:"oz,omitempty"`
	// Time that this was recorded. If none is provided, the time should be generated
	// by the GRPC service.
	Time          *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=time,proto3" json:"time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Record) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Record) GetKj() float32 {
	if x != nil {
		return x.Kj
	}
	return 0
}

func (x *Record) GetMl() float32 {
	if x != nil {
		return x.Ml
	}
	return 0
}

func (x *Record) GetGrams() float32 {
	if x != nil {
		return x.Grams
	}
	return 0
}

func (x *Record) GetCalories() float32 {
	if x != nil {
		return x.Calories
	}
	return 0
}

func (x *Record) GetFlOz() float32 {
	if x != nil {
		return x.FlOz
	}
	return 0
}

func (x *Record) GetOz() float32 {
	if x != nil {
		return x.Oz
	}
	return 0
}

func (x *Record) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type GetRecordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of the record to find, expected in uuid format
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecordRequest) Reset() {
	*x = GetRecordRequest{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordRequest) ProtoMessage() {}

func (x *GetRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordRequest.ProtoReflect.Descriptor instead.
func (*GetRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRecordResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Found record
	Record        *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecordResponse) Reset() {
	*x = GetRecordResponse{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordResponse) ProtoMessage() {}

func (x *GetRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordResponse.ProtoReflect.Descriptor instead.
func (*GetRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecordResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type GetRecordsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the wanted record.
	//
	// Zero values are ignored in filtering. e.g. a value for kj of
	// 0 means any record will be retrieved, regardless of kj value.
	//
	// Description and name are considered to be "contains" filters,
	// e.g. a "bob" value for the description will find all matching
	// records containing "bob" in the description
	Filter        *Record `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecordsRequest) Reset() {
	*x = GetRecordsRequest{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsRequest) ProtoMessage() {}

func (x *GetRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsRequest.ProtoReflect.Descriptor instead.
func (*GetRecordsRequest) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecordsRequest) GetFilter() *Record {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetRecordsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Found record
	Record        *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecordsResponse) Reset() {
	*x = GetRecordsResponse{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsResponse) ProtoMessage() {}

func (x *GetRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsResponse.ProtoReflect.Descriptor instead.
func (*GetRecordsResponse) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{4}
}

func (x *GetRecordsResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type CreateRecordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Record that is meant to be created
	Record        *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRecordRequest) Reset() {
	*x = CreateRecordRequest{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecordRequest) ProtoMessage() {}

func (x *CreateRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRecordRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type CreateRecordResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Record that was created in the database
	Record        *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRecordResponse) Reset() {
	*x = CreateRecordResponse{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecordResponse) ProtoMessage() {}

func (x *CreateRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRecordResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type UpdateRecordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Record that will take place of the existing record.
	// Zero values will be ignored.
	Record        *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRecordRequest) Reset() {
	*x = UpdateRecordRequest{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordRequest) ProtoMessage() {}

func (x *UpdateRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRecordRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type UpdateRecordResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Updated record for verification uses
	Record        *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRecordResponse) Reset() {
	*x = UpdateRecordResponse{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordResponse) ProtoMessage() {}

func (x *UpdateRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordResponse.ProtoReflect.Descriptor instead.
func (*UpdateRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRecordResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type DeleteRecordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of the record to delete, expected in uuid format
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRecordRequest) Reset() {
	*x = DeleteRecordRequest{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecordRequest) ProtoMessage() {}

func (x *DeleteRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteRecordRequest) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRecordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRecordResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of the record that was just deleted
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRecordResponse) Reset() {
	*x = DeleteRecordResponse{}
	mi := &file_proto_sow_v1_sow_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecordResponse) ProtoMessage() {}

func (x *DeleteRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sow_v1_sow_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecordResponse.ProtoReflect.Descriptor instead.
func (*DeleteRecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_sow_v1_sow_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRecordResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_sow_v1_sow_proto protoreflect.FileDescriptor

var file_proto_sow_v1_sow_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6b,
	0x6a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x6b, 0x6a, 0x12, 0x0e, 0x0a, 0x02, 0x6d,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x6d, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x67, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x13, 0x0a,
	0x05, 0x66, 0x6c, 0x5f, 0x6f, 0x7a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x66, 0x6c,
	0x4f, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x7a, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02,
	0x6f, 0x7a, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x3d,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x3d, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x83, 0x03, 0x0a, 0x0d,
	0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x19, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61,
	0x6c, 0x61, 0x6d, 0x69, 0x74, 0x79, 0x2d, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x70, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x77, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sow_v1_sow_proto_rawDescOnce sync.Once
	file_proto_sow_v1_sow_proto_rawDescData = file_proto_sow_v1_sow_proto_rawDesc
)

func file_proto_sow_v1_sow_proto_rawDescGZIP() []byte {
	file_proto_sow_v1_sow_proto_rawDescOnce.Do(func() {
		file_proto_sow_v1_sow_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sow_v1_sow_proto_rawDescData)
	})
	return file_proto_sow_v1_sow_proto_rawDescData
}

var file_proto_sow_v1_sow_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_sow_v1_sow_proto_goTypes = []any{
	(*Record)(nil),                // 0: sow.v1.Record
	(*GetRecordRequest)(nil),      // 1: sow.v1.GetRecordRequest
	(*GetRecordResponse)(nil),     // 2: sow.v1.GetRecordResponse
	(*GetRecordsRequest)(nil),     // 3: sow.v1.GetRecordsRequest
	(*GetRecordsResponse)(nil),    // 4: sow.v1.GetRecordsResponse
	(*CreateRecordRequest)(nil),   // 5: sow.v1.CreateRecordRequest
	(*CreateRecordResponse)(nil),  // 6: sow.v1.CreateRecordResponse
	(*UpdateRecordRequest)(nil),   // 7: sow.v1.UpdateRecordRequest
	(*UpdateRecordResponse)(nil),  // 8: sow.v1.UpdateRecordResponse
	(*DeleteRecordRequest)(nil),   // 9: sow.v1.DeleteRecordRequest
	(*DeleteRecordResponse)(nil),  // 10: sow.v1.DeleteRecordResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_proto_sow_v1_sow_proto_depIdxs = []int32{
	11, // 0: sow.v1.Record.time:type_name -> google.protobuf.Timestamp
	0,  // 1: sow.v1.GetRecordResponse.record:type_name -> sow.v1.Record
	0,  // 2: sow.v1.GetRecordsRequest.filter:type_name -> sow.v1.Record
	0,  // 3: sow.v1.GetRecordsResponse.record:type_name -> sow.v1.Record
	0,  // 4: sow.v1.CreateRecordRequest.record:type_name -> sow.v1.Record
	0,  // 5: sow.v1.CreateRecordResponse.record:type_name -> sow.v1.Record
	0,  // 6: sow.v1.UpdateRecordRequest.record:type_name -> sow.v1.Record
	0,  // 7: sow.v1.UpdateRecordResponse.record:type_name -> sow.v1.Record
	1,  // 8: sow.v1.FoodRecording.GetRecord:input_type -> sow.v1.GetRecordRequest
	3,  // 9: sow.v1.FoodRecording.GetRecords:input_type -> sow.v1.GetRecordsRequest
	5,  // 10: sow.v1.FoodRecording.CreateRecord:input_type -> sow.v1.CreateRecordRequest
	7,  // 11: sow.v1.FoodRecording.UpdateRecord:input_type -> sow.v1.UpdateRecordRequest
	9,  // 12: sow.v1.FoodRecording.DeleteRecord:input_type -> sow.v1.DeleteRecordRequest
	2,  // 13: sow.v1.FoodRecording.GetRecord:output_type -> sow.v1.GetRecordResponse
	4,  // 14: sow.v1.FoodRecording.GetRecords:output_type -> sow.v1.GetRecordsResponse
	6,  // 15: sow.v1.FoodRecording.CreateRecord:output_type -> sow.v1.CreateRecordResponse
	8,  // 16: sow.v1.FoodRecording.UpdateRecord:output_type -> sow.v1.UpdateRecordResponse
	10, // 17: sow.v1.FoodRecording.DeleteRecord:output_type -> sow.v1.DeleteRecordResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_sow_v1_sow_proto_init() }
func file_proto_sow_v1_sow_proto_init() {
	if File_proto_sow_v1_sow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sow_v1_sow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sow_v1_sow_proto_goTypes,
		DependencyIndexes: file_proto_sow_v1_sow_proto_depIdxs,
		MessageInfos:      file_proto_sow_v1_sow_proto_msgTypes,
	}.Build()
	File_proto_sow_v1_sow_proto = out.File
	file_proto_sow_v1_sow_proto_rawDesc = nil
	file_proto_sow_v1_sow_proto_goTypes = nil
	file_proto_sow_v1_sow_proto_depIdxs = nil
}
