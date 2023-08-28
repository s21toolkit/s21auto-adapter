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
		g.POST("/CompetitionCoalitionGetUserTournament", c.Handle_CompetitionCoalitionGetUserTournament).
			SetOperationId("CompetitionCoalitionGetUserTournament").
			AddParamBody(requests.Variables_CompetitionCoalitionGetUserTournament{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_CompetitionCoalitionGetUserTournament{}, nil)
	})
}

func (a *AdapterController) Handle_CompetitionCoalitionGetUserTournament(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_CompetitionCoalitionGetUserTournament `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().CompetitionCoalitionGetUserTournament(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}