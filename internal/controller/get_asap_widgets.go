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
		g.POST("/GetAsapWidgets", c.Handle_GetAsapWidgets).
			SetOperationId("GetAsapWidgets").
			AddParamBody(requests.GetAsapWidgets_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetAsapWidgets_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetAsapWidgets(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetAsapWidgets_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetAsapWidgets(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}