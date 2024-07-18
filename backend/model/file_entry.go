package model

type FileEntry struct {
	FileEntryIdentifier
	FileEntryData
	FileEntryExtra
}

type FileEntryIdentifier struct {
	Id *string `json:"id" form:"id" validation:"required"`
}

type FileEntryData struct {
	Title       *string `json:"title"`
	MimeType    *string `json:"mimeType"`
	SizeInBytes *int64  `json:"sizeInBytes"`
}

type FileEntryExtra struct {
	CreatedAt *string `json:"createdAt"`
}

type FileEntrySearchParams struct {
	Query    *string `form:"q"`
	MimeType *string `form:"mimeType"`

	Sort  *string `form:"sort"`
	Order *string `form:"order"`

	Take *uint `form:"take"`
	Skip *uint `form:"skip"`
}
