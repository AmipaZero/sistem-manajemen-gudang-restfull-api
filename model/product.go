package model

type Product struct {
	ID       uint        `gorm:"primaryKey" json:"id"`
	Name     string      `json:"name"`
	SKU      string      `gorm:"uniqueIndex" json:"sku"`
	Category string      `json:"category"`
	Unit     string      `json:"unit"`
	Stock    int         `json:"stock"`

	Inbounds  []Inbound  `gorm:"foreignKey:ProductID" json:"inbounds,omitempty"`
	Outbounds []Outbound `gorm:"foreignKey:ProductID" json:"outbounds,omitempty"`
}
