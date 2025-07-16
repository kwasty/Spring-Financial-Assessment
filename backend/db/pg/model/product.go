package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `json:"id" pg:",type:uuid"`

	Name          string `json:"name"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Brand         string `json:"brand"`
	StockQuantity int64  `json:"stock_quantity"`
	SKU           string `json:"sku"`

	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at"`
}
