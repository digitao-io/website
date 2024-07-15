package model

type FileEntry struct {
	FileEntryIdentifier
	FileEntryData
}

type FileEntryIdentifier struct {
	Id string `json:"id"`
}

type FileEntryData struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	MimeType    string `json:"mimeType"`
	SizeInBytes int    `json:"sizeInBytes"`
}
