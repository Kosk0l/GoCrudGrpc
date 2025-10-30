package storage

import (
	"GoCrudGrpc/proto/grpc"
	"context"
)

func (p *Postgres) Create(ctx context.Context, msg *grpc.PostMessageRequest) (*grpc.MessageResponse, error) {

	return &grpc.MessageResponse{

	}, nil
}

func(p *Postgres) Get(ctx context.Context, textID int64) (*grpc.MessageResponse, error) {

	return &grpc.MessageResponse{
		
	}, nil
}

func(p *Postgres) Delete(ctx context.Context, textId int64) error {

	
	return nil
}

func(p *Postgres) Update(ctx context.Context,) (*grpc.MessageResponse, error) {

	return &grpc.MessageResponse{
		
	}, nil
}

