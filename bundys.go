package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"./crud/create"
	"./crud/delete"
	"./crud/read"
	"./crud/update"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dropDatabase  = flag.Bool("drop-database", false, "pass to truncate and drop the shoes table")
	inventoryFile = flag.String("inventory-file", "./inventory.json", "./path/to/inventory.json")
	loadDatabase  = flag.Bool("load-database", false, "pass to load the inventory-file")
	sqlDBHost     = flag.String("sql-db-host", "127.0.0.1", "can set the source of the SQL database, ie. host.domain.com")
	sqlDBName     = flag.String("sql-db-name", "bundys", "can set the SQL database name")
	sqlDBPass     = flag.String("sql-db-pass", "bundys", "set the SQL users password")
	sqlDBPort     = flag.String("sql-db-port", "3306", "set the SQL port to connect to")
	sqlDBUser     = flag.String("sql-db-user", "root", "set the SQL username to connect as")

	dataSourceName  string
	tableName       = "shoes"
	inventoryMapped inventoryStructure
)

type inventoryStructure []struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Color string  `json:"color"`
	Size  int     `json:"size"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
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

func loadInventory(db *sql.DB) {
	// open a connection to the inventory file
	file, err := os.Open(*inventoryFile)
	if err != nil {
		log.Fatal("opening the inventory file returned an error:", err)
	}
	defer file.Close()

	// get the file information
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// create a byte slice object of the size of the file
	bs := make([]byte, stat.Size())

	// put the contents of the file in the byte slice. The byte slice is unicode... so basically it is an array of the unicode character of each character read in.
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	// because the json matches the structure, we can take the byte slice and expand it into the structure.
	err = json.Unmarshal(bs, &inventoryMapped)
	if err != nil {
		log.Panic("error Unmarshalling: ", err)
	}

	for _, shoe := range inventoryMapped {
		create.InsertShoe(db, shoe.Brand, shoe.Model, shoe.Color, shoe.Size, shoe.Price, shoe.Stock)
	}
}

func main() {
	// parses the command-line flags and must be called after all flags are defined and before flags are accessed by the program
	flag.Parse()
	// set the database connection string
	dataSourceName = *sqlDBUser + ":" + *sqlDBPass + "@tcp(" + *sqlDBHost + ":" + *sqlDBPort + ")/" + *sqlDBName
	db := createDBObject()
	defer db.Close()
	read.DBConfirmConnection(db)

	if *loadDatabase == true {
		val, err := read.DoesShoesTableExist(db)
		if err != nil {
			log.Println("read.DoesTableExist returned:", err)
		}
		if val == false {
			create.ShoesTable(db)
			loadInventory(db)
		}
	}

	if *dropDatabase == true {
		val, err := read.DoesShoesTableExist(db)
		if err != nil {
			log.Println("read.DoesTableExist returned:", err)
		}
		if val == true {
			update.TruncateShoesTable(db)
			delete.ShoesTable(db)
		}
	}

	fmt.Println("made it to the end of main")
}
