package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/silentpete/SWE482-1903B-01/server"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dropDatabase  = flag.Bool("drop-database", false, "pass to truncate and drop the shoes table")
	help          = flag.Bool("help", false, "will display this message")
	inventoryFile = flag.String("inventory-file", "./inventory.json", "./path/to/inventory.json")
	loadDatabase  = flag.Bool("load-database", false, "pass to load the inventory-file")
	sqlDBHost     = flag.String("sql-db-host", "127.0.0.1", "can set the source of the SQL database, ie. host.domain.com")
	sqlDBName     = flag.String("sql-db-name", "bundys", "can set the SQL database name")
	sqlDBPass     = flag.String("sql-db-pass", "bundys", "set the SQL users password")
	sqlDBPort     = flag.String("sql-db-port", "3306", "set the SQL port to connect to")
	sqlDBUser     = flag.String("sql-db-user", "root", "set the SQL username to connect as")

	dataSourceName string
	tableName      = "shoes"
)

func main() {
	// parses the command-line flags and must be called after all flags are defined and before flags are accessed by the program
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	log.Println("set the dataSourceName")
	dataSourceName := *sqlDBUser + ":" + *sqlDBPass + "@tcp(" + *sqlDBHost + ":" + *sqlDBPort + ")/" + *sqlDBName

	log.Println("create a global database connection object:", dataSourceName)
	server.CreateObject(dataSourceName)

	log.Println("set database close")
	defer server.DB.Close()

	if *loadDatabase == true {
		server.Load(*inventoryFile)
		os.Exit(0)
	}

	if *dropDatabase == true {
		server.Drop()
		os.Exit(0)
	}

	log.Println("set up handlers")
	log.Println("add /shoesTableExist")
	http.HandleFunc("/shoesTableExist", server.ShoesTableExistHandler)

	log.Println("add /allShoesHandler")
	http.HandleFunc("/allShoes", server.AllShoesHandler)

	log.Println("start the http server (this will block, keeping the server running)")
	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("made it to the end of main")
}
