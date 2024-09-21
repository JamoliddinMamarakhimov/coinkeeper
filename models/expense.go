package models

type Expense struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	UserID      uint    `json:"user_id"`
	AccountID   uint    `json:"account_id"`  // ID карты, с которой сняты деньги
	CategoryID  uint    `json:"category_id"` // ID категории расходов
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
