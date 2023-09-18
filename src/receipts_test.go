package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPointsFromDate(t *testing.T) {
	assert.Equal(t, 6, getPointsFromDate(POINTS_RECEIPT_1))
	assert.Equal(t, 0, getPointsFromDate(POINTS_RECEIPT_2))
}
func TestGetPointsFromTime(t *testing.T) {
	assert.Equal(t, 0, getPointsFromTime(POINTS_RECEIPT_1))
	assert.Equal(t, 10, getPointsFromTime(POINTS_RECEIPT_2))
	assert.Equal(t, 10, getPointsFromTime(LOWER_INBOUNDS_CASE_TIME))
	assert.Equal(t, 0, getPointsFromTime(LOWER_EDGE_CASE_TIME))
	assert.Equal(t, 0, getPointsFromTime(LOWER_OUTBOUNDS_CASE_TIME))
	assert.Equal(t, 10, getPointsFromTime(UPPER_INBOUNDS_CASE_TIME))
	assert.Equal(t, 0, getPointsFromTime(UPPER_EDGE_CASE_TIME))
	assert.Equal(t, 0, getPointsFromTime(UPPER_OUTBOUNDS_CASE_TIME))
}

func TestGetPointsFromItemQtty(t *testing.T) {
	assert.Equal(t, 10, getPointsFromItemQtty(POINTS_RECEIPT_1))
	assert.Equal(t, 10, getPointsFromItemQtty(POINTS_RECEIPT_2))
	assert.Equal(t, 0, getPointsFromItemQtty(ONE_ITEM_CASE_ITEM_QTTY))
	assert.Equal(t, 5, getPointsFromItemQtty(TWO_ITEM_CASE_ITEM_QTTY))
	assert.Equal(t, 15, getPointsFromItemQtty(SEVEN_ITEM_CASE_ITEM_QTTY))
}
func TestGetPointsFromItemShortDesc(t *testing.T) {
	assert.Equal(t, 6, getPointsFromItemShortDesc(POINTS_RECEIPT_1))
	assert.Equal(t, 0, getPointsFromItemShortDesc(POINTS_RECEIPT_2))
	assert.Equal(t, 0, getPointsFromItemShortDesc(ALL_SPACES_CASE_SHORTDESC))
	assert.Equal(t, 0, getPointsFromItemShortDesc(FREE_ITEM_CASE_SHORTDESC))
	assert.Equal(t, 1, getPointsFromItemShortDesc(ONE_CENT_CASE_SHORTDESC))
	assert.Equal(t, 1, getPointsFromItemShortDesc(SPACES_AROUND_CASE_SHORTDESC))
}

func TestGetPointsFromTotal(t *testing.T) {
	assert.Equal(t, 0, getPointsFromTotal(POINTS_RECEIPT_1))
	assert.Equal(t, 75, getPointsFromTotal(POINTS_RECEIPT_2))
}

func TestGetPointsFromRetailerName(t *testing.T) {
	assert.Equal(t, 6, getPointsFromRetailerName(POINTS_RECEIPT_1))
	assert.Equal(t, 14, getPointsFromRetailerName(POINTS_RECEIPT_2))
	assert.Equal(t, 2, getPointsFromRetailerName(SYMBOLS_ON_NAME_CASE_RETAILER))
}

func TestValidateReceipt(t *testing.T) {
	assert.Equal(t, "Invalid Retailer\n", validateReceipt(INVALID_RETAILER_RECEIPT))
	assert.Equal(t, "Key: 'receipt.PurchaseDate' Error:Field validation for 'PurchaseDate' failed on the 'datetime' tag\n", validateReceipt(INVALID_PURCHASE_DATE_RECEIPT))
	assert.Equal(t, "Key: 'receipt.PurchaseTime' Error:Field validation for 'PurchaseTime' failed on the 'datetime' tag\n", validateReceipt(INVALID_PURCHASE_TIME_RECEIPT))
	assert.Equal(t, "Invalid Price on item #0\nInvalid short description on item #0\n", validateReceipt(INVALID_ITEM_RECEIPT))
}
