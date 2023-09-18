package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("receipts/process", postReceipt)
	router.GET("receipts/:id/points", getPointsFromReceipt)

	router.Run("127.0.0.1:8080")
}
