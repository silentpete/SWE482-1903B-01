package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/silentpete/SWE482-1903B-01/bundys/api/delete"
	"github.com/silentpete/SWE482-1903B-01/bundys/api/get"
	"github.com/silentpete/SWE482-1903B-01/bundys/api/post"
	"github.com/silentpete/SWE482-1903B-01/bundys/api/put"
)

var (
	// DB is a pointer to a sql database connection object.
	DB              *sql.DB
	inventoryMapped []inventoryStructure
)

type inventoryStructure struct {
	ID    int     `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Color string  `json:"color"`
	Size  int     `json:"size"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}

// CreateObject is used to create a database object
func CreateObject(dataSourceName string) {
	// prepare the database abstraction for later use
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("sql.Open return an error:", err)
	}
	log.Println("database object created")
	DB = db
}

func loadInventory(invFile string) {
	// open a connection to the inventory file
	file, err := os.Open(invFile)
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
		post.InsertShoe(DB, shoe.Brand, shoe.Model, shoe.Color, shoe.Size, shoe.Price, shoe.Stock)
	}
}

// Load the database.
func Load(invFile string) {
	get.DBConfirmConnection(DB)
	val, err := get.DoesShoesTableExist(DB)
	if err != nil {
		log.Println("get.DoesTableExist returned:", err)
	}
	if val == false {
		post.CreateShoesTable(DB)
		loadInventory(invFile)
	}
}

// Drop the database.
func Drop() {
	get.DBConfirmConnection(DB)
	val, err := get.DoesShoesTableExist(DB)
	if err != nil {
		log.Println("get.DoesTableExist returned:", err)
	}
	if val == true {
		put.TruncateShoesTable(DB)
		delete.ShoesTable(DB)
	}
}

// enableCors is used to set Headers.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

// AllShoesHandler is a handler function to check if the shoes table exists.
func AllShoesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("the request uri:", r.RequestURI)
	get.DBConfirmConnection(DB)
	get.DoesShoesTableExist(DB)
	enableCors(&w)
	rows := get.AllShoes(DB)
	defer rows.Close()

	var shoesStruct []inventoryStructure

	log.Println("start rows.Next")
	for rows.Next() {
		var (
			brand, model, color string
			id, size, stock     int
			price               float32
		)
		if err := rows.Scan(&id, &brand, &model, &color, &size, &price, &stock); err != nil {
			log.Panicln("rows.Scan returned an error: ", err)
		}

		shoeStruct := inventoryStructure{
			ID:    id,
			Brand: brand,
			Model: model,
			Color: color,
			Size:  size,
			Price: price,
			Stock: stock,
		}
		shoesStruct = append(shoesStruct, shoeStruct)
	}
	log.Println("end rows.Next")

	bs, err := json.Marshal(shoesStruct)
	if err != nil {
		fmt.Println("marshaling error:", err)
	}
	log.Println("return results")
	fmt.Fprintf(w, string(bs))
}

// ShoesTableExistHandler is a handler function to check if the shoes table exists.
func ShoesTableExistHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("the request uri:", r.RequestURI)
	get.DBConfirmConnection(DB)
	val, err := get.DoesShoesTableExist(DB)
	if err != nil {
		log.Println("DoesShoesTableExist returned:", err)
	}
	exist := "{\"exist\": " + strconv.FormatBool(val) + "}"
	log.Println("return results")
	fmt.Fprintf(w, exist)
}
