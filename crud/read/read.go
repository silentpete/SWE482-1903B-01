package read

import (
	"database/sql"
	"log"
)

// DoesShoesTableExist checks if a table exists in the database object.
func DoesShoesTableExist(db *sql.DB) (bool, error) {
	_, err := db.Query("SELECT 1 FROM shoes LIMIT 1;")
	if err != nil {
        log.Println("shoes tables does NOT exist")
		return false, err
    }
    log.Println("shoes table exists")
	return true, nil
}

// DBConfirmConnection is used to see if the connection succeeds to the database.
func DBConfirmConnection(db *sql.DB) {
	log.Println("test connection to database")
	err := db.Ping()
	if err != nil {
		log.Fatal("db.Ping returned an error:", err)
	}
	log.Println("test succeeded")
}
