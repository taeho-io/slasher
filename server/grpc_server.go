package server

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/taeho-io/idl/gen/go/slasher"
	"github.com/taeho-io/slasher/server/handler"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type SlasherServer struct {
	slasher.SlasherServer

	cfg Config
}

func New(cfg Config) (*SlasherServer, error) {
	return &SlasherServer{
		cfg: cfg,
	}, nil
}

func Mock() *SlasherServer {
	s, _ := New(MockConfig())
	return s
}

func (s *SlasherServer) Config() Config {
	return s.cfg
}

func (s *SlasherServer) RegisterServer(srv *grpc.Server) {
	slasher.RegisterSlasherServer(srv, s)
}

func (s *SlasherServer) Slash(ctx context.Context, req *slasher.SlashRequest) (*slasher.SlashResponse, error) {
	return handler.Slash()(ctx, req)
}

func NewGRPCServer(cfg Config) (*grpc.Server, error) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	slasherServer, err := New(cfg)
	if err != nil {
		return nil, err
	}
	slasher.RegisterSlasherServer(grpcServer, slasherServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}

func ServeGRPC(addr string, cfg Config) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer, err := NewGRPCServer(cfg)
	if err != nil {
		return err
	}

	return grpcServer.Serve(lis)
}
