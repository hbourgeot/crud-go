package database

import (
	"database/sql"
	"log"
)

func makeCN() (*sql.DB, error) {
	conn := "host=localhost port=5432 user=postgres password=Reyshell dbname=TodoTech sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the DB")
	return db, nil
}
