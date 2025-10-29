package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// 
type Postgres struct {
	pool *pgxpool.Pool // пул соединений
	// Управляет несколькими TCP-соединениями к базе, 
	// чтобы сервис не создавал новое соединение на каждый запрос
}

// Создание подключения/соединения к БД
func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {

	return &Postgres{}, nil
}

// Закрытие соединения
func (p *Postgres) Close() {
	p.pool.Close()
}