package util

import (
	"receipt-processor-challenge/model"
	"math"
	"strconv"
	"strings"
	"unicode"
	"github.com/google/uuid"
)

// CalculatePoints adds up all the points gained from all the rules
func CalculatePoints(receipt *model.Receipt) int {
	points := 0
	
	// Summing up all the points earned based on different rules
	points += CalculatePointsRetailerName(receipt.Retailer)	
	points += CalculatePointsTotalRoundDollar(receipt.Total)	
	points += CalculatePointsTotalMultipleOfQuarter(receipt.Total)	
	points += CalculatePointsItems(receipt.Items)	
	points += CalculatePointsItemDescription(receipt.Items)	
	points += CalculatePointsPurchaseDate(receipt.PurchaseDate)
	points += CalculatePointsPurchaseTime(receipt.PurchaseTime)

	return points
}

// CalculatePointsRetailerName calculates points based on Rule 1.
// Rule 1: One point for every alphanumeric character in the retailer name.
func CalculatePointsRetailerName(retailer string) int {
	points := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}
	return points
}

// CalculatePointsTotalRoundDollar calculates points based on Rule 2.
// Rule 2: 50 points if the total is a round dollar amount with no cents.
func CalculatePointsTotalRoundDollar(total string) int {
	amount, _ := ParsePrice(total)
	if amount == math.Trunc(amount) {
		return 50
	}
	return 0
}

// CalculatePointsTotalMultipleOfQuarter calculates points based on Rule 3.
// Rule 3: 25 points if the total is a multiple of 0.25.
func CalculatePointsTotalMultipleOfQuarter(total string) int {
	amount, _ := ParsePrice(total)

	// Check for multiples of 0.25 using mod operator
	if amount > 0 && math.Mod(amount, 0.25) == 0 {
		return 25
	}
	return 0
}

// CalculatePointsItems calculates points based on Rule 4.
// Rule 4: 5 points for every two items on the receipt.
func CalculatePointsItems(items []model.Item) int {
	count := len(items)
	return count / 2 * 5
}

// CalculatePointsItemDescription calculates points based on Rule 5.
// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2
// and round up to the nearest integer. The result is the number of points earned.
func CalculatePointsItemDescription(items []model.Item) int {
	points := 0
	for _, item := range items {

		// Trimming the external spaces from the item description
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))

		// Checking if trimmed length is a multiple of 3
		if trimmedLength%3 == 0 {
			price, _ := ParsePrice(item.Price)

			// Multiplying the price by 0.2 and rounding up to the nearest integer
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

// CalculatePointsPurchaseDate calculates points based on Rule 6.
// Rule 6: 6 points if the day in the purchase date is odd.
func CalculatePointsPurchaseDate(date string) int {
	day, _ := strconv.Atoi(date[8:10])

	// Check if the day is odd using mod operator
	if day%2 != 0 {
		return 6
	}
	return 0
}

// CalculatePointsPurchaseTime calculates points based on Rule 7.
// Rule 7: 10 points if the time of purchase is after 2:00 pm and before 4:00 pm.
func CalculatePointsPurchaseTime(time string) int {
	hour, _ := strconv.Atoi(time[:2])

	// Checking if the time is after 2:00 pm and before 4:00
	if hour >= 14 && hour < 16 {
		return 10
	}
	return 0
}

// ParsePrice parses a string(given orice in receipts) into a float64 value.
func ParsePrice(priceStr string) (float64, error) {
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}

// GenerateUUID generates a random UUID to be associated with each receipt.
func GenerateUUID() string {
	return uuid.New().String()
}