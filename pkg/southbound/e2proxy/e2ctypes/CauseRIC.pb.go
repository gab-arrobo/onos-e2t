// Code generated by protoc-gen-go. DO NOT EDIT.
// CauseRIC.proto is a deprecated file.

package e2ctypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CauseRICT int32

const (
	CauseRICT_CauseRIC_ran_function_id_Invalid                        CauseRICT = 0
	CauseRICT_CauseRIC_action_not_supported                           CauseRICT = 1
	CauseRICT_CauseRIC_excessive_actions                              CauseRICT = 2
	CauseRICT_CauseRIC_duplicate_action                               CauseRICT = 3
	CauseRICT_CauseRIC_duplicate_event                                CauseRICT = 4
	CauseRICT_CauseRIC_function_resource_limit                        CauseRICT = 5
	CauseRICT_CauseRIC_request_id_unknown                             CauseRICT = 6
	CauseRICT_CauseRIC_inconsistent_action_subsequent_action_sequence CauseRICT = 7
	CauseRICT_CauseRIC_control_message_invalid                        CauseRICT = 8
	CauseRICT_CauseRIC_call_process_id_invalid                        CauseRICT = 9
	CauseRICT_CauseRIC_unspecified                                    CauseRICT = 10
)

var CauseRICT_name = map[int32]string{
	0:  "CauseRIC_ran_function_id_Invalid",
	1:  "CauseRIC_action_not_supported",
	2:  "CauseRIC_excessive_actions",
	3:  "CauseRIC_duplicate_action",
	4:  "CauseRIC_duplicate_event",
	5:  "CauseRIC_function_resource_limit",
	6:  "CauseRIC_request_id_unknown",
	7:  "CauseRIC_inconsistent_action_subsequent_action_sequence",
	8:  "CauseRIC_control_message_invalid",
	9:  "CauseRIC_call_process_id_invalid",
	10: "CauseRIC_unspecified",
}

var CauseRICT_value = map[string]int32{
	"CauseRIC_ran_function_id_Invalid":                        0,
	"CauseRIC_action_not_supported":                           1,
	"CauseRIC_excessive_actions":                              2,
	"CauseRIC_duplicate_action":                               3,
	"CauseRIC_duplicate_event":                                4,
	"CauseRIC_function_resource_limit":                        5,
	"CauseRIC_request_id_unknown":                             6,
	"CauseRIC_inconsistent_action_subsequent_action_sequence": 7,
	"CauseRIC_control_message_invalid":                        8,
	"CauseRIC_call_process_id_invalid":                        9,
	"CauseRIC_unspecified":                                    10,
}

func (x CauseRICT) String() string {
	return proto.EnumName(CauseRICT_name, int32(x))
}

func (CauseRICT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f31811999b8cad0d, []int{0}
}

func init() {
	proto.RegisterEnum("e2ctypes.CauseRICT", CauseRICT_name, CauseRICT_value)
}

func init() { proto.RegisterFile("CauseRIC.proto", fileDescriptor_f31811999b8cad0d) }

var fileDescriptor_f31811999b8cad0d = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbb, 0x4e, 0x43, 0x31,
	0x0c, 0x86, 0x29, 0x85, 0x52, 0x32, 0xa0, 0x28, 0x62, 0x28, 0x97, 0x72, 0x91, 0x98, 0x18, 0x18,
	0x60, 0x60, 0x60, 0xa3, 0x53, 0x57, 0x5e, 0x20, 0x4a, 0x73, 0x5c, 0x64, 0x91, 0x3a, 0x21, 0x76,
	0x0a, 0xbc, 0x0b, 0x0f, 0xc2, 0xe3, 0xa1, 0x53, 0x7a, 0x02, 0x95, 0x18, 0x9d, 0xef, 0x8b, 0xf4,
	0xff, 0xb6, 0x3a, 0x98, 0xb8, 0xc2, 0xf0, 0x34, 0x9d, 0xdc, 0xa4, 0x1c, 0x25, 0x9a, 0x21, 0xdc,
	0x7a, 0xf9, 0x48, 0xc0, 0xd7, 0x9f, 0x7d, 0xa5, 0x3a, 0x68, 0xc5, 0x5c, 0xa9, 0x8b, 0x3a, 0x65,
	0x47, 0x76, 0x5e, 0xc8, 0x0b, 0x46, 0xb2, 0xd8, 0xd8, 0x29, 0x2d, 0x5d, 0xc0, 0x46, 0x6f, 0x99,
	0x4b, 0x35, 0xae, 0x96, 0xfb, 0xe1, 0x14, 0xc5, 0x72, 0x49, 0x29, 0x66, 0x81, 0x46, 0xf7, 0xcc,
	0x99, 0x3a, 0xae, 0x0a, 0xbc, 0x7b, 0x60, 0xc6, 0x25, 0xac, 0x65, 0xd6, 0xdb, 0x66, 0xac, 0x8e,
	0x2a, 0x6f, 0x4a, 0x0a, 0xe8, 0x9d, 0x74, 0x5c, 0xf7, 0xcd, 0xa9, 0x1a, 0xfd, 0x83, 0x61, 0x09,
	0x24, 0x7a, 0x67, 0x23, 0x65, 0x4d, 0x98, 0x81, 0x63, 0xc9, 0x1e, 0x6c, 0xc0, 0x05, 0x8a, 0xde,
	0x35, 0xe7, 0xea, 0xe4, 0xb7, 0x0b, 0xbc, 0x16, 0x60, 0x69, 0x6b, 0x14, 0x7a, 0xa1, 0xf8, 0x46,
	0x7a, 0x60, 0x1e, 0xd4, 0x7d, 0x15, 0x90, 0x7c, 0x24, 0x46, 0x16, 0x20, 0xe9, 0x3a, 0x71, 0x99,
	0x71, 0xfb, 0xef, 0xcf, 0xcb, 0x6a, 0xf4, 0xa0, 0xf7, 0x36, 0x32, 0xf8, 0x48, 0x92, 0x63, 0xb0,
	0x0b, 0x60, 0x76, 0xcf, 0x60, 0x71, 0xbd, 0xa9, 0xe1, 0xa6, 0xe5, 0x42, 0xb0, 0x29, 0xc7, 0x76,
	0x19, 0x6d, 0x90, 0xce, 0xda, 0x37, 0x23, 0x75, 0x58, 0xad, 0x42, 0x9c, 0xc0, 0xe3, 0x1c, 0xa1,
	0xd1, 0xea, 0xb1, 0xff, 0xd5, 0xeb, 0xcd, 0x06, 0xab, 0xa3, 0xdd, 0x7d, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x48, 0xa8, 0x80, 0xa8, 0xc6, 0x01, 0x00, 0x00,
}
