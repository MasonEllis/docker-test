package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Pong")
	})

	router.GET("/randomNumber", func(context *gin.Context) {
		n := rand.Int()
		context.JSON(http.StatusOK, n)
	})

	router.Run(":8080")
}
