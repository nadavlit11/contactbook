package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

var postgresClientOnce sync.Once
var postgresClient PostgresClient
var pConn *sql.DB

type PostgresClient interface {
	GetConn() *sql.DB
}

type PostgresClientImpl struct {
}

func NewPostgresClient() PostgresClient {
	postgresClientOnce.Do(func() {
		postgresClient = &PostgresClientImpl{}
		connect()
	})
	return postgresClient
}

func connect() {
	fmt.Println("Connecting to postgres")
	connStr := "host=db port=5432 user=admin password=admin dbname=contactbook sslmode=disable"
	pConn, _ = sql.Open("postgres", connStr)

	usersQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT
		)`
	_, err := pConn.Exec(usersQuery)
	if err != nil {
		fmt.Println("here1")
	}

	contactsQuery := `
		CREATE TABLE IF NOT EXISTS contacts (
			id SERIAL PRIMARY KEY,
			user_id INT,
			first_name TEXT,
			last_name TEXT,
			phone TEXT,
			address TEXT
		)`
	_, err = pConn.Exec(contactsQuery)
	if err != nil {
		panic(err)
	}
}

func (p *PostgresClientImpl) GetConn() *sql.DB {
	return pConn
}
