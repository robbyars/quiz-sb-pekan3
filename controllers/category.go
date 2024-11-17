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

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	cat, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": cat,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetDetailBookbyCategory(c *gin.Context) {
	var (
		result gin.H
	)
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.Category_id = id
	books, err := repository.GetDetailBookbyCategory(database.DbConnection, &book)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)

}

func GetDetailCategory(c *gin.Context) {
	var cat structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	cat.ID = id
	err := repository.GetDetailCategory(database.DbConnection, &cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cat)

}

func InsertCategory(c *gin.Context) {
	var cat structs.Category

	err := c.BindJSON(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if cat.Created_at.IsZero() {
		cat.Created_at = time.Now()
	}

	username, err := GetUsernameFromAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cat.Created_by = username

	err = repository.InsertCategory(database.DbConnection, cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Kategori berhasil ditambahkan!!"),
	})
}

func UpdateCategory(c *gin.Context) {
	var cat structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if cat.Modified_at.IsZero() {
		cat.Modified_at = time.Now()
	}

	username, err := GetUsernameFromAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cat.Modified_by = username
	cat.ID = id

	err = repository.UpdateCategory(database.DbConnection, cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Kategori dengan ID %d berhasil diupdate", id),
	})
}

func DeleteCategory(c *gin.Context) {
	var cat structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	cat.ID = id
	err := repository.DeleteCategory(database.DbConnection, cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Kategori dengan ID %d berhasil dihapus", id),
	})
}
