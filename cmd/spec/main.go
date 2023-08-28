package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"github.com/s21toolkit/s21adapter/internal/server"
)

func main() {
	root := server.InitServer()

	req := httptest.NewRequest(echo.GET, "/spec/swagger.json", nil)

	rec := httptest.NewRecorder()

	ctx := root.Echo().NewContext(req, rec)

	spec, err := root.(*echoswagger.Root).GetSpec(ctx, "spec/")

	if err != nil {
		panic(err)
	}

	rawSpec, err := json.MarshalIndent(spec, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(rawSpec))
}
