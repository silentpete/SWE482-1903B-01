package put

import (
	"database/sql"
	"log"
)

// TruncateShoesTable will remove all data in the shoes table.
func TruncateShoesTable(db *sql.DB) {
	stmt, err := db.Prepare("TRUNCATE TABLE shoes;")
	if err != nil {
		log.Fatal("truncate shoes table db.Prepare returned an error:", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("truncate db.Exec returned an error:", err)
	} else {
		log.Println("table truncate successfully")
	}
}
