package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sqlDBHost = flag.String("sql-db-host", "127.0.0.1", "can set the source of the SQL database, ie. host.domain.com")
	sqlDBName = flag.String("sql-db-name", "bundys", "can set the SQL database name")
	sqlDBPass = flag.String("sql-db-pass", "bundys", "set the SQL users password")
	sqlDBPort = flag.String("sql-db-port", "3306", "set the SQL port to connect to")
	sqlDBUser = flag.String("sql-db-user", "root", "set the SQL username to connect as")
)

func main() {
	// parses the command-line flags and must be called after all flags are defined and before flags are accessed by the program
	flag.Parse()

	// set the database connection string
	dataSourceName := *sqlDBUser + ":" + *sqlDBPass + "@tcp(" + *sqlDBHost + ":" + *sqlDBPort + ")/" + *sqlDBName

	// prepare the database abstraction for later use
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("sql.Open return an error:", err)
	}
	defer db.Close()

	stats := db.Stats()
	fmt.Println(stats.OpenConnections)

	err = db.Ping()
	if err != nil {
		log.Fatal("db.Ping returned an error:", err)
	}

	fmt.Println("made it to the end of main")
}

// References
// http://go-database-sql.org
