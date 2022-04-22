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
// source: third_party/tink/proto/rsa_ssa_pss.proto

package rsa_ssa_pss_go_proto

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

type RsaSsaPssParams struct {
	// Hash function used in computing hash of the signing message
	// (see https://tools.ietf.org/html/rfc8017#section-9.1.1).
	// Required.
	SigHash common_go_proto.HashType `protobuf:"varint,1,opt,name=sig_hash,json=sigHash,proto3,enum=google.crypto.tink.HashType" json:"sig_hash,omitempty"`
	// Hash function used in MGF1 (a mask generation function based on a
	// hash function) (see https://tools.ietf.org/html/rfc8017#appendix-B.2.1).
	// Required.
	Mgf1Hash common_go_proto.HashType `protobuf:"varint,2,opt,name=mgf1_hash,json=mgf1Hash,proto3,enum=google.crypto.tink.HashType" json:"mgf1_hash,omitempty"`
	// Salt length (see https://tools.ietf.org/html/rfc8017#section-9.1.1)
	// Required.
	SaltLength           int32    `protobuf:"varint,3,opt,name=salt_length,json=saltLength,proto3" json:"salt_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPssParams) Reset()         { *m = RsaSsaPssParams{} }
func (m *RsaSsaPssParams) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPssParams) ProtoMessage()    {}
func (*RsaSsaPssParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_aacfb9eade71ca83, []int{0}
}

func (m *RsaSsaPssParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPssParams.Unmarshal(m, b)
}
func (m *RsaSsaPssParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPssParams.Marshal(b, m, deterministic)
}
func (m *RsaSsaPssParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPssParams.Merge(m, src)
}
func (m *RsaSsaPssParams) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPssParams.Size(m)
}
func (m *RsaSsaPssParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPssParams.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPssParams proto.InternalMessageInfo

func (m *RsaSsaPssParams) GetSigHash() common_go_proto.HashType {
	if m != nil {
		return m.SigHash
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *RsaSsaPssParams) GetMgf1Hash() common_go_proto.HashType {
	if m != nil {
		return m.Mgf1Hash
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *RsaSsaPssParams) GetSaltLength() int32 {
	if m != nil {
		return m.SaltLength
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.RsaSsaPssPublicKey
type RsaSsaPssPublicKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	Params *RsaSsaPssParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// Modulus.
	// Unsigned big integer in bigendian representation.
	N []byte `protobuf:"bytes,3,opt,name=n,proto3" json:"n,omitempty"`
	// Public exponent.
	// Unsigned big integer in bigendian representation.
	E                    []byte   `protobuf:"bytes,4,opt,name=e,proto3" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPssPublicKey) Reset()         { *m = RsaSsaPssPublicKey{} }
func (m *RsaSsaPssPublicKey) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPssPublicKey) ProtoMessage()    {}
func (*RsaSsaPssPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_aacfb9eade71ca83, []int{1}
}

