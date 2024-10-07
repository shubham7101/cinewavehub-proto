// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: enums.proto

package gen

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

type MediaType int32

const (
	MediaType_NONE_MEDIATYPE MediaType = 0
	MediaType_ANIME          MediaType = 1
	MediaType_MANGA          MediaType = 2
	MediaType_MOVIE          MediaType = 3
	MediaType_TV             MediaType = 4
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "NONE_MEDIATYPE",
		1: "ANIME",
		2: "MANGA",
		3: "MOVIE",
		4: "TV",
	}
	MediaType_value = map[string]int32{
		"NONE_MEDIATYPE": 0,
		"ANIME":          1,
		"MANGA":          2,
		"MOVIE":          3,
		"TV":             4,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_NO_STATUS              Status = 0
	Status_STATUS_FINISHEDAIRING  Status = 1
	Status_STATUS_CURRENTLYAIRING Status = 2
	Status_STATUS_NOTYETAIRED     Status = 3
	Status_STATUS_CANCELLED       Status = 4
	Status_STATUS_HIATUS          Status = 5
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "NO_STATUS",
		1: "STATUS_FINISHEDAIRING",
		2: "STATUS_CURRENTLYAIRING",
		3: "STATUS_NOTYETAIRED",
		4: "STATUS_CANCELLED",
		5: "STATUS_HIATUS",
	}
	Status_value = map[string]int32{
		"NO_STATUS":              0,
		"STATUS_FINISHEDAIRING":  1,
		"STATUS_CURRENTLYAIRING": 2,
		"STATUS_NOTYETAIRED":     3,
		"STATUS_CANCELLED":       4,
		"STATUS_HIATUS":          5,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

type Format int32

const (
	Format_NO_FORMAT         Format = 0
	Format_FORMAT_MOVIE      Format = 1
	Format_FORMAT_TV         Format = 2
	Format_FORMAT_OVA        Format = 3
	Format_FORMAT_ONA        Format = 4
	Format_FORMAT_SPECIAL    Format = 5
	Format_FORMAT_MUSIC      Format = 6
	Format_FORMAT_COMMERCIAL Format = 7
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "NO_FORMAT",
		1: "FORMAT_MOVIE",
		2: "FORMAT_TV",
		3: "FORMAT_OVA",
		4: "FORMAT_ONA",
		5: "FORMAT_SPECIAL",
		6: "FORMAT_MUSIC",
		7: "FORMAT_COMMERCIAL",
	}
	Format_value = map[string]int32{
		"NO_FORMAT":         0,
		"FORMAT_MOVIE":      1,
		"FORMAT_TV":         2,
		"FORMAT_OVA":        3,
		"FORMAT_ONA":        4,
		"FORMAT_SPECIAL":    5,
		"FORMAT_MUSIC":      6,
		"FORMAT_COMMERCIAL": 7,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[2].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[2]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{2}
}

type Rating int32

const (
	Rating_NO_RATING    Rating = 0
	Rating_RATING_G     Rating = 1
	Rating_RATING_PG    Rating = 2
	Rating_RATING_PG13  Rating = 3
	Rating_RATING_R     Rating = 4
	Rating_RATING_RPLUS Rating = 5
	Rating_RATING_RX    Rating = 6
)

// Enum value maps for Rating.
var (
	Rating_name = map[int32]string{
		0: "NO_RATING",
		1: "RATING_G",
		2: "RATING_PG",
		3: "RATING_PG13",
		4: "RATING_R",
		5: "RATING_RPLUS",
		6: "RATING_RX",
	}
	Rating_value = map[string]int32{
		"NO_RATING":    0,
		"RATING_G":     1,
		"RATING_PG":    2,
		"RATING_PG13":  3,
		"RATING_R":     4,
		"RATING_RPLUS": 5,
		"RATING_RX":    6,
	}
)

func (x Rating) Enum() *Rating {
	p := new(Rating)
	*p = x
	return p
}

func (x Rating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rating) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[3].Descriptor()
}

func (Rating) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[3]
}

func (x Rating) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rating.Descriptor instead.
func (Rating) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{3}
}

type Season int32

const (
	Season_NO_SEASON     Season = 0
	Season_SEASON_SPRING Season = 1
	Season_SEASON_SUMMER Season = 2
	Season_SEASON_FALL   Season = 3
	Season_SEASON_WINTER Season = 4
)

// Enum value maps for Season.
var (
	Season_name = map[int32]string{
		0: "NO_SEASON",
		1: "SEASON_SPRING",
		2: "SEASON_SUMMER",
		3: "SEASON_FALL",
		4: "SEASON_WINTER",
	}
	Season_value = map[string]int32{
		"NO_SEASON":     0,
		"SEASON_SPRING": 1,
		"SEASON_SUMMER": 2,
		"SEASON_FALL":   3,
		"SEASON_WINTER": 4,
	}
)

func (x Season) Enum() *Season {
	p := new(Season)
	*p = x
	return p
}

func (x Season) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Season) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[4].Descriptor()
}

func (Season) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[4]
}

func (x Season) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Season.Descriptor instead.
func (Season) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{4}
}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63,
	0x69, 0x6e, 0x65, 0x77, 0x61, 0x76, 0x65, 0x68, 0x75, 0x62, 0x2a, 0x48, 0x0a, 0x09, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x4e, 0x45, 0x5f,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x4e, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x4e, 0x47, 0x41, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x56, 0x49, 0x45, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02,
	0x54, 0x56, 0x10, 0x04, 0x2a, 0x8f, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x41, 0x49, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x4c, 0x59, 0x41, 0x49, 0x52,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x59, 0x45, 0x54, 0x41, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x48, 0x49,
	0x41, 0x54, 0x55, 0x53, 0x10, 0x05, 0x2a, 0x95, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4d, 0x4f, 0x56, 0x49, 0x45,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x54, 0x56, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4f, 0x56, 0x41, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4f, 0x4e, 0x41, 0x10,
	0x04, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f,
	0x4d, 0x55, 0x53, 0x49, 0x43, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x52, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x07, 0x2a, 0x74,
	0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x52,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x41, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50,
	0x47, 0x31, 0x33, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x52, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x50,
	0x4c, 0x55, 0x53, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x52, 0x58, 0x10, 0x06, 0x2a, 0x61, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0d,
	0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x46, 0x41,
	0x4c, 0x4c, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x57,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x75, 0x62, 0x68, 0x61, 0x6d, 0x37, 0x31, 0x30,
	0x31, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x77, 0x61, 0x76, 0x65, 0x68, 0x75, 0x62, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_enums_proto_goTypes = []any{
	(MediaType)(0), // 0: cinewavehub.MediaType
	(Status)(0),    // 1: cinewavehub.Status
	(Format)(0),    // 2: cinewavehub.Format
	(Rating)(0),    // 3: cinewavehub.Rating
	(Season)(0),    // 4: cinewavehub.Season
}
var file_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
