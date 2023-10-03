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
		g.POST("/GetPenaltyList", c.Handle_GetPenaltyList).
			SetOperationId("GetPenaltyList").
			AddParamBody(requests.GetPenaltyList_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetPenaltyList_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetPenaltyList(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetPenaltyList_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetPenaltyList(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}