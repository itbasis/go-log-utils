package grpc

import (
	"context"

	logUtils "github.com/itbasis/go-log-utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func GrpcLogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger := log.Logger
		newCtx := logger.WithContext(ctx)

		logger.Trace().Msgf(logUtils.ReceiveRequest, req)

		return handler(newCtx, req)
	}
}
