package models

type Account struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	UserID      uint    `json:"user_id"`
	Card        Card    `json:"card" gorm:"unique"`

	Balance     float64 `json:"balance"`
	Description string  `json:"description"`
}

type Card struct{
    ID uint 
}

// type Account struct {
//     ID      uint    `json:"id"`
//     UserID  uint    `json:"user_id"`
//     Balance float64 `json:"balance"`
// }