package datamodel

import "time"

type File struct {
	FileIdentifier `bson:"inline"`
	FileData       `bson:"inline"`
	FileExtra      `bson:"inline"`
}

type FileIdentifier struct {
	Id *string `json:"id" form:"id" bson:"_id,omitempty"`
}

type FileData struct {
	Title       *string `json:"title" bson:"title,omitempty"`
	MimeType    *string `json:"mimeType" bson:"mimeType,omitempty"`
	SizeInBytes *int64  `json:"sizeInBytes" bson:"sizeInBytes,omitempty"`
}

type FileExtra struct {
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}

type FileSearchParams struct {
	Query    *string `form:"q"`
	MimeType *string `form:"mimeType"`

	Sort  *string `form:"sort"`
	Order *string `form:"order"`

	Take *int64 `form:"take"`
	Skip *int64 `form:"skip"`
}
