// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: proto/shippingservice.proto

package shippingService

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

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// 当前币种描述
type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 货币code 例如：EUR欧元 USD美元
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// 货币单位
	// 例如 currency_code= USD，则1个单位是1美元
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// 数量的纳米（10^-9）单位。
	// 该值必须介于-99999999和+99999999之间。
	// 如果单位为正，纳米必须为正或零。
	// 如果单位为零，纳米可以为正、零或负。
	// 如果单位为负，纳米必须为负或零。
	// 例如，$-1.75表示为 单位＝-1 和 纳米＝-75000000。
	Nanos int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{1}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type GetQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Items   []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetQuoteRequest) Reset() {
	*x = GetQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteRequest) ProtoMessage() {}

func (x *GetQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteRequest.ProtoReflect.Descriptor instead.
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetQuoteRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetQuoteRequest) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostUsd *Money `protobuf:"bytes,1,opt,name=cost_usd,json=costUsd,proto3" json:"cost_usd,omitempty"`
}

func (x *GetQuoteResponse) Reset() {
	*x = GetQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteResponse) ProtoMessage() {}

func (x *GetQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteResponse.ProtoReflect.Descriptor instead.
func (*GetQuoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetQuoteResponse) GetCostUsd() *Money {
	if x != nil {
		return x.CostUsd
	}
	return nil
}

type ShipOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Items   []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ShipOrderRequest) Reset() {
	*x = ShipOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderRequest) ProtoMessage() {}

func (x *ShipOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderRequest.ProtoReflect.Descriptor instead.
func (*ShipOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{5}
}

func (x *ShipOrderRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ShipOrderRequest) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ShipOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingId string `protobuf:"bytes,1,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
}

func (x *ShipOrderResponse) Reset() {
	*x = ShipOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shippingservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderResponse) ProtoMessage() {}

func (x *ShipOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shippingservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderResponse.ProtoReflect.Descriptor instead.
func (*ShipOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_shippingservice_proto_rawDescGZIP(), []int{6}
}

func (x *ShipOrderResponse) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

var File_proto_shippingservice_proto protoreflect.FileDescriptor

var file_proto_shippingservice_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x45,
	0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e,
	0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22,
	0x8f, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x08, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x64,
	0x22, 0x77, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x11, 0x53, 0x68, 0x69,
	0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x32,
	0xba, 0x01, 0x0a, 0x0f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x20, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x09, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shippingservice_proto_rawDescOnce sync.Once
	file_proto_shippingservice_proto_rawDescData = file_proto_shippingservice_proto_rawDesc
)

func file_proto_shippingservice_proto_rawDescGZIP() []byte {
	file_proto_shippingservice_proto_rawDescOnce.Do(func() {
		file_proto_shippingservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shippingservice_proto_rawDescData)
	})
	return file_proto_shippingservice_proto_rawDescData
}

var file_proto_shippingservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_shippingservice_proto_goTypes = []interface{}{
	(*CartItem)(nil),          // 0: shippingService.CartItem
	(*Money)(nil),             // 1: shippingService.Money
	(*Address)(nil),           // 2: shippingService.Address
	(*GetQuoteRequest)(nil),   // 3: shippingService.GetQuoteRequest
	(*GetQuoteResponse)(nil),  // 4: shippingService.GetQuoteResponse
	(*ShipOrderRequest)(nil),  // 5: shippingService.ShipOrderRequest
	(*ShipOrderResponse)(nil), // 6: shippingService.ShipOrderResponse
}
var file_proto_shippingservice_proto_depIdxs = []int32{
	2, // 0: shippingService.GetQuoteRequest.address:type_name -> shippingService.Address
	0, // 1: shippingService.GetQuoteRequest.items:type_name -> shippingService.CartItem
	1, // 2: shippingService.GetQuoteResponse.cost_usd:type_name -> shippingService.Money
	2, // 3: shippingService.ShipOrderRequest.address:type_name -> shippingService.Address
	0, // 4: shippingService.ShipOrderRequest.items:type_name -> shippingService.CartItem
	3, // 5: shippingService.shippingService.GetQuote:input_type -> shippingService.GetQuoteRequest
	5, // 6: shippingService.shippingService.ShipOrder:input_type -> shippingService.ShipOrderRequest
	4, // 7: shippingService.shippingService.GetQuote:output_type -> shippingService.GetQuoteResponse
	6, // 8: shippingService.shippingService.ShipOrder:output_type -> shippingService.ShipOrderResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_shippingservice_proto_init() }
func file_proto_shippingservice_proto_init() {
	if File_proto_shippingservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shippingservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_proto_shippingservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
		file_proto_shippingservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_proto_shippingservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteRequest); i {
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
		file_proto_shippingservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteResponse); i {
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
		file_proto_shippingservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderRequest); i {
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
		file_proto_shippingservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderResponse); i {
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
			RawDescriptor: file_proto_shippingservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shippingservice_proto_goTypes,
		DependencyIndexes: file_proto_shippingservice_proto_depIdxs,
		MessageInfos:      file_proto_shippingservice_proto_msgTypes,
	}.Build()
	File_proto_shippingservice_proto = out.File
	file_proto_shippingservice_proto_rawDesc = nil
	file_proto_shippingservice_proto_goTypes = nil
	file_proto_shippingservice_proto_depIdxs = nil
}