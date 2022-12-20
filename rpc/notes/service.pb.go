// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/notes/service.proto

package notes

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

type Note struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4e92d382e48a99, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Note) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Note) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type CreateNoteParams struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNoteParams) Reset()         { *m = CreateNoteParams{} }
func (m *CreateNoteParams) String() string { return proto.CompactTextString(m) }
func (*CreateNoteParams) ProtoMessage()    {}
func (*CreateNoteParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4e92d382e48a99, []int{1}
}

func (m *CreateNoteParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNoteParams.Unmarshal(m, b)
}
func (m *CreateNoteParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNoteParams.Marshal(b, m, deterministic)
}
func (m *CreateNoteParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNoteParams.Merge(m, src)
}
func (m *CreateNoteParams) XXX_Size() int {
	return xxx_messageInfo_CreateNoteParams.Size(m)
}
func (m *CreateNoteParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNoteParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNoteParams proto.InternalMessageInfo

func (m *CreateNoteParams) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type GetAllNotesParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllNotesParams) Reset()         { *m = GetAllNotesParams{} }
func (m *GetAllNotesParams) String() string { return proto.CompactTextString(m) }
func (*GetAllNotesParams) ProtoMessage()    {}
func (*GetAllNotesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4e92d382e48a99, []int{2}
}

func (m *GetAllNotesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllNotesParams.Unmarshal(m, b)
}
func (m *GetAllNotesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllNotesParams.Marshal(b, m, deterministic)
}
func (m *GetAllNotesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllNotesParams.Merge(m, src)
}
func (m *GetAllNotesParams) XXX_Size() int {
	return xxx_messageInfo_GetAllNotesParams.Size(m)
}
func (m *GetAllNotesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllNotesParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllNotesParams proto.InternalMessageInfo

type GetAllNotesResult struct {
	Notes                []*Note  `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllNotesResult) Reset()         { *m = GetAllNotesResult{} }
func (m *GetAllNotesResult) String() string { return proto.CompactTextString(m) }
func (*GetAllNotesResult) ProtoMessage()    {}
func (*GetAllNotesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4e92d382e48a99, []int{3}
}

func (m *GetAllNotesResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllNotesResult.Unmarshal(m, b)
}
func (m *GetAllNotesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllNotesResult.Marshal(b, m, deterministic)
}
func (m *GetAllNotesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllNotesResult.Merge(m, src)
}
func (m *GetAllNotesResult) XXX_Size() int {
	return xxx_messageInfo_GetAllNotesResult.Size(m)
}
func (m *GetAllNotesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllNotesResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllNotesResult proto.InternalMessageInfo

func (m *GetAllNotesResult) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Note)(nil), "gotwirprpcexample.rpc.notes.Note")
	proto.RegisterType((*CreateNoteParams)(nil), "gotwirprpcexample.rpc.notes.CreateNoteParams")
	proto.RegisterType((*GetAllNotesParams)(nil), "gotwirprpcexample.rpc.notes.GetAllNotesParams")
	proto.RegisterType((*GetAllNotesResult)(nil), "gotwirprpcexample.rpc.notes.GetAllNotesResult")
}

func init() { proto.RegisterFile("rpc/notes/service.proto", fileDescriptor_9c4e92d382e48a99) }

var fileDescriptor_9c4e92d382e48a99 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xc9, 0x76, 0x57, 0xe8, 0x54, 0x44, 0xc7, 0x83, 0x65, 0x45, 0xa8, 0x3d, 0x48, 0x2f,
	0x66, 0x61, 0x3d, 0x78, 0x5e, 0x3d, 0x88, 0x20, 0x22, 0xf5, 0xe6, 0x41, 0x89, 0xe9, 0x20, 0x85,
	0xd6, 0x84, 0x64, 0xd4, 0xfd, 0xb2, 0x7e, 0x17, 0x69, 0x2a, 0xbb, 0xfe, 0x81, 0xea, 0x2d, 0x19,
	0xde, 0xfc, 0xe6, 0xbd, 0x19, 0xd8, 0x73, 0x56, 0xcf, 0x9e, 0x0d, 0x93, 0x9f, 0x79, 0x72, 0xaf,
	0xb5, 0x26, 0x69, 0x9d, 0x61, 0x83, 0xfb, 0x4f, 0x86, 0xdf, 0x6a, 0x67, 0x9d, 0xd5, 0xb4, 0x54,
	0xad, 0x6d, 0x48, 0x3a, 0xab, 0x65, 0x90, 0xe6, 0x97, 0x30, 0xbe, 0x36, 0x4c, 0xb8, 0x05, 0xa3,
	0xba, 0x4a, 0x45, 0x26, 0x8a, 0x49, 0x39, 0xaa, 0x2b, 0x44, 0x18, 0x33, 0x2d, 0x39, 0x1d, 0x65,
	0xa2, 0x88, 0xcb, 0xf0, 0xc6, 0x03, 0x00, 0xed, 0x48, 0x31, 0x55, 0x0f, 0x8a, 0xd3, 0x28, 0x13,
	0x45, 0x54, 0xc6, 0x9f, 0x95, 0x05, 0xe7, 0x47, 0xb0, 0x7d, 0x1e, 0x3e, 0x1d, 0xf0, 0x46, 0x39,
	0xd5, 0xfa, 0x15, 0x46, 0xac, 0x31, 0xf9, 0x2e, 0xec, 0x5c, 0x10, 0x2f, 0x9a, 0xa6, 0xd3, 0xf9,
	0x5e, 0x98, 0x5f, 0x7d, 0x2b, 0x96, 0xe4, 0x5f, 0x1a, 0xc6, 0x53, 0x98, 0x04, 0x97, 0xa9, 0xc8,
	0xa2, 0x22, 0x99, 0x1f, 0xca, 0x81, 0x24, 0xb2, 0x6b, 0x2c, 0x7b, 0xfd, 0xfc, 0x5d, 0xc0, 0x66,
	0x00, 0xdd, 0xf6, 0x9b, 0xc0, 0x7b, 0x80, 0xb5, 0x37, 0x3c, 0x1e, 0x04, 0xfd, 0x0c, 0x31, 0xfd,
	0x7b, 0x2e, 0xb6, 0x90, 0x7c, 0xb1, 0x8f, 0x72, 0xb0, 0xe3, 0x57, 0xfa, 0xe9, 0xbf, 0xf5, 0xfd,
	0x62, 0xce, 0x92, 0xbb, 0x78, 0x75, 0xed, 0xc7, 0x8d, 0x70, 0xe6, 0x93, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x61, 0x9a, 0x61, 0x4c, 0x01, 0x02, 0x00, 0x00,
}