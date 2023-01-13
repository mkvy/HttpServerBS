package server

import (
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/config"
	pb "github.com/mkvy/HttpServerBS/custshopsvc/protofiles"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServer struct {
	cfg config.Config
	srv *grpc.Server
}

func NewGrpcServer(cfg config.Config, handler grpcController) *grpcServer {
	s := grpc.NewServer()
	pb.RegisterCustShopServiceServer(s, &handler)
	return &grpcServer{
		cfg: cfg,
		srv: s,
	}
}

func (s *grpcServer) Start() {
	//todo port grpc
	lis, err := net.Listen("tcp", ":9111")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	err = s.srv.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *grpcServer) Stop() {
	s.srv.Stop()
}
