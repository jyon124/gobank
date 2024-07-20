package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type Postgresstore struct {
	db *sql.DB
}

func NewPostgresstore() (*Postgresstore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgresstore{
		db: db,
	}, nil
}

func (s *Postgresstore) CreateAccount(*Account) error {
	return nil
}

func (s *Postgresstore) UpdateAccount(*Account) error {
	return nil
}

func (s *Postgresstore) DeleteAccount(id int) error {
	return nil
}

func (s *Postgresstore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
