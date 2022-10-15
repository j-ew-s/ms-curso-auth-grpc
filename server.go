package main

import (
	"fmt"
	"log"
	"net"

	"github.com/j-ew-s/ms-curso-auth-grpc/auth"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":5500")

	if err != nil {
		log.Fatalf("Erro ao escutar a porta 5500 : %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := auth.Server{}

	auth.RegisterUserServiceServer(grpcServer, &userService)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(fmt.Sprintf("Erro ao levantar o gRPC Server porta 5500 : %v", err))
	}

}
