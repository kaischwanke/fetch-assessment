// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package model

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Item defines model for Item.
type Item struct {
	// Price The total price payed for this item.
	Price string `json:"price"`

	// ShortDescription The Short Product Description for the item.
	ShortDescription string `json:"shortDescription"`
}

// Receipt defines model for Receipt.
type Receipt struct {
	Items []Item `json:"items"`

	// PurchaseDate The date of the purchase printed on the receipt.
	PurchaseDate openapi_types.Date `json:"purchaseDate"`

	// PurchaseTime The time of the purchase printed on the receipt. 24-hour time expected.
	PurchaseTime string `json:"purchaseTime"`

	// Retailer The name of the retailer or store the receipt is from.
	Retailer string `json:"retailer"`

	// Total The total amount paid on the receipt.
	Total string `json:"total"`
}

// PostReceiptsProcessJSONRequestBody defines body for PostReceiptsProcess for application/json ContentType.
type PostReceiptsProcessJSONRequestBody = Receipt
