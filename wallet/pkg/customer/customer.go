package customer

import (
	"wallet/model"
	"wallet/pkg/wallet"
)

type Repo interface {
	Create(*model.Customer) (int, error)
	All(limit, offset int) ([]model.Customer, int64, error)
	GetByPhone(phone string) (*model.Customer, error)
}

type Service struct {
	repo          Repo
	walletService *wallet.Service
}

func NewService(repo Repo, walletService *wallet.Service) *Service {
	return &Service{
		repo:          repo,
		walletService: walletService,
	}
}

func (s *Service) Create(customer *model.Customer) (*model.CustomerCreateResponse, error) {
	if len(customer.Phone) == 0 {
		return nil, model.ServiceError("customer phone cannot be empty")
	}

	oldCustomer, err := s.repo.GetByPhone(customer.Phone)
	if err != nil {
		return nil, err
	}

	if oldCustomer != nil && oldCustomer.ID > 0 {
		return nil, model.ServiceError("duplicate customer phone")
	}

	customerID, err := s.repo.Create(customer)
	if err != nil {
		return nil, err
	}

	w, err := s.walletService.GetByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	return &model.CustomerCreateResponse{
		ID:       customerID,
		WalletID: w.ID,
	}, nil
}

func (s *Service) GetAll(limit, offset int) ([]model.Customer, int64, error) {
	return s.repo.All(limit, offset)
}
