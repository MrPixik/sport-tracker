package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mrpixik/sport-tracker/internal/auth/storage"
)

const driverName = "pgx"

type Storage struct {
	db *sql.DB
}

func New(storageDSN string) (*Storage, error) {
	const op = "storage.postgres.New"

	//storageDSN = "postgres://pixik:pixikpass2025@localhost:5432/auth-db?sslmode=disable"
	//fmt.Println("Using DSN:", storageDSN)

	db, err := sql.Open(driverName, storageDSN)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(ctx context.Context, login, passHash string) error {
	op := "storage.postgres.CreateUser"
	query := `INSERT INTO users (login, password_hash) VALUES ($1, $2)`

	if _, err := s.db.ExecContext(ctx, query, login, passHash); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return fmt.Errorf("%s: %w", op, storage.ErrUserAlreadyExists)
			}
			return fmt.Errorf("%s: %w", op, err)
		}
	}
	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, login, passHash string) error {
	op := "storage.postgres.DeleteUser"
	query := `DELETE FROM users WHERE login=$1 AND password_hash=$2`

	res, err := s.db.ExecContext(ctx, query, login, passHash)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	} else if rows == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return nil

}
