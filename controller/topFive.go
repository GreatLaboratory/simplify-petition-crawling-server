package controller

import (
	"net/http"

	"github.com/GreatLaboratory/simplify-petition-crawling-server/service"
	"github.com/labstack/echo"
)

// TopFiveController -> top5 청원글 데이터 크롤링
func TopFiveController(c echo.Context) error {
	category := c.Param("category")
	topFiveList := service.TopFiveCrawl(category)
	return c.JSON(http.StatusOK, topFiveList)
}
