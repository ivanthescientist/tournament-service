package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init()  {
	var err error
	DB, err = sql.Open("mysql", "root:1111@tcp(localhost:3306)/tournament_db")

	if err != nil {
		log.Fatal("Error opening DB connection: ", err.Error())
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error testing DB connection", err.Error())
	}

	row := DB.QueryRow("SHOW VARIABLES LIKE 'tx_isolation';")

	var value string

	row.Scan(&value, &value)
	log.Printf("Transaction Isolation level is: %s\n", value)
}