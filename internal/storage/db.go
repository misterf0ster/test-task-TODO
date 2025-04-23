package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Psql *pgxpool.Pool
}

// Открываю коннект
func Open(url string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ошибка проверки соединения с базой данных: %w", err)
	}

	log.Println("Успешно подключились к PostgreSQL")
	return &DB{Psql: pool}, nil
}

// Закрываю соединение
func (db *DB) Close() {
	if db.Psql != nil {
		db.Psql.Close()
		log.Println("Closed")
	}
}
