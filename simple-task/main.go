package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Beee Tain", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jean", Artist: "Gerry Muligan", Price: 18.89},
	{ID: "3", Title: "Sarah aVaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pkng",
		})
	})

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.StaticFile("/", "index.html")

	fmt.Println("Starting HTTP Server on port 8080")
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// updateAlbumByID updates an album by its ID with the data provided in the request body.
func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Find the index of the album in the slice.
	index := -1
	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	var updatedAlbum album

	// Call BindJSON to bind the received JSON to updatedAlbum.
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Update the album in the slice.
	albums[index] = updatedAlbum
	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// deleteAlbumByID deletes an album by its ID.
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range albums {
		if a.ID == id {
			// Remove the album from the slice by slicing it out.
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
