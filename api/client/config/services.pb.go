// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: config/services.proto

package config

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

// Configuration for Spinnaker's microservices.
type Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clouddriver *Clouddriver `protobuf:"bytes,1,opt,name=clouddriver,proto3" json:"clouddriver,omitempty"`
	Echo        *Echo        `protobuf:"bytes,2,opt,name=echo,proto3" json:"echo,omitempty"`
	Front50     *Front50     `protobuf:"bytes,3,opt,name=front50,proto3" json:"front50,omitempty"`
	Orca        *Orca        `protobuf:"bytes,4,opt,name=orca,proto3" json:"orca,omitempty"`
	Gate        *Gate        `protobuf:"bytes,5,opt,name=gate,proto3" json:"gate,omitempty"`
	Fiat        *Fiat        `protobuf:"bytes,6,opt,name=fiat,proto3" json:"fiat,omitempty"`
	Kayenta     *Kayenta     `protobuf:"bytes,7,opt,name=kayenta,proto3" json:"kayenta,omitempty"`
	Rosco       *Rosco       `protobuf:"bytes,8,opt,name=rosco,proto3" json:"rosco,omitempty"`
	Deck        *Deck        `protobuf:"bytes,9,opt,name=deck,proto3" json:"deck,omitempty"`
	DeckEnv     *DeckEnv     `protobuf:"bytes,10,opt,name=deckEnv,proto3" json:"deckEnv,omitempty"`
	Igor        *Igor        `protobuf:"bytes,11,opt,name=igor,proto3" json:"igor,omitempty"`
	Monitoring  *Monitoring  `protobuf:"bytes,12,opt,name=monitoring,proto3" json:"monitoring,omitempty"`
}

func (x *Services) Reset() {
	*x = Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Services) ProtoMessage() {}

func (x *Services) ProtoReflect() protoreflect.Message {
	mi := &file_config_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Services.ProtoReflect.Descriptor instead.
func (*Services) Descriptor() ([]byte, []int) {
	return file_config_services_proto_rawDescGZIP(), []int{0}
}

func (x *Services) GetClouddriver() *Clouddriver {
	if x != nil {
		return x.Clouddriver
	}
	return nil
}

func (x *Services) GetEcho() *Echo {
	if x != nil {
		return x.Echo
	}
	return nil
}

func (x *Services) GetFront50() *Front50 {
	if x != nil {
		return x.Front50
	}
	return nil
}

func (x *Services) GetOrca() *Orca {
	if x != nil {
		return x.Orca
	}
	return nil
}

func (x *Services) GetGate() *Gate {
	if x != nil {
		return x.Gate
	}
	return nil
}

func (x *Services) GetFiat() *Fiat {
	if x != nil {
		return x.Fiat
	}
	return nil
}

func (x *Services) GetKayenta() *Kayenta {
	if x != nil {
		return x.Kayenta
	}
	return nil
}

func (x *Services) GetRosco() *Rosco {
	if x != nil {
		return x.Rosco
	}
	return nil
}

func (x *Services) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

func (x *Services) GetDeckEnv() *DeckEnv {
	if x != nil {
		return x.DeckEnv
	}
	return nil
}

func (x *Services) GetIgor() *Igor {
	if x != nil {
		return x.Igor
	}
	return nil
}

func (x *Services) GetMonitoring() *Monitoring {
	if x != nil {
		return x.Monitoring
	}
	return nil
}

var File_config_services_proto protoreflect.FileDescriptor

var file_config_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x5f,
	0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x35, 0x30, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x69, 0x67, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x73,
	0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x0b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x65,
	0x63, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x04, 0x65,
	0x63, 0x68, 0x6f, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x35, 0x30, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x35, 0x30, 0x52, 0x07, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x35, 0x30, 0x12, 0x26, 0x0a, 0x04, 0x6f, 0x72, 0x63, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x52, 0x04, 0x6f, 0x72, 0x63, 0x61, 0x12, 0x26, 0x0a, 0x04,
	0x67, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x67, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x69, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x46, 0x69, 0x61, 0x74, 0x52, 0x04, 0x66, 0x69, 0x61, 0x74, 0x12, 0x2f, 0x0a, 0x07,
	0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x61, 0x79,
	0x65, 0x6e, 0x74, 0x61, 0x52, 0x07, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x12, 0x29, 0x0a,
	0x05, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x6f, 0x73, 0x63,
	0x6f, 0x52, 0x05, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b,
	0x12, 0x2f, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x76, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x44, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x76, 0x52, 0x07, 0x64, 0x65, 0x63, 0x6b, 0x45, 0x6e,
	0x76, 0x12, 0x26, 0x0a, 0x04, 0x69, 0x67, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49,
	0x67, 0x6f, 0x72, 0x52, 0x04, 0x69, 0x67, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_services_proto_rawDescOnce sync.Once
	file_config_services_proto_rawDescData = file_config_services_proto_rawDesc
)

