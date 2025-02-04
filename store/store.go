package store

import (
	"errors"
	"fetch-assessment/model"
	"github.com/google/uuid"
)

var ErrReceiptNotFound = errors.New("receipt not found")

// ReceiptStore uses a simple in-memory map for data storage
type ReceiptStore struct {
	receipts map[string]model.Receipt
}

func NewReceiptStore() *ReceiptStore {
	return &ReceiptStore{
		receipts: make(map[string]model.Receipt),
	}
}

func (r *ReceiptStore) Store(receipt model.Receipt) (uuid.UUID, error) {
	receiptID, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}
	r.receipts[receiptID.String()] = receipt
	return receiptID, nil
}

func (r *ReceiptStore) GetReceipt(id string) *model.Receipt {
	receipt, ok := r.receipts[id]
	if !ok {
		return nil
	}
	return &receipt
}
