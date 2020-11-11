package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/GreatLaboratory/simplify-petition-crawling-server/lib"
	"github.com/GreatLaboratory/simplify-petition-crawling-server/model"
)

// WordCloudCrawl -> 워드클라우드를 위한 분야별 제목 크롤링
func WordCloudCrawl(category string) []string {
	ch := make(chan []string)

	apiURL := "https://www1.president.go.kr"
	resource := "/api/petitions/list"
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

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	baseURL := u.String() // "https://www1.president.go.kr/api/petitions/list"
	req, err := http.NewRequest(http.MethodPost, baseURL, strings.NewReader(reqData.Encode()))
	lib.CheckErr(err)
	// req.Header.Add("Referer", "https://www1.president.go.kr/petitions/?only=1&page=1&order=1&c="+category)
	// req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
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

	totalNum, err := strconv.Atoi(data.Total)
	pageNum := (totalNum / 7) + 1
	fmt.Println("카테고리", category, "==>", pageNum)
	lib.CheckErr(err)

	var extractedTitleList []string
	for i := 1; i <= 10; i++ { // pageNum로 하면 ip가 막혀서... 막히지 않는 최소 페이지수인 10으로 제한
		go getPageData(baseURL, i, category, ch)
	}
	for i := 0; i < 10; i++ { // pageNum로 하면 ip가 막혀서... 막히지 않는 최소 페이지수인 10으로 제한
		titleList := <-ch
		// fmt.Println(i+1, "==>", titleList)
		extractedTitleList = append(extractedTitleList, titleList...)
	}

	return extractedTitleList
}

func getPageData(baseURL string, page int, category string, ch chan<- []string) {
	reqData := url.Values{}
	formData := model.FormData{
		C:     category,
		Only:  "2",
		Page:  strconv.Itoa(page),
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

	var result []string
	for _, item := range data.Item {
		result = append(result, item.Title)
	}
	ch <- result
}
