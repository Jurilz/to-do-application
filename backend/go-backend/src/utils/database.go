package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	//"github.com/jmoiron/sqlx"
)


const (
	HOST		= "localhost"
	PORT 		= "5432"
	DB_USER 	= "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME 	= "tasks"
)


//DB set up

func OpenDBConnection() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s " +
		"password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("DB Connection established...")

	return db
}