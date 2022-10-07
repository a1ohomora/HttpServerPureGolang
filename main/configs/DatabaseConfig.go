package configs

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")

	dbURI := fmt.Sprintf("user=%s password=%s dbname=%s  sslmode=disable", username, password, dbName)

	var err error
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}
