package controller

import "github.com/pangpanglabs/echoswagger/v2"

var methodRegistry []func(echoswagger.ApiGroup) = make([]func(echoswagger.ApiGroup), 0)

func registerMethod(callback func(echoswagger.ApiGroup)) {
	methodRegistry = append(methodRegistry, callback)
}

type AdapterController struct{}

func (AdapterController) Init(g echoswagger.ApiGroup) {
	g.SetDescription("s21adapter API")

	for _, methodInitializer := range methodRegistry {
		methodInitializer(g)
	}
}
