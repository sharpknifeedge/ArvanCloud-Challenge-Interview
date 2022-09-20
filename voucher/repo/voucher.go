package repo

import (
	"errors"
	"fmt"
	"voucher/model"
	"voucher/pkg/voucher"

	"gorm.io/gorm"
)

type voucherRepo struct {
	db *gorm.DB
}

func NewVoucherRepo(db *gorm.DB) voucher.Repo {
	return &voucherRepo{
		db: db,
	}
}

func (r *voucherRepo) Create(v *model.Voucher) error {
	return r.db.Create(v).Error
}

func (r *voucherRepo) Use(customerID int, voucherCode string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var v model.Voucher
		if err := tx.Where("code=?", voucherCode).First(&v).Error; err != nil {
			return err
		}

		v.Remaining -= 1
		if err := tx.Save(&v).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.VoucherCustomer{
			VoucherID:  v.ID,
			CustomerID: customerID,
		}).Error; err != nil {
			return err
		}

		q := fmt.Sprintf("update wallets set balance=balance+%d where customer_id=?", v.Value)
		if err := tx.Exec(q, customerID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *voucherRepo) GetByCode(code string) (*model.Voucher, error) {
	var v model.Voucher
	if err := r.db.Where("code=?", code).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &v, nil
}

func (r *voucherRepo) CustomerUsedVouchers(code string, customerID int) ([]model.VoucherCustomer, error) {
	q := `
		select vc.* from voucher_customers vc
		inner join vouchers v on vc.voucher_id=v.id
		where v.code=? and vc.customer_id=?
	`

	var result []model.VoucherCustomer
	if err := r.db.Raw(q, code, customerID).Scan(&result).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return result, nil
}

func (r *voucherRepo) CustomerUsedVouchersDetails(code string, limit, offset int) ([]model.VoucherCustomerDetail, int64, error) {
	q := `
		select v.code, c.phone from voucher_customers vc
		inner join vouchers v on vc.voucher_id=v.id
		inner join customers c on vc.customer_id=c.id
		where v.code=?
		limit ? offset ?
	`

	var (
		result []model.VoucherCustomerDetail
		total  int64
	)

	v, err := r.GetByCode(code)
	if err != nil {
		return nil, 0, err
	}

	if err := r.db.Table("voucher_customers").Where("voucher_id=?", v.ID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Raw(q, code, limit, offset).Scan(&result).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	return result, total, nil
}
