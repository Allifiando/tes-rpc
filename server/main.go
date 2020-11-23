package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"santara.co.id/set/market-common-service-grpc/helpers"
	"santara.co.id/set/market-common-service-grpc/pkg/proto/common"
)

type server struct {
	common.UnimplementedCommonServiceServer
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error getting env, not comming through %v", err)
	}
	log.Println("Server running ...")
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	common.RegisterCommonServiceServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))
}

func (s *server) GetBearer(ctx context.Context, request *common.RequestGetToken) (*common.ApiResponse, error) {
	log.Println(fmt.Sprintf("Request: GetUser"))
	data, err := helpers.Auth(request.GetToken())
	if err != nil {
		return &common.ApiResponse{Data: nil, Message: "Unauthorized"}, nil
	}
	return &common.ApiResponse{Data: data, Message: "Authorized"}, nil
}

func (s *server) GetPin(ctx context.Context, request *common.RequestGetPin) (*common.ApiResponse, error) {
	log.Println(fmt.Sprintf("Request: GetUser"))
	data, err := helpers.Auth(request.GetToken())
	if err != nil {
		return &common.ApiResponse{Data: nil, Message: "Unauthorized"}, nil
	}

	pin := request.GetPin()

	err = bcrypt.CompareHashAndPassword([]byte(data.Pin), []byte(pin))

	if err != nil {
		return &common.ApiResponse{Data: nil, Message: "PIN Mismatch"}, nil
	}

	return &common.ApiResponse{Data: data, Message: "PIN Valid"}, nil

}
