package handler

import (
	"echo_framework/config"
	internal "echo_framework/internal/userdto"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
)

// Get all users
func GetAllUsers(c echo.Context) error {
	query := "SELECT id, name, email FROM users"
	rows, err := config.Pool.Query(context.Background(), query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}
	defer rows.Close()

	var users []internal.Users
	for rows.Next() {
		var user internal.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}
	return c.JSON(http.StatusOK, users)
}
