// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws"
	azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/azure"
	consul "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/consul"
	faultinjection "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/faultinjection"
	grpc "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/grpc"
	grpc_web "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/grpc_web"
	hcm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/hcm"
	kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"
	rest "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/rest"
	retries "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/retries"
	static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/static"
	stats "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/stats"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Plugin-specific configuration that lives on listeners
// Each ListenerPlugin object contains configuration for a specific plugin
// Note to developers: new Listener Plugins must be added to this struct
// to be usable by Gloo.
type ListenerPlugins struct {
	GrpcWeb                       *grpc_web.GrpcWeb                  `protobuf:"bytes,1,opt,name=grpc_web,json=grpcWeb,proto3" json:"grpc_web,omitempty"`
	HttpConnectionManagerSettings *hcm.HttpConnectionManagerSettings `protobuf:"bytes,2,opt,name=http_connection_manager_settings,json=httpConnectionManagerSettings,proto3" json:"http_connection_manager_settings,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                           `json:"-"`
	XXX_unrecognized              []byte                             `json:"-"`
	XXX_sizecache                 int32                              `json:"-"`
}

func (m *ListenerPlugins) Reset()         { *m = ListenerPlugins{} }
func (m *ListenerPlugins) String() string { return proto.CompactTextString(m) }
func (*ListenerPlugins) ProtoMessage()    {}
func (*ListenerPlugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{0}
}
func (m *ListenerPlugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerPlugins.Unmarshal(m, b)
}
func (m *ListenerPlugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerPlugins.Marshal(b, m, deterministic)
}
func (m *ListenerPlugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerPlugins.Merge(m, src)
}
func (m *ListenerPlugins) XXX_Size() int {
	return xxx_messageInfo_ListenerPlugins.Size(m)
}
func (m *ListenerPlugins) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerPlugins.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerPlugins proto.InternalMessageInfo

func (m *ListenerPlugins) GetGrpcWeb() *grpc_web.GrpcWeb {
	if m != nil {
		return m.GrpcWeb
	}
	return nil
}

func (m *ListenerPlugins) GetHttpConnectionManagerSettings() *hcm.HttpConnectionManagerSettings {
	if m != nil {
		return m.HttpConnectionManagerSettings
	}
	return nil
}

// Plugin-specific configuration that lives on virtual hosts
// Each VirtualHostPlugin object contains configuration for a specific plugin
// Note to developers: new Virtual Host Plugins must be added to this struct
// to be usable by Gloo.
type VirtualHostPlugins struct {
	Extensions           *Extensions          `protobuf:"bytes,1,opt,name=extensions,proto3" json:"extensions,omitempty"`
	Retries              *retries.RetryPolicy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	Stats                *stats.Stats         `protobuf:"bytes,10,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VirtualHostPlugins) Reset()         { *m = VirtualHostPlugins{} }
func (m *VirtualHostPlugins) String() string { return proto.CompactTextString(m) }
func (*VirtualHostPlugins) ProtoMessage()    {}
func (*VirtualHostPlugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{1}
}
func (m *VirtualHostPlugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHostPlugins.Unmarshal(m, b)
}
func (m *VirtualHostPlugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHostPlugins.Marshal(b, m, deterministic)
}
func (m *VirtualHostPlugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHostPlugins.Merge(m, src)
}
func (m *VirtualHostPlugins) XXX_Size() int {
	return xxx_messageInfo_VirtualHostPlugins.Size(m)
}
func (m *VirtualHostPlugins) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHostPlugins.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHostPlugins proto.InternalMessageInfo

func (m *VirtualHostPlugins) GetExtensions() *Extensions {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *VirtualHostPlugins) GetRetries() *retries.RetryPolicy {
	if m != nil {
		return m.Retries
	}
	return nil
}

func (m *VirtualHostPlugins) GetStats() *stats.Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Plugin-specific configuration that lives on routes
// Each RoutePlugin object contains configuration for a specific plugin
// Note to developers: new Route Plugins must be added to this struct
// to be usable by Gloo.
type RoutePlugins struct {
	Transformations      *transformation.RouteTransformations `protobuf:"bytes,1,opt,name=transformations,proto3" json:"transformations,omitempty"`
	Faults               *faultinjection.RouteFaults          `protobuf:"bytes,2,opt,name=faults,proto3" json:"faults,omitempty"`
	PrefixRewrite        *transformation.PrefixRewrite        `protobuf:"bytes,3,opt,name=prefix_rewrite,json=prefixRewrite,proto3" json:"prefix_rewrite,omitempty"`
	Timeout              *time.Duration                       `protobuf:"bytes,4,opt,name=timeout,proto3,stdduration" json:"timeout,omitempty"`
	Retries              *retries.RetryPolicy                 `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	Extensions           *Extensions                          `protobuf:"bytes,6,opt,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *RoutePlugins) Reset()         { *m = RoutePlugins{} }
func (m *RoutePlugins) String() string { return proto.CompactTextString(m) }
func (*RoutePlugins) ProtoMessage()    {}
func (*RoutePlugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{2}
}
func (m *RoutePlugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutePlugins.Unmarshal(m, b)
}
func (m *RoutePlugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutePlugins.Marshal(b, m, deterministic)
}
func (m *RoutePlugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutePlugins.Merge(m, src)
}
func (m *RoutePlugins) XXX_Size() int {
	return xxx_messageInfo_RoutePlugins.Size(m)
}
func (m *RoutePlugins) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutePlugins.DiscardUnknown(m)
}

var xxx_messageInfo_RoutePlugins proto.InternalMessageInfo

func (m *RoutePlugins) GetTransformations() *transformation.RouteTransformations {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *RoutePlugins) GetFaults() *faultinjection.RouteFaults {
	if m != nil {
		return m.Faults
	}
	return nil
}

func (m *RoutePlugins) GetPrefixRewrite() *transformation.PrefixRewrite {
	if m != nil {
		return m.PrefixRewrite
	}
	return nil
}

func (m *RoutePlugins) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RoutePlugins) GetRetries() *retries.RetryPolicy {
	if m != nil {
		return m.Retries
	}
	return nil
}

