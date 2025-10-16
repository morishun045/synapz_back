// main

package main

import (
	
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var jwtSecret = []byte()

func main() {
	e := echo.New()
	e.Use()
}
