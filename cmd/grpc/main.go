package main

import (
	"log"
	"net"
	pb "GoCrudGrpc/proto/grpc"
	"google.golang.org/grpc"
	grpcServer "GoCrudGrpc/internal/grpc"
)


func main() {
	lis, err := net.Listen("tcp", ":44044") // запуск слушателя порта
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // Иннициализация сервера

	log.Println("Server is working")
	pb.RegisterGRPCServer(s, &grpcServer.ServerAPI{}) // реализация АПИ // регистрация сервера

	if err := s.Serve(lis); err != nil { 
		log.Fatalf("failed to serve: %v", err)
	}
}