func (m *RoutePlugins) GetExtensions() *Extensions {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// Configuration for Destinations that are tied to the UpstreamSpec or ServiceSpec on that destination
type DestinationSpec struct {
	// Note to developers: new DestinationSpecs must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to DestinationType:
	//	*DestinationSpec_Aws
	//	*DestinationSpec_Azure
	//	*DestinationSpec_Rest
	//	*DestinationSpec_Grpc
	DestinationType      isDestinationSpec_DestinationType `protobuf_oneof:"destination_type"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{3}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

type isDestinationSpec_DestinationType interface {
	isDestinationSpec_DestinationType()
	Equal(interface{}) bool
}

type DestinationSpec_Aws struct {
	Aws *aws.DestinationSpec `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}
type DestinationSpec_Azure struct {
	Azure *azure.DestinationSpec `protobuf:"bytes,2,opt,name=azure,proto3,oneof"`
}
type DestinationSpec_Rest struct {
	Rest *rest.DestinationSpec `protobuf:"bytes,3,opt,name=rest,proto3,oneof"`
}
type DestinationSpec_Grpc struct {
	Grpc *grpc.DestinationSpec `protobuf:"bytes,4,opt,name=grpc,proto3,oneof"`
}

func (*DestinationSpec_Aws) isDestinationSpec_DestinationType()   {}
func (*DestinationSpec_Azure) isDestinationSpec_DestinationType() {}
func (*DestinationSpec_Rest) isDestinationSpec_DestinationType()  {}
func (*DestinationSpec_Grpc) isDestinationSpec_DestinationType()  {}

func (m *DestinationSpec) GetDestinationType() isDestinationSpec_DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return nil
}

func (m *DestinationSpec) GetAws() *aws.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *DestinationSpec) GetAzure() *azure.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *DestinationSpec) GetRest() *rest.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Rest); ok {
		return x.Rest
	}
	return nil
}

