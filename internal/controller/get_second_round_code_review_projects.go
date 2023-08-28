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
		g.POST("/GetSecondRoundCodeReviewProjects", c.Handle_GetSecondRoundCodeReviewProjects).
			AddParamBody(requests.Variables_GetSecondRoundCodeReviewProjects{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_GetSecondRoundCodeReviewProjects{}, nil)
	})
}

func (a *AdapterController) Handle_GetSecondRoundCodeReviewProjects(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_GetSecondRoundCodeReviewProjects `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().GetSecondRoundCodeReviewProjects(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}