package sitemodel

type Page struct {
	PageIdentifier `bson:"inline"`
	Title          *string                 `json:"title" bson:"title,omitempty" binding:"required,min=1"`
	Description    *string                 `json:"description" bson:"description,omitempty" binding:"required,min=1"`
	UrlPattern     *string                 `json:"urlPattern" bson:"urlPattern,omitempty" binding:"required,uri,startswith=/"`
	Details        *map[string]interface{} `json:"details" bson:"details,omitempty" binding:"required"`
}

type PageIdentifier struct {
	Key *string `json:"key" form:"key" bson:"key,omitempty" binding:"required,key"`
}
