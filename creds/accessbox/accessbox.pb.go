// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: creds/accessbox/accessbox.proto

package accessbox

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

type AccessBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerPublicKey []byte            `protobuf:"bytes,1,opt,name=ownerPublicKey,proto3" json:"ownerPublicKey,omitempty"`
	Gates          []*AccessBox_Gate `protobuf:"bytes,2,rep,name=gates,proto3" json:"gates,omitempty"`
}

func (x *AccessBox) Reset() {
	*x = AccessBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creds_accessbox_accessbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessBox) ProtoMessage() {}

func (x *AccessBox) ProtoReflect() protoreflect.Message {
	mi := &file_creds_accessbox_accessbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessBox.ProtoReflect.Descriptor instead.
func (*AccessBox) Descriptor() ([]byte, []int) {
	return file_creds_accessbox_accessbox_proto_rawDescGZIP(), []int{0}
}

func (x *AccessBox) GetOwnerPublicKey() []byte {
	if x != nil {
		return x.OwnerPublicKey
	}
	return nil
}

func (x *AccessBox) GetGates() []*AccessBox_Gate {
	if x != nil {
		return x.Gates
	}
	return nil
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey    []byte `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	BearerToken  []byte `protobuf:"bytes,2,opt,name=bearerToken,proto3" json:"bearerToken,omitempty"`
	SessionToken []byte `protobuf:"bytes,3,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creds_accessbox_accessbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_creds_accessbox_accessbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_creds_accessbox_accessbox_proto_rawDescGZIP(), []int{1}
}

func (x *Tokens) GetAccessKey() []byte {
	if x != nil {
		return x.AccessKey
	}
	return nil
}

func (x *Tokens) GetBearerToken() []byte {
	if x != nil {
		return x.BearerToken
	}
	return nil
}

func (x *Tokens) GetSessionToken() []byte {
	if x != nil {
		return x.SessionToken
	}
	return nil
}

type AccessBox_Gate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens        []byte `protobuf:"bytes,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
	GatePublicKey []byte `protobuf:"bytes,2,opt,name=gatePublicKey,proto3" json:"gatePublicKey,omitempty"`
}

func (x *AccessBox_Gate) Reset() {
	*x = AccessBox_Gate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creds_accessbox_accessbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessBox_Gate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessBox_Gate) ProtoMessage() {}

func (x *AccessBox_Gate) ProtoReflect() protoreflect.Message {
	mi := &file_creds_accessbox_accessbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessBox_Gate.ProtoReflect.Descriptor instead.
func (*AccessBox_Gate) Descriptor() ([]byte, []int) {
	return file_creds_accessbox_accessbox_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AccessBox_Gate) GetTokens() []byte {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *AccessBox_Gate) GetGatePublicKey() []byte {
	if x != nil {
		return x.GatePublicKey
	}
	return nil
}

var File_creds_accessbox_accessbox_proto protoreflect.FileDescriptor

var file_creds_accessbox_accessbox_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x62, 0x6f,
	0x78, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x78, 0x22, 0xaa, 0x01, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x67, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x78, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x78, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x52, 0x05, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x1a, 0x44, 0x0a, 0x04, 0x47, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x67, 0x61, 0x74, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x6c, 0x0a, 0x06, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d, 0x73, 0x33, 0x2d, 0x67, 0x77, 0x2f, 0x63, 0x72, 0x65, 0x64,
	0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x6f, 0x78, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creds_accessbox_accessbox_proto_rawDescOnce sync.Once
	file_creds_accessbox_accessbox_proto_rawDescData = file_creds_accessbox_accessbox_proto_rawDesc
)

func file_creds_accessbox_accessbox_proto_rawDescGZIP() []byte {
	file_creds_accessbox_accessbox_proto_rawDescOnce.Do(func() {
		file_creds_accessbox_accessbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_creds_accessbox_accessbox_proto_rawDescData)
	})
	return file_creds_accessbox_accessbox_proto_rawDescData
}

var file_creds_accessbox_accessbox_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_creds_accessbox_accessbox_proto_goTypes = []interface{}{
	(*AccessBox)(nil),      // 0: accessbox.AccessBox
	(*Tokens)(nil),         // 1: accessbox.Tokens
	(*AccessBox_Gate)(nil), // 2: accessbox.AccessBox.Gate
}
var file_creds_accessbox_accessbox_proto_depIdxs = []int32{
	2, // 0: accessbox.AccessBox.gates:type_name -> accessbox.AccessBox.Gate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_creds_accessbox_accessbox_proto_init() }
func file_creds_accessbox_accessbox_proto_init() {
	if File_creds_accessbox_accessbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_creds_accessbox_accessbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessBox); i {
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
		file_creds_accessbox_accessbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
		file_creds_accessbox_accessbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessBox_Gate); i {
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
			RawDescriptor: file_creds_accessbox_accessbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creds_accessbox_accessbox_proto_goTypes,
		DependencyIndexes: file_creds_accessbox_accessbox_proto_depIdxs,
		MessageInfos:      file_creds_accessbox_accessbox_proto_msgTypes,
	}.Build()
	File_creds_accessbox_accessbox_proto = out.File
	file_creds_accessbox_accessbox_proto_rawDesc = nil
	file_creds_accessbox_accessbox_proto_goTypes = nil
	file_creds_accessbox_accessbox_proto_depIdxs = nil
}
