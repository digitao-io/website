package sitemodel

type Page struct {
	PageIdentifier `bson:"inline"`
	Title          *string                 `json:"title" bson:"title,omitempty"`
	Description    *string                 `json:"description" bson:"description,omitempty"`
	UrlPattern     *string                 `json:"urlPattern" bson:"urlPattern,omitempty"`
	Details        *map[string]interface{} `json:"details" bson:"details,omitempty"`
}

type PageIdentifier struct {
	Key *string `json:"key" form:"key" bson:"key,omitempty"`
}
