package main

import (
	"github.com/James-nyagatare/gin-bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Run()
}
