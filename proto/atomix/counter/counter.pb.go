// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/counter/counter.proto

package counter

import (
	context "context"
	fmt "fmt"
	headers "github.com/atomix/atomix-k8s-controller/proto/atomix/headers"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type CreateResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type CloseRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Delete               bool                   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{2}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CloseRequest) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

type CloseResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{3}
}

func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func (m *CloseResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type IncrementRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Delta                int64                  `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IncrementRequest) Reset()         { *m = IncrementRequest{} }
func (m *IncrementRequest) String() string { return proto.CompactTextString(m) }
func (*IncrementRequest) ProtoMessage()    {}
func (*IncrementRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{4}
}

func (m *IncrementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementRequest.Unmarshal(m, b)
}
func (m *IncrementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementRequest.Marshal(b, m, deterministic)
}
func (m *IncrementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementRequest.Merge(m, src)
}
func (m *IncrementRequest) XXX_Size() int {
	return xxx_messageInfo_IncrementRequest.Size(m)
}
func (m *IncrementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementRequest proto.InternalMessageInfo

func (m *IncrementRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *IncrementRequest) GetDelta() int64 {
	if m != nil {
		return m.Delta
	}
	return 0
}

type IncrementResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	PreviousValue        int64                   `protobuf:"varint,2,opt,name=previous_value,json=previousValue,proto3" json:"previous_value,omitempty"`
	NextValue            int64                   `protobuf:"varint,3,opt,name=next_value,json=nextValue,proto3" json:"next_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *IncrementResponse) Reset()         { *m = IncrementResponse{} }
func (m *IncrementResponse) String() string { return proto.CompactTextString(m) }
func (*IncrementResponse) ProtoMessage()    {}
func (*IncrementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{5}
}

func (m *IncrementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementResponse.Unmarshal(m, b)
}
func (m *IncrementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementResponse.Marshal(b, m, deterministic)
}
func (m *IncrementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementResponse.Merge(m, src)
}
func (m *IncrementResponse) XXX_Size() int {
	return xxx_messageInfo_IncrementResponse.Size(m)
}
func (m *IncrementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementResponse proto.InternalMessageInfo

func (m *IncrementResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *IncrementResponse) GetPreviousValue() int64 {
	if m != nil {
		return m.PreviousValue
	}
	return 0
}

func (m *IncrementResponse) GetNextValue() int64 {
	if m != nil {
		return m.NextValue
	}
	return 0
}

type DecrementRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Delta                int64                  `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DecrementRequest) Reset()         { *m = DecrementRequest{} }
func (m *DecrementRequest) String() string { return proto.CompactTextString(m) }
func (*DecrementRequest) ProtoMessage()    {}
func (*DecrementRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{6}
}

func (m *DecrementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecrementRequest.Unmarshal(m, b)
}
func (m *DecrementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecrementRequest.Marshal(b, m, deterministic)
}
func (m *DecrementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecrementRequest.Merge(m, src)
}
func (m *DecrementRequest) XXX_Size() int {
	return xxx_messageInfo_DecrementRequest.Size(m)
}
func (m *DecrementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecrementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecrementRequest proto.InternalMessageInfo

func (m *DecrementRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DecrementRequest) GetDelta() int64 {
	if m != nil {
		return m.Delta
	}
	return 0
}

type DecrementResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	PreviousValue        int64                   `protobuf:"varint,2,opt,name=previous_value,json=previousValue,proto3" json:"previous_value,omitempty"`
	NextValue            int64                   `protobuf:"varint,3,opt,name=next_value,json=nextValue,proto3" json:"next_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DecrementResponse) Reset()         { *m = DecrementResponse{} }
func (m *DecrementResponse) String() string { return proto.CompactTextString(m) }
func (*DecrementResponse) ProtoMessage()    {}
func (*DecrementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{7}
}

func (m *DecrementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecrementResponse.Unmarshal(m, b)
}
func (m *DecrementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecrementResponse.Marshal(b, m, deterministic)
}
func (m *DecrementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecrementResponse.Merge(m, src)
}
func (m *DecrementResponse) XXX_Size() int {
	return xxx_messageInfo_DecrementResponse.Size(m)
}
func (m *DecrementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecrementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecrementResponse proto.InternalMessageInfo

func (m *DecrementResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DecrementResponse) GetPreviousValue() int64 {
	if m != nil {
		return m.PreviousValue
	}
	return 0
}

func (m *DecrementResponse) GetNextValue() int64 {
	if m != nil {
		return m.NextValue
	}
	return 0
}

type GetRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{8}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Value                int64                   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{9}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SetRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Value                int64                  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{10}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SetRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SetResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	PreviousValue        int64                   `protobuf:"varint,2,opt,name=previous_value,json=previousValue,proto3" json:"previous_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{11}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SetResponse) GetPreviousValue() int64 {
	if m != nil {
		return m.PreviousValue
	}
	return 0
}

type CheckAndSetRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Expect               int64                  `protobuf:"varint,2,opt,name=expect,proto3" json:"expect,omitempty"`
	Update               int64                  `protobuf:"varint,3,opt,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CheckAndSetRequest) Reset()         { *m = CheckAndSetRequest{} }
func (m *CheckAndSetRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAndSetRequest) ProtoMessage()    {}
func (*CheckAndSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{12}
}

func (m *CheckAndSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAndSetRequest.Unmarshal(m, b)
}
func (m *CheckAndSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAndSetRequest.Marshal(b, m, deterministic)
}
func (m *CheckAndSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAndSetRequest.Merge(m, src)
}
func (m *CheckAndSetRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAndSetRequest.Size(m)
}
func (m *CheckAndSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAndSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAndSetRequest proto.InternalMessageInfo

func (m *CheckAndSetRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CheckAndSetRequest) GetExpect() int64 {
	if m != nil {
		return m.Expect
	}
	return 0
}

func (m *CheckAndSetRequest) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

type CheckAndSetResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Succeeded            bool                    `protobuf:"varint,2,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CheckAndSetResponse) Reset()         { *m = CheckAndSetResponse{} }
func (m *CheckAndSetResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAndSetResponse) ProtoMessage()    {}
func (*CheckAndSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc83a5450a4ecc2, []int{13}
}

func (m *CheckAndSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAndSetResponse.Unmarshal(m, b)
}
func (m *CheckAndSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAndSetResponse.Marshal(b, m, deterministic)
}
func (m *CheckAndSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAndSetResponse.Merge(m, src)
}
func (m *CheckAndSetResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAndSetResponse.Size(m)
}
func (m *CheckAndSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAndSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAndSetResponse proto.InternalMessageInfo

func (m *CheckAndSetResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CheckAndSetResponse) GetSucceeded() bool {
	if m != nil {
		return m.Succeeded
	}
	return false
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "atomix.counter.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "atomix.counter.CreateResponse")
	proto.RegisterType((*CloseRequest)(nil), "atomix.counter.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "atomix.counter.CloseResponse")
	proto.RegisterType((*IncrementRequest)(nil), "atomix.counter.IncrementRequest")
	proto.RegisterType((*IncrementResponse)(nil), "atomix.counter.IncrementResponse")
	proto.RegisterType((*DecrementRequest)(nil), "atomix.counter.DecrementRequest")
	proto.RegisterType((*DecrementResponse)(nil), "atomix.counter.DecrementResponse")
	proto.RegisterType((*GetRequest)(nil), "atomix.counter.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "atomix.counter.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "atomix.counter.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "atomix.counter.SetResponse")
	proto.RegisterType((*CheckAndSetRequest)(nil), "atomix.counter.CheckAndSetRequest")
	proto.RegisterType((*CheckAndSetResponse)(nil), "atomix.counter.CheckAndSetResponse")
}

func init() { proto.RegisterFile("atomix/counter/counter.proto", fileDescriptor_6fc83a5450a4ecc2) }

var fileDescriptor_6fc83a5450a4ecc2 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x37, 0xd6, 0x06, 0x7b, 0x6a, 0x83, 0x8e, 0xeb, 0xb2, 0xc6, 0x76, 0xa9, 0x23, 0xc2,
	0x5e, 0x55, 0x58, 0xd1, 0x6b, 0x35, 0xb2, 0xd9, 0xbd, 0x93, 0x04, 0x44, 0x11, 0x59, 0x62, 0x72,
	0x60, 0xc3, 0x66, 0x33, 0x31, 0x99, 0x94, 0x82, 0x4f, 0xe1, 0x63, 0xf8, 0x96, 0x92, 0xcc, 0x4c,
	0x36, 0x4d, 0xda, 0xde, 0x8c, 0xca, 0x5e, 0x85, 0x73, 0xe6, 0xef, 0x77, 0xce, 0xdf, 0x09, 0x7f,
	0x60, 0x1a, 0x70, 0x76, 0x1d, 0xaf, 0x5e, 0x86, 0xac, 0x4c, 0x39, 0xe6, 0xea, 0xb9, 0xc8, 0x72,
	0xc6, 0x19, 0xb1, 0xc4, 0xe9, 0x42, 0x76, 0x6d, 0xa5, 0xbe, 0xc4, 0x20, 0xc2, 0xbc, 0x50, 0x4f,
	0xa1, 0xa6, 0xa7, 0x30, 0x71, 0x72, 0x0c, 0x38, 0x7a, 0xf8, 0xa3, 0xc4, 0x82, 0x93, 0xd7, 0x60,
	0x0a, 0xc5, 0xa1, 0x31, 0x37, 0x8e, 0xc7, 0x27, 0xb3, 0x85, 0xe4, 0xa9, 0xdf, 0x49, 0xe1, 0x59,
	0x5d, 0x7a, 0x52, 0x4c, 0xcf, 0xc0, 0x52, 0x9c, 0x22, 0x63, 0x69, 0x81, 0xe4, 0x4d, 0x07, 0x74,
	0xd4, 0x07, 0x09, 0x65, 0x87, 0xf4, 0x0d, 0xee, 0x3b, 0x09, 0x2b, 0x34, 0x17, 0x22, 0x07, 0x60,
	0x46, 0x98, 0x20, 0xc7, 0xc3, 0x3b, 0x73, 0xe3, 0xf8, 0x9e, 0x27, 0x2b, 0xea, 0xc2, 0x44, 0xe2,
	0x35, 0xf7, 0xbc, 0x80, 0x07, 0xe7, 0x69, 0x98, 0xe3, 0x35, 0xa6, 0x5c, 0x73, 0xd7, 0x7d, 0x18,
	0x46, 0x98, 0xf0, 0xa0, 0x5e, 0x75, 0xe0, 0x89, 0x82, 0xfe, 0x32, 0xe0, 0x61, 0x6b, 0x82, 0xde,
	0xba, 0xe4, 0x05, 0x58, 0x59, 0x8e, 0xcb, 0x98, 0x95, 0xc5, 0xc5, 0x32, 0x48, 0x4a, 0x94, 0xc3,
	0x26, 0xaa, 0xfb, 0xa9, 0x6a, 0x92, 0x19, 0x40, 0x8a, 0x2b, 0x2e, 0x25, 0x83, 0x5a, 0x32, 0xaa,
	0x3a, 0xf5, 0x71, 0x65, 0xfa, 0x03, 0xfe, 0x6b, 0xd3, 0xad, 0x09, 0xb7, 0xc2, 0xb4, 0x03, 0xe0,
	0xa2, 0xa6, 0x5d, 0xfa, 0x15, 0xc6, 0x35, 0x44, 0xd3, 0xd1, 0x3e, 0x0c, 0xdb, 0x46, 0x44, 0x41,
	0xbf, 0x00, 0xf8, 0xf8, 0x17, 0x2e, 0x64, 0x03, 0x3a, 0x81, 0xb1, 0x8f, 0xff, 0xeb, 0x26, 0xe8,
	0x4f, 0x20, 0xce, 0x25, 0x86, 0x57, 0xef, 0xd2, 0x48, 0xdf, 0xd0, 0x01, 0x98, 0xb8, 0xca, 0x30,
	0xe4, 0x72, 0x96, 0xac, 0xaa, 0x7e, 0x99, 0x45, 0x01, 0x57, 0x57, 0x2d, 0x2b, 0x7a, 0x05, 0x8f,
	0xd6, 0x86, 0x6b, 0x5a, 0x9e, 0xc2, 0xa8, 0x28, 0xc3, 0x10, 0x31, 0xc2, 0x48, 0x86, 0xd0, 0x4d,
	0xe3, 0xe4, 0xf7, 0x5d, 0xb0, 0x1c, 0x11, 0xd1, 0x3e, 0xe6, 0xcb, 0x38, 0x44, 0x72, 0x0e, 0xa6,
	0xc8, 0x50, 0xd2, 0x18, 0x54, 0xd1, 0xbe, 0x96, 0xd1, 0xf6, 0xd1, 0xb6, 0x63, 0xb1, 0x07, 0xdd,
	0x23, 0xa7, 0x30, 0xac, 0x53, 0x8e, 0x4c, 0x7b, 0xd2, 0x56, 0xb6, 0xda, 0xb3, 0x2d, 0xa7, 0x0d,
	0xe7, 0x2d, 0x0c, 0x7c, 0xe4, 0xc4, 0xee, 0xea, 0x6e, 0x2e, 0xc7, 0x7e, 0xba, 0xf1, 0xac, 0x4d,
	0x70, 0x37, 0x11, 0xdc, 0x1d, 0x04, 0x77, 0x8d, 0xe0, 0xc1, 0xa8, 0x89, 0x41, 0x32, 0xef, 0x6a,
	0xbb, 0x19, 0x6c, 0x3f, 0xdb, 0xa1, 0x68, 0x33, 0x9b, 0x94, 0xe9, 0x33, 0xbb, 0x11, 0xd7, 0x67,
	0xf6, 0x22, 0x8a, 0xee, 0x91, 0xcf, 0x30, 0x6e, 0xbd, 0x3e, 0x84, 0xf6, 0xfe, 0xdb, 0xde, 0x8b,
	0x6d, 0x3f, 0xdf, 0xa9, 0x51, 0xe4, 0xf7, 0x4f, 0xe0, 0x71, 0xcc, 0x94, 0x34, 0xc8, 0x62, 0x25,
	0xff, 0x68, 0x7c, 0x37, 0xeb, 0xcf, 0xf8, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x07,
	0x98, 0x0c, 0x14, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Increment(ctx context.Context, in *IncrementRequest, opts ...grpc.CallOption) (*IncrementResponse, error)
	Decrement(ctx context.Context, in *DecrementRequest, opts ...grpc.CallOption) (*DecrementResponse, error)
	CheckAndSet(ctx context.Context, in *CheckAndSetRequest, opts ...grpc.CallOption) (*CheckAndSetResponse, error)
}

type counterServiceClient struct {
	cc *grpc.ClientConn
}

func NewCounterServiceClient(cc *grpc.ClientConn) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) Increment(ctx context.Context, in *IncrementRequest, opts ...grpc.CallOption) (*IncrementResponse, error) {
	out := new(IncrementResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) Decrement(ctx context.Context, in *DecrementRequest, opts ...grpc.CallOption) (*DecrementResponse, error) {
	out := new(DecrementResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/Decrement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) CheckAndSet(ctx context.Context, in *CheckAndSetRequest, opts ...grpc.CallOption) (*CheckAndSetResponse, error) {
	out := new(CheckAndSetResponse)
	err := c.cc.Invoke(ctx, "/atomix.counter.CounterService/CheckAndSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
type CounterServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Increment(context.Context, *IncrementRequest) (*IncrementResponse, error)
	Decrement(context.Context, *DecrementRequest) (*DecrementResponse, error)
	CheckAndSet(context.Context, *CheckAndSetRequest) (*CheckAndSetResponse, error)
}

// UnimplementedCounterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (*UnimplementedCounterServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCounterServiceServer) Close(ctx context.Context, req *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedCounterServiceServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedCounterServiceServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCounterServiceServer) Increment(ctx context.Context, req *IncrementRequest) (*IncrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (*UnimplementedCounterServiceServer) Decrement(ctx context.Context, req *DecrementRequest) (*DecrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrement not implemented")
}
func (*UnimplementedCounterServiceServer) CheckAndSet(ctx context.Context, req *CheckAndSetRequest) (*CheckAndSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndSet not implemented")
}

func RegisterCounterServiceServer(s *grpc.Server, srv CounterServiceServer) {
	s.RegisterService(&_CounterService_serviceDesc, srv)
}

func _CounterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Increment(ctx, req.(*IncrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_Decrement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Decrement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/Decrement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Decrement(ctx, req.(*DecrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_CheckAndSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).CheckAndSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.counter.CounterService/CheckAndSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).CheckAndSet(ctx, req.(*CheckAndSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CounterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.counter.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CounterService_Create_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _CounterService_Close_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CounterService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CounterService_Get_Handler,
		},
		{
			MethodName: "Increment",
			Handler:    _CounterService_Increment_Handler,
		},
		{
			MethodName: "Decrement",
			Handler:    _CounterService_Decrement_Handler,
		},
		{
			MethodName: "CheckAndSet",
			Handler:    _CounterService_CheckAndSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/counter/counter.proto",
}
