package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Coin struct'ı CoinGecko API'sinden gelen verileri temsil eder
type Coin struct {
	ID     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"current_price"`
}

// PingHandler, /ping rotasına hizmet verir
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

// ListCoinsHandler, CoinGecko API'sinden verileri alarak /list rotasına hizmet verir
func ListCoinsHandler(c *gin.Context) {
	// CoinGecko API'sine GET isteği göndermek için HTTP istemcisi oluşturun
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "CoinGecko API'ye bağlanırken bir hata oluştu",
		})
		return
	}
	defer resp.Body.Close()

	// API yanıtını JSON olarak ayrıştırın
	var coins []Coin
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "API yanıtını ayrıştırırken bir hata oluştu",
		})
		return
	}

	// Çekilen verileri JSON olarak yanıtlayın
	c.JSON(http.StatusOK, coins)
}

// FetchCoinsHandler, CoinGecko API'sinden verileri alarak /fetchCoins rotasına hizmet verir
func FetchCoinsHandler(c *gin.Context) {
	// CoinGecko API'sine GET isteği göndermek için HTTP istemcisi oluşturun
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "CoinGecko API'ye bağlanırken bir hata oluştu",
		})
		return
	}
	defer resp.Body.Close()

	// API yanıtını JSON olarak ayrıştırın
	var coins []Coin
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "API yanıtını ayrıştırırken bir hata oluştu",
		})
		return
	}

	// Çekilen verileri JSON olarak yanıtlayın
	c.JSON(http.StatusOK, coins)
}
