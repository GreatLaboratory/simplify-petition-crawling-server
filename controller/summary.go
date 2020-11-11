package controller

import (
	"net/http"

	"github.com/GreatLaboratory/simplify-petition-crawling-server/service"
	"github.com/labstack/echo"
)

// SummaryController -> 요약 데이터 크롤링
func SummaryController(c echo.Context) error {
	petitionID := c.Param("petitionID")
	petition := service.SummaryCrawl(petitionID)
	return c.JSON(http.StatusOK, petition)
}
