package model

type Content struct {
	ContentIdentifier
	ContentData
}

type ContentIdentifier struct {
	Id *string `json:"id" form:"id"`
}

type ContentData struct {
	Type      *string   `json:"type"`
	Title     *string   `json:"title"`
	CreatedAt *string   `json:"createdAt"`
	UpdatedAt *string   `json:"updatedAt"`
	TagKeys   *[]string `json:"tagKeys"`
	Summary   *string   `json:"summary"`
	Content   *string   `json:"content"`
}
