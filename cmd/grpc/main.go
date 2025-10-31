package main

import (
	grpcServer "GoCrudGrpc/internal/grpc"
	"GoCrudGrpc/internal/storage"
	pb "GoCrudGrpc/proto/grpc"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"google.golang.org/grpc"
)


func main() {

	// Базовый контекст
	// Только на этапе инициализации — когда создаём подключение к базе (pgxpool.NewWithConfig).
	// нужен свой собственный контекст, чтобы:
	// 	1. Если БД не отвечает 5 секунд, не висеть вечно;
	// 	2. Корректно завершить подключение при запуске;
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ключ подключения
	dsn := "host=localhost user=postgresCrud password=qwerty dbname=postgresCrud port=5433 sslmode=disable"
	// host=localhost user=postgres password=password dbname=postgres port=port sslmode=disable

	// Создание подключение
	db, err := storage.NewPostgres(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	log.Println("Connected to Postgres")

	//======================================================================//

	// Иннициализация слушателя порта
	lis, err := net.Listen("tcp", ":44044") 
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Иннициализация сервера
	s := grpc.NewServer()

	// Передача подключение к БД в слой gRPC
	// реализация АПИ // регистрация сервера
	serverAPI := grpcServer.NewServerAPI(db)
	pb.RegisterGRPCServer(s, serverAPI)

	// Запуск Сервера в отдельной горутине
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("Server is working")
	// Ожидание сигналов завершения (Ctrl+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
