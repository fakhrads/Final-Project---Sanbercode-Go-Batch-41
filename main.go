package main

import (
	"database/sql"
	"final-project/database"
	"final-project/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

var (
	DB  *sql.DB
	err error
)

// Custom error handler.
func apiErrorHandler(c *fiber.Ctx, err error) error {
	// Custom error handling logic here.
	// ...
	fmt.Println("apiErrorHandler", err)
	// Return from handler
	return fiber.DefaultErrorHandler(c, err)
}

func main() {
	// environment configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load file configuration")
	} else {
		fmt.Println("Success load file configuration")
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
		fmt.Println("DB Connection failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
