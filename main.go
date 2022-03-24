package main

import (
	"echo-mongo-api/configs" //add this
	"echo-mongo-api/routers"
	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()

	// run database
	configs.ConnectDB()

	// routes
	routers.UserRoute(e)

	e.Logger.Fatal(e.Start(":6000"))
}
