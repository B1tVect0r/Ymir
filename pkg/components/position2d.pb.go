// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0-devel
// 	protoc        v3.19.1
// source: pkg/components/position2d.proto

package components

import (
	vector2 "github.com/B1tVect0r/ymir/pkg/types/math/vector2"
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

type Position2D struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *vector2.Vector2 `protobuf:"bytes,1,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (x *Position2D) Reset() {
	*x = Position2D{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_components_position2d_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position2D) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position2D) ProtoMessage() {}

func (x *Position2D) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_components_position2d_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position2D.ProtoReflect.Descriptor instead.
func (*Position2D) Descriptor() ([]byte, []int) {
	return file_pkg_components_position2d_proto_rawDescGZIP(), []int{0}
}

func (x *Position2D) GetPosition() *vector2.Vector2 {
	if x != nil {
		return x.Position
	}
	return nil
}

var File_pkg_components_position2d_proto protoreflect.FileDescriptor

var file_pkg_components_position2d_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x24, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x32, 0x2f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x44, 0x12, 0x2c, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x31,
	0x74, 0x56, 0x65, 0x63, 0x74, 0x30, 0x72, 0x2f, 0x79, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_components_position2d_proto_rawDescOnce sync.Once
	file_pkg_components_position2d_proto_rawDescData = file_pkg_components_position2d_proto_rawDesc
)

func file_pkg_components_position2d_proto_rawDescGZIP() []byte {
	file_pkg_components_position2d_proto_rawDescOnce.Do(func() {
		file_pkg_components_position2d_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_components_position2d_proto_rawDescData)
	})
	return file_pkg_components_position2d_proto_rawDescData
}

var file_pkg_components_position2d_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_components_position2d_proto_goTypes = []interface{}{
	(*Position2D)(nil),      // 0: components.Position2D
	(*vector2.Vector2)(nil), // 1: vector2.Vector2
}
var file_pkg_components_position2d_proto_depIdxs = []int32{
	1, // 0: components.Position2D.Position:type_name -> vector2.Vector2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_components_position2d_proto_init() }
func file_pkg_components_position2d_proto_init() {
	if File_pkg_components_position2d_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_components_position2d_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position2D); i {
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
			RawDescriptor: file_pkg_components_position2d_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_components_position2d_proto_goTypes,
		DependencyIndexes: file_pkg_components_position2d_proto_depIdxs,
		MessageInfos:      file_pkg_components_position2d_proto_msgTypes,
	}.Build()
	File_pkg_components_position2d_proto = out.File
	file_pkg_components_position2d_proto_rawDesc = nil
	file_pkg_components_position2d_proto_goTypes = nil
	file_pkg_components_position2d_proto_depIdxs = nil
}
