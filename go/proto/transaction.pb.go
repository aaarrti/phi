// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: transaction.proto

package proto

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

type ValueMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scale         int32 `protobuf:"varint,1,opt,name=Scale,proto3" json:"Scale,omitempty"`
	UnscaledValue int64 `protobuf:"varint,2,opt,name=UnscaledValue,proto3" json:"UnscaledValue,omitempty"`
}

func (x *ValueMessage) Reset() {
	*x = ValueMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueMessage) ProtoMessage() {}

func (x *ValueMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueMessage.ProtoReflect.Descriptor instead.
func (*ValueMessage) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ValueMessage) GetScale() int32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

func (x *ValueMessage) GetUnscaledValue() int64 {
	if x != nil {
		return x.UnscaledValue
	}
	return 0
}

type AmountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode string        `protobuf:"bytes,1,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	Value        *ValueMessage `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *AmountMessage) Reset() {
	*x = AmountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmountMessage) ProtoMessage() {}

func (x *AmountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmountMessage.ProtoReflect.Descriptor instead.
func (*AmountMessage) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *AmountMessage) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *AmountMessage) GetValue() *ValueMessage {
	if x != nil {
		return x.Value
	}
	return nil
}

type DatesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booked string `protobuf:"bytes,1,opt,name=Booked,proto3" json:"Booked,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *DatesMessage) Reset() {
	*x = DatesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatesMessage) ProtoMessage() {}

func (x *DatesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatesMessage.ProtoReflect.Descriptor instead.
func (*DatesMessage) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *DatesMessage) GetBooked() string {
	if x != nil {
		return x.Booked
	}
	return ""
}

func (x *DatesMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TinkTransactionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID   string         `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	ID          string         `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Amount      *AmountMessage `protobuf:"bytes,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Dates       *DatesMessage  `protobuf:"bytes,4,opt,name=Dates,proto3" json:"Dates,omitempty"`
	Reference   string         `protobuf:"bytes,5,opt,name=Reference,proto3" json:"Reference,omitempty"`
	Description string         `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	Status      string         `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *TinkTransactionMessage) Reset() {
	*x = TinkTransactionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinkTransactionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinkTransactionMessage) ProtoMessage() {}

func (x *TinkTransactionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinkTransactionMessage.ProtoReflect.Descriptor instead.
func (*TinkTransactionMessage) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *TinkTransactionMessage) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *TinkTransactionMessage) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TinkTransactionMessage) GetAmount() *AmountMessage {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TinkTransactionMessage) GetDates() *DatesMessage {
	if x != nil {
		return x.Dates
	}
	return nil
}

func (x *TinkTransactionMessage) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *TinkTransactionMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TinkTransactionMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ProviderMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinancialInstitutionId string `protobuf:"bytes,1,opt,name=FinancialInstitutionId,proto3" json:"FinancialInstitutionId,omitempty"`
	DisplayName            string `protobuf:"bytes,2,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
}

func (x *ProviderMessage) Reset() {
	*x = ProviderMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderMessage) ProtoMessage() {}

func (x *ProviderMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderMessage.ProtoReflect.Descriptor instead.
func (*ProviderMessage) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *ProviderMessage) GetFinancialInstitutionId() string {
	if x != nil {
		return x.FinancialInstitutionId
	}
	return ""
}

func (x *ProviderMessage) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type AccountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinancialInstitutionId string `protobuf:"bytes,1,opt,name=FinancialInstitutionId,proto3" json:"FinancialInstitutionId,omitempty"`
	ID                     string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                   string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *AccountMessage) Reset() {
	*x = AccountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMessage) ProtoMessage() {}

func (x *AccountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountMessage.ProtoReflect.Descriptor instead.
func (*AccountMessage) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *AccountMessage) GetFinancialInstitutionId() string {
	if x != nil {
		return x.FinancialInstitutionId
	}
	return ""
}

func (x *AccountMessage) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *AccountMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x68, 0x69, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x55, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x55, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x5c, 0x0a, 0x0d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x68, 0x69, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x3c, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf3,
	0x01, 0x0a, 0x16, 0x54, 0x69, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x68, 0x69, 0x2e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x68, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x6b, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69,
	0x61, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x6c, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c,
	0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x6e,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0xd2, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x57, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x68,
	0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x1b, 0x2e, 0x70, 0x68, 0x69, 0x2e, 0x54, 0x69, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x3a,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12,
	0x2e, 0x70, 0x68, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x68, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x68, 0x69, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e,
	0x70, 0x68, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x30, 0x01, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x73, 0x74, 0x62, 0x2f, 0x70, 0x68, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transaction_proto_goTypes = []interface{}{
	(*ValueMessage)(nil),           // 0: phi.ValueMessage
	(*AmountMessage)(nil),          // 1: phi.AmountMessage
	(*DatesMessage)(nil),           // 2: phi.DatesMessage
	(*TinkTransactionMessage)(nil), // 3: phi.TinkTransactionMessage
	(*ProviderMessage)(nil),        // 4: phi.ProviderMessage
	(*AccountMessage)(nil),         // 5: phi.AccountMessage
	(*StringMessage)(nil),          // 6: phi.StringMessage
}
var file_transaction_proto_depIdxs = []int32{
	0, // 0: phi.AmountMessage.Value:type_name -> phi.ValueMessage
	1, // 1: phi.TinkTransactionMessage.Amount:type_name -> phi.AmountMessage
	2, // 2: phi.TinkTransactionMessage.Dates:type_name -> phi.DatesMessage
	6, // 3: phi.TransactionGWService.GetTransactions:input_type -> phi.StringMessage
	6, // 4: phi.TransactionGWService.GetProviders:input_type -> phi.StringMessage
	6, // 5: phi.TransactionGWService.GetAccounts:input_type -> phi.StringMessage
	3, // 6: phi.TransactionGWService.GetTransactions:output_type -> phi.TinkTransactionMessage
	4, // 7: phi.TransactionGWService.GetProviders:output_type -> phi.ProviderMessage
	5, // 8: phi.TransactionGWService.GetAccounts:output_type -> phi.AccountMessage
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueMessage); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmountMessage); i {
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
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatesMessage); i {
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
		file_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinkTransactionMessage); i {
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
		file_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderMessage); i {
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
		file_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountMessage); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
