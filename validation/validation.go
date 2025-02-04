package validation

import (
	"errors"
	"fetch-assessment/model"
	"regexp"
	"strconv"
	"strings"
)

// ValidateReceipt Validate the receipt. All properties are being validated, except for purchaseDate,
// since the json parser would have failed if the date format was invalid.
func ValidateReceipt(receipt model.Receipt) (bool, error) {

	if len(receipt.Retailer) == 0 {
		return false, errors.New("retailer name is required")
	}

	if !isValidPurchaseTime(receipt.PurchaseTime) {
		return false, errors.New("purchase time must be in the format HH:MM")
	}

	if !validateItems(receipt.Items) {
		return false, errors.New("invalid items")
	}

	if !validatePrice(receipt.Total) {
		return false, errors.New("invalid total")
	}

	return true, nil
}

func validateItems(items []model.Item) bool {

	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		if len(item.ShortDescription) == 0 {
			return false
		}
		if !validatePrice(item.Price) {
			return false
		}
	}
	return true
}

func validatePrice(price string) bool {
	return regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(price)
}

func isValidPurchaseTime(purchaseTime string) bool {
	// Split the string by ":"
	parts := strings.Split(purchaseTime, ":")
	if len(parts) != 2 {
		return false
	}

	// Check if both parts are numeric
	hour, err1 := strconv.Atoi(parts[0])
	minute, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return false
	}

	// Validate the range of hour and minute
	return hour >= 0 && hour <= 23 && minute >= 0 && minute <= 59
}
