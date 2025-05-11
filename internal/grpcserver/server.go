package grpcserver

import (
	"calculator/api"
	"calculator/internal/calculator"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedCalculatorServer
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterCalculatorServer(s, &Server{})

	log.Printf("GRPC server started on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func (s *Server) Calculate(ctx context.Context, req *api.CalculationRequest) (*api.CalculationResponse, error) {
	result, err := calculator.Evaluate(req.Expression)
	if err != nil {
		return &api.CalculationResponse{Error: err.Error()}, nil
	}
	return &api.CalculationResponse{Result: result}, nil
}
