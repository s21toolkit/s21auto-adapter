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
		g.POST("/GetUserNotifications", c.Handle_GetUserNotifications).
			SetOperationId("GetUserNotifications").
			AddParamBody(requests.GetUserNotifications_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetUserNotifications_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetUserNotifications(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetUserNotifications_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetUserNotifications(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}