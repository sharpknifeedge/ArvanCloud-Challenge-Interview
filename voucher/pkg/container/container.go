package container

import "voucher/pkg/voucher"

type Container struct {
	VoucherService voucher.Service
}

func NewContainer(vs voucher.Service) Container {
	return Container{
		VoucherService: vs,
	}
}
