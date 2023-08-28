package server

import (
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/s21toolkit/s21adapter/internal/controller"
)

func InitServer() echoswagger.ApiRoot {
	e := echo.New()

	se := echoswagger.New(e, "spec/", nil)

	controller.AdapterController{}.Init(se.Group("s21adapter", "/adapter"))

	return se
}

func Serve(root echoswagger.ApiRoot, address string) {
	e := root.Echo()

	e.Logger.Fatal(e.Start(address))
}
