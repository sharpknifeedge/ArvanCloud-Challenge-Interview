package model

type Voucher struct {
	ID        int    `json:"id" gorm:"primarykey"`
	Code      string `json:"code"`
	Remaining int    `json:"remaining"`
	Value     int    `json:"value"`
	MaxUse    uint   `json:"max_use"`
}

type VoucherCustomer struct {
	ID         int `gorm:"primarykey"`
	VoucherID  int
	CustomerID int
}

type VoucherCustomerDetail struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}
