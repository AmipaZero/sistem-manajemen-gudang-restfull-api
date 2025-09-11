package model

import "time"

type Role string

const (
	Admin Role = "admin"
	Staff Role = "staff"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Role         Role      `gorm:"type:enum('admin','staff');not null" json:"role"`
	Token        *string   `gorm:"type:text" json:"token,omitempty"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
