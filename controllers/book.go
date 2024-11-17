package controllers

import (
	"fmt"
	"net/http"
	"quiz-sb-pekan3/database"
	"quiz-sb-pekan3/repository"
	"quiz-sb-pekan3/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	book, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if book.Release_year < 1980 || book.Release_year > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.Total_page > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	if book.Created_at.IsZero() {
		book.Created_at = time.Now()
	}
	username, err := GetUsernameFromAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	book.Created_by = username

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		// Jika terjadi error, kirimkan pesan error yang lebih deskriptif
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Buku berhasil ditambahkan!!"),
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if book.Modified_at.IsZero() {
		book.Modified_at = time.Now()
	}
	if book.Release_year < 1980 || book.Release_year > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.Total_page > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	username, err := GetUsernameFromAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	book.Modified_by = username
	book.ID = id

	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		// Jika terjadi error, kirimkan pesan error yang lebih deskriptif
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Buku dengan ID %d berhasil diupdate", id),
	})
}

func GetDetailBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repository.GetDetailBook(database.DbConnection, &book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)

}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Buku dengan ID %d berhasil dihapus", id),
	})
}
