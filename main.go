package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Confessions", Artist: "User", Price: 56.99},
	{ID: "2", Title: "Nothing was the same", Artist: "Drake", Price: 17.99},
	{ID: "3", Title: "The Emancipation of Mimi", Artist: "Mariah Carey", Price: 39.99},
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedData album

	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i].Title = updatedData.Title
			albums[i].Artist = updatedData.Artist
			albums[i].Price = updatedData.Price
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}
