// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: data_service.proto

package data_service

import (
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

type InstrumentType int32

const (
	InstrumentType_Stock InstrumentType = 0
	InstrumentType_Index InstrumentType = 1
)

// Enum value maps for InstrumentType.
var (
	InstrumentType_name = map[int32]string{
		0: "Stock",
		1: "Index",
	}
	InstrumentType_value = map[string]int32{
		"Stock": 0,
		"Index": 1,
	}
)

func (x InstrumentType) Enum() *InstrumentType {
	p := new(InstrumentType)
	*p = x
	return p
}

func (x InstrumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstrumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_data_service_proto_enumTypes[0].Descriptor()
}

func (InstrumentType) Type() protoreflect.EnumType {
	return &file_data_service_proto_enumTypes[0]
}

func (x InstrumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstrumentType.Descriptor instead.
func (InstrumentType) EnumDescriptor() ([]byte, []int) {
	return file_data_service_proto_rawDescGZIP(), []int{0}
}

type Instrument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scrip          string         `protobuf:"bytes,1,opt,name=Scrip,proto3" json:"Scrip,omitempty"`
	KiteToken      string         `protobuf:"bytes,2,opt,name=KiteToken,proto3" json:"KiteToken,omitempty"`
	ExchangeToken  string         `protobuf:"bytes,3,opt,name=ExchangeToken,proto3" json:"ExchangeToken,omitempty"`
	Exchange       string         `protobuf:"bytes,4,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
	InstrumentType InstrumentType `protobuf:"varint,5,opt,name=instrumentType,proto3,enum=data_service.InstrumentType" json:"instrumentType,omitempty"`
}

func (x *Instrument) Reset() {
	*x = Instrument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instrument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instrument) ProtoMessage() {}

func (x *Instrument) ProtoReflect() protoreflect.Message {
	mi := &file_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instrument.ProtoReflect.Descriptor instead.
func (*Instrument) Descriptor() ([]byte, []int) {
	return file_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *Instrument) GetScrip() string {
	if x != nil {
		return x.Scrip
	}
	return ""
}

func (x *Instrument) GetKiteToken() string {
	if x != nil {
		return x.KiteToken
	}
	return ""
}

func (x *Instrument) GetExchangeToken() string {
	if x != nil {
		return x.ExchangeToken
	}
	return ""
}

func (x *Instrument) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *Instrument) GetInstrumentType() InstrumentType {
	if x != nil {
		return x.InstrumentType
	}
	return InstrumentType_Stock
}

type Instruments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Instrument `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Instruments) Reset() {
	*x = Instruments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruments) ProtoMessage() {}

func (x *Instruments) ProtoReflect() protoreflect.Message {
	mi := &file_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruments.ProtoReflect.Descriptor instead.
func (*Instruments) Descriptor() ([]byte, []int) {
	return file_data_service_proto_rawDescGZIP(), []int{1}
}

func (x *Instruments) GetItems() []*Instrument {
	if x != nil {
		return x.Items
	}
	return nil
}

type LTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ltp            float64        `protobuf:"fixed64,1,opt,name=Ltp,proto3" json:"Ltp,omitempty"`
	Open           float64        `protobuf:"fixed64,2,opt,name=Open,proto3" json:"Open,omitempty"`
	Close          float64        `protobuf:"fixed64,3,opt,name=Close,proto3" json:"Close,omitempty"`
	Low            float64        `protobuf:"fixed64,4,opt,name=Low,proto3" json:"Low,omitempty"`
	High           float64        `protobuf:"fixed64,5,opt,name=High,proto3" json:"High,omitempty"`
	ExchangeToken  string         `protobuf:"bytes,6,opt,name=ExchangeToken,proto3" json:"ExchangeToken,omitempty"`
	Scrip          string         `protobuf:"bytes,7,opt,name=Scrip,proto3" json:"Scrip,omitempty"`
	InstrumentType InstrumentType `protobuf:"varint,8,opt,name=InstrumentType,proto3,enum=data_service.InstrumentType" json:"InstrumentType,omitempty"`
	E              *Error         `protobuf:"bytes,9,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *LTP) Reset() {
	*x = LTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LTP) ProtoMessage() {}

func (x *LTP) ProtoReflect() protoreflect.Message {
	mi := &file_data_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LTP.ProtoReflect.Descriptor instead.
func (*LTP) Descriptor() ([]byte, []int) {
	return file_data_service_proto_rawDescGZIP(), []int{2}
}

func (x *LTP) GetLtp() float64 {
	if x != nil {
		return x.Ltp
	}
	return 0
}

func (x *LTP) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *LTP) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *LTP) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *LTP) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *LTP) GetExchangeToken() string {
	if x != nil {
		return x.ExchangeToken
	}
	return ""
}

func (x *LTP) GetScrip() string {
	if x != nil {
		return x.Scrip
	}
	return ""
}

func (x *LTP) GetInstrumentType() InstrumentType {
	if x != nil {
		return x.InstrumentType
	}
	return InstrumentType_Stock
}

func (x *LTP) GetE() *Error {
	if x != nil {
		return x.E
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_data_service_proto_msgTypes[3]
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
	return file_data_service_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_data_service_proto protoreflect.FileDescriptor

var file_data_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x72, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x63, 0x72, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x4b, 0x69, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4b, 0x69, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a,
	0x0b, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x8c, 0x02, 0x0a,
	0x03, 0x4c, 0x54, 0x50, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x4c, 0x74, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4c,
	0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x48, 0x69, 0x67, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x12, 0x44, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x01, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x01, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2a, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x01, 0x32, 0x58, 0x0a, 0x0a, 0x4c, 0x54,
	0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x54, 0x50, 0x66, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x11, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x54, 0x50,
	0x22, 0x00, 0x30, 0x01, 0x32, 0x5f, 0x0a, 0x11, 0x4c, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x16, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x54, 0x50, 0x66, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x11,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x54,
	0x50, 0x22, 0x00, 0x30, 0x01, 0x42, 0x94, 0x01, 0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65,
	0x76, 0x75, 0x6c, 0x63, 0x61, 0x6e, 0x2e, 0x72, 0x65, 0x76, 0x75, 0x6c, 0x63, 0x61, 0x6e, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x10, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x75, 0x6c, 0x63, 0x61, 0x6e, 0x2f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_service_proto_rawDescOnce sync.Once
	file_data_service_proto_rawDescData = file_data_service_proto_rawDesc
)

func file_data_service_proto_rawDescGZIP() []byte {
	file_data_service_proto_rawDescOnce.Do(func() {
		file_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_service_proto_rawDescData)
	})
	return file_data_service_proto_rawDescData
}

var file_data_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_service_proto_goTypes = []interface{}{
	(InstrumentType)(0), // 0: data_service.InstrumentType
	(*Instrument)(nil),  // 1: data_service.Instrument
	(*Instruments)(nil), // 2: data_service.Instruments
	(*LTP)(nil),         // 3: data_service.LTP
	(*Error)(nil),       // 4: data_service.Error
}
var file_data_service_proto_depIdxs = []int32{
	0, // 0: data_service.Instrument.instrumentType:type_name -> data_service.InstrumentType
	1, // 1: data_service.Instruments.items:type_name -> data_service.Instrument
	0, // 2: data_service.LTP.InstrumentType:type_name -> data_service.InstrumentType
	4, // 3: data_service.LTP.e:type_name -> data_service.Error
	2, // 4: data_service.LTPService.WatchLTPforInstruments:input_type -> data_service.Instruments
	2, // 5: data_service.LTPServiceManager.WatchLTPforInstruments:input_type -> data_service.Instruments
	3, // 6: data_service.LTPService.WatchLTPforInstruments:output_type -> data_service.LTP
	3, // 7: data_service.LTPServiceManager.WatchLTPforInstruments:output_type -> data_service.LTP
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_data_service_proto_init() }
func file_data_service_proto_init() {
	if File_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instrument); i {
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
		file_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instruments); i {
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
		file_data_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LTP); i {
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
		file_data_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_data_service_proto_goTypes,
		DependencyIndexes: file_data_service_proto_depIdxs,
		EnumInfos:         file_data_service_proto_enumTypes,
		MessageInfos:      file_data_service_proto_msgTypes,
	}.Build()
	File_data_service_proto = out.File
	file_data_service_proto_rawDesc = nil
	file_data_service_proto_goTypes = nil
	file_data_service_proto_depIdxs = nil
}