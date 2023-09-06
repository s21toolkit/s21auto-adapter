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
		g.POST("/SaveFilledChecklist", c.Handle_SaveFilledChecklist).
			SetOperationId("SaveFilledChecklist").
			AddParamBody(requests.Variables_SaveFilledChecklist{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_SaveFilledChecklist{}, nil)
	})
}

func (a *AdapterController) Handle_SaveFilledChecklist(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_SaveFilledChecklist `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().SaveFilledChecklist(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}