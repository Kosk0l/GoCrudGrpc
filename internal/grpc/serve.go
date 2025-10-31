package grpc

// gRPC слой не должен знать, как именно хранится информация — только делегировать работу вниз.

import (
	"GoCrudGrpc/internal/storage"
	"GoCrudGrpc/proto/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Структура, наследующая интерфейс со всеми методами - хендлерами
type ServerAPI struct {
    grpc.UnimplementedGRPCServer // наследует генерированные данные
	storage *storage.Postgres // зависимость с БД // указатель на память объекта Структуры
}

// Конструктор для ServerAPI // Передает указатель
func NewServerAPI(storage *storage.Postgres) *ServerAPI {
	return &ServerAPI{storage: storage}
}

//===================================================================================================================//
// Обработчики:

// Получить сообщение из БД
func (s *ServerAPI) GetMessage(ctx context.Context, req *grpc.GetMessageRequest) (*grpc.MessageResponse, error) {
	if req.GetTextID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad text id") // Неверные аргументы в запросе.
	}

	// Проходка в Базу Данных
	Response, err := s.storage.Get(ctx, req.TextID) 
	if err != nil {
		return nil, status.Error(codes.Internal, "Error DataBase Get") // Внутренняя ошибка сервера
	}

    return Response, nil
}

// Добавить сообщение в бд
func (s *ServerAPI) CreateMessage(ctx context.Context, req *grpc.PostMessageRequest) (*grpc.MessageResponse, error) {
	if req.GetTextID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad text id")
	}

	if req.GetUserID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad user id")
	}

	if req.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad text")
	}

	Response, err := s.storage.Create(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error DataBase Create") 
	}

	return Response, nil
}

// Обновить сообщение в бд
func (s *ServerAPI) UpdateMessage(ctx context.Context, req *grpc.UpdateMessageRequest) (*grpc.MessageResponse, error) {
	if req.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad text")
	}

	if req.GetTextID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad text id")
	}

	Response, err := s.storage.Update(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error DataBase Create") 
	}

	return Response, nil
}

// Удалить сообщение из бд
func (s *ServerAPI) DeleteMessage(ctx context.Context, req *grpc.DeleteMessageRequest) (*grpc.MessageResponse, error) {
	if req.GetTextID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad text id")
	}

	err := s.storage.Delete(ctx,req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error DataBase Create") 
	}

	return &grpc.MessageResponse{
        TextID: req.TextID,
        Text:   "Message Deleted",
        Status: "Message Deleted",
    }, nil
}