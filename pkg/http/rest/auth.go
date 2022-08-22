package rest

import (
	"contacts-api/pkg/auth"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func authMiddleware(a auth.Service) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if a.ValidUser(username, password) {
			if a.Permission(username, c.Param("id")) {
				return true, nil
			}
		}
		return false, fmt.Errorf("user does not exist or does not have permission")
	})
}
