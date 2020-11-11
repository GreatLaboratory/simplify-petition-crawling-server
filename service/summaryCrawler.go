package service

import (
	"net/http"

	"github.com/GreatLaboratory/simplify-petition-crawling-server/lib"
	"github.com/GreatLaboratory/simplify-petition-crawling-server/model"
	"github.com/PuerkitoBio/goquery"
)

// SummaryCrawl -> 메인 크롤러
func SummaryCrawl(petitionID string) model.PetitionContent {
	var result model.PetitionContent

	pageURL := "https://www1.president.go.kr/petitions/" + petitionID
	res, err := http.Get(pageURL)
	lib.CheckErr(err)
	lib.CheckCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	lib.CheckErr(err)
	question := doc.Find(".View_write").Text()
	answer := doc.Find(".pr_tk25").Text()
	result = model.PetitionContent{
		Question: lib.CleanStr(question),
		Answer:   lib.CleanStr(answer)[13:],
	}
	return result
}
