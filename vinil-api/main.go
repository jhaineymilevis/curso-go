package main

import (
	"fmt"
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
	{ID: "1", Title: "BLUE RAIN", Artist: "GUNS AND ROSES", Price: 12.9},
	{ID: "2", Title: "19", Artist: "ABBA", Price: 10.9},
	{ID: "3", Title: "REVOLUTION", Artist: "REM", Price: 14.9},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(
		http.StatusOK,
		albums,
	)
}

func addAlbum(ctx *gin.Context) {

	var newAlbum album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(
		http.StatusCreated,
		albums,
	)
}

func getAlbumById(ctx *gin.Context) {

	id := ctx.Param("id")

	fmt.Println(id)
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(
				http.StatusOK,
				a,
			)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})

}
func main() {
	router := gin.Default()

	router.GET(
		"/albums",
		getAlbums,
	)
	router.GET(
		"/albums/:id",
		getAlbumById,
	)
	router.POST(
		"/albums",
		addAlbum,
	)

	router.Run("localhost:8080")
}
