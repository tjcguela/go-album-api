package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// schema
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// populating w/ data
var albums = []album{
	{
		ID:     1,
		Title:  "The Misadventures of Billy and Mandy",
		Artist: "Cartoon Network",
		Price:  550.50,
	},
	{
		ID:     2,
		Title:  "Lilo and Stich: The Movie",
		Artist: "Disney Channel",
		Price:  2200.56,
	},
}

// API methods
// GET all albums
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// Get By ID
func getAlbumByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr) // convert ID to Int

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album ID"})
		return
	}

	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// POST new album
func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)

}
