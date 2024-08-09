package sitemodel

type Resource struct {
	ResourceIdentifier `bson:"inline"`
	Title              *string                 `json:"title" bson:"title,omitempty" binding:"required,min=1"`
	Description        *string                 `json:"description" bson:"description,omitempty" binding:"required,min=1"`
	Type               *string                 `json:"type" bson:"type,omitempty" binding:"required,min=1"`
	Details            *map[string]interface{} `json:"details" bson:"details,omitempty" binding:"required"`
}

type ResourceIdentifier struct {
	Key *string `json:"key" form:"key" bson:"key,omitempty" binding:"required,key"`
}

type ResourceSearchParams struct {
	Type *string `form:"type" binding:"omitnil,min=1"`
}
