// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party/tink/proto/hkdf_prf.proto

package hkdf_prf_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/clearblade/tink/go/proto/common_go_proto"
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

type HkdfPrfParams struct {
	Hash common_go_proto.HashType `protobuf:"varint,1,opt,name=hash,proto3,enum=google.crypto.tink.HashType" json:"hash,omitempty"`
	// Salt, optional in RFC 5869. Using "" is equivalent to zeros of length up to
	// the block length of the HMac.
	Salt                 []byte   `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HkdfPrfParams) Reset()         { *m = HkdfPrfParams{} }
func (m *HkdfPrfParams) String() string { return proto.CompactTextString(m) }
func (*HkdfPrfParams) ProtoMessage()    {}
func (*HkdfPrfParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c59dfa5362ca5fc0, []int{0}
}

func (m *HkdfPrfParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HkdfPrfParams.Unmarshal(m, b)
}
func (m *HkdfPrfParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HkdfPrfParams.Marshal(b, m, deterministic)
}
func (m *HkdfPrfParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HkdfPrfParams.Merge(m, src)
}
func (m *HkdfPrfParams) XXX_Size() int {
	return xxx_messageInfo_HkdfPrfParams.Size(m)
}
func (m *HkdfPrfParams) XXX_DiscardUnknown() {
	xxx_messageInfo_HkdfPrfParams.DiscardUnknown(m)
}

var xxx_messageInfo_HkdfPrfParams proto.InternalMessageInfo

func (m *HkdfPrfParams) GetHash() common_go_proto.HashType {
	if m != nil {
		return m.Hash
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *HkdfPrfParams) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

type HkdfPrfKey struct {
	Version              uint32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params               *HkdfPrfParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue             []byte         `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HkdfPrfKey) Reset()         { *m = HkdfPrfKey{} }
func (m *HkdfPrfKey) String() string { return proto.CompactTextString(m) }
func (*HkdfPrfKey) ProtoMessage()    {}
func (*HkdfPrfKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c59dfa5362ca5fc0, []int{1}
}

func (m *HkdfPrfKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HkdfPrfKey.Unmarshal(m, b)
}
func (m *HkdfPrfKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HkdfPrfKey.Marshal(b, m, deterministic)
}
func (m *HkdfPrfKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HkdfPrfKey.Merge(m, src)
}
func (m *HkdfPrfKey) XXX_Size() int {
	return xxx_messageInfo_HkdfPrfKey.Size(m)
}
func (m *HkdfPrfKey) XXX_DiscardUnknown() {
	xxx_messageInfo_HkdfPrfKey.DiscardUnknown(m)
}

var xxx_messageInfo_HkdfPrfKey proto.InternalMessageInfo

func (m *HkdfPrfKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HkdfPrfKey) GetParams() *HkdfPrfParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *HkdfPrfKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type HkdfPrfKeyFormat struct {
	Params               *HkdfPrfParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize              uint32         `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	Version              uint32         `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HkdfPrfKeyFormat) Reset()         { *m = HkdfPrfKeyFormat{} }
func (m *HkdfPrfKeyFormat) String() string { return proto.CompactTextString(m) }
func (*HkdfPrfKeyFormat) ProtoMessage()    {}
func (*HkdfPrfKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c59dfa5362ca5fc0, []int{2}
}

func (m *HkdfPrfKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HkdfPrfKeyFormat.Unmarshal(m, b)
}
func (m *HkdfPrfKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HkdfPrfKeyFormat.Marshal(b, m, deterministic)
}
func (m *HkdfPrfKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HkdfPrfKeyFormat.Merge(m, src)
}
func (m *HkdfPrfKeyFormat) XXX_Size() int {
	return xxx_messageInfo_HkdfPrfKeyFormat.Size(m)
}
func (m *HkdfPrfKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_HkdfPrfKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_HkdfPrfKeyFormat proto.InternalMessageInfo

func (m *HkdfPrfKeyFormat) GetParams() *HkdfPrfParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *HkdfPrfKeyFormat) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *HkdfPrfKeyFormat) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*HkdfPrfParams)(nil), "google.crypto.tink.HkdfPrfParams")
	proto.RegisterType((*HkdfPrfKey)(nil), "google.crypto.tink.HkdfPrfKey")
	proto.RegisterType((*HkdfPrfKeyFormat)(nil), "google.crypto.tink.HkdfPrfKeyFormat")
}

