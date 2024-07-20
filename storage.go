package main

import (
	"database/sql"
	"fmt"

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

func (s *Postgresstore) Init() error {
	return s.createAccountTable()
}

func (s *Postgresstore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *Postgresstore) CreateAccount(acc *Account) error {
	query := `insert into account
	(first_name, last_name, number, balance, created_at)
	values ($1, $2, $3, $4, $5)`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreateAt,
	)
	if err != nil {
		return nil
	}

	fmt.Printf("%+v\n", resp)
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
