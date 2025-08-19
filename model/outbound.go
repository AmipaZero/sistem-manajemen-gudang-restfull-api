package model

import "time"

type Outbound struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	SentAt      time.Time `json:"sent_at"`
	Destination string    `json:"destination"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