func (m *DestinationSpec) GetGrpc() *grpc.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DestinationSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DestinationSpec_OneofMarshaler, _DestinationSpec_OneofUnmarshaler, _DestinationSpec_OneofSizer, []interface{}{
		(*DestinationSpec_Aws)(nil),
		(*DestinationSpec_Azure)(nil),
		(*DestinationSpec_Rest)(nil),
		(*DestinationSpec_Grpc)(nil),
	}
}

func _DestinationSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DestinationSpec)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *DestinationSpec_Aws:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *DestinationSpec_Azure:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *DestinationSpec_Rest:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rest); err != nil {
			return err
		}
	case *DestinationSpec_Grpc:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Grpc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DestinationSpec.DestinationType has unexpected type %T", x)
	}
	return nil
}

func _DestinationSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DestinationSpec)
	switch tag {
	case 1: // destination_type.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(aws.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Aws{msg}
		return true, err
	case 2: // destination_type.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(azure.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Azure{msg}
		return true, err
	case 3: // destination_type.rest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(rest.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Rest{msg}
		return true, err
	case 4: // destination_type.grpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(grpc.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Grpc{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DestinationSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DestinationSpec)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *DestinationSpec_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Rest:
		s := proto.Size(x.Rest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Grpc:
		s := proto.Size(x.Grpc)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin.
type UpstreamSpec struct {
	SslConfig *UpstreamSslConfig `protobuf:"bytes,6,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Circuit breakers for this upstream. if not set, the defaults ones from the Gloo settings will be used.
	// if those are not set, [envoy's defaults](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
	// will be used.
	CircuitBreakers    *CircuitBreakerConfig `protobuf:"bytes,7,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	LoadBalancerConfig *LoadBalancerConfig   `protobuf:"bytes,8,opt,name=load_balancer_config,json=loadBalancerConfig,proto3" json:"load_balancer_config,omitempty"`
	ConnectionConfig   *ConnectionConfig     `protobuf:"bytes,9,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	// Note to developers: new Upstream Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to UpstreamType:
	//	*UpstreamSpec_Kube
	//	*UpstreamSpec_Static
	//	*UpstreamSpec_Aws
	//	*UpstreamSpec_Azure
	//	*UpstreamSpec_Consul
	UpstreamType         isUpstreamSpec_UpstreamType `protobuf_oneof:"upstream_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{4}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

type isUpstreamSpec_UpstreamType interface {
	isUpstreamSpec_UpstreamType()
	Equal(interface{}) bool
}

type UpstreamSpec_Kube struct {
	Kube *kubernetes.UpstreamSpec `protobuf:"bytes,1,opt,name=kube,proto3,oneof"`
}
type UpstreamSpec_Static struct {
	Static *static.UpstreamSpec `protobuf:"bytes,4,opt,name=static,proto3,oneof"`
}
type UpstreamSpec_Aws struct {
	Aws *aws.UpstreamSpec `protobuf:"bytes,2,opt,name=aws,proto3,oneof"`
}
type UpstreamSpec_Azure struct {
	Azure *azure.UpstreamSpec `protobuf:"bytes,3,opt,name=azure,proto3,oneof"`
}
type UpstreamSpec_Consul struct {
	Consul *consul.UpstreamSpec `protobuf:"bytes,5,opt,name=consul,proto3,oneof"`
}

func (*UpstreamSpec_Kube) isUpstreamSpec_UpstreamType()   {}
func (*UpstreamSpec_Static) isUpstreamSpec_UpstreamType() {}
func (*UpstreamSpec_Aws) isUpstreamSpec_UpstreamType()    {}
func (*UpstreamSpec_Azure) isUpstreamSpec_UpstreamType()  {}
func (*UpstreamSpec_Consul) isUpstreamSpec_UpstreamType() {}

func (m *UpstreamSpec) GetUpstreamType() isUpstreamSpec_UpstreamType {
	if m != nil {
		return m.UpstreamType
	}
	return nil
}

func (m *UpstreamSpec) GetSslConfig() *UpstreamSslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *UpstreamSpec) GetCircuitBreakers() *CircuitBreakerConfig {
	if m != nil {
		return m.CircuitBreakers
	}
	return nil
}

func (m *UpstreamSpec) GetLoadBalancerConfig() *LoadBalancerConfig {
	if m != nil {
		return m.LoadBalancerConfig
	}
	return nil
}

func (m *UpstreamSpec) GetConnectionConfig() *ConnectionConfig {
	if m != nil {
		return m.ConnectionConfig
	}
	return nil
}

func (m *UpstreamSpec) GetKube() *kubernetes.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *UpstreamSpec) GetStatic() *static.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Static); ok {
		return x.Static
	}
	return nil
}

