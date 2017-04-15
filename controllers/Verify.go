package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/xml"
)

func Verify() echo.HandlerFunc {
	type Status struct {
		XMLName		xml.Name	`xml:"status"`
		Verified		int		`xml:"verified"`
	}

	status := &Status{Verified: 1}

	return func(c echo.Context) error {
		return c.XML(http.StatusOK, status)
	}
}