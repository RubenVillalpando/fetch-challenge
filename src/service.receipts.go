package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var receipts = make(map[string]receipt)

func postReceipt(c *gin.Context) {
	var newReceipt receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := validateReceipt(newReceipt); err != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid"})
		return
	}

	id := generateId(newReceipt)
	receipts[id] = newReceipt
	fmt.Println(receipts)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getPointsFromReceipt(c *gin.Context) {
	id := c.Param("id")
	receipt, exists := receipts[id]
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"description": "No receipt found for that id"})
		return
	}
	points := getTotalPoints(receipt)
	c.JSON(http.StatusOK, gin.H{"points": points})
}
