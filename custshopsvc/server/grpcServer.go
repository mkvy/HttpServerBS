package server

import (
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/config"
	pb "github.com/mkvy/HttpServerBS/shared/protofiles"
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
	lis, err := net.Listen("tcp", ":"+s.cfg.GrpcServer.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	err = s.srv.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("Grpc server is running")
}

func (s *grpcServer) Stop() {
	log.Println("Grpc server stops")
	s.srv.Stop()
}
