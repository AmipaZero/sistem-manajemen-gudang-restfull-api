package domain

import "time"

type Inbound struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	ReceivedAt time.Time `json:"received_at"`
	Supplier   string    `json:"supplier"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
