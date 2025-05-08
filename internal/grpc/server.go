package grpc

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/kupetss/calculator_golang_final/proto"
	"google.golang.org/grpc"
)

type Server struct {
	db *sql.DB
	proto.UnimplementedCalculatorServiceServer
}

func InitCalculatorService(db *sql.DB) {
	// Инициализация может включать подготовку БД
}

func StartGRPCServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *Server) Calculate(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	result, err := evaluateExpression(req.Expression)
	if err != nil {
		return &proto.CalculationResponse{
			Error: err.Error(),
		}, nil
	}

	// Сохраняем в историю
	if err := saveToHistory(req.UserId, req.Expression, result, s.db); err != nil {
		log.Printf("Failed to save history: %v", err)
	}

	return &proto.CalculationResponse{
		Result: result,
	}, nil
}

func (s *Server) GetHistory(ctx context.Context, req *proto.HistoryRequest) (*proto.HistoryResponse, error) {
	history, err := getHistoryFromDB(req.UserId, s.db)
	if err != nil {
		return nil, err
	}

	return &proto.HistoryResponse{
		Expressions: history,
	}, nil
}

// Вспомогательные функции
func evaluateExpression(expr string) (string, error) {
	// Реализация вычислений
	return "", nil
}

func saveToHistory(userID int32, expr, result string, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO calculations(user_id, expression, result) VALUES(?, ?, ?)",
		userID, expr, result,
	)
	return err
}

func getHistoryFromDB(userID int32, db *sql.DB) ([]string, error) {
	rows, err := db.Query(
		"SELECT expression FROM calculations WHERE user_id = ? ORDER BY created_at DESC",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []string
	for rows.Next() {
		var expr string
		if err := rows.Scan(&expr); err != nil {
			return nil, err
		}
		history = append(history, expr)
	}

	return history, nil
}
