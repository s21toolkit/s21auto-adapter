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
		g.POST("/GetCampusCurrentUser", c.Handle_GetCampusCurrentUser).
			AddParamBody(requests.Variables_GetCampusCurrentUser{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_GetCampusCurrentUser{}, nil)
	})
}

func (a *AdapterController) Handle_GetCampusCurrentUser(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_GetCampusCurrentUser `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetCampusCurrentUser(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}