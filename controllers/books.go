package controllers

import (
	"net/http"

	"github.com/James-nyagatare/gin-bookstore/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct{
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct{
	Title string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context){
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context){
	var input CreateBookInput
	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func UpdateBook(c *gin.Context){
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	var input UpdateBookInput
	validationErr := c.ShouldBindJSON(&input)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context){
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})

}

func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }