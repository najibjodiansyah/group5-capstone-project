package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(connectionString string) (*sql.DB, error) {
	return sql.Open("mysql", connectionString)
}
