package model

// FormData -> https://www1.president.go.kr/api/petitions/list||rank 으로 들어가는 request form data
type FormData struct {
	C     string `json:"c"`
	Only  string `json:"only"`
	Page  string `json:"page"`
	Order string `json:"order"`
}

// Responsedata -> https://www1.president.go.kr/api/petitions/list||rank의 응답 json 골격
type Responsedata struct {
	Status string `json:"status"`
	Total  string `json:"total"`
	Page   int    `json:"page"`
	Item   []Item `json:"item"`
}

// Item -> 청원 게시물 정보
type Item struct {
	ID        string `json:"id"`
	PagingID  int    `json:"paging_id"`
	Title     string `json:"title"`
	Agreement string `json:"agreement"`
	Category  string `json:"category"`
	Created   string `json:"created"`
	Finished  string `json:"finished"`
	Provider  string `json:"provider"`
}

// PetitionContent -> 청원 내용
type PetitionContent struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
