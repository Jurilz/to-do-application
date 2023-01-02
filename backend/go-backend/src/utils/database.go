package utils

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg/v10"
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

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")

	options := &pg.Options{
		User:		DB_USER,
		Password:	DB_PASSWORD,
		Addr: 		address,
		Database: 	DB_NAME,
		PoolSize:   50,
	}

	con = pg.Connect(options)
	if con == nil {
		log.Fatal("cannot connect to postgres")
	}
	return
}