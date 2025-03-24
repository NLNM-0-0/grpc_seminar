// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: seminar/order/v1/order.proto

package orderv1

import (
	_ "gateway/gen/google/api"
	v1 "gateway/gen/seminar/product/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     []byte                 `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Products      []*Order_OrderProduct  `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetCreatedAt() []byte {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetProducts() []*Order_OrderProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type PostOrderRequest struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	UserId        string                           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Products      []*PostOrderRequest_OrderProduct `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostOrderRequest) Reset() {
	*x = PostOrderRequest{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOrderRequest) ProtoMessage() {}

func (x *PostOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOrderRequest.ProtoReflect.Descriptor instead.
func (*PostOrderRequest) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *PostOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostOrderRequest) GetProducts() []*PostOrderRequest_OrderProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type PostOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostOrderResponse) Reset() {
	*x = PostOrderResponse{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOrderResponse) ProtoMessage() {}

func (x *PostOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOrderResponse.ProtoReflect.Descriptor instead.
func (*PostOrderResponse) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *PostOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type RateOrderRequest struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Products      []*v1.RateProductByIdsRequest `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateOrderRequest) Reset() {
	*x = RateOrderRequest{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateOrderRequest) ProtoMessage() {}

func (x *RateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateOrderRequest.ProtoReflect.Descriptor instead.
func (*RateOrderRequest) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{5}
}

func (x *RateOrderRequest) GetProducts() []*v1.RateProductByIdsRequest {
	if x != nil {
		return x.Products
	}
	return nil
}

type RateOrderResponse struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Ratings       []*v1.RateProductByIdsResponse `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateOrderResponse) Reset() {
	*x = RateOrderResponse{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateOrderResponse) ProtoMessage() {}

func (x *RateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateOrderResponse.ProtoReflect.Descriptor instead.
func (*RateOrderResponse) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{6}
}

func (x *RateOrderResponse) GetRatings() []*v1.RateProductByIdsResponse {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type Order_OrderProduct struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      uint32                 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order_OrderProduct) Reset() {
	*x = Order_OrderProduct{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order_OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order_OrderProduct) ProtoMessage() {}

func (x *Order_OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order_OrderProduct.ProtoReflect.Descriptor instead.
func (*Order_OrderProduct) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Order_OrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Order_OrderProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Order_OrderProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order_OrderProduct) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type PostOrderRequest_OrderProduct struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      uint32                 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostOrderRequest_OrderProduct) Reset() {
	*x = PostOrderRequest_OrderProduct{}
	mi := &file_seminar_order_v1_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostOrderRequest_OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOrderRequest_OrderProduct) ProtoMessage() {}

func (x *PostOrderRequest_OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_seminar_order_v1_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOrderRequest_OrderProduct.ProtoReflect.Descriptor instead.
func (*PostOrderRequest_OrderProduct) Descriptor() ([]byte, []int) {
	return file_seminar_order_v1_order_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PostOrderRequest_OrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *PostOrderRequest_OrderProduct) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_seminar_order_v1_order_proto protoreflect.FileDescriptor

var file_seminar_order_v1_order_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa7, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x73, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xc3, 0x01, 0x0a, 0x10, 0x50,
	0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x65, 0x6d,
	0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x49, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x42, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x6d,
	0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x10, 0x52, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x5b, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x32, 0xed, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x70, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x22, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x42, 0x9e, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x6d,
	0x69, 0x6e, 0x61, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x73, 0x65, 0x6d,
	0x69, 0x6e, 0x61, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x4f, 0x58, 0xaa, 0x02, 0x10, 0x53, 0x65,
	0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x10, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1c, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x5c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x12, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x3a, 0x3a, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_seminar_order_v1_order_proto_rawDescOnce sync.Once
	file_seminar_order_v1_order_proto_rawDescData []byte
)

func file_seminar_order_v1_order_proto_rawDescGZIP() []byte {
	file_seminar_order_v1_order_proto_rawDescOnce.Do(func() {
		file_seminar_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_seminar_order_v1_order_proto_rawDesc), len(file_seminar_order_v1_order_proto_rawDesc)))
	})
	return file_seminar_order_v1_order_proto_rawDescData
}

var file_seminar_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_seminar_order_v1_order_proto_goTypes = []any{
	(*Order)(nil),                         // 0: seminar.order.v1.Order
	(*PostOrderRequest)(nil),              // 1: seminar.order.v1.PostOrderRequest
	(*PostOrderResponse)(nil),             // 2: seminar.order.v1.PostOrderResponse
	(*GetOrderRequest)(nil),               // 3: seminar.order.v1.GetOrderRequest
	(*GetOrderResponse)(nil),              // 4: seminar.order.v1.GetOrderResponse
	(*RateOrderRequest)(nil),              // 5: seminar.order.v1.RateOrderRequest
	(*RateOrderResponse)(nil),             // 6: seminar.order.v1.RateOrderResponse
	(*Order_OrderProduct)(nil),            // 7: seminar.order.v1.Order.OrderProduct
	(*PostOrderRequest_OrderProduct)(nil), // 8: seminar.order.v1.PostOrderRequest.OrderProduct
	(*v1.RateProductByIdsRequest)(nil),    // 9: seminar.product.v1.RateProductByIdsRequest
	(*v1.RateProductByIdsResponse)(nil),   // 10: seminar.product.v1.RateProductByIdsResponse
}
var file_seminar_order_v1_order_proto_depIdxs = []int32{
	7,  // 0: seminar.order.v1.Order.products:type_name -> seminar.order.v1.Order.OrderProduct
	8,  // 1: seminar.order.v1.PostOrderRequest.products:type_name -> seminar.order.v1.PostOrderRequest.OrderProduct
	0,  // 2: seminar.order.v1.PostOrderResponse.order:type_name -> seminar.order.v1.Order
	0,  // 3: seminar.order.v1.GetOrderResponse.order:type_name -> seminar.order.v1.Order
	9,  // 4: seminar.order.v1.RateOrderRequest.products:type_name -> seminar.product.v1.RateProductByIdsRequest
	10, // 5: seminar.order.v1.RateOrderResponse.ratings:type_name -> seminar.product.v1.RateProductByIdsResponse
	1,  // 6: seminar.order.v1.OrderService.PostOrder:input_type -> seminar.order.v1.PostOrderRequest
	5,  // 7: seminar.order.v1.OrderService.RateOrder:input_type -> seminar.order.v1.RateOrderRequest
	2,  // 8: seminar.order.v1.OrderService.PostOrder:output_type -> seminar.order.v1.PostOrderResponse
	6,  // 9: seminar.order.v1.OrderService.RateOrder:output_type -> seminar.order.v1.RateOrderResponse
	8,  // [8:10] is the sub-list for method output_type
	6,  // [6:8] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_seminar_order_v1_order_proto_init() }
func file_seminar_order_v1_order_proto_init() {
	if File_seminar_order_v1_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_seminar_order_v1_order_proto_rawDesc), len(file_seminar_order_v1_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seminar_order_v1_order_proto_goTypes,
		DependencyIndexes: file_seminar_order_v1_order_proto_depIdxs,
		MessageInfos:      file_seminar_order_v1_order_proto_msgTypes,
	}.Build()
	File_seminar_order_v1_order_proto = out.File
	file_seminar_order_v1_order_proto_goTypes = nil
	file_seminar_order_v1_order_proto_depIdxs = nil
}
