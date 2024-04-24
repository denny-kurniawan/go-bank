package main

import (
	"database/sql"
	"fmt"
	"go-bank/database"
	"go-bank/routers"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// env configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Environment loaded")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error: Could not establish a connection with the database")
		panic(err)
	} else {
		fmt.Println("Connected to the database")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// router
	// var PORT = ":" + os.Getenv("PORT")
	var PORT = ":" + "8080"
	routers.StartServer().Run(PORT)
}
