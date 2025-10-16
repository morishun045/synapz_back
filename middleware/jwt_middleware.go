// jwtのMiddleware

package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
)

var jwtSecret = []byte("YOUR_SECRET_KEY")

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtSecret,
		ContextKey: "user",
		TokenLookup: "header: Authorization",
	})
}