func init() {
	proto.RegisterFile("proto/hkdf_prf.proto", fileDescriptor_c59dfa5362ca5fc0)
}

var fileDescriptor_c59dfa5362ca5fc0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x1c, 0xc5, 0x59, 0x15, 0xad, 0x7f, 0x19, 0x31, 0xa7, 0xad, 0x3c, 0x98, 0x11, 0x08, 0xc1, 0x6e,
	0xd8, 0xa9, 0xab, 0x87, 0x30, 0x84, 0x58, 0x36, 0xf3, 0xd0, 0x65, 0x19, 0xd7, 0x59, 0x67, 0x18,
	0xc7, 0x19, 0x66, 0x46, 0x61, 0x3c, 0x74, 0xe8, 0xa3, 0xf4, 0x49, 0xc3, 0xd1, 0xc8, 0x68, 0x3b,
	0x74, 0xdb, 0x07, 0xef, 0xbd, 0xdf, 0xfb, 0xef, 0xc0, 0xb5, 0xa5, 0x4c, 0x4f, 0x33, 0x85, 0xb5,
	0x75, 0xb1, 0x65, 0x0b, 0x1e, 0x2b, 0x2d, 0xad, 0x8c, 0x29, 0x9f, 0x16, 0x99, 0xd2, 0x45, 0xe4,
	0x25, 0x42, 0x33, 0x29, 0x67, 0x73, 0x12, 0xe5, 0xda, 0x29, 0x2b, 0xa3, 0x8d, 0xf1, 0xfc, 0xea,
	0x8f, 0x68, 0x2e, 0x85, 0x90, 0x8b, 0x6d, 0xb0, 0xf3, 0x02, 0xcd, 0x01, 0x9f, 0x16, 0x89, 0x2e,
	0x12, 0xac, 0xb1, 0x30, 0xe8, 0x16, 0x6a, 0x14, 0x1b, 0x1a, 0x06, 0xed, 0xa0, 0x7b, 0xd2, 0x6b,
	0x45, 0xbf, 0x8b, 0xa3, 0x01, 0x36, 0x74, 0xe4, 0x14, 0x49, 0xbd, 0x13, 0x21, 0xa8, 0x19, 0x3c,
	0xb7, 0x61, 0xa5, 0x1d, 0x74, 0x8f, 0x53, 0xff, 0xdd, 0x79, 0x03, 0xd8, 0xd5, 0x0e, 0x89, 0x43,
	0x21, 0x34, 0x56, 0x44, 0x1b, 0x26, 0x17, 0xbe, 0xb6, 0x99, 0x7e, 0x49, 0x74, 0x0f, 0x75, 0xe5,
	0xb9, 0x3e, 0x7d, 0xd4, 0xbb, 0x2c, 0xe5, 0xed, 0x0f, 0x4c, 0x77, 0x01, 0x74, 0x01, 0x87, 0x9c,
	0xb8, 0x6c, 0x85, 0xe7, 0x4b, 0x12, 0x56, 0x3d, 0xfb, 0x80, 0x13, 0x37, 0xde, 0xe8, 0xce, 0x7b,
	0x00, 0xa7, 0xdf, 0x03, 0x1e, 0xa4, 0x16, 0xd8, 0xee, 0xc1, 0x82, 0xff, 0xc2, 0xce, 0x60, 0xd3,
	0x9d, 0x19, 0xb6, 0x26, 0x7e, 0x69, 0x33, 0x6d, 0x70, 0xe2, 0x9e, 0xd9, 0x9a, 0xec, 0x1f, 0x57,
	0xfd, 0x71, 0x5c, 0x7f, 0x0c, 0xad, 0x5c, 0x8a, 0x32, 0x88, 0xff, 0xf7, 0x49, 0xf0, 0x7a, 0x33,
	0x63, 0x96, 0x2e, 0x27, 0x51, 0x2e, 0x45, 0xbc, 0xb5, 0x95, 0xbd, 0x71, 0xe6, 0xe5, 0x47, 0xa5,
	0x3e, 0x7a, 0x7c, 0x1a, 0x26, 0xfd, 0x49, 0xdd, 0xeb, 0xbb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x29, 0x6a, 0x90, 0xdc, 0x1c, 0x02, 0x00, 0x00,
}
