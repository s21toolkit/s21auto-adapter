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
		g.POST("/CalendarGetEvents", c.Handle_CalendarGetEvents).
			AddParamBody(requests.Variables_CalendarGetEvents{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_CalendarGetEvents{}, nil)
	})
}

func (a *AdapterController) Handle_CalendarGetEvents(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_CalendarGetEvents `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().CalendarGetEvents(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}