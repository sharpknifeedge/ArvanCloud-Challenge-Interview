package handlers

import (
	"net/http"
	"strconv"
	"voucher/model"
	"voucher/pkg/voucher"

	"github.com/labstack/echo/v4"
)

func CreateVoucher(service voucher.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var v model.Voucher

		if err := c.Bind(&v); err != nil {
			return err
		}

		if err := service.Create(&v); err != nil {
			return writeErr(c, err)
		}

		return c.NoContent(http.StatusOK)
	}
}

func VoucherCustomersList(service voucher.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

		list, total, err := service.VoucherCustomersList(c.Param("code"), pageSize, (page-1)*pageSize)
		if err != nil {
			return writeErr(c, err)
		}

		return c.JSON(http.StatusOK, model.CreatePagination(list, total, page, pageSize))
	}
}

func UseVoucher(service voucher.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			CustomerID int    `json:"customer_id"`
			Code       string `json:"code"`
		}

		if err := c.Bind(&req); err != nil {
			return err
		}

		if err := service.Use(req.CustomerID, req.Code); err != nil {
			return writeErr(c, err)
		}

		return c.NoContent(http.StatusNoContent)
	}
}
