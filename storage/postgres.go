package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connectionString string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Проверка и создание таблицы, если она не существует
	err = createURLsTableIfNotExists(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create URLs table: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) SaveURL(shortURL, originalURL string) error {
	exists, err := s.CheckURLExists(originalURL)
	if exists {
		return fmt.Errorf("URL already exists in the database")
	}

	_, err = s.db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortURL, originalURL)
	if err != nil {
		return fmt.Errorf("failed to save URL to database: %w", err)
	}
	return nil
}

func (s *PostgresStorage) GetURL(shortURL string) (string, bool, error) {
	var originalURL string
	err := s.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, nil
		}
		return "", false, fmt.Errorf("failed to get URL from database: %w", err)
	}
	return originalURL, true, nil
}

func (s *PostgresStorage) Close() {
	err := s.db.Close()
	if err != nil {
		log.Printf("failed to close database connection: %v", err)
	}
}
func createURLsTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			short_url VARCHAR(10) PRIMARY KEY,
			original_url varchar(100) NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create URLs table: %w", err)
	}
	return nil
}

func (s *PostgresStorage) CheckURLExists(originalURL string) (bool, error) {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM urls WHERE original_url = $1", originalURL).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check URL existence: %w", err)
	}
	return count > 0, nil
}
