// user用のhandler

package handlers

import (
	
	"net/http"
	"synapz_back/lib"

	"github.com/labstack/echo/v4"

)

func Profile (c echo.Context) error {
	claims, err := lib.GetClaimsFromContext(c)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": claims.UserID,
	})
}
