package database

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"log"
)

func InitialiseDatabase() {
	fmt.Println("Initialising Database")
	db, err := sql.Open("sqlite", "./movarrdb.db")
	if err != nil {
		log.Fatal("Cannot open database {}", err.Error())
	}
	CreateTables(db)
}
