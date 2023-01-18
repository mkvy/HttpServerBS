package server

import (
	"context"
	"github.com/mkvy/HttpServerBS/custshopsvc/service"
	pb "github.com/mkvy/HttpServerBS/shared/protofiles"
	"log"
)

type grpcController struct {
	s service.Service
}

func NewGrpcController(svc service.Service) *grpcController {
	return &grpcController{s: svc}
}

func (gc *grpcController) CreateService(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println("Grpc Controller create request")
	ret, errMsg := gc.s.Create(request.Jsondata, request.ModelType)
	res := pb.CreateResponse{
		Id:    ret,
		Error: errMsg,
	}
	return &res, nil
}

func (gc *grpcController) UpdateService(ctx context.Context, request *pb.UpdateRequest) (*pb.ErrorResponse, error) {
	log.Println("Grpc Controller update request")
	errMsg := gc.s.Update(request.Jsondata, request.Id, request.ModelType)
	res := pb.ErrorResponse{
		Error: errMsg,
	}
	return &res, nil
}

func (gc *grpcController) DeleteService(ctx context.Context, request *pb.DeleteRequest) (*pb.ErrorResponse, error) {
	log.Println("Grpc Controller update request")
	errMsg := gc.s.Delete(request.Id, request.ModelType)
	res := pb.ErrorResponse{
		Error: errMsg,
	}
	return &res, nil
}

func (gc *grpcController) GetByIDService(ctx context.Context, request *pb.GetByIDRequest) (*pb.ModelResponse, error) {
	log.Println("Grpc Controller GetById request")
	ret, errMsg := gc.s.GetById(request.Id, request.Field, request.ModelType)
	res := pb.ModelResponse{
		Jsondata: ret,
		Error:    errMsg,
	}
	return &res, nil
}

func (gc *grpcController) GetByParameters(ctx context.Context, request *pb.GetByParametersRequest) (*pb.ModelResponse, error) {
	log.Println("Grpc Controller GetByParameters request")
	ret, errMsg := gc.s.GetByParameters(request.SearchParam, request.Field, request.ModelType)
	res := pb.ModelResponse{
		Jsondata: ret,
		Error:    errMsg,
	}
	return &res, nil
}
