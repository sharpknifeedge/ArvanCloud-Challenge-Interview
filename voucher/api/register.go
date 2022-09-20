package api

import (
	"fmt"
	"voucher/api/handlers"
	"voucher/pkg/container"
	"voucher/utils"

	"github.com/labstack/echo/v4"
)

func Run(c container.Container) {
	e := echo.New()

	vouchersAPI := e.Group("/vouchers")
	addRoutes(c, vouchersAPI)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", utils.EnvInt("APP_PORT", 8080))))
}

func addRoutes(c container.Container, voucher *echo.Group) {
	voucher.POST("", handlers.CreateVoucher(c.VoucherService))
	voucher.GET("/used/:code", handlers.VoucherCustomersList(c.VoucherService))
	voucher.PUT("/use", handlers.UseVoucher(c.VoucherService))
}
