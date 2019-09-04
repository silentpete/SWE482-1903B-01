// Package delete is used to delete a resource.
package delete

import (
	"database/sql"
	"log"
)

// ShoesTable is used to drop the shoes table
func ShoesTable(db *sql.DB) {
	stmt, err := db.Prepare("DROP TABLE IF EXISTS shoes;")
	if err != nil {
		log.Fatal("delete shoes table db.Prepare returned an error:", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("db.Exec returned an error:", err)
	} else {
		log.Println("table dropped successfully")
	}
}
