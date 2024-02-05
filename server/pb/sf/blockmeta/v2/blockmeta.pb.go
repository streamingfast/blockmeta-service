// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sf/blockmeta/v2/blockmeta.proto

package pbbmsrv

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP(), []int{0}
}

// Block Requests
type NumToIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNum uint64 `protobuf:"varint,1,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
}

func (x *NumToIDReq) Reset() {
	*x = NumToIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumToIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumToIDReq) ProtoMessage() {}

func (x *NumToIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumToIDReq.ProtoReflect.Descriptor instead.
func (*NumToIDReq) Descriptor() ([]byte, []int) {
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP(), []int{1}
}

func (x *NumToIDReq) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

type IDToNumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockID string `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
}

func (x *IDToNumReq) Reset() {
	*x = IDToNumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDToNumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDToNumReq) ProtoMessage() {}

func (x *IDToNumReq) ProtoReflect() protoreflect.Message {
	mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDToNumReq.ProtoReflect.Descriptor instead.
func (*IDToNumReq) Descriptor() ([]byte, []int) {
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP(), []int{2}
}

func (x *IDToNumReq) GetBlockID() string {
	if x != nil {
		return x.BlockID
	}
	return ""
}

// Block & BlockByTime Responses
type BlockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num  uint64                 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Time *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *BlockResp) Reset() {
	*x = BlockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResp) ProtoMessage() {}

func (x *BlockResp) ProtoReflect() protoreflect.Message {
	mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResp.ProtoReflect.Descriptor instead.
func (*BlockResp) Descriptor() ([]byte, []int) {
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP(), []int{3}
}

func (x *BlockResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlockResp) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BlockResp) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

// BlockByTime Requests
type TimeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TimeReq) Reset() {
	*x = TimeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeReq) ProtoMessage() {}

func (x *TimeReq) ProtoReflect() protoreflect.Message {
	mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeReq.ProtoReflect.Descriptor instead.
func (*TimeReq) Descriptor() ([]byte, []int) {
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP(), []int{4}
}

func (x *TimeReq) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type RelativeTimeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Inclusive bool                   `protobuf:"varint,2,opt,name=inclusive,proto3" json:"inclusive,omitempty"`
}

func (x *RelativeTimeReq) Reset() {
	*x = RelativeTimeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelativeTimeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelativeTimeReq) ProtoMessage() {}

func (x *RelativeTimeReq) ProtoReflect() protoreflect.Message {
	mi := &file_sf_blockmeta_v2_blockmeta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelativeTimeReq.ProtoReflect.Descriptor instead.
func (*RelativeTimeReq) Descriptor() ([]byte, []int) {
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP(), []int{5}
}

func (x *RelativeTimeReq) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *RelativeTimeReq) GetInclusive() bool {
	if x != nil {
		return x.Inclusive
	}
	return false
}

var File_sf_blockmeta_v2_blockmeta_proto protoreflect.FileDescriptor

var file_sf_blockmeta_v2_blockmeta_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x66, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76,
	0x32, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x0a,
	0x4e, 0x75, 0x6d, 0x54, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x22, 0x26, 0x0a, 0x0a, 0x49, 0x44, 0x54, 0x6f, 0x4e, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x22, 0x5d,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x39, 0x0a,
	0x07, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x32, 0xcb, 0x01, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x54, 0x6f, 0x49, 0x44, 0x12, 0x1b,
	0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32,
	0x2e, 0x4e, 0x75, 0x6d, 0x54, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x66,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x07, 0x49, 0x44, 0x54, 0x6f, 0x4e,
	0x75, 0x6d, 0x12, 0x1b, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x44, 0x54, 0x6f, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x04, 0x48,
	0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x66,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x32, 0xd8, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x02, 0x41, 0x74, 0x12, 0x18, 0x2e,
	0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x05, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73,
	0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x06, 0x42, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x3b, 0x70, 0x62, 0x62, 0x6d, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_blockmeta_v2_blockmeta_proto_rawDescOnce sync.Once
	file_sf_blockmeta_v2_blockmeta_proto_rawDescData = file_sf_blockmeta_v2_blockmeta_proto_rawDesc
)

