package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var receipts = make(map[string]receipt)
var validate *validator.Validate
var retailerRegex = regexp.MustCompile("^\\S+$")
var totalRegex = regexp.MustCompile("^\\d+\\.\\d{2}$")
var shortDescRegex = regexp.MustCompile("^[\\w\\s\\-]+$")
var priceRegex = regexp.MustCompile("^\\d+\\.\\d{2}$")

func postReceipt(c *gin.Context) {
	var newReceipt receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := validateReceipt(newReceipt); err != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
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

const digits = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func generateRandChars(n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		builder.WriteByte(digits[rand.Intn(len(digits))])
	}
	return builder.String()
}

func validateReceipt(rec receipt) string {
	validate = validator.New()
	err := validate.Struct(&rec)
	builder := strings.Builder{}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			builder.WriteString(err.Error())
			builder.WriteString("\n")
		}
	}

	validRetailer := retailerRegex.MatchString(rec.Retailer)
	validTotal := totalRegex.MatchString(rec.Total)

	if !validRetailer {
		builder.WriteString("Invalid Retailer\n")
	}
	if !validTotal {
		builder.WriteString("Invalid Total\n")
	}
	for i, item := range rec.Items {
		if !priceRegex.MatchString(item.Price) {
			builder.WriteString(fmt.Sprintf("Invalid Price on item #%d\n", i))
		}
		if !shortDescRegex.MatchString(item.ShortDescription) {
			builder.WriteString(fmt.Sprintf("Invalid short description on item #%d\n", i))
		}
	}

	return builder.String()
}
