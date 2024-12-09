package main

import (
	"echo_framework/config"
	internal "echo_framework/internal/middleware"
	handler "echo_framework/internal/userhandler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	u := e.Group("/users")
	u.Use(internal.JwtMid)
	u.GET("/", handler.GetAllUsers)
	//u.GET("/:id", handler.GetUserById)

	//? BISA PAKE INI JUGA
	//e.POST("/login", handler.LoginUser)
	//e.POST("/register", handler.Register)

	g := e.Group("/auth")
	g.POST("/login", handler.LoginUser)
	g.POST("/register", handler.Register)

	e.Logger.Fatal(e.Start(":8080"))

}
