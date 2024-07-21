package datamodel

import "time"

type Article struct {
	ArticleIdentifier `bson:"inline"`
	ArticleData       `bson:"inline"`
	ArticleExtra      `bson:"inline"`
}

type ArticleIdentifier struct {
	Id *string `json:"id" form:"id" bson:"_id,omitempty"`
}

type ArticleData struct {
	Title   *string   `json:"title" bson:"title,omitempty"`
	TagKeys *[]string `json:"tagKeys" bson:"tagKeys,omitempty"`
	Summary *string   `json:"summary" bson:"summary,omitempty"`
	Content *string   `json:"content" bson:"content,omitempty"`
}

type ArticleExtra struct {
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}

type ArticleSearchParams struct {
	Query *string   `form:"q"`
	Tags  *[]string `form:"tag"`

	Sort  *string `form:"sort"`
	Order *string `form:"order"`

	Skip *int64 `form:"skip"`
	Take *int64 `form:"take"`
}
