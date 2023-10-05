// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: chan_config.proto

package proto

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	msp "github.com/hyperledger/fabric-protos-go/msp"
	orderer "github.com/hyperledger/fabric-protos-go/orderer"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PolicyKey int32

const (
	PolicyKey_POLICY_KEY_UNDEFINED             PolicyKey = 0
	PolicyKey_POLICY_KEY_READERS               PolicyKey = 1
	PolicyKey_POLICY_KEY_WRITERS               PolicyKey = 2
	PolicyKey_POLICY_KEY_LIFECYCLE_ENDORSEMENT PolicyKey = 3
	PolicyKey_POLICY_KEY_ENDORSEMENT           PolicyKey = 4
)

// Enum value maps for PolicyKey.
var (
	PolicyKey_name = map[int32]string{
		0: "POLICY_KEY_UNDEFINED",
		1: "POLICY_KEY_READERS",
		2: "POLICY_KEY_WRITERS",
		3: "POLICY_KEY_LIFECYCLE_ENDORSEMENT",
		4: "POLICY_KEY_ENDORSEMENT",
	}
	PolicyKey_value = map[string]int32{
		"POLICY_KEY_UNDEFINED":             0,
		"POLICY_KEY_READERS":               1,
		"POLICY_KEY_WRITERS":               2,
		"POLICY_KEY_LIFECYCLE_ENDORSEMENT": 3,
		"POLICY_KEY_ENDORSEMENT":           4,
	}
)

func (x PolicyKey) Enum() *PolicyKey {
	p := new(PolicyKey)
	*p = x
	return p
}

func (x PolicyKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyKey) Descriptor() protoreflect.EnumDescriptor {
	return file_chan_config_proto_enumTypes[0].Descriptor()
}

func (PolicyKey) Type() protoreflect.EnumType {
	return &file_chan_config_proto_enumTypes[0]
}

func (x PolicyKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyKey.Descriptor instead.
func (PolicyKey) EnumDescriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{0}
}

type CertType int32

const (
	CertType_CERT_TYPE_UNDEFINED    CertType = 0
	CertType_CERT_TYPE_CA           CertType = 1
	CertType_CERT_TYPE_INTERMEDIATE CertType = 2
	CertType_CERT_TYPE_ADMIN        CertType = 3
)

// Enum value maps for CertType.
var (
	CertType_name = map[int32]string{
		0: "CERT_TYPE_UNDEFINED",
		1: "CERT_TYPE_CA",
		2: "CERT_TYPE_INTERMEDIATE",
		3: "CERT_TYPE_ADMIN",
	}
	CertType_value = map[string]int32{
		"CERT_TYPE_UNDEFINED":    0,
		"CERT_TYPE_CA":           1,
		"CERT_TYPE_INTERMEDIATE": 2,
		"CERT_TYPE_ADMIN":        3,
	}
)

func (x CertType) Enum() *CertType {
	p := new(CertType)
	*p = x
	return p
}

func (x CertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertType) Descriptor() protoreflect.EnumDescriptor {
	return file_chan_config_proto_enumTypes[1].Descriptor()
}

func (CertType) Type() protoreflect.EnumType {
	return &file_chan_config_proto_enumTypes[1]
}

func (x CertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertType.Descriptor instead.
func (CertType) EnumDescriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{1}
}

type ChannelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications              map[string]*ApplicationConfig     `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Orderers                  map[string]*OrdererConfig         `protobuf:"bytes,2,rep,name=orderers,proto3" json:"orderers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrdererBatchSize          *orderer.BatchSize                `protobuf:"bytes,3,opt,name=orderer_batch_size,json=ordererBatchSize,proto3" json:"orderer_batch_size,omitempty"`
	OrdererBatchTimeout       string                            `protobuf:"bytes,4,opt,name=orderer_batch_timeout,json=ordererBatchTimeout,proto3" json:"orderer_batch_timeout,omitempty"`
	OrdererConsensusType      *orderer.ConsensusType            `protobuf:"bytes,5,opt,name=orderer_consensus_type,json=ordererConsensusType,proto3" json:"orderer_consensus_type,omitempty"`
	Consortium                string                            `protobuf:"bytes,6,opt,name=consortium,proto3" json:"consortium,omitempty"`
	HashingAlgorithm          string                            `protobuf:"bytes,7,opt,name=hashing_algorithm,json=hashingAlgorithm,proto3" json:"hashing_algorithm,omitempty"`
	BlockDataHashingStructure *common.BlockDataHashingStructure `protobuf:"bytes,8,opt,name=block_data_hashing_structure,json=blockDataHashingStructure,proto3" json:"block_data_hashing_structure,omitempty"`
	Capabilities              *common.Capabilities              `protobuf:"bytes,9,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	Policy                    map[string]*Policy                `protobuf:"bytes,10,rep,name=policy,proto3" json:"policy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChannelConfig) Reset() {
	*x = ChannelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chan_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelConfig) ProtoMessage() {}

