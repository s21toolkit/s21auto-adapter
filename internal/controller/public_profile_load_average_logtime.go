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
		g.POST("/PublicProfileLoadAverageLogtime", c.Handle_PublicProfileLoadAverageLogtime).
			AddParamBody(requests.Variables_PublicProfileLoadAverageLogtime{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_PublicProfileLoadAverageLogtime{}, nil)
	})
}

func (a *AdapterController) Handle_PublicProfileLoadAverageLogtime(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_PublicProfileLoadAverageLogtime `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().PublicProfileLoadAverageLogtime(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}