func (m *RsaSsaPssPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPssPublicKey.Unmarshal(m, b)
}
func (m *RsaSsaPssPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPssPublicKey.Marshal(b, m, deterministic)
}
func (m *RsaSsaPssPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPssPublicKey.Merge(m, src)
}
func (m *RsaSsaPssPublicKey) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPssPublicKey.Size(m)
}
func (m *RsaSsaPssPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPssPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPssPublicKey proto.InternalMessageInfo

func (m *RsaSsaPssPublicKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RsaSsaPssPublicKey) GetParams() *RsaSsaPssParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *RsaSsaPssPublicKey) GetN() []byte {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *RsaSsaPssPublicKey) GetE() []byte {
	if m != nil {
		return m.E
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.RsaSsaPssPrivateKey
type RsaSsaPssPrivateKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	PublicKey *RsaSsaPssPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Private exponent.
	// Unsigned big integer in bigendian representation.
	// Required.
	D []byte `protobuf:"bytes,3,opt,name=d,proto3" json:"d,omitempty"`
	// The following parameters are used to optimize RSA signature computation.
	// The prime factor p of n.
	// Unsigned big integer in bigendian representation.
	// Required.
	P []byte `protobuf:"bytes,4,opt,name=p,proto3" json:"p,omitempty"`
	// The prime factor q of n.
	// Unsigned big integer in bigendian representation.
	// Required.
	Q []byte `protobuf:"bytes,5,opt,name=q,proto3" json:"q,omitempty"`
	// d mod (p - 1).
	// Unsigned big integer in bigendian representation.
	// Required.
	Dp []byte `protobuf:"bytes,6,opt,name=dp,proto3" json:"dp,omitempty"`
	// d mod (q - 1).
	// Unsigned big integer in bigendian representation.
	// Required.
	Dq []byte `protobuf:"bytes,7,opt,name=dq,proto3" json:"dq,omitempty"`
	// Chinese Remainder Theorem coefficient q^(-1) mod p.
	// Unsigned big integer in bigendian representation.
	// Required.
	Crt                  []byte   `protobuf:"bytes,8,opt,name=crt,proto3" json:"crt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPssPrivateKey) Reset()         { *m = RsaSsaPssPrivateKey{} }
func (m *RsaSsaPssPrivateKey) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPssPrivateKey) ProtoMessage()    {}
func (*RsaSsaPssPrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_aacfb9eade71ca83, []int{2}
}

func (m *RsaSsaPssPrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPssPrivateKey.Unmarshal(m, b)
}
func (m *RsaSsaPssPrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPssPrivateKey.Marshal(b, m, deterministic)
}
func (m *RsaSsaPssPrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPssPrivateKey.Merge(m, src)
}
func (m *RsaSsaPssPrivateKey) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPssPrivateKey.Size(m)
}
func (m *RsaSsaPssPrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPssPrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPssPrivateKey proto.InternalMessageInfo

func (m *RsaSsaPssPrivateKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RsaSsaPssPrivateKey) GetPublicKey() *RsaSsaPssPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *RsaSsaPssPrivateKey) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *RsaSsaPssPrivateKey) GetP() []byte {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *RsaSsaPssPrivateKey) GetQ() []byte {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *RsaSsaPssPrivateKey) GetDp() []byte {
	if m != nil {
		return m.Dp
	}
	return nil
}

func (m *RsaSsaPssPrivateKey) GetDq() []byte {
	if m != nil {
		return m.Dq
	}
	return nil
}

func (m *RsaSsaPssPrivateKey) GetCrt() []byte {
	if m != nil {
		return m.Crt
	}
	return nil
}

type RsaSsaPssKeyFormat struct {
	// Required.
	Params *RsaSsaPssParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// Required.
	ModulusSizeInBits uint32 `protobuf:"varint,2,opt,name=modulus_size_in_bits,json=modulusSizeInBits,proto3" json:"modulus_size_in_bits,omitempty"`
	// Required.
	PublicExponent       []byte   `protobuf:"bytes,3,opt,name=public_exponent,json=publicExponent,proto3" json:"public_exponent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPssKeyFormat) Reset()         { *m = RsaSsaPssKeyFormat{} }
func (m *RsaSsaPssKeyFormat) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPssKeyFormat) ProtoMessage()    {}
func (*RsaSsaPssKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_aacfb9eade71ca83, []int{3}
}

func (m *RsaSsaPssKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPssKeyFormat.Unmarshal(m, b)
}
func (m *RsaSsaPssKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPssKeyFormat.Marshal(b, m, deterministic)
}
func (m *RsaSsaPssKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPssKeyFormat.Merge(m, src)
}
func (m *RsaSsaPssKeyFormat) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPssKeyFormat.Size(m)
}
func (m *RsaSsaPssKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPssKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPssKeyFormat proto.InternalMessageInfo

func (m *RsaSsaPssKeyFormat) GetParams() *RsaSsaPssParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *RsaSsaPssKeyFormat) GetModulusSizeInBits() uint32 {
	if m != nil {
		return m.ModulusSizeInBits
	}
	return 0
}

func (m *RsaSsaPssKeyFormat) GetPublicExponent() []byte {
	if m != nil {
		return m.PublicExponent
	}
	return nil
}

func init() {
	proto.RegisterType((*RsaSsaPssParams)(nil), "google.crypto.tink.RsaSsaPssParams")
	proto.RegisterType((*RsaSsaPssPublicKey)(nil), "google.crypto.tink.RsaSsaPssPublicKey")
	proto.RegisterType((*RsaSsaPssPrivateKey)(nil), "google.crypto.tink.RsaSsaPssPrivateKey")
	proto.RegisterType((*RsaSsaPssKeyFormat)(nil), "google.crypto.tink.RsaSsaPssKeyFormat")
}

func init() {
	proto.RegisterFile("proto/rsa_ssa_pss.proto", fileDescriptor_aacfb9eade71ca83)
}

