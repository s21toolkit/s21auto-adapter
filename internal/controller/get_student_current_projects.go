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
		g.POST("/GetStudentCurrentProjects", c.Handle_GetStudentCurrentProjects).
			SetOperationId("GetStudentCurrentProjects").
			AddParamBody(requests.GetStudentCurrentProjects_Variables{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.GetStudentCurrentProjects_Data{}, nil)
	})
}

func (a *AdapterController) Handle_GetStudentCurrentProjects(c echo.Context) (err error) {
	var data struct {
		Variables requests.GetStudentCurrentProjects_Variables `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetStudentCurrentProjects(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}