package voucher

import (
	"sync"
	"voucher/consts"
	"voucher/model"
)

type Repo interface {
	Create(v *model.Voucher) error
	Use(customerID int, voucherCode string) error
	GetByCode(code string) (*model.Voucher, error)
	CustomerUsedVouchers(code string, customerID int) ([]model.VoucherCustomer, error)
	CustomerUsedVouchersDetails(code string, limit, offset int) ([]model.VoucherCustomerDetail, int64, error)
}

type Service struct {
	repo Repo
	l    sync.Locker
}

func NewService(r Repo, l sync.Locker) *Service {
	return &Service{
		repo: r,
		l:    l,
	}
}

func (s *Service) Create(v *model.Voucher) error {
	if len(v.Code) < consts.VoucherCodeMinLength {
		return model.ServiceError(consts.ErrInvalidVoucherCode)
	}

	if v.Remaining == 0 {
		return model.ServiceError(consts.ErrInvalidVoucherRemaining)
	}

	old, err := s.repo.GetByCode(v.Code)
	if err != nil {
		return err
	}

	if old != nil && old.ID > 0 {
		return model.ServiceError(consts.ErrDuplicateVoucherCode)
	}

	return s.repo.Create(v)
}

func (s *Service) GetByCode(code string) (*model.Voucher, error) {
	return s.repo.GetByCode(code)
}

func (s *Service) Use(customerID int, voucherCode string) error {
	v, err := s.repo.GetByCode(voucherCode)
	if err != nil {
		return err
	}

	if v == nil || v.ID == 0 {
		return model.ServiceError(consts.ErrVoucherNotFound)
	}

	if v.Remaining == 0 {
		return model.ServiceError(consts.ErrVoucherExpired)
	}

	codes, err := s.repo.CustomerUsedVouchers(voucherCode, customerID)
	if err != nil {
		return err
	}

	if len(codes) >= int(v.MaxUse) {
		return model.ServiceError(consts.ErrAlreadyUsedVoucher)
	}

	s.l.Lock()
	defer s.l.Unlock()

	if err := s.repo.Use(customerID, voucherCode); err != nil {
		return err
	}

	return nil
}

func (s *Service) VoucherCustomersList(voucherCode string, limit, offset int) ([]model.VoucherCustomerDetail, int64, error) {
	v, err := s.repo.GetByCode(voucherCode)
	if err != nil {
		return nil, 0, err
	}

	if v == nil || v.ID == 0 {
		return nil, 0, model.ServiceError(consts.ErrVoucherNotFound)
	}

	return s.repo.CustomerUsedVouchersDetails(voucherCode, limit, offset)
}
