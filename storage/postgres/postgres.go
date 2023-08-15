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
		fmt.Println("can't connect to DataBase Link: ", err)
		return nil, err
	}
	if err := DB.Ping(context.Background()); err != nil {
		fmt.Println("can't ping DataBase Link: ", err)
		return nil, err
	}
	fmt.Println("connected to postgres")
	return DB, nil
}
