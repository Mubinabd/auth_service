// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: restaurant_protos/retaurant.proto

package genproto

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

type CreateRestaurantReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *CreateRestaurantReq) Reset() {
	*x = CreateRestaurantReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_retaurant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRestaurantReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestaurantReq) ProtoMessage() {}

func (x *CreateRestaurantReq) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_retaurant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestaurantReq.ProtoReflect.Descriptor instead.
func (*CreateRestaurantReq) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_retaurant_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRestaurantReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRestaurantReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRestaurantReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateRestaurantReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateRestaurantReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_retaurant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_retaurant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_retaurant_proto_rawDescGZIP(), []int{1}
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Restaurant) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Restaurant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Restaurants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurants []*Restaurant `protobuf:"bytes,1,rep,name=Restaurants,proto3" json:"Restaurants,omitempty"`
}

func (x *Restaurants) Reset() {
	*x = Restaurants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_retaurant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurants) ProtoMessage() {}

func (x *Restaurants) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_retaurant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurants.ProtoReflect.Descriptor instead.
func (*Restaurants) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_retaurant_proto_rawDescGZIP(), []int{2}
}

func (x *Restaurants) GetRestaurants() []*Restaurant {
	if x != nil {
		return x.Restaurants
	}
	return nil
}

type AddressFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *AddressFilter) Reset() {
	*x = AddressFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_retaurant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressFilter) ProtoMessage() {}

func (x *AddressFilter) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_retaurant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressFilter.ProtoReflect.Descriptor instead.
func (*AddressFilter) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_retaurant_proto_rawDescGZIP(), []int{3}
}

func (x *AddressFilter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_restaurant_protos_retaurant_proto protoreflect.FileDescriptor

var file_restaurant_protos_retaurant_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x7e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x4e, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12,
	0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73,
	0x22, 0x29, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xad, 0x03, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x55, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x57, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restaurant_protos_retaurant_proto_rawDescOnce sync.Once
	file_restaurant_protos_retaurant_proto_rawDescData = file_restaurant_protos_retaurant_proto_rawDesc
)

func file_restaurant_protos_retaurant_proto_rawDescGZIP() []byte {
	file_restaurant_protos_retaurant_proto_rawDescOnce.Do(func() {
		file_restaurant_protos_retaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_protos_retaurant_proto_rawDescData)
	})
	return file_restaurant_protos_retaurant_proto_rawDescData
}

var file_restaurant_protos_retaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_restaurant_protos_retaurant_proto_goTypes = []interface{}{
	(*CreateRestaurantReq)(nil), // 0: restaurant_protos.CreateRestaurantReq
	(*Restaurant)(nil),          // 1: restaurant_protos.Restaurant
	(*Restaurants)(nil),         // 2: restaurant_protos.Restaurants
	(*AddressFilter)(nil),       // 3: restaurant_protos.AddressFilter
	(*ById)(nil),                // 4: restaurant_protos.ById
	(*Void)(nil),                // 5: restaurant_protos.Void
}
var file_restaurant_protos_retaurant_proto_depIdxs = []int32{
	1, // 0: restaurant_protos.Restaurants.Restaurants:type_name -> restaurant_protos.Restaurant
	0, // 1: restaurant_protos.RestaurantService.CreateRestaurant:input_type -> restaurant_protos.CreateRestaurantReq
	0, // 2: restaurant_protos.RestaurantService.UpdateRestaurant:input_type -> restaurant_protos.CreateRestaurantReq
	4, // 3: restaurant_protos.RestaurantService.DeleteRestaurant:input_type -> restaurant_protos.ById
	4, // 4: restaurant_protos.RestaurantService.GetRestaurant:input_type -> restaurant_protos.ById
	3, // 5: restaurant_protos.RestaurantService.GetAllRestaurants:input_type -> restaurant_protos.AddressFilter
	5, // 6: restaurant_protos.RestaurantService.CreateRestaurant:output_type -> restaurant_protos.Void
	5, // 7: restaurant_protos.RestaurantService.UpdateRestaurant:output_type -> restaurant_protos.Void
	5, // 8: restaurant_protos.RestaurantService.DeleteRestaurant:output_type -> restaurant_protos.Void
	1, // 9: restaurant_protos.RestaurantService.GetRestaurant:output_type -> restaurant_protos.Restaurant
	2, // 10: restaurant_protos.RestaurantService.GetAllRestaurants:output_type -> restaurant_protos.Restaurants
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_restaurant_protos_retaurant_proto_init() }
func file_restaurant_protos_retaurant_proto_init() {
	if File_restaurant_protos_retaurant_proto != nil {
		return
	}
	file_restaurant_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_restaurant_protos_retaurant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRestaurantReq); i {
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
		file_restaurant_protos_retaurant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_restaurant_protos_retaurant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurants); i {
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
		file_restaurant_protos_retaurant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressFilter); i {
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
			RawDescriptor: file_restaurant_protos_retaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restaurant_protos_retaurant_proto_goTypes,
		DependencyIndexes: file_restaurant_protos_retaurant_proto_depIdxs,
		MessageInfos:      file_restaurant_protos_retaurant_proto_msgTypes,
	}.Build()
	File_restaurant_protos_retaurant_proto = out.File
	file_restaurant_protos_retaurant_proto_rawDesc = nil
	file_restaurant_protos_retaurant_proto_goTypes = nil
	file_restaurant_protos_retaurant_proto_depIdxs = nil
}
