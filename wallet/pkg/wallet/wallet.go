package wallet

import "wallet/model"

type Repo interface {
	GetByCustomerID(customerID int) (*model.Wallet, error)
}

type Service struct {
	repo Repo
}

func NewService(r Repo) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetByCustomerID(customerID int) (*model.Wallet, error) {
	return s.repo.GetByCustomerID(customerID)
}
