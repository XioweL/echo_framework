package handler

import (
	"echo_framework/config"
	internal "echo_framework/internal/userdto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

var jwtSecret = []byte("12345")

func LoginUser(c echo.Context) error {
	var req internal.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	var user internal.User
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	err := config.Pool.QueryRow(context.Background(), query, req.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid email or password"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Generate Token"})
	}
	return c.JSON(http.StatusOK, internal.LoginResponse{Token: tokenString})
}
