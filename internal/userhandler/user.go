package handler

import (
	"echo_framework/config"
	internal "echo_framework/internal/userdto"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"net/http"
)

func Register(c echo.Context) error {
	var req internal.RegisterUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid generate password"})
	}
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var userID int
	err = config.Pool.QueryRow(context.Background(), query, req.Name, req.Email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Register failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Register success"})
}
