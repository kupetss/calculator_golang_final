package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/kupetss/calculator_golang_final/internal/proto"
	"google.golang.org/grpc"
)

type Server struct {
	db *sql.DB
	proto.UnimplementedCalculatorServiceServer
}

func StartGRPCServer(db *sql.DB) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{db: db})

	log.Printf("gRPC server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Calculate(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	result, err := evaluateExpression(req.Expression)
	if err != nil {
		return &proto.CalculationResponse{
			Error: err.Error(),
		}, nil
	}

	// Сохраняем в БД
	err = saveToDB(req.UserId, req.Expression, result, s.db)
	if err != nil {
		return nil, err
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

// evaluateExpression - вычисление математического выражения
func evaluateExpression(expr string) (string, error) {
	// Простейшая реализация (замените на свою логику)
	switch expr {
	case "2+2":
		return "4", nil
	case "2*2":
		return "4", nil
	default:
		return "", fmt.Errorf("unsupported expression")
	}
}

// saveToDB - сохранение результата в базу данных
func saveToDB(userID int32, expr string, result string, db *sql.DB) error {
	// Простая реализация для SQLite
	_, err := db.Exec(
		"INSERT INTO calculations (user_id, expression, result) VALUES (?, ?, ?)",
		userID, expr, result,
	)
	return err
}

// getHistoryFromDB - получение истории вычислений
func getHistoryFromDB(userID int32, db *sql.DB) ([]string, error) {
	rows, err := db.Query(
		"SELECT expression FROM calculations WHERE user_id = ?",
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
