package internal

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var JwtMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:  []byte("12345"),
	Claims:      &jwt.StandardClaims{},
	TokenLookup: "header:Authorization",
	AuthScheme:  "Bearer",
	ErrorHandlerWithContext: func(c echo.Context, err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authentication")
	},
})
