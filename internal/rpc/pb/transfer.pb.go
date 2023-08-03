// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/transfer.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request
type BalanceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserToken string `protobuf:"bytes,1,opt,name=userToken,proto3" json:"userToken,omitempty"` // User token to identify the requester. (Required)
}

func (x *BalanceUpdate) Reset() {
	*x = BalanceUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceUpdate) ProtoMessage() {}

func (x *BalanceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceUpdate.ProtoReflect.Descriptor instead.
func (*BalanceUpdate) Descriptor() ([]byte, []int) {
	return file_proto_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceUpdate) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserToken     string `protobuf:"bytes,1,opt,name=userToken,proto3" json:"userToken,omitempty"`         // User token to identify the requester. (Required)
	TransferSum   int32  `protobuf:"varint,2,opt,name=transferSum,proto3" json:"transferSum,omitempty"`    // Amount to transfer. (Required)
	RecipientName string `protobuf:"bytes,3,opt,name=recipientName,proto3" json:"recipientName,omitempty"` // Recipient's name. (Required)
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferRequest) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

func (x *TransferRequest) GetTransferSum() int32 {
	if x != nil {
		return x.TransferSum
	}
	return 0
}

func (x *TransferRequest) GetRecipientName() string {
	if x != nil {
		return x.RecipientName
	}
	return ""
}

type OperationRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserToken string `protobuf:"bytes,1,opt,name=userToken,proto3" json:"userToken,omitempty"` // User token to identify the requester. (Required)
}

func (x *OperationRequestList) Reset() {
	*x = OperationRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationRequestList) ProtoMessage() {}

func (x *OperationRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationRequestList.ProtoReflect.Descriptor instead.
func (*OperationRequestList) Descriptor() ([]byte, []int) {
	return file_proto_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *OperationRequestList) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

// Response
type TransferInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferSum   int32                  `protobuf:"varint,1,opt,name=transferSum,proto3" json:"transferSum,omitempty"`    // Amount transferred. (Required)
	RecipientName string                 `protobuf:"bytes,2,opt,name=recipientName,proto3" json:"recipientName,omitempty"` // Recipient's name. (Required)
	TimeCompleted *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timeCompleted,proto3" json:"timeCompleted,omitempty"` // Timestamp of when the transfer was completed. (Required)
}

func (x *TransferInfo) Reset() {
	*x = TransferInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferInfo) ProtoMessage() {}

func (x *TransferInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferInfo.ProtoReflect.Descriptor instead.
func (*TransferInfo) Descriptor() ([]byte, []int) {
	return file_proto_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *TransferInfo) GetTransferSum() int32 {
	if x != nil {
		return x.TransferSum
	}
	return 0
}

func (x *TransferInfo) GetRecipientName() string {
	if x != nil {
		return x.RecipientName
	}
	return ""
}

func (x *TransferInfo) GetTimeCompleted() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCompleted
	}
	return nil
}

type TransferStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`            // Status of the transfer operation. (Required)
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"` // Error message in case of failure. (Optional)
}

func (x *TransferStatusResponse) Reset() {
	*x = TransferStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferStatusResponse) ProtoMessage() {}

func (x *TransferStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferStatusResponse.ProtoReflect.Descriptor instead.
func (*TransferStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *TransferStatusResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *TransferStatusResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type OperationResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*TransferInfo `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"` // List of transfer operations. (Required)
}

func (x *OperationResponseList) Reset() {
	*x = OperationResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponseList) ProtoMessage() {}

func (x *OperationResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponseList.ProtoReflect.Descriptor instead.
func (*OperationResponseList) Descriptor() ([]byte, []int) {
	return file_proto_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *OperationResponseList) GetOperations() []*TransferInfo {
	if x != nil {
		return x.Operations
	}
	return nil
}

var File_proto_transfer_proto protoreflect.FileDescriptor

var file_proto_transfer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x77, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x53, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a,
	0x14, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0d,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x54,
	0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x73, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x32, 0x7f, 0x0a, 0x10, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b,
	0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transfer_proto_rawDescOnce sync.Once
	file_proto_transfer_proto_rawDescData = file_proto_transfer_proto_rawDesc
)

func file_proto_transfer_proto_rawDescGZIP() []byte {
	file_proto_transfer_proto_rawDescOnce.Do(func() {
		file_proto_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transfer_proto_rawDescData)
	})
	return file_proto_transfer_proto_rawDescData
}

var file_proto_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_transfer_proto_goTypes = []interface{}{
	(*BalanceUpdate)(nil),          // 0: transfer.BalanceUpdate
	(*TransferRequest)(nil),        // 1: transfer.TransferRequest
	(*OperationRequestList)(nil),   // 2: transfer.OperationRequestList
	(*TransferInfo)(nil),           // 3: transfer.TransferInfo
	(*TransferStatusResponse)(nil), // 4: transfer.TransferStatusResponse
	(*OperationResponseList)(nil),  // 5: transfer.OperationResponseList
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_proto_transfer_proto_depIdxs = []int32{
	6, // 0: transfer.TransferInfo.timeCompleted:type_name -> google.protobuf.Timestamp
	3, // 1: transfer.OperationResponseList.operations:type_name -> transfer.TransferInfo
	1, // 2: transfer.TransferService.Transfer:input_type -> transfer.TransferRequest
	2, // 3: transfer.OperationService.OperationList:input_type -> transfer.OperationRequestList
	4, // 4: transfer.TransferService.Transfer:output_type -> transfer.TransferStatusResponse
	5, // 5: transfer.OperationService.OperationList:output_type -> transfer.OperationResponseList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_transfer_proto_init() }
func file_proto_transfer_proto_init() {
	if File_proto_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceUpdate); i {
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
		file_proto_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_proto_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationRequestList); i {
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
		file_proto_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferInfo); i {
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
		file_proto_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferStatusResponse); i {
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
		file_proto_transfer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationResponseList); i {
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
			RawDescriptor: file_proto_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_transfer_proto_goTypes,
		DependencyIndexes: file_proto_transfer_proto_depIdxs,
		MessageInfos:      file_proto_transfer_proto_msgTypes,
	}.Build()
	File_proto_transfer_proto = out.File
	file_proto_transfer_proto_rawDesc = nil
	file_proto_transfer_proto_goTypes = nil
	file_proto_transfer_proto_depIdxs = nil
}
