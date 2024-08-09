package datamodel

import "time"

type File struct {
	FileIdentifier `bson:"inline"`
	FileData       `bson:"inline"`
	FileExtra      `bson:"inline"`
}

type FileIdentifier struct {
	Id *string `json:"id" form:"id" bson:"_id,omitempty" binding:"required,uuid"`
}

type FileData struct {
	Title       *string `json:"title" bson:"title,omitempty" binding:"required,min=1"`
	MimeType    *string `json:"mimeType" bson:"mimeType,omitempty" binding:"required,min=1"`
	SizeInBytes *int64  `json:"sizeInBytes" bson:"sizeInBytes,omitempty" binding:"required,min=0"`
}

type FileExtra struct {
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}

type FileSearchParams struct {
	Query    *string `form:"q" binding:"omitnil,min=1"`
	MimeType *string `form:"mimeType" binding:"omitnil,min=1"`

	Sort  *string `form:"sort" binding:"required"`
	Order *string `form:"order" binding:"required,oneof=ASC DESC"`

	Take *int64 `form:"take" binding:"required,gt=0"`
	Skip *int64 `form:"skip" binding:"required,gte=0"`
}
