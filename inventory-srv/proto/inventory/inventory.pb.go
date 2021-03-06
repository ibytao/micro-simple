// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: inventory/inventory.proto

package mu_micro_book_srv_inventory

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Inv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId      int64 `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	UnitPrice   int64 `protobuf:"varint,3,opt,name=unitPrice,proto3" json:"unitPrice,omitempty"`
	Stock       int64 `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Version     int64 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	CreatedTime int64 `protobuf:"varint,6,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime int64 `protobuf:"varint,7,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
}

func (x *Inv) Reset() {
	*x = Inv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inv) ProtoMessage() {}

func (x *Inv) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inv.ProtoReflect.Descriptor instead.
func (*Inv) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Inv) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Inv) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *Inv) GetUnitPrice() int64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *Inv) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Inv) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Inv) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *Inv) GetUpdatedTime() int64 {
	if x != nil {
		return x.UpdatedTime
	}
	return 0
}

type InvHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId int64 `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	UserId int64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	State  int64 `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *InvHistory) Reset() {
	*x = InvHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvHistory) ProtoMessage() {}

func (x *InvHistory) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvHistory.ProtoReflect.Descriptor instead.
func (*InvHistory) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *InvHistory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InvHistory) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *InvHistory) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *InvHistory) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId       int64 `protobuf:"varint,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	UserId       int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	HistoryId    int64 `protobuf:"varint,3,opt,name=historyId,proto3" json:"historyId,omitempty"`
	HistoryState int32 `protobuf:"varint,4,opt,name=historyState,proto3" json:"historyState,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *Request) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Request) GetHistoryId() int64 {
	if x != nil {
		return x.HistoryId
	}
	return 0
}

func (x *Request) GetHistoryState() int32 {
	if x != nil {
		return x.HistoryState
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   *Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Inv     *Inv        `protobuf:"bytes,3,opt,name=inv,proto3" json:"inv,omitempty"`
	InvH    *InvHistory `protobuf:"bytes,4,opt,name=invH,proto3" json:"invH,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_inventory_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Response) GetInv() *Inv {
	if x != nil {
		return x.Inv
	}
	return nil
}

func (x *Response) GetInvH() *InvHistory {
	if x != nil {
		return x.InvH
	}
	return nil
}

var File_inventory_inventory_proto protoreflect.FileDescriptor

var file_inventory_inventory_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d, 0x75, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x03, 0x49, 0x6e, 0x76,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x6e, 0x69,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x0a, 0x49, 0x6e,
	0x76, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x33,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x22, 0x7b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0xcf, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x32, 0x0a, 0x03, 0x69, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76,
	0x52, 0x03, 0x69, 0x6e, 0x76, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x6e, 0x76, 0x48, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x49, 0x6e, 0x76, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x69, 0x6e,
	0x76, 0x48, 0x32, 0xbc, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x55, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x24, 0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x12, 0x24, 0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x75, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_inventory_proto_rawDescOnce sync.Once
	file_inventory_inventory_proto_rawDescData = file_inventory_inventory_proto_rawDesc
)

func file_inventory_inventory_proto_rawDescGZIP() []byte {
	file_inventory_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_inventory_proto_rawDescData)
	})
	return file_inventory_inventory_proto_rawDescData
}

var file_inventory_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_inventory_inventory_proto_goTypes = []interface{}{
	(*Inv)(nil),        // 0: mu.micro.book.srv.inventory.Inv
	(*InvHistory)(nil), // 1: mu.micro.book.srv.inventory.InvHistory
	(*Error)(nil),      // 2: mu.micro.book.srv.inventory.Error
	(*Request)(nil),    // 3: mu.micro.book.srv.inventory.Request
	(*Response)(nil),   // 4: mu.micro.book.srv.inventory.Response
}
var file_inventory_inventory_proto_depIdxs = []int32{
	2, // 0: mu.micro.book.srv.inventory.Response.error:type_name -> mu.micro.book.srv.inventory.Error
	0, // 1: mu.micro.book.srv.inventory.Response.inv:type_name -> mu.micro.book.srv.inventory.Inv
	1, // 2: mu.micro.book.srv.inventory.Response.invH:type_name -> mu.micro.book.srv.inventory.InvHistory
	3, // 3: mu.micro.book.srv.inventory.Inventory.Sell:input_type -> mu.micro.book.srv.inventory.Request
	3, // 4: mu.micro.book.srv.inventory.Inventory.Confirm:input_type -> mu.micro.book.srv.inventory.Request
	4, // 5: mu.micro.book.srv.inventory.Inventory.Sell:output_type -> mu.micro.book.srv.inventory.Response
	4, // 6: mu.micro.book.srv.inventory.Inventory.Confirm:output_type -> mu.micro.book.srv.inventory.Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_inventory_inventory_proto_init() }
func file_inventory_inventory_proto_init() {
	if File_inventory_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inv); i {
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
		file_inventory_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvHistory); i {
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
		file_inventory_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_inventory_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_inventory_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_inventory_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_inventory_proto_msgTypes,
	}.Build()
	File_inventory_inventory_proto = out.File
	file_inventory_inventory_proto_rawDesc = nil
	file_inventory_inventory_proto_goTypes = nil
	file_inventory_inventory_proto_depIdxs = nil
}
