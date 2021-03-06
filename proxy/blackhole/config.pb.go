// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/blackhole/config.proto
// DO NOT EDIT!

/*
Package blackhole is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/blackhole/config.proto

It has these top-level messages:
	NoneResponse
	HTTPResponse
	Config
*/
package blackhole

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_loader "v2ray.com/core/common/loader"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NoneResponse struct {
}

func (m *NoneResponse) Reset()                    { *m = NoneResponse{} }
func (m *NoneResponse) String() string            { return proto.CompactTextString(m) }
func (*NoneResponse) ProtoMessage()               {}
func (*NoneResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HTTPResponse struct {
}

func (m *HTTPResponse) Reset()                    { *m = HTTPResponse{} }
func (m *HTTPResponse) String() string            { return proto.CompactTextString(m) }
func (*HTTPResponse) ProtoMessage()               {}
func (*HTTPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Config struct {
	Response *v2ray_core_common_loader.TypedSettings `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Config) GetResponse() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*NoneResponse)(nil), "v2ray.core.proxy.blackhole.NoneResponse")
	proto.RegisterType((*HTTPResponse)(nil), "v2ray.core.proxy.blackhole.HTTPResponse")
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.blackhole.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/blackhole/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x28, 0xca, 0xaf, 0xa8,
	0xd4, 0x4f, 0xca, 0x49, 0x4c, 0xce, 0xce, 0xc8, 0xcf, 0x49, 0xd5, 0x4f, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0x29, 0x2e, 0x4a, 0xd5, 0x03, 0x2b,
	0xd4, 0x83, 0x2b, 0x94, 0x52, 0x47, 0x33, 0x28, 0x39, 0x3f, 0x37, 0x37, 0x3f, 0x4f, 0x3f, 0x27,
	0x3f, 0x31, 0x25, 0xb5, 0x48, 0xbf, 0xa4, 0xb2, 0x20, 0x15, 0x62, 0x88, 0x12, 0x1f, 0x17, 0x8f,
	0x5f, 0x7e, 0x5e, 0x6a, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x88, 0xef, 0x11, 0x12,
	0x12, 0x00, 0xe7, 0xfb, 0x72, 0xb1, 0x39, 0x83, 0x2d, 0x15, 0x72, 0xe6, 0xe2, 0x28, 0x82, 0x8a,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xeb, 0x21, 0xb9, 0x00, 0x62, 0x83, 0x1e, 0xc4,
	0x06, 0xbd, 0x90, 0xca, 0x82, 0xd4, 0x94, 0xe0, 0xd4, 0x92, 0x92, 0xcc, 0xbc, 0xf4, 0xe2, 0x20,
	0xb8, 0x46, 0x27, 0x0b, 0x2e, 0xb9, 0xe4, 0xfc, 0x5c, 0x3d, 0xdc, 0x2e, 0x77, 0xe2, 0x86, 0x58,
	0x17, 0x00, 0x72, 0x5d, 0x14, 0x27, 0x5c, 0x3c, 0x89, 0x0d, 0xec, 0x5e, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x27, 0x0a, 0x95, 0x10, 0x23, 0x01, 0x00, 0x00,
}
