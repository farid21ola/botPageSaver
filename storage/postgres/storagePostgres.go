package postgres

import (
	"botPageSaver/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	DB *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) PickRandom(ctx context.Context, userName string) (*storage.Page, error) {
	page := storage.Page{
		URL:      "",
		UserName: userName,
	}

	q := `SELECT "URL" FROM "Page" WHERE "UserName"=$1 ORDER BY RANDOM()`

	if err := s.DB.QueryRow(ctx, q, userName).Scan(&page.URL); err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	s.Remove(ctx, &page)
	return &page, nil
}

func (s *Storage) Remove(ctx context.Context, p *storage.Page) error {
	q := `DELETE FROM "Page" WHERE "URL" = $1 AND "UserName" = $2`

	if _, err := s.DB.Exec(ctx, q, p.URL, p.UserName); err != nil {
		return fmt.Errorf("can't remove page: %w", err)
	}
	return nil
}

func (s *Storage) Save(ctx context.Context, p *storage.Page) error {
	q := `INSERT INTO "Page" ("URL", "UserName") VALUES ($1,$2)`

	if _, err := s.DB.Exec(ctx, q, p.URL, p.UserName); err != nil {
		return err
	}
	return nil
}

func (s *Storage) IsExists(ctx context.Context, p *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM "Page" WHERE "URL" = $1 AND "UserName" = $2`

	var count int

	if err := s.DB.QueryRow(ctx, q, p.URL, p.UserName).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check if page exists: %w", err)
	}

	return count > 0, nil
}
