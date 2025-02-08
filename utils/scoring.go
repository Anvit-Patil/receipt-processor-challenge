package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"
)

// Will compute points based on receipt rules.
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1. One point per alphanumeric character in the retailer name
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))
	total, _ := strconv.ParseFloat(receipt.Total, 64)

	// 2. 50 points if total is a round number (e.g., "10.00")
	if total == math.Floor(total) {
		points += 50
	}

	// 3. 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// 5. Item description length check
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. 6 points if purchase day is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 7. 10 points if time is between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points
}
