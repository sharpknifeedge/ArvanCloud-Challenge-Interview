package dep

import (
	"sync"
	"voucher/pkg/container"
	"voucher/pkg/voucher"
	"voucher/repo"
)

func Init() (container.Container, error) {
	db, err := repo.Connect()
	if err != nil {
		return container.Container{}, err
	}

	return container.Container{
		VoucherService: *voucher.NewService(repo.NewVoucherRepo(db), &sync.Mutex{}),
	}, nil
}
