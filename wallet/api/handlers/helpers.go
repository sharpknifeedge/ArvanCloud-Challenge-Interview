package handlers

import (
	"net/http"
	"wallet/model"

	"github.com/labstack/echo/v4"
)

func writeErr(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(model.ServiceError); ok {
		return c.JSON(http.StatusOK, &model.ErrorResponse{Message: err.Error()})
	}

	return err
}
