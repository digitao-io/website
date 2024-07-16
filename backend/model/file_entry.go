package model

type FileEntry struct {
	FileEntryIdentifier
	FileEntryData
}

type FileEntryIdentifier struct {
	Id *string `json:"id" form:"id" validation:"required"`
}

type FileEntryData struct {
	Title       *string `json:"title"`
	MimeType    *string `json:"mimeType"`
	SizeInBytes *uint   `json:"sizeInBytes"`
}
