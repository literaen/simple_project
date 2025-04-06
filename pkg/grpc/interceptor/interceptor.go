package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// проверяет, что запрос не пустой
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be empty")
	}

	start := time.Now()

	// Выполнение основного хендлера
	res, err := handler(ctx, req)

	// Логирование данных о запросе
	duration := time.Since(start)
	log.Printf("Method: %s | Duration: %v | Error: %v", info.FullMethod, duration, err)

	return res, err
}
