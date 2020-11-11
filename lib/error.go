package lib

import (
	"fmt"
	"net/http"
)

// CheckErr -> 에러의 내용 체크
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

// CheckCode -> 에러의 코드 체크
func CheckCode(res *http.Response) {
	if res.StatusCode != 200 {
		fmt.Println("Request Fail With StatusCode : ", res.StatusCode)
	}
}
