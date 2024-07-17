package model

type Content struct {
	ContentIdentifier
	ContentData
	ContentExtra
}

type ContentIdentifier struct {
	Id *string `json:"id" form:"id"`
}

type ContentData struct {
	Type    *string   `json:"type"`
	Title   *string   `json:"title"`
	TagKeys *[]string `json:"tagKeys"`
	Summary *string   `json:"summary"`
	Content *string   `json:"content"`
}

type ContentExtra struct {
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
}

type ContentSearchParams struct {
	Query *string   `form:"q"`
	Type  *string   `form:"type"`
	Tags  *[]string `form:"tag"`

	Sort  *string `form:"sort"`
	Order *string `form:"order"`

	Skip *uint `form:"skip"`
	Take *uint `form:"take"`
}
