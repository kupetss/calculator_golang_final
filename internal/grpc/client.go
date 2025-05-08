package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/kupetss/calculator_golang_final/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallCalculate(expr string, userID int) (string, error) {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	c := proto.NewCalculatorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Calculate(ctx, &proto.CalculationRequest{
		Expression: expr,
		UserId:     int32(userID),
	})
	if err != nil {
		return "", err
	}

	if res.Error != "" {
		return "", errors.New(res.Error)
	}

	return res.Result, nil
}
