// Package post is used to create a new resource
package post

import (
	"database/sql"
	"log"
)

// CreateShoesTable creates the shoes table
func CreateShoesTable(db *sql.DB) {
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

// InsertShoe will insert the inventory into the shoes table
func InsertShoe(db *sql.DB, brand, model, color string, size int, price float32, stock int) {
	// will need the database and a map/shoes structure
	stmt, err := db.Prepare("INSERT INTO shoes(brand,model,color,size,price,stock) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal("INSERT INTO returned error:", err)
	}
	res, err := stmt.Exec(brand, model, color, size, price, stock)
	if err != nil {
		log.Fatal("statement exec returned an error:", err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
