package utils

import (
	"fetch-assessment/model"
	"reflect"
	"runtime"
	"strconv"
	"unicode"
)

func StripNonAlphanumeric(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func GetFunctionName(rule func(model.Receipt) int) any {
	funcName := runtime.FuncForPC(reflect.ValueOf(rule).Pointer()).Name()
	return funcName
}

func ParseTotal(receipt model.Receipt) (float64, error) {
	return strconv.ParseFloat(receipt.Total, 64)
}