func (x *ChannelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_chan_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelConfig.ProtoReflect.Descriptor instead.
func (*ChannelConfig) Descriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelConfig) GetApplications() map[string]*ApplicationConfig {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *ChannelConfig) GetOrderers() map[string]*OrdererConfig {
	if x != nil {
		return x.Orderers
	}
	return nil
}

func (x *ChannelConfig) GetOrdererBatchSize() *orderer.BatchSize {
	if x != nil {
		return x.OrdererBatchSize
	}
	return nil
}

func (x *ChannelConfig) GetOrdererBatchTimeout() string {
	if x != nil {
		return x.OrdererBatchTimeout
	}
	return ""
}

func (x *ChannelConfig) GetOrdererConsensusType() *orderer.ConsensusType {
	if x != nil {
		return x.OrdererConsensusType
	}
	return nil
}

func (x *ChannelConfig) GetConsortium() string {
	if x != nil {
		return x.Consortium
	}
	return ""
}

func (x *ChannelConfig) GetHashingAlgorithm() string {
	if x != nil {
		return x.HashingAlgorithm
	}
	return ""
}

func (x *ChannelConfig) GetBlockDataHashingStructure() *common.BlockDataHashingStructure {
	if x != nil {
		return x.BlockDataHashingStructure
	}
	return nil
}

func (x *ChannelConfig) GetCapabilities() *common.Capabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *ChannelConfig) GetPolicy() map[string]*Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type MSP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config *msp.FabricMSPConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Policy map[string]*Policy   `protobuf:"bytes,3,rep,name=policy,proto3" json:"policy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MSP) Reset() {
	*x = MSP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chan_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSP) ProtoMessage() {}

func (x *MSP) ProtoReflect() protoreflect.Message {
	mi := &file_chan_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSP.ProtoReflect.Descriptor instead.
func (*MSP) Descriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{1}
}

func (x *MSP) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MSP) GetConfig() *msp.FabricMSPConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *MSP) GetPolicy() map[string]*Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type ApplicationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Msp         *MSP               `protobuf:"bytes,2,opt,name=msp,proto3" json:"msp,omitempty"`
	AnchorPeers []*peer.AnchorPeer `protobuf:"bytes,3,rep,name=anchor_peers,json=anchorPeers,proto3" json:"anchor_peers,omitempty"`
}

func (x *ApplicationConfig) Reset() {
	*x = ApplicationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chan_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationConfig) ProtoMessage() {}

func (x *ApplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_chan_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationConfig.ProtoReflect.Descriptor instead.
func (*ApplicationConfig) Descriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationConfig) GetMsp() *MSP {
	if x != nil {
		return x.Msp
	}
	return nil
}

func (x *ApplicationConfig) GetAnchorPeers() []*peer.AnchorPeer {
	if x != nil {
		return x.AnchorPeers
	}
	return nil
}

type OrdererConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Msp       *MSP     `protobuf:"bytes,2,opt,name=msp,proto3" json:"msp,omitempty"`
	Endpoints []string `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *OrdererConfig) Reset() {
	*x = OrdererConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chan_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdererConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdererConfig) ProtoMessage() {}

func (x *OrdererConfig) ProtoReflect() protoreflect.Message {
	mi := &file_chan_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdererConfig.ProtoReflect.Descriptor instead.
func (*OrdererConfig) Descriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{3}
}

func (x *OrdererConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrdererConfig) GetMsp() *MSP {
	if x != nil {
		return x.Msp
	}
	return nil
}

