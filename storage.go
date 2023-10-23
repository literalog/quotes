package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateQuote(*Quote) error
	GetQuoteByID(string) (*Quote, error)
	GetQuotes() ([]*Quote, error)
	UpdateQuote(*Quote) error
	DeleteQuote(string) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {

	connStr := ""

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil

}

func (s *PostgresStorage) Init() error {
	return nil
}

func (s *PostgresStorage) createQuoteTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS quotes (
		id UUID PRIMARY KEY,
		author VARCHAR(50) NOT NULL,
		quote TEXT NOT NULL,
	)`

	_, err := s.db.Exec(query)

	return err

}

func (s *PostgresStorage) CreateQuote(q *Quote) error {
	query := `
	INSERT INTO quotes (author, quote)
	VALUES ($1, $2)`

	_, err := s.db.Query(query, q.Author, q.Text)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) GetQuoteByID(string) (*Quote, error) {
	return nil, nil
}

func (s *PostgresStorage) GetQuotes() ([]*Quote, error) {

	rows, err := s.db.Query("SELECT * FROM quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotes := []*Quote{}

	for rows.Next() {
		q := new(Quote)
		err := rows.Scan(
			&q.Id,
			&q.Author,
			&q.Text)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, q)
	}

	return quotes, nil
}

func (s *PostgresStorage) UpdateQuote(*Quote) error {
	return nil
}

func (s *PostgresStorage) DeleteQuote(string) error {
	return nil
}
