package storage

import (
	"context"
	"fmt"
	"time"

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
	config, err := pgxpool.ParseConfig(dsn) 
	if err != nil {
		return nil, fmt.Errorf("failed parse Config Postgres pgxpool: %v", err)
	}

	// настройка таймаута
	config.MaxConns = 10 // Максимальное количество открытых соединений
	config.MaxConnLifetime = time.Hour // Максимальное время жизни соединения

	// Подключение
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed config connection: %v", err)
	}

	return &Postgres{
		pool: pool,
	}, nil
}

// Закрытие соединения
func (p *Postgres) Close() {
	p.pool.Close()
}

// CREATE TABLE messages (
//     text_id BIGINT PRIMARY KEY,
//     user_id BIGINT NOT NULL,
//     text TEXT NOT NULL,
//     status TEXT NOT NULL DEFAULT 'ok',
// );