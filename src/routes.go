package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("receipts/process", postReceipt)
	router.GET("receipts/:id/points", getPointsFromReceipt)

	router.Run("localhost:8080")
}
