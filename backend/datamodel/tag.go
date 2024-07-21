package datamodel

type Tag struct {
	TagIdentifier `bson:"inline"`
	Name          *string `json:"name" bson:"name,omitempty"`
}

type TagIdentifier struct {
	Key *string `json:"key" form:"key" bson:"key,omitempty"`
}