func (x *OrdererConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Policy:
	//
	//	*Policy_Implicit
	//	*Policy_SignaturePolicy
	Policy isPolicy_Policy `protobuf_oneof:"policy"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chan_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_chan_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{4}
}

func (m *Policy) GetPolicy() isPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *Policy) GetImplicit() *common.ImplicitMetaPolicy {
	if x, ok := x.GetPolicy().(*Policy_Implicit); ok {
		return x.Implicit
	}
	return nil
}

func (x *Policy) GetSignaturePolicy() *common.SignaturePolicyEnvelope {
	if x, ok := x.GetPolicy().(*Policy_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

type isPolicy_Policy interface {
	isPolicy_Policy()
}

type Policy_Implicit struct {
	Implicit *common.ImplicitMetaPolicy `protobuf:"bytes,1,opt,name=implicit,proto3,oneof"`
}

type Policy_SignaturePolicy struct {
	SignaturePolicy *common.SignaturePolicyEnvelope `protobuf:"bytes,2,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

func (*Policy_Implicit) isPolicy_Policy() {}

func (*Policy_SignaturePolicy) isPolicy_Policy() {}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sha256 hash
	Fingerprint []byte   `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Data        []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Type        CertType `protobuf:"varint,3,opt,name=type,proto3,enum=hlfsdkgo.proto.CertType" json:"type,omitempty"`
	MspId       string   `protobuf:"bytes,4,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	MspName     string   `protobuf:"bytes,5,opt,name=msp_name,json=mspName,proto3" json:"msp_name,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chan_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_chan_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_chan_config_proto_rawDescGZIP(), []int{5}
}

func (x *Certificate) GetFingerprint() []byte {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

func (x *Certificate) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Certificate) GetType() CertType {
	if x != nil {
		return x.Type
	}
	return CertType_CERT_TYPE_UNDEFINED
}

func (x *Certificate) GetMspId() string {
	if x != nil {
		return x.MspId
	}
	return ""
}

func (x *Certificate) GetMspName() string {
	if x != nil {
		return x.MspName
	}
	return ""
}

var File_chan_config_proto protoreflect.FileDescriptor

var file_chan_config_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x73, 0x70, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66,
	0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x07, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x53, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x65, 0x72, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x4c, 0x0a, 0x16, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x75, 0x6d, 0x12, 0x2b, 0x0a, 0x11, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x62, 0x0a, 0x1c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x19, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x62, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5a, 0x0a, 0x0d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd3, 0x01, 0x0a, 0x03, 0x4d,
	0x53, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x62,
	0x72, 0x69, 0x63, 0x4d, 0x53, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x53, 0x50, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x51, 0x0a,
	0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x85, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x53, 0x50, 0x52, 0x03, 0x6d, 0x73,
	0x70, 0x12, 0x35, 0x0a, 0x0c, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0b, 0x61, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0x68, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x03, 0x6d, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x6c, 0x66,
	0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x53, 0x50, 0x52,
	0x03, 0x6d, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x38, 0x0a,
	0x08, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x08, 0x69,
	0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22,
	0xa3, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x97, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x45, 0x52, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f,
	0x4b, 0x45, 0x59, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x53, 0x10, 0x02, 0x12, 0x24, 0x0a,
	0x20, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x4c, 0x49, 0x46, 0x45,
	0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x53, 0x45, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x53, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x2a,
	0x66, 0x0a, 0x08, 0x43, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x43,
	0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x41, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2f,
	0x68, 0x6c, 0x66, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chan_config_proto_rawDescOnce sync.Once
	file_chan_config_proto_rawDescData = file_chan_config_proto_rawDesc
)

func file_chan_config_proto_rawDescGZIP() []byte {
	file_chan_config_proto_rawDescOnce.Do(func() {
		file_chan_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_chan_config_proto_rawDescData)
	})
	return file_chan_config_proto_rawDescData
}

