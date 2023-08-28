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
		g.POST("/BonusesGetBadgesWithFakePublicProfile", c.Handle_BonusesGetBadgesWithFakePublicProfile).
			AddParamBody(requests.Variables_BonusesGetBadgesWithFakePublicProfile{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_BonusesGetBadgesWithFakePublicProfile{}, nil)
	})
}

func (a *AdapterController) Handle_BonusesGetBadgesWithFakePublicProfile(c echo.Context) (err error) {
	var data struct {
		Variables requests.Variables_BonusesGetBadgesWithFakePublicProfile `json:"variables"`
	}

	err = json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return
	}

	res, err := a.client.R().BonusesGetBadgesWithFakePublicProfile(data.Variables)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}
