package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"./crud/create"
	"./crud/delete"
	"./crud/read"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sqlDBHost = flag.String("sql-db-host", "127.0.0.1", "can set the source of the SQL database, ie. host.domain.com")
	sqlDBName = flag.String("sql-db-name", "bundys", "can set the SQL database name")
	sqlDBPass = flag.String("sql-db-pass", "bundys", "set the SQL users password")
	sqlDBPort = flag.String("sql-db-port", "3306", "set the SQL port to connect to")
	sqlDBUser = flag.String("sql-db-user", "root", "set the SQL username to connect as")

	dataSourceName string
	tableName      = "shoes"
)

type ShoesMap []struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
	Size  int    `json:"size"`
	Price string `json:"price"`
	Stock int    `json:"stock"`
}

func createDBObject() (db *sql.DB) {
	// prepare the database abstraction for later use
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("sql.Open return an error:", err)
	}
	log.Println("database object created")
	return db
}

func main() {
	// parses the command-line flags and must be called after all flags are defined and before flags are accessed by the program
	flag.Parse()

	// set the database connection string
	dataSourceName = *sqlDBUser + ":" + *sqlDBPass + "@tcp(" + *sqlDBHost + ":" + *sqlDBPort + ")/" + *sqlDBName

	db := createDBObject()
	defer db.Close()

	// stats := db.Stats()
	// log.Println("stats.OpenConnections:", stats.OpenConnections)

	read.DBConfirmConnection(db)

	val, err := read.DoesShoesTableExist(db)
	if err != nil {
		log.Println("read.DoesTableExist returned:", err)
	}
	if val == false {
		create.ShoesTable(db)
		// load table
	} else {
		delete.ShoesTable(db)
	}

	fmt.Println("made it to the end of main")
}
