// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/http/router.proto

package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_filter1 "github.com/envoyproxy/go-control-plane/api/filter"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Router struct {
	DynamicStats   *google_protobuf1.BoolValue       `protobuf:"bytes,1,opt,name=dynamic_stats,json=dynamicStats" json:"dynamic_stats,omitempty"`
	StartChildSpan bool                              `protobuf:"varint,2,opt,name=start_child_span,json=startChildSpan" json:"start_child_span,omitempty"`
	UpstreamLog    []*envoy_api_v2_filter1.AccessLog `protobuf:"bytes,3,rep,name=upstream_log,json=upstreamLog" json:"upstream_log,omitempty"`
}

func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
func (*Router) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Router) GetDynamicStats() *google_protobuf1.BoolValue {
	if m != nil {
		return m.DynamicStats
	}
	return nil
}

func (m *Router) GetStartChildSpan() bool {
	if m != nil {
		return m.StartChildSpan
	}
	return false
}

func (m *Router) GetUpstreamLog() []*envoy_api_v2_filter1.AccessLog {
	if m != nil {
		return m.UpstreamLog
	}
	return nil
}

func init() {
	proto.RegisterType((*Router)(nil), "envoy.api.v2.filter.http.Router")
}

func init() { proto.RegisterFile("api/filter/http/router.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x85, 0x22, 0x49, 0x15, 0xc9, 0x29, 0x04, 0x29, 0xc1, 0x53, 0x4e, 0xb3, 0x10,
	0x7f, 0x80, 0x54, 0xaf, 0x3d, 0xa5, 0xe0, 0x35, 0x4c, 0xd3, 0xed, 0x76, 0x61, 0x9b, 0x19, 0x76,
	0x27, 0x95, 0xfe, 0x31, 0x7f, 0x9f, 0x64, 0x63, 0xc1, 0x83, 0xc7, 0x79, 0xf3, 0xbe, 0x37, 0xf3,
	0xd2, 0x67, 0x64, 0xab, 0x8e, 0xd6, 0x89, 0xf6, 0xea, 0x24, 0xc2, 0xca, 0xd3, 0x28, 0xda, 0x03,
	0x7b, 0x12, 0xca, 0x0b, 0x3d, 0x5c, 0xe8, 0x0a, 0xc8, 0x16, 0x2e, 0x0d, 0xcc, 0x36, 0x98, 0x6c,
	0x65, 0xf9, 0x87, 0xc3, 0xbe, 0xd7, 0x21, 0x38, 0x32, 0x33, 0x55, 0xae, 0x0d, 0x91, 0x71, 0x5a,
	0xc5, 0x69, 0x3f, 0x1e, 0xd5, 0x97, 0x47, 0x66, 0xed, 0xc3, 0xbc, 0x7f, 0xf9, 0x4e, 0xd2, 0x65,
	0x1b, 0xcf, 0xe4, 0x6f, 0xe9, 0xc3, 0xe1, 0x3a, 0xe0, 0xd9, 0xf6, 0x5d, 0x10, 0x94, 0x50, 0x24,
	0x55, 0x52, 0x67, 0x4d, 0x09, 0x73, 0x04, 0xdc, 0x22, 0xe0, 0x9d, 0xc8, 0x7d, 0xa2, 0x1b, 0x75,
	0xbb, 0xfa, 0x05, 0x76, 0x93, 0x3f, 0xaf, 0xd3, 0xa7, 0x20, 0xe8, 0xa5, 0xeb, 0x4f, 0xd6, 0x1d,
	0xba, 0xc0, 0x38, 0x14, 0x77, 0x55, 0x52, 0xdf, 0xb7, 0x8f, 0x51, 0xff, 0x98, 0xe4, 0x1d, 0xe3,
	0x90, 0x6f, 0xd2, 0xd5, 0xc8, 0x41, 0xbc, 0xc6, 0x73, 0xe7, 0xc8, 0x14, 0x8b, 0x6a, 0x51, 0x67,
	0xcd, 0x1a, 0xfe, 0xab, 0xb8, 0x89, 0x8d, 0xb6, 0x64, 0xda, 0xec, 0xc6, 0x6c, 0xc9, 0xec, 0x97,
	0xf1, 0x9d, 0xd7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xbf, 0x42, 0x0f, 0x35, 0x01, 0x00,
	0x00,
}
