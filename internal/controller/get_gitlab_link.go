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
		g.POST("/GetGitlabLink", c.Handle_GetGitlabLink).
			AddParamBody(requests.Variables_GetGitlabLink{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_GetGitlabLink{}, nil)
	})
}

func (a *AdapterController) Handle_GetGitlabLink(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_GetGitlabLink `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetGitlabLink(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}