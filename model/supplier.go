package model

type Supplier struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"not null" json:"name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
}