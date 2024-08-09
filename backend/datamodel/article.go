package datamodel

import "time"

type Article struct {
	ArticleIdentifier `bson:"inline"`
	ArticleData       `bson:"inline"`
	ArticleExtra      `bson:"inline"`
}

type ArticleIdentifier struct {
	Id *string `json:"id" form:"id" bson:"_id,omitempty" binding:"required,uuid"`
}

type ArticleData struct {
	Title        *string   `json:"title" bson:"title,omitempty" binding:"required,min=1"`
	Type         *string   `json:"type" bson:"type,omitempty" binding:"required,min=1"`
	TagKeys      *[]string `json:"tagKeys" bson:"tagKeys,omitempty" binding:"required,unique,dive,min=1"`
	Summary      *string   `json:"summary" bson:"summary,omitempty" binding:"required,min=1"`
	Content      *string   `json:"content" bson:"content,omitempty" binding:"required,min=1"`
	ThumbnailUrl *string   `json:"thumbnailUrl" bson:"thumbnailUrl,omitempty" binding:"required,http_url"`
}

type ArticleExtra struct {
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}

type ArticleSearchParams struct {
	Query *string   `form:"q" binding:"omitnil,min=1"`
	Type  *string   `form:"type" binding:"omitnil,min=1"`
	Tags  *[]string `form:"tag" binding:"omitnil,unique,dive,min=1"`

	Sort  *string `form:"sort" binding:"required"`
	Order *string `form:"order" binding:"required,oneof=ASC DESC"`

	Take *int64 `form:"take" binding:"required,gt=0"`
	Skip *int64 `form:"skip" binding:"required,gte=0"`
}
