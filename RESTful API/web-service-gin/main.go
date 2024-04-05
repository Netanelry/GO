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
    router.GET("/casinos/:id", getCasinoByID)
    router.POST("/casinos", postCasinos)

    router.Run("localhost:8080")
}

// getCasinos responds with the list of all casinos as JSON.
func getCasinos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, casinos)
}

// postCasinos adds a new casino from JSON received in the request body
// to the list of existing casinos.
func postCasinos(c *gin.Context) {
    var newCasino casino

    // Call BindJSON to bind the received JSON to newCasino.
    if err := c.BindJSON(&newCasino); err != nil {
        return
    }

    // Add the new casino to the slice.
    casinos = append(casinos, newCasino)
    c.IndentedJSON(http.StatusCreated, newCasino)
}

// getCasinoByID locates the casino whose ID value matches the id
// parameter sent in the URL path by the client, then returns that
// casino as a response.
func getCasinoByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of casinos, looking for
    // a casino whose ID value matches the parameter.
    for _, a := range casinos {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "casino not found"})
}
