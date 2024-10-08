package controllers

import (
	"belajarGo/db"
	"belajarGo/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, artist, price FROM album")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		albums = append(albums, album)
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	err := db.DB.QueryRow("SELECT id, title, artist, price FROM album WHERE id = $1", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, album)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(sqlStatement, newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&newAlbum.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE album SET title=$1, artist=$2, price=$3 WHERE id=$4`
	res, err := db.DB.Exec(sqlStatement, updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album updated successfully"})
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := `DELETE FROM album WHERE id=$1`
	res, err := db.DB.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}
