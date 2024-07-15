package model

type Page struct {
	PageIdentifier
	PageData
}

type PageIdentifier struct {
	Key string `json:"key"`
}

type PageData struct {
	MenuName  string `json:"menuName"`
	ArticleId string `json:"articledId"`
}
