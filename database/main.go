package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitialiseDatabase() {
	fmt.Println("Initialising Database")
	db, err := sql.Open("sqlite3", "./movarrdb.db")
	if err != nil {
		log.Fatal("Cannot open database {}", err.Error())
	}
	CreateTables(db)
}
