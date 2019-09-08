// Package get is used to retrieve a resource.
package get

import (
	"database/sql"
	"log"
)

// AllShoes is used to get all the show form the shoes table.
func AllShoes(db *sql.DB) *sql.Rows {
	stmt, err := db.Prepare("SELECT * FROM shoes ORDER BY id DESC;")
	if err != nil {
		log.Fatal("all shoes db.Prepare returned an error:", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("all shoes query returned an error: ", err)
	}

	return rows
}

// DoesShoesTableExist checks if a table exists in the database object.
func DoesShoesTableExist(db *sql.DB) (bool, error) {
	log.Println("check for shoes table")
	_, err := db.Query("SELECT 1 FROM shoes LIMIT 1;")
	if err != nil {
		log.Println("shoes tables does NOT exist")
		return false, err
	}

	log.Println("shoes table exists")
	return true, nil
}

// DBConfirmConnection is used to see if the connection succeeds to the database.
func DBConfirmConnection(db *sql.DB) (bool, error) {
	log.Println("test connection to database")
	err := db.Ping()
	if err != nil {
		log.Println("db.Ping returned an error:", err)
		return false, err
	}
	log.Println("connection succeeded")
	return true, nil
}
