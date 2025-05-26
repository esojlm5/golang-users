package main

import (
	"echo-app/internal/db"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var demo int

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(c echo.Context) error {
	demo = 20
	id := c.Param("id")
	fmt.Println("ID:", id) // Print the ID to the console
	return c.String(http.StatusOK, "User ID: "+id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	demo = 40
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func saveUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	fmt.Println("HTTP methods:", http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	return c.JSON(http.StatusOK, user)
}

func main() {
	database, err := db.Connect("myuser", "mypassword", "mydb", "localhost", 5432)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer database.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Pass `database` to your handlers/services here
	log.Println("Connected to DB!")

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	//
	e.Logger.Fatal(e.Start(":1323"))
}
