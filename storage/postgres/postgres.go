package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPoolPostgres(config string) (*pgxpool.Pool, error) {
	url := config
	DB, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Println("не удалось подключиться к DataBase Link: ", err)
		return nil, err
	}
	if err := DB.Ping(context.Background()); err != nil {
		fmt.Println("не удалось пингануть к DataBase Link: ", err)
		return nil, err
	}
	fmt.Println("подключён успешно postgres")
	return DB, nil
}