func (m *UpstreamSpec) GetAws() *aws.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *UpstreamSpec) GetAzure() *azure.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *UpstreamSpec) GetConsul() *consul.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Consul); ok {
		return x.Consul
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpstreamSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpstreamSpec_OneofMarshaler, _UpstreamSpec_OneofUnmarshaler, _UpstreamSpec_OneofSizer, []interface{}{
		(*UpstreamSpec_Kube)(nil),
		(*UpstreamSpec_Static)(nil),
		(*UpstreamSpec_Aws)(nil),
		(*UpstreamSpec_Azure)(nil),
		(*UpstreamSpec_Consul)(nil),
	}
}

func _UpstreamSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpstreamSpec)
	// upstream_type
	switch x := m.UpstreamType.(type) {
	case *UpstreamSpec_Kube:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Kube); err != nil {
			return err
		}
	case *UpstreamSpec_Static:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Static); err != nil {
			return err
		}
	case *UpstreamSpec_Aws:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *UpstreamSpec_Azure:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *UpstreamSpec_Consul:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Consul); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UpstreamSpec.UpstreamType has unexpected type %T", x)
	}
	return nil
}

func _UpstreamSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpstreamSpec)
	switch tag {
	case 1: // upstream_type.kube
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Kube{msg}
		return true, err
	case 4: // upstream_type.static
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(static.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Static{msg}
		return true, err
	case 2: // upstream_type.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(aws.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Aws{msg}
		return true, err
	case 3: // upstream_type.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(azure.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Azure{msg}
		return true, err
	case 5: // upstream_type.consul
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(consul.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Consul{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UpstreamSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpstreamSpec)
	// upstream_type
	switch x := m.UpstreamType.(type) {
	case *UpstreamSpec_Kube:
		s := proto.Size(x.Kube)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Static:
		s := proto.Size(x.Static)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Consul:
		s := proto.Size(x.Consul)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ListenerPlugins)(nil), "gloo.solo.io.ListenerPlugins")
	proto.RegisterType((*VirtualHostPlugins)(nil), "gloo.solo.io.VirtualHostPlugins")
	proto.RegisterType((*RoutePlugins)(nil), "gloo.solo.io.RoutePlugins")
	proto.RegisterType((*DestinationSpec)(nil), "gloo.solo.io.DestinationSpec")
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.solo.io.UpstreamSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins.proto", fileDescriptor_ae47d2df5fad2a45)
}

var fileDescriptor_ae47d2df5fad2a45 = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4d, 0x73, 0xdc, 0x34,
	0x18, 0xc7, 0xd9, 0xbc, 0x6c, 0x5a, 0x91, 0xb2, 0x41, 0xd3, 0xc3, 0x92, 0x81, 0x34, 0xec, 0x01,
	0x9a, 0x32, 0x95, 0x21, 0xcc, 0xb4, 0xd0, 0x99, 0x42, 0xd8, 0x0d, 0x25, 0x03, 0xe9, 0x90, 0x71,
	0x78, 0x29, 0x5c, 0x76, 0xb4, 0x8a, 0xd6, 0x51, 0xe3, 0xb5, 0x3c, 0x92, 0x9c, 0x34, 0x9c, 0xf8,
	0x18, 0x1c, 0xf8, 0x00, 0x5c, 0xf8, 0x0a, 0x7c, 0x02, 0xee, 0x1c, 0x99, 0xe1, 0x93, 0x30, 0x92,
	0x1e, 0x3b, 0xb6, 0x59, 0x3a, 0xbb, 0xde, 0x1c, 0xd6, 0x2b, 0x5b, 0xcf, 0xff, 0x67, 0x49, 0xcf,
	0x5f, 0x7a, 0x8c, 0x1e, 0x45, 0xc2, 0x9c, 0x66, 0x23, 0xc2, 0xe4, 0x24, 0xd0, 0x32, 0x96, 0xf7,
	0x85, 0x0c, 0xa2, 0x58, 0xca, 0x20, 0x55, 0xf2, 0x39, 0x67, 0x46, 0xfb, 0x3b, 0x9a, 0x8a, 0xe0,
	0xfc, 0x83, 0x20, 0x8d, 0xb3, 0x48, 0x24, 0x9a, 0xa4, 0x4a, 0x1a, 0x89, 0xd7, 0x6d, 0x17, 0xb1,
	0x2a, 0x22, 0xe4, 0xe6, 0x9b, 0x91, 0x94, 0x51, 0xcc, 0x03, 0xd7, 0x37, 0xca, 0xc6, 0x81, 0x36,
	0x2a, 0x63, 0xc6, 0xc7, 0x6e, 0xde, 0x8e, 0x64, 0x24, 0x5d, 0x33, 0xb0, 0x2d, 0x78, 0xfa, 0x60,
	0xae, 0xb7, 0x6b, 0x1d, 0x83, 0xee, 0xf1, 0x5c, 0x3a, 0xfe, 0xc2, 0xf0, 0x44, 0x0b, 0x99, 0x0f,
	0x7c, 0xb3, 0x3f, 0x97, 0x9c, 0x09, 0xc5, 0x32, 0x61, 0x86, 0x23, 0xc5, 0xe9, 0x19, 0x57, 0xc0,
	0xd8, 0x9b, 0x8b, 0x11, 0x4b, 0x7a, 0x32, 0x1c, 0xd1, 0x98, 0x26, 0xac, 0x20, 0xcc, 0x37, 0x09,
	0x26, 0x93, 0x84, 0x33, 0x23, 0x64, 0xd2, 0x68, 0x12, 0x90, 0xb9, 0x80, 0x5e, 0xb8, 0x1f, 0x30,
	0xf6, 0x1b, 0x31, 0x14, 0xd7, 0xc6, 0x5d, 0x16, 0xa2, 0x44, 0x2a, 0x65, 0xee, 0x02, 0x94, 0xc3,
	0xc6, 0x94, 0xe1, 0x05, 0x1f, 0x15, 0x8d, 0x85, 0x56, 0xe7, 0x94, 0x4d, 0xec, 0x0f, 0x18, 0x4f,
	0x9a, 0xad, 0xf0, 0x4f, 0x99, 0xe2, 0xfe, 0x0a, 0x9c, 0x83, 0x46, 0x1c, 0x26, 0x13, 0x9d, 0xc5,
	0xf0, 0x07, 0xa4, 0xa3, 0x46, 0xa4, 0xb3, 0x6c, 0xc4, 0x55, 0xc2, 0x0d, 0x2f, 0x37, 0x81, 0xf8,
	0x65, 0x43, 0x07, 0x18, 0x25, 0x78, 0xf1, 0xbf, 0xd0, 0x3c, 0xb5, 0xa1, 0x46, 0x30, 0xf8, 0x5b,
	0x68, 0xe5, 0x2d, 0x02, 0xae, 0xc0, 0x79, 0xd6, 0x88, 0x63, 0x14, 0x4d, 0xf4, 0x58, 0xaa, 0x09,
	0xb5, 0xdb, 0x2d, 0x48, 0x15, 0x1f, 0x8b, 0x17, 0x43, 0xc5, 0x2f, 0x94, 0x30, 0xfc, 0x3a, 0xc9,
	0xd5, 0x5b, 0x20, 0x7f, 0xdd, 0x88, 0x3c, 0xa6, 0x59, 0x6c, 0x44, 0xf2, 0xdc, 0x1f, 0x11, 0xfe,
	0x16, 0x80, 0x5b, 0xf5, 0x83, 0xf9, 0x24, 0x53, 0xa5, 0x17, 0xf6, 0xfe, 0x6a, 0xa1, 0xce, 0xa1,
	0xd0, 0x86, 0x27, 0x5c, 0x1d, 0x79, 0x1c, 0xfe, 0x0c, 0xdd, 0xc8, 0x37, 0x54, 0xb7, 0xb5, 0xdd,
	0xba, 0xfb, 0xea, 0xee, 0x3b, 0xe4, 0x6a, 0x87, 0x41, 0x15, 0x28, 0x1f, 0xff, 0xe4, 0x0b, 0x95,
	0xb2, 0xef, 0xf9, 0x28, 0x5c, 0x8b, 0x7c, 0x03, 0xff, 0xdc, 0x42, 0xdb, 0xa7, 0xc6, 0xa4, 0xc3,
	0xab, 0x93, 0x6b, 0x38, 0xa1, 0x09, 0x8d, 0xb8, 0x1a, 0x6a, 0x6e, 0x8c, 0x48, 0x22, 0xdd, 0x5d,
	0x72, 0xec, 0x87, 0xc4, 0x6d, 0xba, 0x69, 0xd8, 0x03, 0x63, 0xd2, 0x41, 0x01, 0x78, 0xea, 0xf5,
	0xc7, 0x20, 0x0f, 0xdf, 0x3a, 0x7d, 0x59, 0x77, 0xef, 0xcf, 0x16, 0xc2, 0xdf, 0x09, 0x65, 0x32,
	0x1a, 0x1f, 0x48, 0x6d, 0xf2, 0xc9, 0x7d, 0x84, 0xd0, 0x55, 0x49, 0x80, 0xe9, 0x75, 0xab, 0xaf,
	0xfd, 0xbc, 0xe8, 0x0f, 0x4b, 0xb1, 0x78, 0x80, 0xd6, 0xc0, 0xf2, 0xdd, 0x55, 0x27, 0xdb, 0x21,
	0xc5, 0x16, 0x98, 0x36, 0xfa, 0x90, 0x1b, 0x75, 0x79, 0x24, 0x63, 0xc1, 0x2e, 0xc3, 0x5c, 0x89,
	0x1f, 0xa2, 0x55, 0xe7, 0xd1, 0x2e, 0x72, 0x88, 0xb7, 0x09, 0x38, 0x76, 0x1a, 0xe0, 0xd8, 0x76,
	0x85, 0x3e, 0xbe, 0xf7, 0xc7, 0x32, 0x5a, 0x0f, 0x65, 0x66, 0x78, 0x3e, 0x91, 0x67, 0xa8, 0x53,
	0xb5, 0x50, 0x3e, 0x1b, 0x42, 0x78, 0x72, 0x2e, 0x2f, 0x09, 0x4d, 0x05, 0x39, 0xdf, 0x25, 0x63,
	0x11, 0x1b, 0xae, 0x88, 0x5d, 0x2c, 0xe2, 0x00, 0xdf, 0x54, 0x55, 0x61, 0x1d, 0x83, 0x3f, 0x45,
	0x6d, 0x67, 0xa1, 0x3c, 0x43, 0xef, 0x12, 0x70, 0xd4, 0xd4, 0x59, 0x5a, 0xe4, 0x13, 0x17, 0x1e,
	0x82, 0x0c, 0xff, 0x80, 0x5e, 0xab, 0xee, 0x9b, 0xee, 0xb2, 0x03, 0xed, 0x92, 0xba, 0xe9, 0xa7,
	0x11, 0x8f, 0x9c, 0x34, 0xf4, 0xca, 0xf0, 0x56, 0x5a, 0xbe, 0xc5, 0x1f, 0xa3, 0x35, 0x23, 0x26,
	0x5c, 0x66, 0xa6, 0xbb, 0xe2, 0x98, 0x6f, 0x10, 0xef, 0x70, 0x92, 0x3b, 0x9c, 0xec, 0x83, 0xc3,
	0xfb, 0x2b, 0xbf, 0xfc, 0x7d, 0xa7, 0x15, 0xe6, 0xf1, 0xd7, 0x93, 0xbf, 0xaa, 0x7d, 0xda, 0xb3,
	0xdb, 0xa7, 0xf7, 0xfb, 0x12, 0xea, 0xec, 0x73, 0x6d, 0x44, 0xe2, 0x46, 0x77, 0x9c, 0x72, 0x86,
	0x1f, 0xa3, 0x65, 0x7a, 0x91, 0xe7, 0x6d, 0x87, 0xb8, 0xda, 0x3c, 0x6d, 0x28, 0x35, 0xdd, 0xc1,
	0x2b, 0xa1, 0xd5, 0xe1, 0x01, 0x5a, 0x75, 0xa5, 0x06, 0xf2, 0xf4, 0x1e, 0x81, 0xc2, 0x33, 0x1b,
	0xc2, 0x6b, 0xf1, 0x1e, 0x5a, 0xb1, 0xe5, 0x1c, 0x52, 0x74, 0x8f, 0xf8, 0xda, 0x3e, 0x1b, 0xc2,
	0x29, 0x2d, 0xc1, 0xee, 0x7b, 0x48, 0xc8, 0x3d, 0xe2, 0xeb, 0xfa, 0x8c, 0x04, 0x1b, 0xdc, 0xc7,
	0x68, 0xe3, 0xe4, 0xaa, 0x6b, 0x68, 0x2e, 0x53, 0xde, 0xfb, 0x75, 0x15, 0xad, 0x7f, 0x9b, 0x6a,
	0xa3, 0x38, 0x9d, 0xb8, 0xc5, 0xfa, 0x04, 0x21, 0xad, 0x63, 0x7b, 0xa2, 0x8c, 0x45, 0x04, 0x4b,
	0x7f, 0xa7, 0xca, 0x2f, 0xe2, 0x75, 0x3c, 0x70, 0x61, 0xe1, 0x4d, 0x9d, 0x37, 0xf1, 0x53, 0xb4,
	0x51, 0xfb, 0x9a, 0xd3, 0xdd, 0x35, 0x47, 0xe9, 0x55, 0x29, 0x03, 0x1f, 0xd5, 0xf7, 0x41, 0x00,
	0xea, 0xb0, 0xca, 0x53, 0x8d, 0x43, 0x74, 0xbb, 0xf2, 0x61, 0x97, 0x0f, 0xec, 0x86, 0x43, 0x6e,
	0x57, 0x91, 0x87, 0x92, 0x9e, 0xf4, 0x21, 0x10, 0x80, 0x38, 0xfe, 0xcf, 0x33, 0xfc, 0x15, 0x7a,
	0xbd, 0x74, 0x60, 0x02, 0xf0, 0xa6, 0x03, 0x6e, 0xd5, 0xc6, 0x58, 0x84, 0x01, 0x6e, 0x83, 0xd5,
	0x9e, 0xe0, 0x01, 0x5a, 0xb1, 0x15, 0x1f, 0xdc, 0x75, 0x9f, 0x94, 0xcb, 0xff, 0xb4, 0xe4, 0x94,
	0x17, 0xdb, 0x66, 0xc6, 0xc6, 0xe3, 0x01, 0x6a, 0xfb, 0xe2, 0x0c, 0xd9, 0xdd, 0x21, 0x79, 0xad,
	0x9e, 0x01, 0x01, 0x52, 0xfc, 0xc8, 0xdb, 0x7c, 0x09, 0x6a, 0xc9, 0xff, 0xda, 0xbc, 0x26, 0x77,
	0x1e, 0xdf, 0xcb, 0x3d, 0xee, 0xfd, 0x79, 0xf7, 0x65, 0x1e, 0xaf, 0xe9, 0xc1, 0xe0, 0x03, 0xd4,
	0xf6, 0xdf, 0x51, 0xc5, 0xb6, 0xcf, 0x3f, 0xab, 0x66, 0x99, 0x82, 0x8f, 0xed, 0x77, 0xd0, 0xad,
	0x0c, 0x7a, 0x9c, 0x3d, 0xfb, 0x0f, 0x7e, 0xfb, 0x67, 0xab, 0xf5, 0xe3, 0xfb, 0xb3, 0xd5, 0xeb,
	0xf4, 0x2c, 0x82, 0x9a, 0x3d, 0x6a, 0xbb, 0x73, 0xea, 0xc3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0xb1, 0x09, 0x7d, 0x7c, 0x0d, 0x00, 0x00,
}

