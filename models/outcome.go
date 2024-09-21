package models

import "time"

type Outcome struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Category    string    `json:"category" gorm:"not null"`
	Amount      float32   `json:"amount" gorm:"not null"`
	User        User      `json:"-" gorm:"foreignKey:UserID;references:ID"`
	UserID      uint      `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
	IsDeleted   bool      `json:"-" gorm:"default:false"`
}
