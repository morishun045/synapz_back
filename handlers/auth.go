// authのHandler

package handlers

import (

	"net/http"
	"time"
	"synapz_back/lib"

	"github.com/labstack/echo/v4"

)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username != "foo" || password != "bar" {
		return echo.ErrUnauthorized
	}
	token, err := utils.GenerateJWT(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
