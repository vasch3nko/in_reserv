package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(host, port, user, password, dbname, sslmode string) (*Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Migrate(path string) error {
	const dialect = "postgres"

	goose.SetLogger(goose.NopLogger())
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	return goose.Up(r.db, path)
}

// Close closes the db connection
func (r *Repository) Close() error {
	return r.db.Close()
}
