// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: vod/business/vod_callback.proto

package business

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_vod_business_vod_callback_proto protoreflect.FileDescriptor

var file_vod_business_vod_callback_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f,
	0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x42, 0xcf, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x6f, 0x64,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42,
	0x0b, 0x56, 0x6f, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x01, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f,
	0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x20, 0x56, 0x6f, 0x6c,
	0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x23,
	0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_vod_business_vod_callback_proto_goTypes = []interface{}{}
var file_vod_business_vod_callback_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vod_business_vod_callback_proto_init() }
func file_vod_business_vod_callback_proto_init() {
	if File_vod_business_vod_callback_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vod_business_vod_callback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_callback_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_callback_proto_depIdxs,
	}.Build()
	File_vod_business_vod_callback_proto = out.File
	file_vod_business_vod_callback_proto_rawDesc = nil
	file_vod_business_vod_callback_proto_goTypes = nil
	file_vod_business_vod_callback_proto_depIdxs = nil
}
