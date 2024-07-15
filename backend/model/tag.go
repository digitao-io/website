package model

type Tag struct {
	TagIdentifier
	TagData
}

type TagIdentifier struct {
	Key string `json:"key"`
}

type TagData struct {
	Name string `json:"name"`
}


