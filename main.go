package main

import (
	"github.com/GreatLaboratory/simplify-petition-crawling-server/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.GET("/wordCloud/:category", controller.WordCloudController)
	e.GET("/topFivePosts/:category", controller.TopFiveController)
	e.GET("/summary/:petitionID", controller.SummaryController)

	e.Logger.Fatal(e.Start(":1323"))
}
