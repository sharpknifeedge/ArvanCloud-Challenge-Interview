package dep

import (
	"wallet/pkg/container"
	"wallet/pkg/customer"
	"wallet/pkg/wallet"
	"wallet/repo"
)

func Init() (container.Container, error) {
	db, err := repo.Connect()
	if err != nil {
		return container.Container{}, err
	}

	walletRepo := repo.NewWalletRepo(db)
	customerRepo := repo.NewCustomerRepo(db)

	walletService := wallet.NewService(walletRepo)
	customerService := customer.NewService(customerRepo, walletService)

	return container.NewContainer(*walletService, *customerService), nil
}
