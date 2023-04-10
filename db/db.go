package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Database(host, user, pass, dbname, port string) (*sql.DB, error) {
	link := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, port)
	db, err := sql.Open("postgres", link)
	if err != nil {
		return nil, err
	}

	return db, nil
}
