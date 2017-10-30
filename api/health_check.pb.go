// Code generated by protoc-gen-go.
// source: api/health_check.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
}
var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}
func (HealthStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type HealthCheck struct {
	Timeout            *google_protobuf1.Duration   `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
	Interval           *google_protobuf1.Duration   `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	IntervalJitter     *google_protobuf1.Duration   `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter" json:"interval_jitter,omitempty"`
	UnhealthyThreshold *google_protobuf.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold" json:"healthy_threshold,omitempty"`
	AltPort            *google_protobuf.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort" json:"alt_port,omitempty"`
	ReuseConnection    *google_protobuf.BoolValue   `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_RedisHealthCheck_
	HealthChecker isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,oneof"`
}
type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,oneof"`
}
type HealthCheck_RedisHealthCheck_ struct {
	RedisHealthCheck *HealthCheck_RedisHealthCheck `protobuf:"bytes,10,opt,name=redis_health_check,json=redisHealthCheck,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker()  {}
func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker()   {}
func (*HealthCheck_RedisHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetTimeout() *google_protobuf1.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *google_protobuf1.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *google_protobuf1.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyThreshold() *google_protobuf.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *google_protobuf.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *google_protobuf.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *google_protobuf.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetRedisHealthCheck() *HealthCheck_RedisHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_RedisHealthCheck_); ok {
		return x.RedisHealthCheck
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheck_OneofMarshaler, _HealthCheck_OneofUnmarshaler, _HealthCheck_OneofSizer, []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_RedisHealthCheck_)(nil),
	}
}

func _HealthCheck_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheck)
	// health_checker
	switch x := m.HealthChecker.(type) {
	case *HealthCheck_HttpHealthCheck_:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpHealthCheck); err != nil {
			return err
		}
	case *HealthCheck_TcpHealthCheck_:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TcpHealthCheck); err != nil {
			return err
		}
	case *HealthCheck_RedisHealthCheck_:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RedisHealthCheck); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HealthCheck.HealthChecker has unexpected type %T", x)
	}
	return nil
}

func _HealthCheck_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheck)
	switch tag {
	case 8: // health_checker.http_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_HttpHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_HttpHealthCheck_{msg}
		return true, err
	case 9: // health_checker.tcp_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_TcpHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_TcpHealthCheck_{msg}
		return true, err
	case 10: // health_checker.redis_health_check
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheck_RedisHealthCheck)
		err := b.DecodeMessage(msg)
		m.HealthChecker = &HealthCheck_RedisHealthCheck_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheck_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheck)
	// health_checker
	switch x := m.HealthChecker.(type) {
	case *HealthCheck_HttpHealthCheck_:
		s := proto.Size(x.HttpHealthCheck)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheck_TcpHealthCheck_:
		s := proto.Size(x.TcpHealthCheck)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheck_RedisHealthCheck_:
		s := proto.Size(x.RedisHealthCheck)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
}

func (m *HealthCheck_Payload) Reset()                    { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()               {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 0} }

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,oneof"`
}
type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload()   {}
func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheck_Payload_OneofMarshaler, _HealthCheck_Payload_OneofUnmarshaler, _HealthCheck_Payload_OneofSizer, []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

func _HealthCheck_Payload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheck_Payload)
	// payload
	switch x := m.Payload.(type) {
	case *HealthCheck_Payload_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *HealthCheck_Payload_Binary:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Binary)
	case nil:
	default:
		return fmt.Errorf("HealthCheck_Payload.Payload has unexpected type %T", x)
	}
	return nil
}

func _HealthCheck_Payload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheck_Payload)
	switch tag {
	case 1: // payload.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &HealthCheck_Payload_Text{x}
		return true, err
	case 2: // payload.binary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &HealthCheck_Payload_Binary{x}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheck_Payload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheck_Payload)
	// payload
	switch x := m.Payload.(type) {
	case *HealthCheck_Payload_Text:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *HealthCheck_Payload_Binary:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Binary)))
		n += len(x.Binary)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HealthCheck_HttpHealthCheck struct {
	Host        string               `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Path        string               `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Send        *HealthCheck_Payload `protobuf:"bytes,3,opt,name=send" json:"send,omitempty"`
	Receive     *HealthCheck_Payload `protobuf:"bytes,4,opt,name=receive" json:"receive,omitempty"`
	ServiceName string               `protobuf:"bytes,5,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()                    { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()               {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 1} }

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type HealthCheck_TcpHealthCheck struct {
	Send    *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send" json:"send,omitempty"`
	Receive []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive" json:"receive,omitempty"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()                    { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()               {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 2} }

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
}

