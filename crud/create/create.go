package create

import (
	"database/sql"
	"log"
)

// ShoesTable creates the shoes table
func ShoesTable(db *sql.DB) {
	stmt, err := db.Prepare(
		"CREATE Table shoes" +
			"(id int NOT NULL AUTO_INCREMENT," +
			"brand varchar(128)," +
			"model varchar(128)," +
			"color varchar(128)," +
			"size int," +
			"price decimal(7,2)," +
			"stock int," +
			"PRIMARY KEY (id));")
	if err != nil {
		log.Fatal("db.Prepare returned an error:", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("db.Exec returned an error:", err)
	} else {
		log.Println("table created successfully")
	}
}
