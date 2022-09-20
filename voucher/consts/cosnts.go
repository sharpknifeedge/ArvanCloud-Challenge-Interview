package consts

const (
	ErrInvalidVoucherCode      = "voucher code length should be greater than 4"
	ErrInvalidVoucherRemaining = "voucher remaining field should be greater than zero"
	ErrDuplicateVoucherCode    = "duplicate voucher code"
	ErrVoucherExpired          = "voucher expired"
	ErrAlreadyUsedVoucher      = "already used voucher"
	ErrVoucherNotFound         = "voucher not found"
)

const VoucherCodeMinLength = 4
