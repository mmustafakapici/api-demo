package main

import (
	"api-demo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// /ping rotası
	r.GET("/ping", handlers.PingHandler)

	// /list rotası
	r.GET("/list", handlers.ListCoinsHandler)

	// /fetchCoins rotası
	r.GET("/fetchCoins", handlers.FetchCoinsHandler)

	r.Run(":8080")
}
