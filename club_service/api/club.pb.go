// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: club.proto

package api

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[0]
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
	return file_club_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Club struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	OpenTime string   `protobuf:"bytes,3,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
	Currency string   `protobuf:"bytes,4,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Prices   []*Price `protobuf:"bytes,5,rep,name=Prices,proto3" json:"Prices,omitempty"`
}

func (x *Club) Reset() {
	*x = Club{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Club) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Club) ProtoMessage() {}

func (x *Club) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Club.ProtoReflect.Descriptor instead.
func (*Club) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{1}
}

func (x *Club) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Club) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Club) GetOpenTime() string {
	if x != nil {
		return x.OpenTime
	}
	return ""
}

func (x *Club) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Club) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PricePeriod    *PricePeriod `protobuf:"bytes,1,opt,name=PricePeriod,proto3" json:"PricePeriod,omitempty"`
	ValuePerMinute float32      `protobuf:"fixed32,2,opt,name=ValuePerMinute,proto3" json:"ValuePerMinute,omitempty"`
	Currency       int64        `protobuf:"varint,3,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{2}
}

func (x *Price) GetPricePeriod() *PricePeriod {
	if x != nil {
		return x.PricePeriod
	}
	return nil
}

func (x *Price) GetValuePerMinute() float32 {
	if x != nil {
		return x.ValuePerMinute
	}
	return 0
}

func (x *Price) GetCurrency() int64 {
	if x != nil {
		return x.Currency
	}
	return 0
}

type PricePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int32 `protobuf:"varint,1,opt,name=From,proto3" json:"From,omitempty"`
	To   int32 `protobuf:"varint,2,opt,name=To,proto3" json:"To,omitempty"`
}

func (x *PricePeriod) Reset() {
	*x = PricePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricePeriod) ProtoMessage() {}

func (x *PricePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricePeriod.ProtoReflect.Descriptor instead.
func (*PricePeriod) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{3}
}

func (x *PricePeriod) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PricePeriod) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[4]
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
	return file_club_proto_rawDescGZIP(), []int{4}
}

var File_club_proto protoreflect.FileDescriptor

var file_club_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a,
	0x04, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x70, 0x65,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x22, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x7f, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x65, 0x72, 0x4d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x31, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x54, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x33, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x24, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6c, 0x75, 0x62, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_club_proto_rawDescOnce sync.Once
	file_club_proto_rawDescData = file_club_proto_rawDesc
)

func file_club_proto_rawDescGZIP() []byte {
	file_club_proto_rawDescOnce.Do(func() {
		file_club_proto_rawDescData = protoimpl.X.CompressGZIP(file_club_proto_rawDescData)
	})
	return file_club_proto_rawDescData
}

var file_club_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_club_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: api.Request
	(*Club)(nil),        // 1: api.Club
	(*Price)(nil),       // 2: api.Price
	(*PricePeriod)(nil), // 3: api.PricePeriod
	(*Empty)(nil),       // 4: api.Empty
}
var file_club_proto_depIdxs = []int32{
	2, // 0: api.Club.Prices:type_name -> api.Price
	3, // 1: api.Price.PricePeriod:type_name -> api.PricePeriod
	0, // 2: api.ClubService.GetById:input_type -> api.Request
	1, // 3: api.ClubService.GetById:output_type -> api.Club
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_club_proto_init() }
func file_club_proto_init() {
	if File_club_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_club_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_club_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Club); i {
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
		file_club_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_club_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricePeriod); i {
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
		file_club_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_club_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_club_proto_goTypes,
		DependencyIndexes: file_club_proto_depIdxs,
		MessageInfos:      file_club_proto_msgTypes,
	}.Build()
	File_club_proto = out.File
	file_club_proto_rawDesc = nil
	file_club_proto_goTypes = nil
	file_club_proto_depIdxs = nil
}
