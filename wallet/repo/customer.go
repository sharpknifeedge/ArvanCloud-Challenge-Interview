package repo

import (
	"errors"
	"wallet/model"
	"wallet/pkg/customer"

	"gorm.io/gorm"
)

func NewCustomerRepo(db *gorm.DB) customer.Repo {
	return &customerRepo{db: db}
}

type customerRepo struct {
	db *gorm.DB
}

func (r *customerRepo) Create(c *model.Customer) (int, error) {
	var id int

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(c).Error; err != nil {
			return err
		}

		id = c.ID

		return tx.Create(&model.Wallet{CustomerID: id}).Error
	})

	return id, err
}

func (r *customerRepo) All(limit int, offset int) ([]model.Customer, int64, error) {
	var (
		result []model.Customer
		total  int64
	)

	if err := r.db.Table("customers").Count(&total).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&result).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	return result, total, nil
}

func (r *customerRepo) GetByPhone(phone string) (*model.Customer, error) {
	var c model.Customer
	if err := r.db.Where("phone=?", phone).First(&c).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &c, nil
}
