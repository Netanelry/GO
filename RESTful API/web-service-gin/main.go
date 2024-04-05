package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type casino struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Location string  `json:"location"`
    Rank  float64 `json:"rank"`
}

// casinos slice to seed record data.
var casinos = []casino{
    {ID: "1", Name: "Bellagio", Location: "Las Vegas, USA", Rank: 98.8},
    {ID: "2", Name: "The Venetian Macao", Location: "Macau, China", Rank: 86.36},
    {ID: "3", Name: "Monte Carlo Casino", Location: "Monte Carlo, Monaco", Rank: 75},
}

func main() {
    router := gin.Default()
    router.GET("/casinos", getCasinos)

    router.Run("localhost:8080")
}

// getCasinos responds with the list of all casinos as JSON.
func getCasinos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, casinos)
}
