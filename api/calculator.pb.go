package api

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CalculationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Expression    string                 `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculationRequest) Reset() {
	*x = CalculationRequest{}
	mi := &file_api_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CalculationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CalculationRequest) ProtoMessage() {}
func (x *CalculationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CalculationRequest) Descriptor() ([]byte, []int) {
	return file_api_calculator_proto_rawDescGZIP(), []int{0}
}
func (x *CalculationRequest) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type CalculationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculationResponse) Reset() {
	*x = CalculationResponse{}
	mi := &file_api_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CalculationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CalculationResponse) ProtoMessage() {}
func (x *CalculationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CalculationResponse) Descriptor() ([]byte, []int) {
	return file_api_calculator_proto_rawDescGZIP(), []int{1}
}
func (x *CalculationResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}
func (x *CalculationResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_api_calculator_proto protoreflect.FileDescriptor

const file_api_calculator_proto_rawDesc = "" +
	"\n" +
	"\x14api/calculator.proto\x12\n" +
	"calculator\"4\n" +
	"\x12CalculationRequest\x12\x1e\n" +
	"\n" +
	"expression\x18\x01 \x01(\tR\n" +
	"expression\"C\n" +
	"\x13CalculationResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06result\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error2Z\n" +
	"\n" +
	"Calculator\x12L\n" +
	"\tCalculate\x12\x1e.calculator.CalculationRequest\x1a\x1f.calculator.CalculationResponseB\x10Z\x0ecalculator/apib\x06proto3"

var (
	file_api_calculator_proto_rawDescOnce sync.Once
	file_api_calculator_proto_rawDescData []byte
)

func file_api_calculator_proto_rawDescGZIP() []byte {
	file_api_calculator_proto_rawDescOnce.Do(func() {
		file_api_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_calculator_proto_rawDesc), len(file_api_calculator_proto_rawDesc)))
	})
	return file_api_calculator_proto_rawDescData
}

var file_api_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_calculator_proto_goTypes = []any{
	(*CalculationRequest)(nil),  // 0: calculator.CalculationRequest
	(*CalculationResponse)(nil), // 1: calculator.CalculationResponse
}
var file_api_calculator_proto_depIdxs = []int32{
	0,
	1,
	1,
	0,
	0,
	0,
	0,
}

func init() { file_api_calculator_proto_init() }
func file_api_calculator_proto_init() {
	if File_api_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_calculator_proto_rawDesc), len(file_api_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calculator_proto_goTypes,
		DependencyIndexes: file_api_calculator_proto_depIdxs,
		MessageInfos:      file_api_calculator_proto_msgTypes,
	}.Build()
	File_api_calculator_proto = out.File
	file_api_calculator_proto_goTypes = nil
	file_api_calculator_proto_depIdxs = nil
}
