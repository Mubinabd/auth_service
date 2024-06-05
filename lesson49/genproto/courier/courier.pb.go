// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: courier.proto

package courier

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

type TakeOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsPaid  bool   `protobuf:"varint,3,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TakeOrder) Reset() {
	*x = TakeOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeOrder) ProtoMessage() {}

func (x *TakeOrder) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeOrder.ProtoReflect.Descriptor instead.
func (*TakeOrder) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{0}
}

func (x *TakeOrder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TakeOrder) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *TakeOrder) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DeliverOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     string     `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	IsHotYet bool       `protobuf:"varint,3,opt,name=is_hot_yet,json=isHotYet,proto3" json:"is_hot_yet,omitempty"`
}

func (x *DeliverOrder) Reset() {
	*x = DeliverOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverOrder) ProtoMessage() {}

func (x *DeliverOrder) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverOrder.ProtoReflect.Descriptor instead.
func (*DeliverOrder) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{1}
}

func (x *DeliverOrder) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *DeliverOrder) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *DeliverOrder) GetIsHotYet() bool {
	if x != nil {
		return x.IsHotYet
	}
	return false
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_courier_proto protoreflect.FileDescriptor

var file_courier_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x09, 0x54, 0x61, 0x6b, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x70, 0x61, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1c,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x68, 0x6f, 0x74, 0x5f, 0x79, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x59, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x48, 0x0a, 0x0e, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x07, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_courier_proto_rawDescOnce sync.Once
	file_courier_proto_rawDescData = file_courier_proto_rawDesc
)

func file_courier_proto_rawDescGZIP() []byte {
	file_courier_proto_rawDescOnce.Do(func() {
		file_courier_proto_rawDescData = protoimpl.X.CompressGZIP(file_courier_proto_rawDescData)
	})
	return file_courier_proto_rawDescData
}

var file_courier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_courier_proto_goTypes = []interface{}{
	(*TakeOrder)(nil),    // 0: courier.TakeOrder
	(*DeliverOrder)(nil), // 1: courier.DeliverOrder
	(*Product)(nil),      // 2: courier.Product
}
var file_courier_proto_depIdxs = []int32{
	2, // 0: courier.DeliverOrder.products:type_name -> courier.Product
	0, // 1: courier.CourierService.Deliver:input_type -> courier.TakeOrder
	1, // 2: courier.CourierService.Deliver:output_type -> courier.DeliverOrder
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_courier_proto_init() }
func file_courier_proto_init() {
	if File_courier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_courier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeOrder); i {
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
		file_courier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverOrder); i {
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
		file_courier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_courier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courier_proto_goTypes,
		DependencyIndexes: file_courier_proto_depIdxs,
		MessageInfos:      file_courier_proto_msgTypes,
	}.Build()
	File_courier_proto = out.File
	file_courier_proto_rawDesc = nil
	file_courier_proto_goTypes = nil
	file_courier_proto_depIdxs = nil
}
