package main

import (
	"database/sql"
	"fmt"
	"os"

	"sanbercode-golang-final-project-nando/pkg/database"
	"sanbercode-golang-final-project-nando/pkg/routers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV Configuration
	err = godotenv.Load("pkg/config/.env")
	if err != nil {
		fmt.Println("\nFailed load file environment")
	} else {
		fmt.Println("\nSuccess read file environment")
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
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	routers.StartServer().Run(":" + os.Getenv("PORT"))
}