func file_sf_blockmeta_v2_blockmeta_proto_rawDescGZIP() []byte {
	file_sf_blockmeta_v2_blockmeta_proto_rawDescOnce.Do(func() {
		file_sf_blockmeta_v2_blockmeta_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_blockmeta_v2_blockmeta_proto_rawDescData)
	})
	return file_sf_blockmeta_v2_blockmeta_proto_rawDescData
}

var file_sf_blockmeta_v2_blockmeta_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sf_blockmeta_v2_blockmeta_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: sf.blockmeta.v2.Empty
	(*NumToIDReq)(nil),            // 1: sf.blockmeta.v2.NumToIDReq
	(*IDToNumReq)(nil),            // 2: sf.blockmeta.v2.IDToNumReq
	(*BlockResp)(nil),             // 3: sf.blockmeta.v2.BlockResp
	(*TimeReq)(nil),               // 4: sf.blockmeta.v2.TimeReq
	(*RelativeTimeReq)(nil),       // 5: sf.blockmeta.v2.RelativeTimeReq
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_sf_blockmeta_v2_blockmeta_proto_depIdxs = []int32{
	6, // 0: sf.blockmeta.v2.BlockResp.time:type_name -> google.protobuf.Timestamp
	6, // 1: sf.blockmeta.v2.TimeReq.time:type_name -> google.protobuf.Timestamp
	6, // 2: sf.blockmeta.v2.RelativeTimeReq.time:type_name -> google.protobuf.Timestamp
	1, // 3: sf.blockmeta.v2.Block.NumToID:input_type -> sf.blockmeta.v2.NumToIDReq
	2, // 4: sf.blockmeta.v2.Block.IDToNum:input_type -> sf.blockmeta.v2.IDToNumReq
	0, // 5: sf.blockmeta.v2.Block.Head:input_type -> sf.blockmeta.v2.Empty
	4, // 6: sf.blockmeta.v2.BlockByTime.At:input_type -> sf.blockmeta.v2.TimeReq
	5, // 7: sf.blockmeta.v2.BlockByTime.After:input_type -> sf.blockmeta.v2.RelativeTimeReq
	5, // 8: sf.blockmeta.v2.BlockByTime.Before:input_type -> sf.blockmeta.v2.RelativeTimeReq
	3, // 9: sf.blockmeta.v2.Block.NumToID:output_type -> sf.blockmeta.v2.BlockResp
	3, // 10: sf.blockmeta.v2.Block.IDToNum:output_type -> sf.blockmeta.v2.BlockResp
	3, // 11: sf.blockmeta.v2.Block.Head:output_type -> sf.blockmeta.v2.BlockResp
	3, // 12: sf.blockmeta.v2.BlockByTime.At:output_type -> sf.blockmeta.v2.BlockResp
	3, // 13: sf.blockmeta.v2.BlockByTime.After:output_type -> sf.blockmeta.v2.BlockResp
	3, // 14: sf.blockmeta.v2.BlockByTime.Before:output_type -> sf.blockmeta.v2.BlockResp
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sf_blockmeta_v2_blockmeta_proto_init() }
func file_sf_blockmeta_v2_blockmeta_proto_init() {
	if File_sf_blockmeta_v2_blockmeta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_blockmeta_v2_blockmeta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_sf_blockmeta_v2_blockmeta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumToIDReq); i {
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
		file_sf_blockmeta_v2_blockmeta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDToNumReq); i {
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
		file_sf_blockmeta_v2_blockmeta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResp); i {
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
		file_sf_blockmeta_v2_blockmeta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeReq); i {
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
		file_sf_blockmeta_v2_blockmeta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelativeTimeReq); i {
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
			RawDescriptor: file_sf_blockmeta_v2_blockmeta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_sf_blockmeta_v2_blockmeta_proto_goTypes,
		DependencyIndexes: file_sf_blockmeta_v2_blockmeta_proto_depIdxs,
		MessageInfos:      file_sf_blockmeta_v2_blockmeta_proto_msgTypes,
	}.Build()
	File_sf_blockmeta_v2_blockmeta_proto = out.File
	file_sf_blockmeta_v2_blockmeta_proto_rawDesc = nil
	file_sf_blockmeta_v2_blockmeta_proto_goTypes = nil
	file_sf_blockmeta_v2_blockmeta_proto_depIdxs = nil
}