func file_config_services_proto_rawDescGZIP() []byte {
	file_config_services_proto_rawDescOnce.Do(func() {
		file_config_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_services_proto_rawDescData)
	})
	return file_config_services_proto_rawDescData
}

var file_config_services_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_services_proto_goTypes = []interface{}{
	(*Services)(nil),    // 0: proto.config.Services
	(*Clouddriver)(nil), // 1: proto.config.Clouddriver
	(*Echo)(nil),        // 2: proto.config.Echo
	(*Front50)(nil),     // 3: proto.config.Front50
	(*Orca)(nil),        // 4: proto.config.Orca
	(*Gate)(nil),        // 5: proto.config.Gate
	(*Fiat)(nil),        // 6: proto.config.Fiat
	(*Kayenta)(nil),     // 7: proto.config.Kayenta
	(*Rosco)(nil),       // 8: proto.config.Rosco
	(*Deck)(nil),        // 9: proto.config.Deck
	(*DeckEnv)(nil),     // 10: proto.config.DeckEnv
	(*Igor)(nil),        // 11: proto.config.Igor
	(*Monitoring)(nil),  // 12: proto.config.Monitoring
}
var file_config_services_proto_depIdxs = []int32{
	1,  // 0: proto.config.Services.clouddriver:type_name -> proto.config.Clouddriver
	2,  // 1: proto.config.Services.echo:type_name -> proto.config.Echo
	3,  // 2: proto.config.Services.front50:type_name -> proto.config.Front50
	4,  // 3: proto.config.Services.orca:type_name -> proto.config.Orca
	5,  // 4: proto.config.Services.gate:type_name -> proto.config.Gate
	6,  // 5: proto.config.Services.fiat:type_name -> proto.config.Fiat
	7,  // 6: proto.config.Services.kayenta:type_name -> proto.config.Kayenta
	8,  // 7: proto.config.Services.rosco:type_name -> proto.config.Rosco
	9,  // 8: proto.config.Services.deck:type_name -> proto.config.Deck
	10, // 9: proto.config.Services.deckEnv:type_name -> proto.config.DeckEnv
	11, // 10: proto.config.Services.igor:type_name -> proto.config.Igor
	12, // 11: proto.config.Services.monitoring:type_name -> proto.config.Monitoring
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_config_services_proto_init() }
func file_config_services_proto_init() {
	if File_config_services_proto != nil {
		return
	}
	file_config_clouddriver_proto_init()
	file_config_deck_proto_init()
	file_config_deck_env_proto_init()
	file_config_echo_proto_init()
	file_config_fiat_proto_init()
	file_config_front50_proto_init()
	file_config_gate_proto_init()
	file_config_igor_proto_init()
	file_config_kayenta_proto_init()
	file_config_orca_proto_init()
	file_config_rosco_proto_init()
	file_config_monitoring_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Services); i {
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
			RawDescriptor: file_config_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_services_proto_goTypes,
		DependencyIndexes: file_config_services_proto_depIdxs,
		MessageInfos:      file_config_services_proto_msgTypes,
	}.Build()
	File_config_services_proto = out.File
	file_config_services_proto_rawDesc = nil
	file_config_services_proto_goTypes = nil
	file_config_services_proto_depIdxs = nil
}
