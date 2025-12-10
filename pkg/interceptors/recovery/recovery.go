package recovery

import (
	"context"
	"runtime/debug"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryServerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		defer func() {
			if r := recover(); r != nil {
				logger.Error(
					"panic recovered in gRPC handler",
					zap.String("method", info.FullMethod),
					zap.Any("panic", r),
					zap.String("stacktrace", string(debug.Stack())),
				)

				err = status.Errorf(codes.Internal, "panic: %v", r)
			}
		}()

		return handler(ctx, req)
	}
}
