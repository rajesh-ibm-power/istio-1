// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/certs.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Certificates struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{0}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	CaCert               []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

type CertificateDetails struct {
	Path                 string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SerialNumber         string                  `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	SubjectAltNames      []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	DaysUntilExpiration  uint64                  `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	ValidFrom            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ExpirationTime       *timestamp.Timestamp    `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{2}
}

func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateDetails.Unmarshal(m, b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return xxx_messageInfo_CertificateDetails.Size(m)
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *timestamp.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	//	*SubjectAlternateName_IpAddress
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{3}
}

func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectAlternateName.Unmarshal(m, b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return xxx_messageInfo_SubjectAlternateName.Size(m)
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}

type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

type SubjectAlternateName_IpAddress struct {
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_IpAddress) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *SubjectAlternateName) GetIpAddress() string {
	if x, ok := m.GetName().(*SubjectAlternateName_IpAddress); ok {
		return x.IpAddress
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
		(*SubjectAlternateName_IpAddress)(nil),
	}
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v3.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v3.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v3.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v3.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v3/certs.proto", fileDescriptor_51228a64c985b2dc) }

var fileDescriptor_51228a64c985b2dc = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x8e, 0x12, 0x4d,
	0x14, 0xfd, 0x1a, 0xf8, 0x30, 0xdc, 0xc1, 0x19, 0x2d, 0x35, 0xe9, 0xe0, 0x0f, 0x4c, 0x6b, 0xc6,
	0x8e, 0x89, 0xdd, 0x09, 0xb8, 0x11, 0x17, 0x06, 0x50, 0xe3, 0x6a, 0x42, 0x5a, 0x5d, 0x77, 0x2e,
	0xdd, 0x05, 0x94, 0xe9, 0xae, 0xea, 0x54, 0x55, 0x13, 0xd8, 0x19, 0x57, 0x3e, 0x82, 0xf1, 0x3d,
	0xdc, 0xb8, 0x37, 0x71, 0xeb, 0x1b, 0x99, 0xea, 0x66, 0x66, 0x60, 0x86, 0x88, 0xbb, 0xaa, 0x73,
	0xee, 0xb9, 0x3f, 0xe7, 0x5e, 0x68, 0x51, 0xbe, 0x10, 0x2b, 0x1f, 0xe3, 0x94, 0x71, 0x7f, 0xd1,
	0xf3, 0x23, 0x2a, 0xb5, 0xf2, 0x32, 0x29, 0xb4, 0x20, 0x87, 0x05, 0xe7, 0x15, 0x9c, 0xb7, 0xe8,
	0xb5, 0xda, 0x33, 0x21, 0x66, 0x09, 0xf5, 0x0b, 0x76, 0x92, 0x4f, 0x7d, 0xcd, 0x52, 0xaa, 0x34,
	0xa6, 0x59, 0x29, 0x68, 0xdd, 0xcf, 0xe3, 0x0c, 0x7d, 0xe4, 0x5c, 0x68, 0xd4, 0x4c, 0x70, 0xe5,
	0x2b, 0x8d, 0x3a, 0x5f, 0xe7, 0x6b, 0x1d, 0x5f, 0xa1, 0x17, 0x54, 0x2a, 0x26, 0x38, 0xe3, 0xb3,
	0x32, 0xc4, 0x59, 0x42, 0x73, 0x44, 0xa5, 0x66, 0x53, 0x16, 0xa1, 0xa6, 0x8a, 0xbc, 0x84, 0x66,
	0xb4, 0xf1, 0xb7, 0xad, 0x4e, 0xd5, 0x3d, 0xe8, 0xde, 0xf5, 0xb6, 0x3b, 0xf3, 0x36, 0x34, 0xc1,
	0x96, 0xa0, 0xff, 0xf8, 0xdb, 0xcf, 0x2f, 0x0f, 0x1c, 0xe8, 0x6c, 0x09, 0xba, 0x98, 0x64, 0x73,
	0xdc, 0x54, 0x29, 0xe7, 0xbb, 0x05, 0x07, 0x1b, 0x00, 0x79, 0x01, 0xd7, 0x22, 0x0c, 0x4d, 0xae,
	0x75, 0x51, 0xe7, 0x2f, 0x45, 0x5f, 0x51, 0x8d, 0x2c, 0x51, 0x41, 0x3d, 0x42, 0x83, 0x92, 0x01,
	0x80, 0x51, 0x86, 0xd1, 0x1c, 0x19, 0xb7, 0x2b, 0xff, 0xac, 0x6f, 0x18, 0xd5, 0xc8, 0x88, 0xfa,
	0x27, 0xa6, 0xf1, 0x63, 0x68, 0xef, 0x69, 0xdc, 0xf9, 0x5c, 0x05, 0x72, 0x35, 0x13, 0x21, 0x50,
	0xcb, 0x50, 0xcf, 0x6d, 0xab, 0x63, 0xb9, 0x8d, 0xa0, 0x78, 0x93, 0x87, 0x70, 0x5d, 0x51, 0xc9,
	0x30, 0x09, 0x79, 0x9e, 0x4e, 0xa8, 0xb4, 0x2b, 0x05, 0xd9, 0x2c, 0xc1, 0xd3, 0x02, 0x23, 0x63,
	0xb8, 0xa9, 0xf2, 0xc9, 0x47, 0x1a, 0xe9, 0x10, 0x13, 0x1d, 0x72, 0x4c, 0xa9, 0xb2, 0xab, 0xc5,
	0x04, 0x8f, 0x2e, 0x4f, 0xf0, 0xae, 0x0c, 0x1c, 0x24, 0x9a, 0x4a, 0x8e, 0x9a, 0x9e, 0x62, 0x4a,
	0x83, 0x23, 0x75, 0x8e, 0x9a, 0xbf, 0x22, 0x5d, 0xb8, 0x13, 0xe3, 0x4a, 0x85, 0x39, 0xd7, 0x2c,
	0x09, 0xe9, 0x32, 0x63, 0xb2, 0x58, 0xbf, 0x5d, 0xeb, 0x58, 0x6e, 0x2d, 0xb8, 0x65, 0xc8, 0x0f,
	0x86, 0x7b, 0x7d, 0x4e, 0x91, 0xe7, 0x00, 0x0b, 0x4c, 0x58, 0x1c, 0x4e, 0xa5, 0x48, 0xed, 0xff,
	0x3b, 0x96, 0x7b, 0xd0, 0x6d, 0x79, 0xe5, 0xfd, 0x79, 0x67, 0xf7, 0xe7, 0xbd, 0x3f, 0xbb, 0xbf,
	0xa0, 0x51, 0x44, 0xbf, 0x91, 0x22, 0x25, 0x23, 0x38, 0xba, 0xa8, 0x11, 0x9a, 0x13, 0xb5, 0xeb,
	0x7b, 0xf5, 0x87, 0x17, 0x12, 0x03, 0xf6, 0x9f, 0x1a, 0xf7, 0x5d, 0x38, 0xd9, 0xe3, 0xfe, 0xda,
	0x6d, 0xe7, 0xab, 0x05, 0xb7, 0x77, 0x99, 0x41, 0x08, 0x54, 0x63, 0xae, 0xca, 0x2d, 0xbc, 0xfd,
	0x2f, 0x30, 0x1f, 0x83, 0xe5, 0x92, 0x95, 0xe6, 0x1b, 0x2c, 0x97, 0x8c, 0xb4, 0x01, 0x58, 0x16,
	0x62, 0x1c, 0x4b, 0xaa, 0x8c, 0xdd, 0x25, 0xd5, 0x60, 0xd9, 0xa0, 0x84, 0xfa, 0xbe, 0x69, 0xe8,
	0x09, 0xb8, 0xbb, 0x1a, 0xda, 0x55, 0x79, 0x58, 0x87, 0x9a, 0xd9, 0xdd, 0xf0, 0xd9, 0x8f, 0x4f,
	0xbf, 0x7e, 0xd7, 0x2b, 0x37, 0x2a, 0x70, 0x8f, 0x89, 0x72, 0x81, 0x99, 0x14, 0xcb, 0xd5, 0xa5,
	0x5d, 0x0e, 0xc1, 0x8c, 0xa5, 0xc6, 0xc6, 0x9a, 0xb1, 0x35, 0xa9, 0x17, 0x1e, 0xf5, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0x84, 0x4f, 0xb0, 0x1f, 0x04, 0x00, 0x00,
}