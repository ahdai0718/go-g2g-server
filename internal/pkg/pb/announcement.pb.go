// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: announcement.proto

package pb

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

type Announcement struct {
	state           protoimpl.MessageState
	Title           string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content         string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields   protoimpl.UnknownFields
	CreatedAt       int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	MarqueeShowTime int64 `protobuf:"varint,9,opt,name=marquee_show_time,json=marqueeShowTime,proto3" json:"marquee_show_time,omitempty"`
	sizeCache       protoimpl.SizeCache
	Status          int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	MarqueeStatus   int32 `protobuf:"varint,5,opt,name=marquee_status,json=marqueeStatus,proto3" json:"marquee_status,omitempty"`
	MarqueeTimes    int32 `protobuf:"varint,6,opt,name=marquee_times,json=marqueeTimes,proto3" json:"marquee_times,omitempty"`
	MarqueeCycle    int32 `protobuf:"varint,7,opt,name=marquee_cycle,json=marqueeCycle,proto3" json:"marquee_cycle,omitempty"`
	MarqueeInterval int32 `protobuf:"varint,8,opt,name=marquee_interval,json=marqueeInterval,proto3" json:"marquee_interval,omitempty"`
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announcement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_announcement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_announcement_proto_rawDescGZIP(), []int{0}
}

func (x *Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Announcement) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Announcement) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Announcement) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Announcement) GetMarqueeStatus() int32 {
	if x != nil {
		return x.MarqueeStatus
	}
	return 0
}

func (x *Announcement) GetMarqueeTimes() int32 {
	if x != nil {
		return x.MarqueeTimes
	}
	return 0
}

func (x *Announcement) GetMarqueeCycle() int32 {
	if x != nil {
		return x.MarqueeCycle
	}
	return 0
}

func (x *Announcement) GetMarqueeInterval() int32 {
	if x != nil {
		return x.MarqueeInterval
	}
	return 0
}

func (x *Announcement) GetMarqueeShowTime() int64 {
	if x != nil {
		return x.MarqueeShowTime
	}
	return 0
}

var File_announcement_proto protoreflect.FileDescriptor

var file_announcement_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xbd, 0x02, 0x0a, 0x0c, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x71, 0x75, 0x65,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x72, 0x71, 0x75,
	0x65, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x72,
	0x71, 0x75, 0x65, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x11,
	0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x72, 0x71, 0x75, 0x65, 0x65,
	0x53, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x68, 0x64, 0x61,
	0x64, 0x61, 0x2f, 0x67, 0x32, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_announcement_proto_rawDescOnce sync.Once
	file_announcement_proto_rawDescData = file_announcement_proto_rawDesc
)

func file_announcement_proto_rawDescGZIP() []byte {
	file_announcement_proto_rawDescOnce.Do(func() {
		file_announcement_proto_rawDescData = protoimpl.X.CompressGZIP(file_announcement_proto_rawDescData)
	})
	return file_announcement_proto_rawDescData
}

var file_announcement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_announcement_proto_goTypes = []interface{}{
	(*Announcement)(nil), // 0: pb.Announcement
}
var file_announcement_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_announcement_proto_init() }
func file_announcement_proto_init() {
	if File_announcement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_announcement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement); i {
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
			RawDescriptor: file_announcement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_announcement_proto_goTypes,
		DependencyIndexes: file_announcement_proto_depIdxs,
		MessageInfos:      file_announcement_proto_msgTypes,
	}.Build()
	File_announcement_proto = out.File
	file_announcement_proto_rawDesc = nil
	file_announcement_proto_goTypes = nil
	file_announcement_proto_depIdxs = nil
}
