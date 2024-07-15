package model

type Article struct {
	ArticleIdentifier
	ArticleData
}

type ArticleIdentifier struct {
	Id string `json:"id"`
}

type ArticleData struct {
	Type      string   `json:"type"`
	Title     string   `json:"title"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	TagKeys   []string `json:"tagKeys"`
	Summary   string   `json:"summary"`
	Content   string   `json:"content"`
}
