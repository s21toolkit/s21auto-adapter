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
		g.POST("/GetPenaltyReasons", c.Handle_GetPenaltyReasons).
			SetOperationId("GetPenaltyReasons").
			AddParamBody(requests.GetPenaltyReasons_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetPenaltyReasons_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetPenaltyReasons(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetPenaltyReasons_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetPenaltyReasons(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}