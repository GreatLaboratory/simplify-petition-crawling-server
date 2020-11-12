package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/GreatLaboratory/simplify-petition-crawling-server/lib"
	"github.com/GreatLaboratory/simplify-petition-crawling-server/model"
)

// TopFiveCrawl -> 분야별 top5 청원글 목록 크롤링
func TopFiveCrawl(category string) []model.Item {
	apiURL := "https://www1.president.go.kr"
	resource := "/api/petitions/rank"
	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	baseURL := u.String() // "https://www1.president.go.kr/api/petitions/rank"

	reqData := url.Values{}
	formData := model.FormData{
		C:     category,
		Only:  "2",
		Page:  "1",
		Order: "1",
	}
	reqData.Set("c", formData.C)
	reqData.Set("only", formData.Only)
	reqData.Set("page", formData.Page)
	reqData.Set("order", formData.Order)

	req, err := http.NewRequest(http.MethodPost, baseURL, strings.NewReader(reqData.Encode()))
	lib.CheckErr(err)

	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqData.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Client객체에서 Request 실행
	client := &http.Client{}
	res, err := client.Do(req)
	lib.CheckErr(err)
	defer res.Body.Close()

	var data model.Responsedata
	resBody, err := ioutil.ReadAll(res.Body)
	lib.CheckErr(err)
	json.Unmarshal(resBody, &data)

	var result []model.Item
	for _, item := range data.Item {
		result = append(result, item)
	}

	return result
}
