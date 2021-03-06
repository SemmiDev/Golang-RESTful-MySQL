package config

import (
	"database/sql"
	"fmt"
)

const (
	username string = "root"
	password string = ""
	database string = "gomysql"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
