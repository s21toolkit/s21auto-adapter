package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/s21toolkit/s21client/requests"
)

func init() {
	registerMethod(func(g echoswagger.ApiGroup, c *AdapterController) {
		g.POST("/RemoveP2P", c.Handle_RemoveP2P).
			SetOperationId("RemoveP2P").
			AddParamBody(requests.Variables_RemoveP2P{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_RemoveP2P{}, nil)
	})
}

func (a *AdapterController) Handle_RemoveP2P(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_RemoveP2P `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().RemoveP2P(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}