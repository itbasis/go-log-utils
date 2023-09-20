package grpc

import (
	"context"
	"fmt"

	logUtils "github.com/itbasis/go-log-utils/v2"
	"github.com/juju/zaputil/zapctx"
	"google.golang.org/grpc"
)

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger := zapctx.Logger(ctx)

		logger.Debug(fmt.Sprintf(logUtils.ReceiveRequest, req))

		if logger == zapctx.Default {
			return handler(ctx, req)
		}

		newCtx := zapctx.WithLogger(ctx, logger)

		return handler(newCtx, req)
	}
}
