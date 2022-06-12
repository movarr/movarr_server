package database

import (
	"database/sql"
	"log"
)

func createContactTable(db *sql.DB) {
	const cmd string = `
  CREATE TABLE IF NOT EXISTS contacts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name TEXT,
  last_name TEXT
);
 `

	if _, err := db.Exec(cmd); err != nil {
		log.Fatal("Error creating contacts table: ", err.Error())
	}

}

func createNumbersTable(db *sql.DB) {
	const cmd string = `
  CREATE TABLE IF NOT EXISTS numbers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number INTEGER NOT NULL,
    contact_id INTEGER NOT NULL,
    FOREIGNKEY(contact_id) REFERENCES contacts(id)
);
  `
	if _, err := db.Exec(cmd); err != nil {
		log.Fatal("Error creating backups table: ", err.Error())
	}
}

func createBackupTable(db *sql.DB) {
	const cmd string = `
  CREATE TABLE IF NOT EXISTS backup (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date INTEGER NOT NULL
);
`

	if _, err := db.Exec(cmd); err != nil {
		log.Fatal("Error creating backup table: ", err.Error())
	}
}

func createMessagesTable(db *sql.DB) {
	const cmd string = `
  CREATE TABLE IF NOT EXISTS message (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    address TEXT NOT NULL,
    thread_id INTEGER NOT NULL,
    person INTEGER NOT NULL,
    date INTEGER NOT NULL,
    date_sent INTEGER NOT NULL,
    protocol INTEGER NOT NULL,
    read INTEGER NOT NULL,
    status INTEGER NOT NULL,
    type INTEGER NOT NULL,
    reply_path_present INTEGER NOT NULL,
    subject TEXT,
    body TEXT,
    service_center TEXT,
    locked INTEGER NOT NULL,
    error_code INTEGER NOT NULL,
    seen INTEGER NOT NULL,
  );
  `
	if _, err := db.Exec(cmd); err != nil {
		log.Fatal("Error creating message table: ", err.Error())
	}
}

func CreateTables(db *sql.DB) {
	createContactTable(db)
	createNumbersTable(db)
	createMessagesTable(db)
	createBackupTable(db)
}