func (m *HealthCheck_RedisHealthCheck) Reset()                    { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()               {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 3} }

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v2.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.api.v2.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.api.v2.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.api.v2.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.api.v2.HealthCheck.RedisHealthCheck")
	proto.RegisterEnum("envoy.api.v2.HealthStatus", HealthStatus_name, HealthStatus_value)
}

func init() { proto.RegisterFile("api/health_check.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0x93, 0x36, 0x6f, 0x1d, 0x4f, 0xf2, 0x26, 0xee, 0x82, 0x90, 0x89, 0x10, 0x6a, 0x39,
	0x95, 0x1e, 0x5c, 0xa9, 0x55, 0xc5, 0x81, 0x53, 0xd3, 0x56, 0x38, 0x40, 0xdd, 0xca, 0x24, 0x54,
	0x70, 0xb1, 0xb6, 0xce, 0x50, 0x1b, 0x5c, 0xaf, 0xb5, 0x1e, 0x07, 0x72, 0xe7, 0xc3, 0xf0, 0x71,
	0xf8, 0x48, 0xc8, 0x6b, 0x3b, 0xca, 0x1f, 0x55, 0x0d, 0xdc, 0x76, 0x66, 0x9f, 0xe7, 0xb7, 0xe3,
	0xdd, 0xc7, 0xf0, 0x84, 0x27, 0xe1, 0x41, 0x80, 0x3c, 0xa2, 0xc0, 0xf3, 0x03, 0xf4, 0xbf, 0x59,
	0x89, 0x14, 0x24, 0x58, 0x1b, 0xe3, 0x89, 0x98, 0x5a, 0x3c, 0x09, 0xad, 0xc9, 0x61, 0xef, 0xf9,
	0xad, 0x10, 0xb7, 0x11, 0x1e, 0xa8, 0xbd, 0x9b, 0xec, 0xcb, 0xc1, 0x38, 0x93, 0x9c, 0x42, 0x11,
	0x17, 0xea, 0xd5, 0xfd, 0xef, 0x92, 0x27, 0x09, 0xca, 0xb4, 0xd8, 0x7f, 0xf1, 0x4b, 0x87, 0x96,
	0xad, 0x0e, 0x39, 0xcd, 0xcf, 0x60, 0x47, 0xa0, 0x51, 0x78, 0x87, 0x22, 0x23, 0xb3, 0xbe, 0x53,
	0xdf, 0x6b, 0x1d, 0x3e, 0xb5, 0x0a, 0x82, 0x55, 0x11, 0xac, 0xb3, 0xf2, 0x04, 0xb7, 0x52, 0xb2,
	0x63, 0x68, 0x86, 0x31, 0xa1, 0x9c, 0xf0, 0xc8, 0xdc, 0x78, 0xc8, 0x35, 0x93, 0xb2, 0x3e, 0x74,
	0xab, 0xb5, 0xf7, 0x35, 0x24, 0x42, 0x69, 0x6e, 0x3e, 0xe4, 0xee, 0x54, 0x8e, 0xb7, 0xca, 0xc0,
	0x2e, 0xe0, 0x51, 0x16, 0x17, 0xb7, 0x34, 0xf5, 0x28, 0x90, 0x98, 0x06, 0x22, 0x1a, 0x9b, 0x0d,
	0xc5, 0x79, 0xb6, 0xc2, 0x19, 0x0d, 0x62, 0x3a, 0x3a, 0xfc, 0xc8, 0xa3, 0x0c, 0x5d, 0x36, 0x33,
	0x0e, 0x2b, 0x1f, 0x1b, 0xc0, 0xf6, 0x2a, 0xec, 0xbf, 0x35, 0x60, 0xc6, 0x0a, 0xea, 0x15, 0x34,
	0x79, 0x44, 0x5e, 0x22, 0x24, 0x99, 0x5b, 0x6b, 0x10, 0x34, 0x1e, 0xd1, 0x95, 0x90, 0xc4, 0xce,
	0xc1, 0x90, 0x98, 0xa5, 0xe8, 0xf9, 0x22, 0x8e, 0xd1, 0xcf, 0x3f, 0xdb, 0xd4, 0x14, 0xa0, 0xb7,
	0x02, 0xe8, 0x0b, 0x11, 0x15, 0xf6, 0xae, 0xf2, 0x9c, 0xce, 0x2c, 0xec, 0x1a, 0xb6, 0x03, 0xa2,
	0xc4, 0x9b, 0x8f, 0x90, 0xd9, 0x54, 0x9c, 0x97, 0xd6, 0x7c, 0x86, 0xac, 0xb9, 0xf7, 0xb7, 0x6c,
	0xa2, 0x64, 0xae, 0xb6, 0x6b, 0x6e, 0x37, 0x58, 0x6c, 0xb1, 0x21, 0x18, 0xe4, 0x2f, 0x71, 0x75,
	0xc5, 0xdd, 0xbb, 0x9f, 0x3b, 0xf4, 0x97, 0xb0, 0x1d, 0x5a, 0xe8, 0xb0, 0xcf, 0xc0, 0x24, 0x8e,
	0xc3, 0x74, 0x91, 0x0b, 0x8a, 0xbb, 0x7f, 0x3f, 0xd7, 0xcd, 0x3d, 0x8b, 0x64, 0x43, 0x2e, 0xf5,
	0x7a, 0x67, 0xa0, 0x5d, 0xf1, 0x69, 0x24, 0xf8, 0x98, 0x3d, 0x86, 0x06, 0xe1, 0x8f, 0x22, 0xdc,
	0xba, 0x5d, 0x73, 0x55, 0xc5, 0x4c, 0xd8, 0xba, 0x09, 0x63, 0x2e, 0xa7, 0x2a, 0xbe, 0x6d, 0xbb,
	0xe6, 0x96, 0x75, 0x5f, 0x07, 0x2d, 0x29, 0xac, 0xbd, 0xdf, 0x75, 0xe8, 0x2e, 0x5d, 0x0f, 0x63,
	0xd0, 0x08, 0x44, 0x5a, 0xe2, 0x5c, 0xb5, 0xce, 0x7b, 0x09, 0xa7, 0x40, 0xa1, 0x74, 0x57, 0xad,
	0xd9, 0x31, 0x34, 0x52, 0x8c, 0xc7, 0x65, 0xbe, 0x77, 0xef, 0xff, 0x9e, 0x72, 0x4e, 0x57, 0xc9,
	0xd9, 0x6b, 0xd0, 0x24, 0xfa, 0x18, 0x4e, 0xb0, 0x4c, 0xf4, 0x1a, 0xce, 0xca, 0xc1, 0x76, 0xa1,
	0x9d, 0xa2, 0x9c, 0x84, 0x3e, 0x7a, 0x31, 0xbf, 0x43, 0x15, 0x63, 0xdd, 0x6d, 0x95, 0x3d, 0x87,
	0xdf, 0x61, 0xef, 0x67, 0x1d, 0x3a, 0x8b, 0x2f, 0x33, 0x9b, 0xb4, 0xfe, 0xcf, 0x93, 0x6e, 0xec,
	0x6c, 0xfe, 0xdd, 0xa4, 0x3d, 0x06, 0xc6, 0xf2, 0x3b, 0xf6, 0x0d, 0xe8, 0xcc, 0x27, 0x01, 0xe5,
	0xbe, 0x0b, 0xed, 0x42, 0xf0, 0x81, 0x38, 0x65, 0x29, 0x6b, 0x81, 0x36, 0x72, 0xde, 0x39, 0x97,
	0xd7, 0x8e, 0x51, 0xcb, 0x0b, 0xfb, 0xfc, 0xe4, 0xfd, 0xd0, 0xfe, 0x64, 0xd4, 0xd9, 0xff, 0xa0,
	0x8f, 0x9c, 0xaa, 0xdc, 0x60, 0x6d, 0x68, 0x9e, 0xb9, 0x27, 0x03, 0x67, 0xe0, 0xbc, 0x31, 0x36,
	0x73, 0xe5, 0x70, 0x70, 0x71, 0x7e, 0x39, 0x1a, 0x1a, 0x8d, 0x9b, 0x2d, 0xf5, 0x27, 0x1d, 0xfd,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x9f, 0x50, 0x3b, 0x6d, 0x05, 0x00, 0x00,
}