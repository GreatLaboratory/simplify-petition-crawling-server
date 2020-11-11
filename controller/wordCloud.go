package controller

import (
	"net/http"

	"github.com/GreatLaboratory/simplify-petition-crawling-server/service"
	"github.com/labstack/echo"
)

// WordCloudController -> 워드클라우드 데이터 크롤링
func WordCloudController(c echo.Context) error {
	category := c.Param("category")
	titleList := service.WordCloudCrawl(category)
	var result string
	for _, str := range titleList {
		result += str + " "
	}
	return c.String(http.StatusOK, result)
}
