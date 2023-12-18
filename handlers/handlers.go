// handlers.go dosyası api-demo projesinin rotalarını tanımlar.
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Coin struct'ı CoinGecko API'sinden gelen verileri temsil eder
type Coin struct {
	ID        string  `json:"id"`
	Rank      int     `json:"market_cap_rank"`
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Price     float64 `json:"current_price"`
	Image     string  `json:"image"`
	Change    float64 `json:"price_change_percentage_24h"`
	MarketCap float64 `json:"market_cap"`
}

// CoinDataCache, Coingecko verilerini önbellekte tutacak bir yapı
type CoinDataCache struct {
	Data        []Coin
	LastUpdated time.Time
}

var coinCache CoinDataCache

// HomeHandler, / rotasına hizmet verir
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// HelloHandler, /hello rotasına hizmet verir
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, onur!",
	})
}

// PingHandler, /ping rotasına hizmet verir
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

// ListCoinsHandler, önbellekteki verileri kullanarak /list rotasına hizmet verir
func ListCoinsHandler(c *gin.Context) {
	if len(coinCache.Data) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veri bulunamadı. Lütfen daha sonra tekrar deneyin.",
		})
		return
	}
	c.JSON(http.StatusOK, coinCache.Data)
}

// FetchCoinDataPeriodically, Coingecko API'den veri çekme işlemini belirli aralıklarla sürekli olarak yapacak bir fonksiyon
func FetchCoinDataPeriodically(interval time.Duration) {
	for {
		// Coingecko API'sine GET isteği göndermek için HTTP istemcisi oluşturun
		resp, err := http.Get("https://api.coingecko.com/apzi/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
		if err != nil {
			// Hata oluştuğunda log kaydı alabilirsiniz
			// log.Printf("CoinGecko API'ye bağlanırken bir hata oluştu: %v", err)
			continue
		}
		defer resp.Body.Close()

		// API yanıtını JSON olarak ayrıştırın
		var coins []Coin
		if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
			// Hata oluştuğunda log kaydı alabilirsiniz
			// log.Printf("API yanıtını ayrıştırırken bir hata oluştu: %v", err)
			continue
		}

		// Verileri önbellekte saklayın
		coinCache.Data = coins
		coinCache.LastUpdated = time.Now()

		// Belirli bir süre bekle
		time.Sleep(interval)
	}
}
