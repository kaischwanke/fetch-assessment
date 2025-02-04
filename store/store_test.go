package store

import (
	"fetch-assessment/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {

	store := NewReceiptStore()

	assert.Equal(t, 0, len(store.receipts))

	uuid, err := store.Store(model.Receipt{
		Retailer: "test",
	})
	assert.NoError(t, err)

	assert.Equal(t, 1, len(store.receipts))

	item := store.GetReceipt(uuid.String())

	assert.Equal(t, 1, len(store.receipts))

	assert.Equal(t, "test", item.Retailer)

}
