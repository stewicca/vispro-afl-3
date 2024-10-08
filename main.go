package main

import (
	"belajarGo/controllers"
	"belajarGo/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
