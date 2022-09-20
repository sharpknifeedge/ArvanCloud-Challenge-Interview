package handlers

import (
	"net/http"
	"strconv"
	"wallet/model"
	"wallet/pkg/customer"
	"wallet/pkg/wallet"

	"github.com/labstack/echo/v4"
)

func CreateCustomer(service customer.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cs model.Customer
		if err := c.Bind(&cs); err != nil {
			return err
		}

		resp, err := service.Create(&cs)
		if err != nil {
			return writeErr(c, err)
		}

		return c.JSON(http.StatusCreated, resp)
	}
}

func GetCustomers(service customer.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

		customers, total, err := service.GetAll(pageSize, (page-1)*pageSize)
		if err != nil {
			return writeErr(c, err)
		}

		return c.JSON(http.StatusOK, model.CreatePagination(customers, total, page, pageSize))
	}
}

func CustomerWallet(service wallet.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID, err := strconv.Atoi(c.Param("customerID"))
		if err != nil {
			return err
		}

		w, err := service.GetByCustomerID(customerID)
		if err != nil {
			return writeErr(c, err)
		}

		return c.JSON(http.StatusOK, w)
	}
}
