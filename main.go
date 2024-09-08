package main

import (
	"fmt"
	"gin-simple-api/databases"
	"gin-simple-api/src/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// menghilangkan mode debug
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// menggunakan recovery middleware
	router.Use(gin.Recovery())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	fmt.Printf("Server is running on http://localhost:%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	routes.AuthRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":" + port)
}
