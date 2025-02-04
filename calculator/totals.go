package calculator

import (
	"fetch-assessment/model"
	"fetch-assessment/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const wholeDollarRulePoints = 50
const quarterDollarsRulePoints = 25
const oddDayRulePoints = 6
const afternoonTimeRulePoints = 10

// rules is a slice of functions, each defining a scoring rule applied to a Receipt.
var rules = []func(model.Receipt) int{
	retailerNamePointsRule,
	checkWholeDollarTotalRule,
	checkQuarterDollarTotalRule,
	itemPairPointsRule,
	itemsDescriptionRule,
	oddDayPointsRule,
	afternoonTimePointsRule,
}

// CalculateTotals computes the total points for a given receipt by applying a list of scoring rules.
// It iterates through predefined rules, calculates points for each rule, and returns the total points.
// The function accepts a model.Receipt as input and returns an integer representing the calculated points.
func CalculateTotals(receipt model.Receipt) int {
	points := 0
	for _, rule := range rules {
		additionalPoint := rule(receipt)
		// to be converted into debug log for production system
		fmt.Printf("%s points added by rule %s\n", strconv.Itoa(additionalPoint), utils.GetFunctionName(rule))
		points += additionalPoint
	}
	return points
}

func retailerNamePointsRule(receipt model.Receipt) int {
	nameClean := utils.StripNonAlphanumeric(receipt.Retailer)
	points := len(nameClean)
	return points
}

func checkWholeDollarTotalRule(receipt model.Receipt) int {
	totalString := receipt.Total
	if strings.HasSuffix(totalString, ".00") {
		return wholeDollarRulePoints
	}
	return 0
}

func checkQuarterDollarTotalRule(receipt model.Receipt) int {
	// parsing errors can be ignored due to preceding validation rules
	total, _ := utils.ParseTotal(receipt)
	if math.Mod(total, 0.25) == 0 {
		return quarterDollarsRulePoints
	}
	return 0
}

func itemPairPointsRule(receipt model.Receipt) int {
	return (len(receipt.Items) / 2) * 5
}

func itemsDescriptionRule(receipt model.Receipt) int {
	points := 0
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points

}

func oddDayPointsRule(receipt model.Receipt) int {
	_, _, day := receipt.PurchaseDate.Date()
	if day%2 != 0 {
		return oddDayRulePoints
	}
	return 0
}

func afternoonTimePointsRule(receipt model.Receipt) int {
	split := strings.Split(receipt.PurchaseTime, ":")
	hour, _ := strconv.Atoi(split[0])
	// edge case: 2:00pm is considered after 2pm, 3:59pm is the last applicable time before 4pm
	if hour >= 14 && hour < 16 {
		return afternoonTimeRulePoints
	}
	return 0
}
