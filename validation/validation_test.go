package validation

import (
	"fetch-assessment/model"
	"testing"
)

func TestValidateReceipt(t *testing.T) {
	tests := []struct {
		name    string
		receipt model.Receipt
		want    bool
		wantErr string
	}{
		{
			name: "valid receipt",
			receipt: model.Receipt{
				Retailer: "Retailer",
				Items: []model.Item{{
					Price:            "1.00",
					ShortDescription: "first",
				}, {
					Price:            "2.00",
					ShortDescription: "second",
				}},
				PurchaseTime: "15:30",
				Total:        "25.50",
			},
			want: true,
		},
		{
			name: "missing retailer",
			receipt: model.Receipt{
				Retailer:     "",
				Items:        []model.Item{{}, {}},
				PurchaseTime: "15:30",
				Total:        "25.50",
			},
			want:    false,
			wantErr: "retailer name is required",
		},
		{
			name: "missing items",
			receipt: model.Receipt{
				Retailer:     "",
				Items:        make([]model.Item, 0),
				PurchaseTime: "15:30",
				Total:        "25.50",
			},
			want:    false,
			wantErr: "retailer name is required",
		},
		{
			name: "invalid time format",
			receipt: model.Receipt{
				Retailer:     "Retailer",
				Items:        []model.Item{{}, {}},
				PurchaseTime: "1530",
				Total:        "25.50",
			},
			want:    false,
			wantErr: "purchase time must be in the format HH:MM",
		},
		{
			name: "invalid items",
			receipt: model.Receipt{
				Retailer: "Retailer",
				Items: []model.Item{{
					Price:            "1.00",
					ShortDescription: "",
				}},
				PurchaseTime: "15:30",
				Total:        "25.50",
			},
			want:    false,
			wantErr: "invalid items",
		},
		{
			name: "invalid total",
			receipt: model.Receipt{
				Retailer: "Retailer",
				Items: []model.Item{{
					Price:            "1.00",
					ShortDescription: "first",
				}, {
					Price:            "2.00",
					ShortDescription: "second",
				}},
				PurchaseTime: "15:30",
				Total:        "invalid",
			},
			want:    false,
			wantErr: "invalid total",
		},
		{
			name: "invalid time",
			receipt: model.Receipt{
				Retailer: "Retailer",
				Items: []model.Item{{
					Price:            "1.00",
					ShortDescription: "first",
				}, {
					Price:            "2.00",
					ShortDescription: "second",
				}},
				PurchaseTime: "invalid",
				Total:        "100.00",
			},
			want:    false,
			wantErr: "purchase time must be in the format HH:MM",
		},
		{
			name: "invalid time - minutes",
			receipt: model.Receipt{
				Retailer: "Retailer",
				Items: []model.Item{{
					Price:            "1.00",
					ShortDescription: "first",
				}, {
					Price:            "2.00",
					ShortDescription: "second",
				}},
				PurchaseTime: "10:x",
				Total:        "100.00",
			},
			want:    false,
			wantErr: "purchase time must be in the format HH:MM",
		},
		{
			name: "invalid time - hours",
			receipt: model.Receipt{
				Retailer: "Retailer",
				Items: []model.Item{{
					Price:            "1.00",
					ShortDescription: "first",
				}, {
					Price:            "2.00",
					ShortDescription: "second",
				}},
				PurchaseTime: "99:10",
				Total:        "100.00",
			},
			want:    false,
			wantErr: "purchase time must be in the format HH:MM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateReceipt(tt.receipt)
			if got != tt.want {
				t.Errorf("ValidateReceipt() = %v, want %v", got, tt.want)
			}
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("ValidateReceipt() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if tt.wantErr != "" {
				t.Errorf("ValidateReceipt() error = nil, wantErr %v", tt.wantErr)
			}
		})
	}
}
