package model

type Customer struct {
	ID    int    `json:"id" gorm:"primarykey"`
	Phone string `json:"phone"`
}

type Wallet struct {
	ID         int `json:"id" gorm:"primarykey"`
	CustomerID int `json:"customer_id"`
	Balance    int `json:"balance"`
}
