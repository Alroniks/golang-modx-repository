package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/xml"
)

func Index() echo.HandlerFunc {
	type Root struct {
		XMLName	xml.Name	`xml:"root"`
		Message	string	`xml:"message"`
	}

	root := &Root{Message: "Not found"}

	return func(c echo.Context) error {
		return c.XML(http.StatusOK, root)
	}
}
