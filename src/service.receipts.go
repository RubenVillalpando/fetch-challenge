package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var receipts = make(map[string]receipt)
var validate *validator.Validate
var retailerRegex = regexp.MustCompile(`^\S+$`)
var totalRegex = regexp.MustCompile(`^\d+\.\d{2}$`)
var shortDescRegex = regexp.MustCompile(`^[\w\s\-]+$`)
var priceRegex = regexp.MustCompile(`^\d+\.\d{2}$`)

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

type PointGetter func(receipt) int

func getPointsFromReceipt(c *gin.Context) {

	id := c.Param("id")
	receipt := receipts[id]

	points := getTotalPoints(receipt)

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

var functionsForPoints = []PointGetter{
	getPointsFromRetailerName,
	getPointsFromTotal,
	getPointsFromItemQtty,
	getPointsFromItemShortDesc,
	getPointsFromDate,
	getPointsFromTime,
}

func getTotalPoints(rec receipt) int {
	points := 0
	for _, fn := range functionsForPoints {
		points += getPointsFromFunc(fn, rec)
	}
	return points
}

func getPointsFromFunc(fn PointGetter, rec receipt) int {
	return fn(rec)
}

func getPointsFromRetailerName(rec receipt) int {
	points := 0
	for _, char := range rec.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}
	return points
}

func getPointsFromTotal(rec receipt) int {
	points := 0
	switch cents := rec.Total[len(rec.Total)-2 : len(rec.Total)]; cents {
	case "00":
		points = 75
	case "25":
		points = 25
	case "50":
		points = 25
	case "75":
		points = 25
	default:
		points = 0
	}
	return points
}

func getPointsFromItemQtty(rec receipt) int {
	return (len(rec.Items) / 2) * 5
}

func getPointsFromItemShortDesc(rec receipt) int {
	points := 0
	for _, item := range rec.Items {
		trimmedLength := len(strings.Trim(item.ShortDescription, " \t\r\n"))
		if trimmedLength%3 == 0 && trimmedLength != 0 {
			if price, err := strconv.ParseFloat(item.Price, 32); err == nil {
				points += int(math.Ceil(price * 0.2))

			}
		}
	}
	return points
}

func getPointsFromDate(rec receipt) int {
	points := 0
	date, dateErr := time.Parse("2006-01-02", rec.PurchaseDate)
	if dateErr == nil {
		day := date.Day()
		if day%2 == 1 {
			points += 6
		}
	}
	return points
}

func getPointsFromTime(rec receipt) int {
	points := 0
	time, timeErr := time.Parse("15:04", rec.PurchaseTime)
	if timeErr == nil {
		hour := time.Hour()
		minute := time.Minute()
		timeAfter2pm := hour == 14 && minute > 0 || hour > 14
		timeBefore4pm := hour < 16
		if timeAfter2pm && timeBefore4pm {
			points += 10
		}
	}
	return points
}
