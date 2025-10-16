// jwt生成用

package lib

import (
	
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

)

var jwtSecret = []byte("YOUR_SECRET_KEY");

type JwtCustomClaims struct {
	UserID string `json: "user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT (UserID string) (string, error) {
	claims := &JwtCustomClaims {
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GetClaimsFromContext(c echo.Context) (*JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token);
	if user == nil {
		return nil, errors.New("Invalid Token")
	}
	claims, ok := user.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, errors.New("Invalid Claims")
	}
	return claims, nil
}
