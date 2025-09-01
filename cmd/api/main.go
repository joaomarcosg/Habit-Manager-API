package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	user := os.Getenv("HABIT_MANAGER_USER")
	password := os.Getenv("HABIT_MANAGER_PASSWORD")
	host := "localhost"
	port := os.Getenv("HABIT_MANAGER_PORT")
	dbName := os.Getenv("HABIT_MANAGER_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting to MySQL: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("database did not responsd: %v", err)
	}

	log.Println("connected to MySQL!")
}
