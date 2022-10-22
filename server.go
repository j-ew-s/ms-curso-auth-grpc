package main

import (
	"fmt"
	"log"
	"net"

	"github.com/j-ew-s/ms-curso-auth-grpc/auth-services"
	"github.com/j-ew-s/ms-curso-auth-grpc/database"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":5500")

	if err != nil {
		log.Fatalf("Erro ao escutar a porta 5500 : %v", err)
	}

	sqlCommand := SetDataBase()

	grpcServer := grpc.NewServer()

	userService := auth.AuthService{
		DBConn: sqlCommand,
	}

	auth.RegisterUserServiceServer(grpcServer, &userService)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(fmt.Sprintf("Erro ao levantar o gRPC Server porta 5500 : %v", err))
	}

}

func SetDataBase() *database.SQLCommand {

	conn := database.SetDrivers("mysql", "user_management")

	sqlCommand := &database.SQLCommand{
		SqlConn: conn,
	}

	err := sqlCommand.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return sqlCommand

}
