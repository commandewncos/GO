package main // Main Package

// Default Import
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struc Album
type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Slice Albun Datas
var albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Get Albums by Function
func getInformationAlbuns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Main Function
func main() {
	router := gin.Default()
	router.GET("/albums", getInformationAlbuns)
	router.Run("localhost:8080")

}
