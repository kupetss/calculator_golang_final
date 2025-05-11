package api

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9
const (
	Calculator_Calculate_FullMethodName = "/calculator.Calculator/Calculate"
)

type CalculatorClient interface {
	Calculate(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
}
type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}
func (c *calculatorClient) Calculate(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, Calculator_Calculate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CalculatorServer interface {
	Calculate(context.Context, *CalculationRequest) (*CalculationResponse, error)
	mustEmbedUnimplementedCalculatorServer()
}
type UnimplementedCalculatorServer struct{}

func (UnimplementedCalculatorServer) Calculate(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}
func (UnimplementedCalculatorServer) testEmbeddedByValue()                    {}

type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Calculator_ServiceDesc, srv)
}
func _Calculator_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Calculate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Calculate(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Calculator_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/calculator.proto",
}
