package storage

import (
	"GoCrudGrpc/proto/grpc"
	"context"
	"fmt"
)

// Вставляем сообщение в БД
func (p *Postgres) Create(ctx context.Context, req *grpc.PostMessageRequest) (*grpc.MessageResponse, error) {
	// string для sql запроса
	sql := `
		INSERT INTO messages (text_id, user_id, text)
		VALUES ($1, $2, $3)
	`

	// Выполняется SQL-запрос с передачей аргументов:
	// Exec: Берёт одно соединение из пула; Выполняет команду; Автоматически возвращает соединение в пул
	cmd, err := p.pool.Exec(ctx, sql, req.TextID, req.UserID, req.Text)
	if err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}

	// Проверка: INSERT должен вставить хотя бы 1 строку
	if cmd.RowsAffected() == 0 {
        return nil, fmt.Errorf("create: no rows affected")
    }

	return &grpc.MessageResponse{
        TextID: req.TextID,
        Text:   req.Text,
        Status: "created",
    }, nil
}

// Получаем сообщение из БД
func(p *Postgres) Get(ctx context.Context, textID int64) (*grpc.MessageResponse, error) {
	sql := `SELECT text_id, user_id, text FROM messages WHERE text_id = $1`

	var (
        id  	int64
        userID  int64
        text    string
    )

	row := p.pool.QueryRow(ctx, sql, id)
	err := row.Scan(&id, &userID, &text)
	if err != nil {
		return nil, fmt.Errorf("failed row: %v", err)
	}

	return &grpc.MessageResponse{
        TextID: id,
        Text: text,
        Status: "Getted",
    }, nil
}

// Удаляем сообщение из БД
func(p *Postgres) Delete(ctx context.Context, req grpc.DeleteMessageRequest) error {
	sql := `DELETE FROM messages WHERE text_id = $1`

	cmd, err := p.pool.Exec(ctx, sql, req.TextID)
	if err != nil {
		return fmt.Errorf("failed delete message: %v", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("failed delete row: %v", err)
	}

	return nil
}

// Обновляем сообщение в БД
func(p *Postgres) Update(ctx context.Context, req *grpc.UpdateMessageRequest) (*grpc.MessageResponse, error) {
	sql := `
		UPDATE messages
        SET text = $2
        WHERE text_id = $1
        RETURNING text_id, user_id, text
	`

	var (
        id     int64
        userID int64
        text   string
    )

	row := p.pool.QueryRow(ctx, sql, req.TextID, req.Text)
	err := row.Scan(&id, &userID, &text)
	if err != nil {
		return nil, fmt.Errorf("failed row: %v", err)
	}

	return &grpc.MessageResponse{
        TextID: req.TextID,
        Text:   req.Text,
        Status: "updated",
    }, nil
}