var fileDescriptor_aacfb9eade71ca83 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0xe5, 0x8e, 0xb5, 0xdd, 0xb7, 0xae, 0x03, 0xc3, 0x21, 0x42, 0x93, 0x98, 0x3a, 0x09,
	0x7a, 0x4a, 0xb4, 0x71, 0x40, 0x88, 0xdb, 0xa4, 0x21, 0xd0, 0x38, 0x54, 0x29, 0x27, 0x2e, 0x96,
	0x9b, 0x98, 0xc4, 0x5a, 0x62, 0xbb, 0xfe, 0xdc, 0x89, 0xec, 0x0f, 0xf0, 0x4f, 0xb8, 0xf0, 0x6f,
	0xf8, 0x45, 0xc8, 0x4e, 0xc2, 0x26, 0xd8, 0x26, 0x76, 0x4a, 0xde, 0x37, 0xdf, 0x1b, 0x3d, 0xef,
	0x97, 0x18, 0xe6, 0xae, 0x94, 0x36, 0x67, 0x86, 0x5b, 0xd7, 0x24, 0x4e, 0xaa, 0x8b, 0xc4, 0x58,
	0xed, 0x74, 0x62, 0x91, 0x33, 0x44, 0xce, 0x0c, 0x62, 0x1c, 0x1c, 0x4a, 0x0b, 0xad, 0x8b, 0x4a,
	0xc4, 0x99, 0x6d, 0x8c, 0xd3, 0xb1, 0x9f, 0x7d, 0x7e, 0x74, 0x47, 0x3a, 0xd3, 0x75, 0xad, 0x55,
	0x1b, 0x9c, 0xfd, 0x20, 0xb0, 0x9f, 0x22, 0x5f, 0x22, 0x5f, 0x20, 0x2e, 0xb8, 0xe5, 0x35, 0xd2,
	0x37, 0x30, 0x46, 0x59, 0xb0, 0x92, 0x63, 0x19, 0x91, 0x43, 0x32, 0x9f, 0x9e, 0x1c, 0xc4, 0xff,
	0xbe, 0x3f, 0xfe, 0xc0, 0xb1, 0xfc, 0xdc, 0x18, 0x91, 0x8e, 0x50, 0x16, 0x5e, 0xd0, 0xb7, 0xb0,
	0x53, 0x17, 0x5f, 0x8f, 0xdb, 0xe4, 0xe0, 0x3f, 0x92, 0x63, 0x3f, 0x1e, 0xa2, 0x2f, 0x60, 0x17,
	0x79, 0xe5, 0x58, 0x25, 0x54, 0xe1, 0xca, 0x68, 0xeb, 0x90, 0xcc, 0xb7, 0x53, 0xf0, 0xd6, 0xa7,
	0xe0, 0xcc, 0xbe, 0x13, 0xa0, 0xd7, 0xa0, 0x9b, 0x55, 0x25, 0xb3, 0x73, 0xd1, 0xd0, 0x08, 0x46,
	0x97, 0xc2, 0xa2, 0xd4, 0x2a, 0xa0, 0xee, 0xa5, 0xbd, 0xa4, 0xef, 0x60, 0x68, 0x42, 0x9f, 0x40,
	0xb2, 0x7b, 0x72, 0x74, 0x1b, 0xc9, 0x5f, 0xd5, 0xd3, 0x2e, 0x42, 0x27, 0x40, 0x54, 0x80, 0x98,
	0xa4, 0x44, 0x79, 0x25, 0xa2, 0x47, 0xad, 0x12, 0xb3, 0x5f, 0x04, 0x9e, 0x5e, 0xe7, 0xac, 0xbc,
	0xe4, 0x4e, 0xdc, 0x8f, 0x72, 0x06, 0x60, 0x02, 0x31, 0xbb, 0x10, 0x4d, 0x87, 0xf3, 0xf2, 0x7e,
	0x9c, 0xbe, 0x60, 0xba, 0x63, 0xfe, 0x74, 0x9d, 0x00, 0xc9, 0x7b, 0xa8, 0xdc, 0x2b, 0xd3, 0x43,
	0x19, 0xaf, 0xd6, 0xd1, 0x76, 0xab, 0xd6, 0x74, 0x0a, 0x83, 0xdc, 0x44, 0xc3, 0x20, 0x07, 0xb9,
	0x09, 0x7a, 0x1d, 0x8d, 0x3a, 0xbd, 0xa6, 0x8f, 0x61, 0x2b, 0xb3, 0x2e, 0x1a, 0x07, 0xc3, 0xdf,
	0xce, 0x7e, 0xde, 0x5c, 0xef, 0xb9, 0x68, 0xde, 0x6b, 0x5b, 0x73, 0x77, 0x63, 0x89, 0xe4, 0xe1,
	0x4b, 0x4c, 0xe0, 0x59, 0xad, 0xf3, 0x4d, 0xb5, 0x41, 0x86, 0xf2, 0x4a, 0x30, 0xa9, 0xd8, 0x4a,
	0xba, 0xf6, 0x7b, 0xec, 0xa5, 0x4f, 0xba, 0x67, 0x4b, 0x79, 0x25, 0x3e, 0xaa, 0x53, 0xe9, 0x90,
	0xbe, 0x82, 0xfd, 0x6e, 0x4f, 0xe2, 0x9b, 0xd1, 0x4a, 0x28, 0xd7, 0xd5, 0x9d, 0xb6, 0xf6, 0x59,
	0xe7, 0x9e, 0x2e, 0xe1, 0x20, 0xd3, 0xf5, 0x6d, 0x2c, 0xe1, 0xaf, 0x5e, 0x90, 0x2f, 0xc7, 0x85,
	0x74, 0xe5, 0x66, 0x15, 0x67, 0xba, 0x4e, 0xda, 0xb1, 0x3b, 0x0e, 0x10, 0x2b, 0x34, 0x0b, 0xe6,
	0x6a, 0x18, 0x2e, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x09, 0x87, 0xba, 0xa7, 0x76, 0x03,
	0x00, 0x00,
}
