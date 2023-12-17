package main

import (
	"api-demo/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	// / rotası
	r.GET("/", handlers.HomeHandler)

	// /ping rotası
	r.GET("/ping", handlers.PingHandler)

	// /list rotası
	r.GET("/list", handlers.ListCoinsHandler)

	//r.Run(":8080")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run("0.0.0.0:" + port))
}
