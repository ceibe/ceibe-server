package main

import (
	"github.com/ceibe/ceibe-server/api/rest/controllers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Debug = true

	controllers.SetupUserController(e)

	e.Logger.Fatal(e.Start(":1323"))
}
