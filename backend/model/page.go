package model

type Page struct {
	PageIdentifier
	PageData
}

type PageIdentifier struct {
	Key *string `json:"key" form:"key" validation:"required"`
}

type PageData struct {
	MenuName  *string `json:"menuName"`
	ContentId *string `json:"contentId"`
}
