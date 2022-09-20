package repo

import (
	"wallet/model"
	"wallet/pkg/wallet"

	"gorm.io/gorm"
)

func NewWalletRepo(db *gorm.DB) wallet.Repo {
	return &walletRepo{db: db}
}

type walletRepo struct {
	db *gorm.DB
}

func (r *walletRepo) GetByCustomerID(customerID int) (*model.Wallet, error) {
	var w model.Wallet
	return &w, r.db.Where("customer_id=?", customerID).First(&w).Error
}
