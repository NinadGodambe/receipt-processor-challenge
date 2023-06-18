package util

import (
	"receipt-processor-challenge/model"
	"testing"
	"github.com/stretchr/testify/assert"
)


// TestCalculatePointsRetailerName tests the function CalculatePointsRetailerName()
func TestCalculatePointsRetailerName(t *testing.T) {
	points := CalculatePointsRetailerName("Target")
	assert.Equal(t, 6, points, "Expected 6 points")
}

// TestCalculatePointsTotalRoundDollar tests the function CalculatePointsTotalRoundDollar()
func TestCalculatePointsTotalRoundDollar(t *testing.T) {
	points := CalculatePointsTotalRoundDollar("35.00")
	assert.Equal(t, 50, points, "Expected 50 points")
}

// TestCalculatePointsTotalMultipleOfQuarter tests the function CalculatePointsTotalMultipleOfQuarter()
func TestCalculatePointsTotalMultipleOfQuarter(t *testing.T) {
	points := CalculatePointsTotalMultipleOfQuarter("10.50")
	assert.Equal(t, 25, points, "Expected 25 points")
}

// TestCalculatePointsItems tests the function CalculatePointsItems()
func TestCalculatePointsItems(t *testing.T) {
	items := []model.Item{
		{ShortDescription: "Item 1", Price: "10.00"},
		{ShortDescription: "Item 2", Price: "5.00"},
		{ShortDescription: "Item 3", Price: "2.50"},
		{ShortDescription: "Item 4", Price: "3.00"},
	}

	points := CalculatePointsItems(items)
	assert.Equal(t, 10, points, "Expected 10 points")
}

// TestCalculatePointsItemDescription tests the function CalculatePointsItemDescription()
func TestCalculatePointsItemDescription(t *testing.T) {
	items := []model.Item{
		{ShortDescription: "Paper Towels", Price: "10.00"},
		{ShortDescription: "Fruit", Price: "5.00"},
	}

	points := CalculatePointsItemDescription(items)
	assert.Equal(t, 2, points, "Expected 3 points")
}

// TestCalculatePointsPurchaseDate tests the function CalculatePointsPurchaseDate()
func TestCalculatePointsPurchaseDate(t *testing.T) {
	points := CalculatePointsPurchaseDate("2022-01-01")
	assert.Equal(t, 6, points, "Expected 6 points")
}

// TestCalculatePointsPurchaseTime tests the function CalculatePointsPurchaseTime()
func TestCalculatePointsPurchaseTime(t *testing.T) {
	points := CalculatePointsPurchaseTime("14:30")
	assert.Equal(t, 10, points, "Expected 10 points")
}
