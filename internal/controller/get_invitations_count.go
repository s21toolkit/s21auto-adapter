package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/s21toolkit/s21client/requests"
)

func init() {
	registerMethod(func(g echoswagger.ApiGroup) {
		g.POST("/GetInvitationsCount", Handle_GetInvitationsCount).
			AddParamBody(requests.Variables_GetInvitationsCount{}, "variables", "Request variables", true).
			AddResponse(http.StatusOK, "Success", requests.Data_GetInvitationsCount{}, nil)
	})
}

func Handle_GetInvitationsCount(c echo.Context) error {
	return nil
}
