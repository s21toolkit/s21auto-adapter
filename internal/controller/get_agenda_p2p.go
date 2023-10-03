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
		g.POST("/GetAgendaP2P", c.Handle_GetAgendaP2P).
			SetOperationId("GetAgendaP2P").
			AddParamBody(requests.GetAgendaP2P_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetAgendaP2P_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetAgendaP2P(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetAgendaP2P_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetAgendaP2P(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}