package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Yamako76/backend-sample-golang/infra"
	"github.com/Yamako76/backend-sample-golang/infra/mysql"
	"github.com/Yamako76/backend-sample-golang/usecase"
)

func main() {
	e := echo.New()

	db, err := infra.NewDB()
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました: %v", err)
	}

	userRepository := mysql.NewUserRepository(db)

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.GET("/users/:id", func(c echo.Context) error {
		return usecase.GetUser(c, userRepository)
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatalf("サーバーにエラーが発生しました: %v", err)
	}
}
