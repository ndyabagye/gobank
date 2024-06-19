package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account(
		id	serial primary key,
		first_name varchar(50),
		last_name  varchar(50),
		number serial,
		balance serial,
		created_at timestamp,
		updated_at timestamp
		)`

	_, err := s.db.Exec(query)
	return err
}

func (p *PostgresStore) CreateAccount(a *Account) error {
	return nil
}

func (p *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (p *PostgresStore) UpdateAccount(a *Account) error {
	return nil
}

func (p *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
