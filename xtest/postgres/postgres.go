package postgres

import (
	"database/sql"

	// Import postgresql driver
	_ "github.com/lib/pq"
)

const (
	maxIdleConns = 4
	maxOpenConns = 8
)

// Xtest for postgresql
type Xtest struct {
	db *sql.DB
}

// New client for postgres xtest backend
func New(db *sql.DB) *Xtest {
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	return &Xtest{
		db: db,
	}
}

// TestPingDefaultDSN ping postgres backend
func (x *Xtest) TestPingDefaultDSN() (string, error) {
	err := x.db.Ping()
	if err != nil {
		return "Ping failed!", err
	}
	return "Ping success!", nil
}

// TestPingNewDSN ping postgres backend
func (x *Xtest) TestPingNewDSN(dsn string) (string, error) {
	postgresClient, err := sql.Open("postgres", dsn)
	if err != nil {
		return "Postgres db open failed!", err
	}

	err = postgresClient.Ping()
	postgresClient.Close()
	if err != nil {
		return "Ping failed!", err
	}
	return "Ping success!", nil
}
