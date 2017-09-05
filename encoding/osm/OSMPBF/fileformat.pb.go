// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileformat.proto

/*
Package OSMPBF is a generated protocol buffer package.

It is generated from these files:
	fileformat.proto
	osmformat.proto

It has these top-level messages:
	Blob
	BlobHeader
	HeaderBlock
	HeaderBBox
	PrimitiveBlock
	PrimitiveGroup
	StringTable
	Info
	DenseInfo
	ChangeSet
	Node
	DenseNodes
	Way
	Relation
*/
package OSMPBF

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Blob struct {
	Raw     []byte `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
	RawSize *int32 `protobuf:"varint,2,opt,name=raw_size,json=rawSize" json:"raw_size,omitempty"`
	// Possible compressed versions of the data.
	ZlibData []byte `protobuf:"bytes,3,opt,name=zlib_data,json=zlibData" json:"zlib_data,omitempty"`
	// PROPOSED feature for LZMA compressed data. SUPPORT IS NOT REQUIRED.
	LzmaData []byte `protobuf:"bytes,4,opt,name=lzma_data,json=lzmaData" json:"lzma_data,omitempty"`
	// Formerly used for bzip2 compressed data. Depreciated in 2010.
	OBSOLETEBzip2Data []byte `protobuf:"bytes,5,opt,name=OBSOLETE_bzip2_data,json=OBSOLETEBzip2Data" json:"OBSOLETE_bzip2_data,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *Blob) Reset()                    { *m = Blob{} }
func (m *Blob) String() string            { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()               {}
func (*Blob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Blob) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Blob) GetRawSize() int32 {
	if m != nil && m.RawSize != nil {
		return *m.RawSize
	}
	return 0
}

func (m *Blob) GetZlibData() []byte {
	if m != nil {
		return m.ZlibData
	}
	return nil
}

func (m *Blob) GetLzmaData() []byte {
	if m != nil {
		return m.LzmaData
	}
	return nil
}

func (m *Blob) GetOBSOLETEBzip2Data() []byte {
	if m != nil {
		return m.OBSOLETEBzip2Data
	}
	return nil
}

type BlobHeader struct {
	Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Indexdata        []byte  `protobuf:"bytes,2,opt,name=indexdata" json:"indexdata,omitempty"`
	Datasize         *int32  `protobuf:"varint,3,req,name=datasize" json:"datasize,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BlobHeader) Reset()                    { *m = BlobHeader{} }
func (m *BlobHeader) String() string            { return proto.CompactTextString(m) }
func (*BlobHeader) ProtoMessage()               {}
func (*BlobHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlobHeader) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *BlobHeader) GetIndexdata() []byte {
	if m != nil {
		return m.Indexdata
	}
	return nil
}

func (m *BlobHeader) GetDatasize() int32 {
	if m != nil && m.Datasize != nil {
		return *m.Datasize
	}
	return 0
}

func init() {
	proto.RegisterType((*Blob)(nil), "OSMPBF.Blob")
	proto.RegisterType((*BlobHeader)(nil), "OSMPBF.BlobHeader")
}

func init() { proto.RegisterFile("fileformat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcd, 0x4e, 0x84, 0x30,
	0x14, 0x85, 0x53, 0x7e, 0x14, 0x6e, 0x5c, 0x8c, 0x75, 0x83, 0x3f, 0x0b, 0x32, 0x2e, 0x64, 0xc5,
	0x62, 0x1e, 0xa1, 0x71, 0xcc, 0x2c, 0x34, 0x18, 0x70, 0xe5, 0x86, 0x5c, 0x42, 0xc7, 0x34, 0x01,
	0x4a, 0x4a, 0x13, 0x1c, 0xde, 0xc6, 0x37, 0x35, 0xb7, 0x44, 0x67, 0x75, 0x4f, 0xcf, 0x77, 0x4e,
	0xdb, 0x0b, 0x9b, 0xa3, 0xea, 0xe4, 0x51, 0x9b, 0x1e, 0x6d, 0x3e, 0x1a, 0x6d, 0x35, 0xbf, 0x28,
	0xaa, 0xb7, 0x77, 0xf1, 0xb2, 0xfd, 0x61, 0x10, 0x88, 0x4e, 0x37, 0x7c, 0x03, 0xbe, 0xc1, 0x39,
	0x61, 0x29, 0xcb, 0xae, 0x4a, 0x92, 0xfc, 0x16, 0x22, 0x83, 0x73, 0x3d, 0xa9, 0x45, 0x26, 0x5e,
	0xca, 0xb2, 0xb0, 0xbc, 0x34, 0x38, 0x57, 0x6a, 0x91, 0xfc, 0x1e, 0xe2, 0xa5, 0x53, 0x4d, 0xdd,
	0xa2, 0xc5, 0xc4, 0x77, 0x95, 0x88, 0x8c, 0x67, 0xb4, 0x48, 0xb0, 0x5b, 0x7a, 0x5c, 0x61, 0xb0,
	0x42, 0x32, 0x1c, 0xdc, 0xc1, 0x4d, 0x21, 0xaa, 0xe2, 0x75, 0xff, 0xb1, 0xaf, 0x9b, 0x45, 0x8d,
	0xbb, 0x35, 0x16, 0x52, 0x4c, 0x78, 0x09, 0x2b, 0xaf, 0xff, 0xb0, 0x20, 0x4a, 0x9d, 0xed, 0x27,
	0x00, 0x7d, 0xf1, 0x20, 0xb1, 0x95, 0x86, 0x73, 0x08, 0xec, 0x69, 0x94, 0x09, 0x4b, 0xbd, 0x2c,
	0x2e, 0x9d, 0xe6, 0x0f, 0x10, 0xab, 0xa1, 0x95, 0xdf, 0xee, 0x2e, 0xcf, 0x3d, 0x79, 0x36, 0xf8,
	0x1d, 0x44, 0x34, 0xdd, 0x22, 0x7e, 0xea, 0x65, 0x61, 0xf9, 0x7f, 0x16, 0x4f, 0xf0, 0xa8, 0xcd,
	0x57, 0xae, 0x47, 0x39, 0x4c, 0xd6, 0x48, 0x69, 0x7b, 0x1c, 0x73, 0x3d, 0xf5, 0x7a, 0x52, 0x13,
	0xcd, 0x46, 0x0d, 0x68, 0x4e, 0x07, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x59, 0x97, 0x03, 0xcf,
	0x43, 0x01, 0x00, 0x00,
}