package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"echo-app/internal/db"
	"echo-app/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env into os environment
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer database.Close()

	repo := user.NewRepository(database)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	e := echo.New()
	handler.RegisterRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "User service is running")
	})

	log.Fatal(e.Start(":3000"))
}
