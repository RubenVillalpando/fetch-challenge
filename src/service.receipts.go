package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
)

var receipts = make(map[string]receipt)

func postReceipt(c *gin.Context) {
	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}
	id := generateId(newReceipt)
	receipts[id] = newReceipt
	fmt.Println(receipts)
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func getPointsFromReceipt(c *gin.Context) {
	points := 32
	c.IndentedJSON(http.StatusOK, gin.H{"points": points})
}

func generateId(rec receipt) string {
	builder := strings.Builder{}
	builder.WriteString("REC-")
	builder.WriteString(rec.PurchaseDate)
	builder.WriteString("-")
	builder.WriteString(strings.ReplaceAll(rec.PurchaseTime, ":", "-"))
	builder.WriteString("-")
	builder.WriteString(generateRandChars(6))
	return builder.String()
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func generateRandChars(n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		builder.WriteByte(letters[rand.Intn(len(letters))])
	}
	return builder.String()
}
