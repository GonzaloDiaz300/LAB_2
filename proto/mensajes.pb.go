// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: proto/mensajes.proto

package proto

import (
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

type ContiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *ContiReq) Reset() {
	*x = ContiReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContiReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContiReq) ProtoMessage() {}

func (x *ContiReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContiReq.ProtoReflect.Descriptor instead.
func (*ContiReq) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{0}
}

func (x *ContiReq) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *ContiReq) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *ContiReq) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Confirmacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respuesta int32 `protobuf:"varint,1,opt,name=respuesta,proto3" json:"respuesta,omitempty"`
}

func (x *Confirmacion) Reset() {
	*x = Confirmacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmacion) ProtoMessage() {}

func (x *Confirmacion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmacion.ProtoReflect.Descriptor instead.
func (*Confirmacion) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{1}
}

func (x *Confirmacion) GetRespuesta() int32 {
	if x != nil {
		return x.Respuesta
	}
	return 0
}

type OMSReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre   string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,3,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *OMSReq) Reset() {
	*x = OMSReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OMSReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OMSReq) ProtoMessage() {}

func (x *OMSReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OMSReq.ProtoReflect.Descriptor instead.
func (*OMSReq) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{2}
}

func (x *OMSReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OMSReq) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *OMSReq) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type DTNResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *DTNResp) Reset() {
	*x = DTNResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DTNResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DTNResp) ProtoMessage() {}

func (x *DTNResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DTNResp.ProtoReflect.Descriptor instead.
func (*DTNResp) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{3}
}

func (x *DTNResp) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *DTNResp) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type ONUReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *ONUReq) Reset() {
	*x = ONUReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ONUReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ONUReq) ProtoMessage() {}

func (x *ONUReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ONUReq.ProtoReflect.Descriptor instead.
func (*ONUReq) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{4}
}

func (x *ONUReq) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type ONUResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persona []string `protobuf:"bytes,1,rep,name=persona,proto3" json:"persona,omitempty"`
}

func (x *ONUResp) Reset() {
	*x = ONUResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ONUResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ONUResp) ProtoMessage() {}

func (x *ONUResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ONUResp.ProtoReflect.Descriptor instead.
func (*ONUResp) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{5}
}

func (x *ONUResp) GetPersona() []string {
	if x != nil {
		return x.Persona
	}
	return nil
}

type OMSONUReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OMSONUReq) Reset() {
	*x = OMSONUReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensajes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OMSONUReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OMSONUReq) ProtoMessage() {}

func (x *OMSONUReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensajes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OMSONUReq.ProtoReflect.Descriptor instead.
func (*OMSONUReq) Descriptor() ([]byte, []int) {
	return file_proto_mensajes_proto_rawDescGZIP(), []int{6}
}

func (x *OMSONUReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_mensajes_proto protoreflect.FileDescriptor

var file_proto_mensajes_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x56, 0x0a, 0x08,
	0x43, 0x6f, 0x6e, 0x74, 0x69, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x22, 0x2c, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x22, 0x4c, 0x0a, 0x06, 0x4f, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f,
	0x22, 0x3d, 0x0a, 0x07, 0x44, 0x54, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22,
	0x20, 0x0a, 0x06, 0x4f, 0x4e, 0x55, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x22, 0x23, 0x0a, 0x07, 0x4f, 0x4e, 0x55, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x22, 0x1b, 0x0a, 0x09, 0x4f, 0x4d, 0x53, 0x4f, 0x4e, 0x55,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xbe, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x61, 0x6d,
	0x62, 0x69, 0x6f, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x72, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72,
	0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x07, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x73, 0x12, 0x0c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x4e, 0x55, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4f, 0x4e, 0x55, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x42, 0x75,
	0x73, 0x63, 0x61, 0x72, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x4d, 0x53, 0x4f,
	0x4e, 0x55, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x54, 0x4e,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6e, 0x7a, 0x61, 0x6c, 0x6f, 0x44, 0x69, 0x61, 0x7a, 0x33, 0x30,
	0x30, 0x2f, 0x4c, 0x41, 0x42, 0x5f, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mensajes_proto_rawDescOnce sync.Once
	file_proto_mensajes_proto_rawDescData = file_proto_mensajes_proto_rawDesc
)

func file_proto_mensajes_proto_rawDescGZIP() []byte {
	file_proto_mensajes_proto_rawDescOnce.Do(func() {
		file_proto_mensajes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mensajes_proto_rawDescData)
	})
	return file_proto_mensajes_proto_rawDescData
}

var file_proto_mensajes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_mensajes_proto_goTypes = []interface{}{
	(*ContiReq)(nil),     // 0: grpc.ContiReq
	(*Confirmacion)(nil), // 1: grpc.Confirmacion
	(*OMSReq)(nil),       // 2: grpc.OMSReq
	(*DTNResp)(nil),      // 3: grpc.DTNResp
	(*ONUReq)(nil),       // 4: grpc.ONUReq
	(*ONUResp)(nil),      // 5: grpc.ONUResp
	(*OMSONUReq)(nil),    // 6: grpc.OMSONUReq
}
var file_proto_mensajes_proto_depIdxs = []int32{
	0, // 0: grpc.Intercambios.Notificar:input_type -> grpc.ContiReq
	2, // 1: grpc.Intercambios.Guardar:input_type -> grpc.OMSReq
	4, // 2: grpc.Intercambios.Nombres:input_type -> grpc.ONUReq
	6, // 3: grpc.Intercambios.Buscar:input_type -> grpc.OMSONUReq
	1, // 4: grpc.Intercambios.Notificar:output_type -> grpc.Confirmacion
	1, // 5: grpc.Intercambios.Guardar:output_type -> grpc.Confirmacion
	5, // 6: grpc.Intercambios.Nombres:output_type -> grpc.ONUResp
	3, // 7: grpc.Intercambios.Buscar:output_type -> grpc.DTNResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mensajes_proto_init() }
func file_proto_mensajes_proto_init() {
	if File_proto_mensajes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mensajes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContiReq); i {
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
		file_proto_mensajes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmacion); i {
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
		file_proto_mensajes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OMSReq); i {
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
		file_proto_mensajes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DTNResp); i {
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
		file_proto_mensajes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ONUReq); i {
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
		file_proto_mensajes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ONUResp); i {
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
		file_proto_mensajes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OMSONUReq); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mensajes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mensajes_proto_goTypes,
		DependencyIndexes: file_proto_mensajes_proto_depIdxs,
		MessageInfos:      file_proto_mensajes_proto_msgTypes,
	}.Build()
	File_proto_mensajes_proto = out.File
	file_proto_mensajes_proto_rawDesc = nil
	file_proto_mensajes_proto_goTypes = nil
	file_proto_mensajes_proto_depIdxs = nil
}
