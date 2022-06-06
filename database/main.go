package database

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitialiseDatabase() {
	fmt.Println("Initialising Database")
	_, err := gorm.Open(sqlite.Open("movarrdb.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error occured while creating database, " + err.Error())
	}
}