var file_chan_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chan_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chan_config_proto_goTypes = []interface{}{
	(PolicyKey)(0),                           // 0: hlfsdkgo.proto.PolicyKey
	(CertType)(0),                            // 1: hlfsdkgo.proto.CertType
	(*ChannelConfig)(nil),                    // 2: hlfsdkgo.proto.ChannelConfig
	(*MSP)(nil),                              // 3: hlfsdkgo.proto.MSP
	(*ApplicationConfig)(nil),                // 4: hlfsdkgo.proto.ApplicationConfig
	(*OrdererConfig)(nil),                    // 5: hlfsdkgo.proto.OrdererConfig
	(*Policy)(nil),                           // 6: hlfsdkgo.proto.Policy
	(*Certificate)(nil),                      // 7: hlfsdkgo.proto.Certificate
	nil,                                      // 8: hlfsdkgo.proto.ChannelConfig.ApplicationsEntry
	nil,                                      // 9: hlfsdkgo.proto.ChannelConfig.OrderersEntry
	nil,                                      // 10: hlfsdkgo.proto.ChannelConfig.PolicyEntry
	nil,                                      // 11: hlfsdkgo.proto.MSP.PolicyEntry
	(*orderer.BatchSize)(nil),                // 12: orderer.BatchSize
	(*orderer.ConsensusType)(nil),            // 13: orderer.ConsensusType
	(*common.BlockDataHashingStructure)(nil), // 14: common.BlockDataHashingStructure
	(*common.Capabilities)(nil),              // 15: common.Capabilities
	(*msp.FabricMSPConfig)(nil),              // 16: msp.FabricMSPConfig
	(*peer.AnchorPeer)(nil),                  // 17: protos.AnchorPeer
	(*common.ImplicitMetaPolicy)(nil),        // 18: common.ImplicitMetaPolicy
	(*common.SignaturePolicyEnvelope)(nil),   // 19: common.SignaturePolicyEnvelope
}
var file_chan_config_proto_depIdxs = []int32{
	8,  // 0: hlfsdkgo.proto.ChannelConfig.applications:type_name -> hlfsdkgo.proto.ChannelConfig.ApplicationsEntry
	9,  // 1: hlfsdkgo.proto.ChannelConfig.orderers:type_name -> hlfsdkgo.proto.ChannelConfig.OrderersEntry
	12, // 2: hlfsdkgo.proto.ChannelConfig.orderer_batch_size:type_name -> orderer.BatchSize
	13, // 3: hlfsdkgo.proto.ChannelConfig.orderer_consensus_type:type_name -> orderer.ConsensusType
	14, // 4: hlfsdkgo.proto.ChannelConfig.block_data_hashing_structure:type_name -> common.BlockDataHashingStructure
	15, // 5: hlfsdkgo.proto.ChannelConfig.capabilities:type_name -> common.Capabilities
	10, // 6: hlfsdkgo.proto.ChannelConfig.policy:type_name -> hlfsdkgo.proto.ChannelConfig.PolicyEntry
	16, // 7: hlfsdkgo.proto.MSP.config:type_name -> msp.FabricMSPConfig
	11, // 8: hlfsdkgo.proto.MSP.policy:type_name -> hlfsdkgo.proto.MSP.PolicyEntry
	3,  // 9: hlfsdkgo.proto.ApplicationConfig.msp:type_name -> hlfsdkgo.proto.MSP
	17, // 10: hlfsdkgo.proto.ApplicationConfig.anchor_peers:type_name -> protos.AnchorPeer
	3,  // 11: hlfsdkgo.proto.OrdererConfig.msp:type_name -> hlfsdkgo.proto.MSP
	18, // 12: hlfsdkgo.proto.Policy.implicit:type_name -> common.ImplicitMetaPolicy
	19, // 13: hlfsdkgo.proto.Policy.signature_policy:type_name -> common.SignaturePolicyEnvelope
	1,  // 14: hlfsdkgo.proto.Certificate.type:type_name -> hlfsdkgo.proto.CertType
	4,  // 15: hlfsdkgo.proto.ChannelConfig.ApplicationsEntry.value:type_name -> hlfsdkgo.proto.ApplicationConfig
	5,  // 16: hlfsdkgo.proto.ChannelConfig.OrderersEntry.value:type_name -> hlfsdkgo.proto.OrdererConfig
	6,  // 17: hlfsdkgo.proto.ChannelConfig.PolicyEntry.value:type_name -> hlfsdkgo.proto.Policy
	6,  // 18: hlfsdkgo.proto.MSP.PolicyEntry.value:type_name -> hlfsdkgo.proto.Policy
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_chan_config_proto_init() }
func file_chan_config_proto_init() {
	if File_chan_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chan_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chan_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chan_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chan_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdererConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chan_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chan_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_chan_config_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Policy_Implicit)(nil),
		(*Policy_SignaturePolicy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chan_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chan_config_proto_goTypes,
		DependencyIndexes: file_chan_config_proto_depIdxs,
		EnumInfos:         file_chan_config_proto_enumTypes,
		MessageInfos:      file_chan_config_proto_msgTypes,
	}.Build()
	File_chan_config_proto = out.File
	file_chan_config_proto_rawDesc = nil
	file_chan_config_proto_goTypes = nil
	file_chan_config_proto_depIdxs = nil
}
