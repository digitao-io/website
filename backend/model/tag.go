package model

type Tag struct {
	TagIdentifier
	TagData
}

type TagIdentifier struct {
	Key *string `json:"key" form:"key" validation:"required"`
}

type TagData struct {
	Name *string `json:"name" validation:"required"`
}