func (this *ListenerPlugins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerPlugins)
	if !ok {
		that2, ok := that.(ListenerPlugins)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.GrpcWeb.Equal(that1.GrpcWeb) {
		return false
	}
	if !this.HttpConnectionManagerSettings.Equal(that1.HttpConnectionManagerSettings) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualHostPlugins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHostPlugins)
	if !ok {
		that2, ok := that.(VirtualHostPlugins)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	if !this.Retries.Equal(that1.Retries) {
		return false
	}
	if !this.Stats.Equal(that1.Stats) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RoutePlugins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoutePlugins)
	if !ok {
		that2, ok := that.(RoutePlugins)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Transformations.Equal(that1.Transformations) {
		return false
	}
	if !this.Faults.Equal(that1.Faults) {
		return false
	}
	if !this.PrefixRewrite.Equal(that1.PrefixRewrite) {
		return false
	}
	if this.Timeout != nil && that1.Timeout != nil {
		if *this.Timeout != *that1.Timeout {
			return false
		}
	} else if this.Timeout != nil {
		return false
	} else if that1.Timeout != nil {
		return false
	}
	if !this.Retries.Equal(that1.Retries) {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.DestinationType == nil {
		if this.DestinationType != nil {
			return false
		}
	} else if this.DestinationType == nil {
		return false
	} else if !this.DestinationType.Equal(that1.DestinationType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Aws)
	if !ok {
		that2, ok := that.(DestinationSpec_Aws)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *DestinationSpec_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Azure)
	if !ok {
		that2, ok := that.(DestinationSpec_Azure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *DestinationSpec_Rest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Rest)
	if !ok {
		that2, ok := that.(DestinationSpec_Rest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Rest.Equal(that1.Rest) {
		return false
	}
	return true
}
func (this *DestinationSpec_Grpc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Grpc)
	if !ok {
		that2, ok := that.(DestinationSpec_Grpc)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Grpc.Equal(that1.Grpc) {
		return false
	}
	return true
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.CircuitBreakers.Equal(that1.CircuitBreakers) {
		return false
	}
	if !this.LoadBalancerConfig.Equal(that1.LoadBalancerConfig) {
		return false
	}
	if !this.ConnectionConfig.Equal(that1.ConnectionConfig) {
		return false
	}
	if that1.UpstreamType == nil {
		if this.UpstreamType != nil {
			return false
		}
	} else if this.UpstreamType == nil {
		return false
	} else if !this.UpstreamType.Equal(that1.UpstreamType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Kube) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Kube)
	if !ok {
		that2, ok := that.(UpstreamSpec_Kube)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Kube.Equal(that1.Kube) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Static) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Static)
	if !ok {
		that2, ok := that.(UpstreamSpec_Static)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Static.Equal(that1.Static) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Aws)
	if !ok {
		that2, ok := that.(UpstreamSpec_Aws)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Azure)
	if !ok {
		that2, ok := that.(UpstreamSpec_Azure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Consul)
	if !ok {
		that2, ok := that.(UpstreamSpec_Consul)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Consul.Equal(that1.Consul) {
		return false
	}
	return true
}
