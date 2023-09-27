// config/database.go
package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v"),
						os.Getenv("DB_USER"),
						os.Getenv("DB_PASSWORD"),
						os.Getenv("DB_HOST"),
						os.Getenv("DB_PORT"),
						os.Getenv("DB_NAME"),
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
}
