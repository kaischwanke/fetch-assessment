package calculator

import (
	"fetch-assessment/model"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCalculationRules(t *testing.T) {
	t.Run("retailerNamePointsRule", func(t *testing.T) {
		require.Equal(t, 6, retailerNamePointsRule(model.Receipt{Retailer: "Target"}))
		require.Equal(t, 13, retailerNamePointsRule(model.Receipt{Retailer: "Bed Bath & Beyond"}))
	})

	t.Run("checkWholeDollarTotalRule", func(t *testing.T) {
		require.Equal(t, 50, checkWholeDollarTotalRule(model.Receipt{Total: "10.00"}))
		require.Equal(t, 0, checkWholeDollarTotalRule(model.Receipt{Total: "10.01"}))
	})

	t.Run("itemPairPointsRule", func(t *testing.T) {
		require.Equal(t, 0, itemPairPointsRule(model.Receipt{Items: make([]model.Item, 0)}))
		require.Equal(t, 0, itemPairPointsRule(model.Receipt{Items: make([]model.Item, 1)}))
		require.Equal(t, 5, itemPairPointsRule(model.Receipt{Items: make([]model.Item, 2)}))
		require.Equal(t, 5, itemPairPointsRule(model.Receipt{Items: make([]model.Item, 3)}))
		require.Equal(t, 10, itemPairPointsRule(model.Receipt{Items: make([]model.Item, 4)}))
	})

	t.Run("checkQuarterDollarTotalRule", func(t *testing.T) {
		require.Equal(t, 25, checkQuarterDollarTotalRule(model.Receipt{Total: "0.00"})) // 0 is considered a multiple for any number
		require.Equal(t, 25, checkQuarterDollarTotalRule(model.Receipt{Total: "0.25"}))
		require.Equal(t, 25, checkQuarterDollarTotalRule(model.Receipt{Total: "1.50"}))
		require.Equal(t, 25, checkQuarterDollarTotalRule(model.Receipt{Total: "2.75"}))
		require.Equal(t, 0, checkQuarterDollarTotalRule(model.Receipt{Total: "2.99"}))
	})

	t.Run("itemsDescriptionRule", func(t *testing.T) {
		require.Equal(t, 0, itemsDescriptionRule(model.Receipt{Items: []model.Item{
			{
				ShortDescription: "xx",
			},
		}}))
		require.Equal(t, 1, itemsDescriptionRule(model.Receipt{Items: []model.Item{
			{
				ShortDescription: "xxx",
				Price:            "1.00",
			},
		}}))
		require.Equal(t, 3, itemsDescriptionRule(model.Receipt{Items: []model.Item{
			{
				ShortDescription: "xxx",
				Price:            "1.00",
			},
			{
				ShortDescription: "yyy",
				Price:            "10.00",
			},
		}}))
	})

	t.Run("oddDayPointsRule", func(t *testing.T) {
		require.Equal(t, 6, oddDayPointsRule(model.Receipt{PurchaseDate: mustParseDate("2025-01-01")}))
		require.Equal(t, 0, oddDayPointsRule(model.Receipt{PurchaseDate: mustParseDate("2025-01-02")}))
		require.Equal(t, 6, oddDayPointsRule(model.Receipt{PurchaseDate: mustParseDate("2025-02-03")}))
		require.Equal(t, 0, oddDayPointsRule(model.Receipt{PurchaseDate: mustParseDate("2025-10-10")}))

	})

	t.Run("afternoonTimePointsRule", func(t *testing.T) {
		require.Equal(t, 0, afternoonTimePointsRule(model.Receipt{PurchaseTime: "13:00"}))
		require.Equal(t, 0, afternoonTimePointsRule(model.Receipt{PurchaseTime: "13:59"}))
		require.Equal(t, 10, afternoonTimePointsRule(model.Receipt{PurchaseTime: "14:00"}))
		require.Equal(t, 10, afternoonTimePointsRule(model.Receipt{PurchaseTime: "14:01"}))
		require.Equal(t, 10, afternoonTimePointsRule(model.Receipt{PurchaseTime: "15:00"}))
		require.Equal(t, 10, afternoonTimePointsRule(model.Receipt{PurchaseTime: "15:59"}))
		require.Equal(t, 0, afternoonTimePointsRule(model.Receipt{PurchaseTime: "10:00"}))
	})

}

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name    string
		receipt model.Receipt
		want    int
	}{
		{
			name: "Example #1 from requirement",
			receipt: model.Receipt{
				Retailer: "Target",
				Total:    "35.35",
				Items: []model.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12"},
				},
				PurchaseDate: mustParseDate("2022-01-01"),
				PurchaseTime: "13:01",
			},
			want: 28,
		},
		{
			name: "Example #2 from requirement",
			receipt: model.Receipt{
				Retailer: "M&M Corner Market",
				Total:    "9.00",
				Items: []model.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				PurchaseDate: mustParseDate("2022-03-20"),
				PurchaseTime: "14:33",
			},
			want: 109,
		},
		{
			name: "Basic points calculation",
			receipt: model.Receipt{
				Retailer: "Walmart",
				Total:    "25.00",
				Items: []model.Item{
					{ShortDescription: "Apple", Price: "2.00"},
				},
				PurchaseDate: mustParseDate("2023-11-15"),
				PurchaseTime: "03:15",
			},
			want: 88,
		},
		{
			name: "Empty receipt",
			receipt: model.Receipt{
				Retailer:     "",
				Total:        "0.00",
				Items:        []model.Item{},
				PurchaseDate: mustParseDate("2023-11-14"),
				PurchaseTime: "01:00",
			},
			want: 75,
		},
		{
			name: "Item with description length multiple of 3",
			receipt: model.Receipt{
				Retailer: "ABC",
				Total:    "8.00",
				Items: []model.Item{
					{ShortDescription: "Grape", Price: "9.00"},
					{ShortDescription: "Melon", Price: "3.00"},
				},
				PurchaseDate: mustParseDate("2023-11-13"),
				PurchaseTime: "05:00",
			},
			want: 89,
		},
		{
			name: "Time not between 2:00 to 4:00",
			receipt: model.Receipt{
				Retailer: "Market",
				Total:    "10.00",
				Items: []model.Item{
					{ShortDescription: "Grape", Price: "9.00"},
				},
				PurchaseDate: mustParseDate("2023-11-14"),
				PurchaseTime: "04:30",
			},
			want: 81,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTotals(tt.receipt)
			require.Equal(t, tt.want, got)
		})
	}

}

func mustParseDate(dateStr string) openapi_types.Date {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err)
	}
	return openapi_types.Date{date}
}
