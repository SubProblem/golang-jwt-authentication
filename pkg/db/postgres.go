package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"	
)

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb() (*PostgresDb, error) {
	connStr := "user=postgres dbname=go-project password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDb{
		db: db,
	}, nil
}

func (pg *PostgresDb) Init() error {
	return pg.CreateUserTable()
}

