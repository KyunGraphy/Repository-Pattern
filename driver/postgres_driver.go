package driver

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	// Using pointer -> Avoid being copied new DB object
	SQL *sql.DB
}

// Create new DB object
var Postgres = &PostgresDB{}

// Connect DB function
func Connect(host, port, user, password, dbname string) (*PostgresDB) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)	
	if err != nil {
		panic(err)
	}			

	// If connect successfully, assign the connection to SQL field in PostgresDB struct
	Postgres.SQL = db
	return Postgres
}