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
		g.POST("/CompetitionCoalitionGetMyCoalitionMembers", c.Handle_CompetitionCoalitionGetMyCoalitionMembers).
			SetOperationId("CompetitionCoalitionGetMyCoalitionMembers").
			AddParamBody(requests.Variables_CompetitionCoalitionGetMyCoalitionMembers{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_CompetitionCoalitionGetMyCoalitionMembers{}, nil)
	})
}

func (a *AdapterController) Handle_CompetitionCoalitionGetMyCoalitionMembers(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_CompetitionCoalitionGetMyCoalitionMembers `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().CompetitionCoalitionGetMyCoalitionMembers(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}