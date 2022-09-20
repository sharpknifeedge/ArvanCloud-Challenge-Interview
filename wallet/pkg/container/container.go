package container

import (
	"wallet/pkg/customer"
	"wallet/pkg/wallet"
)

type Container struct {
	WalletService   wallet.Service
	CustomerService customer.Service
}

func NewContainer(
	walletService wallet.Service,
	customerService customer.Service) Container {

	return Container{
		WalletService:   walletService,
		CustomerService: customerService,
	}
}
