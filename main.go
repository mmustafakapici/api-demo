// main.go dosyası api-demo projesinin giriş noktasıdır.

package main

import (
	"api-demo/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	r := gin.Default()

	// / rotası
	r.GET("/", handlers.HomeHandler)

	r.GET("/hello", handlers.HelloHandler)

	// /ping rotası
	r.GET("/ping", handlers.PingHandler)

	// /list rotası
	r.GET("/list", handlers.ListCoinsHandler)

	// Coingecko veri çekme işlemini belirli aralıklarla sürekli olarak yapacak bir goroutine başlat
	//go handlers.FetchCoinDataPeriodically(60 * time.Second)

	//r.Run(":8080")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(r.Run("0.0.0.0:" + port))
}
