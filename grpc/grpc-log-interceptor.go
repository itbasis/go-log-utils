package grpc

import (
	"context"
	"fmt"

	logUtils "github.com/itbasis/go-log-utils/v2"
	"github.com/juju/zaputil/zapctx"
	"google.golang.org/grpc"
)

func GrpcLogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger := zapctx.Default
		newCtx := zapctx.WithLogger(ctx, logger)

		logger.Debug(fmt.Sprintf(logUtils.ReceiveRequest, req))

		return handler(newCtx, req)
	}
}
