package grpc

import (
	"context"
	"net"

	"github.com/tommyatchiron/togolist/internal/pkg/config"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var Module = fx.Options(
	fx.Provide(newGrpcServer),
	fx.Invoke(runGrpcServer),
)

func newGrpcServer() *grpc.Server {
	ser := grpc.NewServer()
	reflection.Register(ser)
	return ser
}

func runGrpcServer(lifecycle fx.Lifecycle, grpcServer *grpc.Server, config *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", config.Grpc.ListenAddr)
			if err != nil {
				return err
			}
			go grpcServer.Serve(lis)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.Stop()
			return nil
		},
	})
}
