package sitemodel

type Resource struct {
	ResourceIdentifier `bson:"inline"`
	Title              *string                 `json:"title" bson:"title,omitempty"`
	Description        *string                 `json:"description" bson:"description,omitempty"`
	Type               *string                 `json:"type" bson:"type,omitempty"`
	Details            *map[string]interface{} `json:"details" bson:"details,omitempty"`
}

type ResourceIdentifier struct {
	Key *string `json:"key" form:"key" bson:"key,omitempty"`
}

type ResourceSearchParams struct {
	Type *string `form:"type"`
}
