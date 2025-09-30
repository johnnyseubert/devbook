package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/johnnyseubert/devbook/src/config"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseConectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
