package controller

import (
	"os"

	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/s21toolkit/s21client"
)

var methodRegistry []func(echoswagger.ApiGroup, *AdapterController) = make([]func(echoswagger.ApiGroup, *AdapterController), 0)

func registerMethod(callback func(echoswagger.ApiGroup, *AdapterController)) {
	methodRegistry = append(methodRegistry, callback)
}

type AdapterController struct {
	client *s21client.Client
}

func New() AdapterController {
	client := s21client.New(s21client.DefaultAuth(os.Getenv("S21_USERNAME"), os.Getenv("S21_PASSWORD")))

	return AdapterController{
		client: client,
	}
}

func (c *AdapterController) Init(g echoswagger.ApiGroup) {
	g.SetDescription("s21adapter API")

	for _, methodInitializer := range methodRegistry {
		methodInitializer(g, c)
	}
}
