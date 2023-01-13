package client

import (
	"context"
	"encoding/json"
	pb "github.com/mkvy/HttpServerBS/api-gateway/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type CustShopService interface {
	Create(json.RawMessage, string) (string, string)
	Update(json.RawMessage, string, string) string
	Delete(string, string) string
	GetById(string, string, string) (json.RawMessage, string)
	GetByParameters(string, string, string) (json.RawMessage, string)
}

type grpcClient struct {
	client pb.CustShopServiceClient
}

const address = "localhost:9111"

// todo config
func NewGrpcClient() *grpcClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	client := pb.NewCustShopServiceClient(conn)
	return &grpcClient{client: client}
}

func (gc *grpcClient) Create(msg json.RawMessage, modelType string) (string, string) {
	log.Println("GrpcClient Create request")
	req := &pb.CreateRequest{
		Jsondata:  msg,
		ModelType: modelType,
	}
	resp, err := gc.client.CreateService(context.Background(), req)
	if err != nil {
		log.Println("Error occured while making grpc call ", err)
		return "", err.Error()
	}
	log.Println("GrpcClient Created with id " + resp.Id)
	return resp.Id, resp.Error
}

func (gc *grpcClient) Update(msg json.RawMessage, id string, modelType string) string {
	log.Println("GrpcClient Update request")
	req := &pb.UpdateRequest{
		Jsondata:  msg,
		Id:        id,
		ModelType: modelType,
	}
	resp, err := gc.client.UpdateService(context.Background(), req)
	if err != nil {
		log.Println("Error occured while making grpc call ", err)
		return err.Error()
	}
	return resp.Error
}

func (gc *grpcClient) Delete(id string, modelType string) string {
	log.Println("GrpcClient Create request")
	req := &pb.DeleteRequest{
		Id:        id,
		ModelType: modelType,
	}
	resp, err := gc.client.DeleteService(context.Background(), req)
	if err != nil {
		log.Println("Error occured while making grpc call ", err)
		return err.Error()
	}
	return resp.Error
}

func (gc *grpcClient) GetById(id string, field string, modelType string) (json.RawMessage, string) {
	log.Println("GrpcClient GetByID request")
	req := &pb.GetByIDRequest{
		Id:        id,
		Field:     field,
		ModelType: modelType,
	}
	resp, err := gc.client.GetByIDService(context.Background(), req)
	if err != nil {
		log.Println("Error occured while making grpc call ", err)
		return nil, err.Error()
	}
	return resp.Jsondata, resp.Error
}

func (gc *grpcClient) GetByParameters(searchParam string, field string, modelType string) (json.RawMessage, string) {
	log.Println("GrpcClient GetByParameters request")
	req := &pb.GetByParametersRequest{
		SearchParam: searchParam,
		Field:       field,
		ModelType:   modelType,
	}
	resp, err := gc.client.GetByParameters(context.Background(), req)
	if err != nil {
		log.Println("Error occured while creating grpc call ", err)
		return nil, err.Error()
	}
	return resp.Jsondata, resp.Error
}
