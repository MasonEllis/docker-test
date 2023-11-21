package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
		log.Printf("The random number is %d\n", n)
		context.JSON(http.StatusOK, n)
	})

	router.Run(":8080")
}
