// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// What is sent through the wire
type ChatMessagePayload struct {
	// Message content
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// MIME type
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Message type
	MessageType string `protobuf:"bytes,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// Sender's clock value for message ordering
	ClockValue           float64  `protobuf:"fixed64,4,opt,name=clock_value,json=clockValue,proto3" json:"clock_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessagePayload) Reset()         { *m = ChatMessagePayload{} }
func (m *ChatMessagePayload) String() string { return proto.CompactTextString(m) }
func (*ChatMessagePayload) ProtoMessage()    {}
func (*ChatMessagePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *ChatMessagePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessagePayload.Unmarshal(m, b)
}
func (m *ChatMessagePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessagePayload.Marshal(b, m, deterministic)
}
func (m *ChatMessagePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessagePayload.Merge(m, src)
}
func (m *ChatMessagePayload) XXX_Size() int {
	return xxx_messageInfo_ChatMessagePayload.Size(m)
}
func (m *ChatMessagePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessagePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessagePayload proto.InternalMessageInfo

func (m *ChatMessagePayload) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ChatMessagePayload) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *ChatMessagePayload) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *ChatMessagePayload) GetClockValue() float64 {
	if m != nil {
		return m.ClockValue
	}
	return 0
}

// ContactUpdatePayload is sent when a user updates its profile
type ContactUpdatePayload struct {
	// Contact display name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Contact profile image, using the data URI scheme (e.g. "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAA...")
	ProfileImage string `protobuf:"bytes,2,opt,name=profile_image,json=profileImage,proto3" json:"profile_image,omitempty"`
	// Contact address
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Contact Firebase Cloud Messaging token
	FcmToken             string   `protobuf:"bytes,4,opt,name=fcm_token,json=fcmToken,proto3" json:"fcm_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactUpdatePayload) Reset()         { *m = ContactUpdatePayload{} }
func (m *ContactUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*ContactUpdatePayload) ProtoMessage()    {}
func (*ContactUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *ContactUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactUpdatePayload.Unmarshal(m, b)
}
func (m *ContactUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactUpdatePayload.Marshal(b, m, deterministic)
}
func (m *ContactUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactUpdatePayload.Merge(m, src)
}
func (m *ContactUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_ContactUpdatePayload.Size(m)
}
func (m *ContactUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ContactUpdatePayload proto.InternalMessageInfo

func (m *ContactUpdatePayload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactUpdatePayload) GetProfileImage() string {
	if m != nil {
		return m.ProfileImage
	}
	return ""
}

func (m *ContactUpdatePayload) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContactUpdatePayload) GetFcmToken() string {
	if m != nil {
		return m.FcmToken
	}
	return ""
}

// Incoming RPC messages
type OneToOneRPC struct {
	Src                  string   `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst                  string   `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneToOneRPC) Reset()         { *m = OneToOneRPC{} }
func (m *OneToOneRPC) String() string { return proto.CompactTextString(m) }
func (*OneToOneRPC) ProtoMessage()    {}
func (*OneToOneRPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *OneToOneRPC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneToOneRPC.Unmarshal(m, b)
}
func (m *OneToOneRPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneToOneRPC.Marshal(b, m, deterministic)
}
func (m *OneToOneRPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneToOneRPC.Merge(m, src)
}
func (m *OneToOneRPC) XXX_Size() int {
	return xxx_messageInfo_OneToOneRPC.Size(m)
}
func (m *OneToOneRPC) XXX_DiscardUnknown() {
	xxx_messageInfo_OneToOneRPC.DiscardUnknown(m)
}

var xxx_messageInfo_OneToOneRPC proto.InternalMessageInfo

func (m *OneToOneRPC) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *OneToOneRPC) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *OneToOneRPC) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ContactUpdateRPC struct {
	Src                  string                `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst                  string                `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	Payload              *ContactUpdatePayload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ContactUpdateRPC) Reset()         { *m = ContactUpdateRPC{} }
func (m *ContactUpdateRPC) String() string { return proto.CompactTextString(m) }
func (*ContactUpdateRPC) ProtoMessage()    {}
func (*ContactUpdateRPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *ContactUpdateRPC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactUpdateRPC.Unmarshal(m, b)
}
func (m *ContactUpdateRPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactUpdateRPC.Marshal(b, m, deterministic)
}
func (m *ContactUpdateRPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactUpdateRPC.Merge(m, src)
}
func (m *ContactUpdateRPC) XXX_Size() int {
	return xxx_messageInfo_ContactUpdateRPC.Size(m)
}
func (m *ContactUpdateRPC) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactUpdateRPC.DiscardUnknown(m)
}

var xxx_messageInfo_ContactUpdateRPC proto.InternalMessageInfo

func (m *ContactUpdateRPC) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *ContactUpdateRPC) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *ContactUpdateRPC) GetPayload() *ContactUpdatePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Incoming messages
type ChatProtocolMessage struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatProtocolMessage) Reset()         { *m = ChatProtocolMessage{} }
func (m *ChatProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ChatProtocolMessage) ProtoMessage()    {}
func (*ChatProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
}

func (m *ChatProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatProtocolMessage.Unmarshal(m, b)
}
func (m *ChatProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatProtocolMessage.Marshal(b, m, deterministic)
}
func (m *ChatProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatProtocolMessage.Merge(m, src)
}
func (m *ChatProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ChatProtocolMessage.Size(m)
}
func (m *ChatProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatProtocolMessage proto.InternalMessageInfo

func (m *ChatProtocolMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatMessagePayload)(nil), "chat.ChatMessagePayload")
	proto.RegisterType((*ContactUpdatePayload)(nil), "chat.ContactUpdatePayload")
	proto.RegisterType((*OneToOneRPC)(nil), "chat.OneToOneRPC")
	proto.RegisterType((*ContactUpdateRPC)(nil), "chat.ContactUpdateRPC")
	proto.RegisterType((*ChatProtocolMessage)(nil), "chat.ChatProtocolMessage")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x59, 0x5b, 0xd4, 0x4e, 0x2a, 0x94, 0xd5, 0x43, 0xd0, 0x83, 0x35, 0x5e, 0x7a, 0xaa,
	0xa0, 0xbe, 0x41, 0x4f, 0x22, 0xd2, 0x12, 0xaa, 0xd7, 0xb0, 0x6e, 0xa6, 0x7f, 0xe8, 0x66, 0x77,
	0xc9, 0xae, 0x42, 0x5f, 0xc0, 0x37, 0xf0, 0x7d, 0x65, 0x36, 0x1b, 0x35, 0xe0, 0xc1, 0xdb, 0x37,
	0x5f, 0x86, 0xf9, 0x7e, 0x93, 0x59, 0x00, 0xb9, 0x11, 0x7e, 0x6a, 0x6b, 0xe3, 0x0d, 0xef, 0x93,
	0xce, 0x3e, 0x19, 0xf0, 0xd9, 0x46, 0xf8, 0x27, 0x74, 0x4e, 0xac, 0x71, 0x21, 0xf6, 0xca, 0x88,
	0x92, 0xa7, 0x70, 0x24, 0x8d, 0xf6, 0xa8, 0x7d, 0xca, 0xc6, 0x6c, 0x32, 0xc8, 0xdb, 0x92, 0x5f,
	0xc1, 0x30, 0xca, 0xc2, 0xef, 0x2d, 0xa6, 0x07, 0xe1, 0x73, 0x12, 0xbd, 0xe5, 0xde, 0x22, 0xb5,
	0x54, 0xcd, 0xb8, 0xa6, 0xa5, 0xd7, 0xb4, 0x44, 0x2f, 0xb4, 0x5c, 0x42, 0x22, 0x95, 0x91, 0xbb,
	0xe2, 0x5d, 0xa8, 0x37, 0x4c, 0xfb, 0x63, 0x36, 0x61, 0x39, 0x04, 0xeb, 0x85, 0x9c, 0xec, 0x83,
	0xc1, 0xd9, 0xcc, 0x68, 0x2f, 0xa4, 0x7f, 0xb6, 0xa5, 0xf0, 0xdf, 0x64, 0x1c, 0xfa, 0x5a, 0x54,
	0x18, 0xb1, 0x82, 0xe6, 0xd7, 0x70, 0x62, 0x6b, 0xb3, 0xda, 0x2a, 0x2c, 0xb6, 0x95, 0x58, 0xb7,
	0x50, 0xc3, 0x68, 0x3e, 0x90, 0x47, 0x2b, 0x89, 0xb2, 0xac, 0xd1, 0xb9, 0x08, 0xd4, 0x96, 0xfc,
	0x02, 0x06, 0x2b, 0x59, 0x15, 0xde, 0xec, 0x50, 0x07, 0x94, 0x41, 0x7e, 0xbc, 0x92, 0xd5, 0x92,
	0xea, 0xec, 0x11, 0x92, 0xb9, 0xc6, 0xa5, 0x99, 0x6b, 0xcc, 0x17, 0x33, 0x3e, 0x82, 0x9e, 0xab,
	0x65, 0x4c, 0x27, 0x49, 0x4e, 0xe9, 0x7c, 0x8c, 0x24, 0x49, 0x49, 0xb6, 0xa1, 0x0d, 0x49, 0xc3,
	0xbc, 0x2d, 0x33, 0x05, 0xa3, 0xce, 0x52, 0xff, 0x9d, 0x78, 0xdf, 0x9d, 0x98, 0xdc, 0x9e, 0x4f,
	0xc3, 0x25, 0xff, 0xfa, 0x43, 0x3f, 0x69, 0x37, 0x70, 0x4a, 0xa7, 0x5d, 0xd0, 0xb9, 0xa5, 0x51,
	0xf1, 0xc4, 0xbf, 0xf1, 0x58, 0x07, 0xef, 0xf5, 0x30, 0xbc, 0x8c, 0xbb, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xad, 0x31, 0x5d, 0x6c, 0x27, 0x02, 0x00, 0x00,
}