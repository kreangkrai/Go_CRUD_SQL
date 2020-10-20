package Controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "192.168.15.202"
var port = 1433
var user = "sa"
var password = "p@ssw0rd"
var database = "Mee"

var db *sql.DB

func Connect() {
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", server, database, user, password, port)
	var err error
	db, err = sql.Open("mssql", connString)
	//defer db.Close()
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	log.Printf("Connected!\n")

}

func CheckVersion() string {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Ping database failed:", err.Error())
	}

	var version string
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&version)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	log.Printf("%s\n", version)

	return version
}
