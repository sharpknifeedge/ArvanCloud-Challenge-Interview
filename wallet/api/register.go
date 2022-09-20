package api

import (
	"fmt"
	"wallet/api/handlers"
	"wallet/pkg/container"
	"wallet/utils"

	"github.com/labstack/echo/v4"
)

func Run(c container.Container) {
	e := echo.New()

	customersAPI := e.Group("/customers")
	addRoutes(c, customersAPI)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", utils.EnvInt("APP_PORT", 8080))))
}

func addRoutes(c container.Container, customers *echo.Group) {
	customers.POST("", handlers.CreateCustomer(c.CustomerService))
	customers.GET("", handlers.GetCustomers(c.CustomerService))
	customers.GET("/wallet/:customerID", handlers.CustomerWallet(c.WalletService))
}
