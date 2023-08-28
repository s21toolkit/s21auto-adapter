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
		g.POST("/GetStudentStageGroupS21", c.Handle_GetStudentStageGroupS21).
			AddParamBody(requests.Variables_GetStudentStageGroupS21{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_GetStudentStageGroupS21{}, nil)
	})
}

func (a *AdapterController) Handle_GetStudentStageGroupS21(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_GetStudentStageGroupS21 `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetStudentStageGroupS21(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}