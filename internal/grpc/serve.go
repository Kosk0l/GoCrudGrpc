package grpc

import (
	"GoCrudGrpc/proto/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Будущая структура для обработки
type Response interface {}

// Структура, наследующая интерфейс со всеми методами - хендлерами
type ServerAPI struct {
    grpc.UnimplementedGRPCServer // наследует генерированные данные
}

//===================================================================================================================//
// Обработчики:

func (s *ServerAPI) GetMessage(ctx context.Context, req *grpc.GetMessageRequest) (*grpc.MessageResponse, error) {
	if req.GetTextID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad text id")
	}

	//..

    return &grpc.MessageResponse{
        TextID: req.TextID,
        Text:   "",
        Status: "",
    }, nil
}

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

	//..

	return &grpc.MessageResponse{
        TextID: req.TextID,
        Text:   "",
        Status: "",
    }, nil
}

func (s *ServerAPI) DeleteMessage(ctx context.Context, req *grpc.DeleteMessageRequest) (*grpc.MessageResponse, error) {
	if req.GetTextID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad text id")
	}

	//..

	return &grpc.MessageResponse{
        TextID: req.TextID,
        Text:   "",
        Status: "",
    }, nil
}