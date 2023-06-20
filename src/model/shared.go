package model

type FilterData struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int32  `json:"limit"`
}

type ConfigKey string

type UploadedData struct {
	Name        string `json:"name"`
	Uploader    string `json:"uploader"`
	URL         string `json:"url"`
	PreviewURL  string `json:"previewURL"`
	Width       string `json:"width"`
	Height      string `json:"height"`
	Size        string `json:"size"`
	ContentType string `json:"contentType"`
}

type UploadedDataRelation struct {
	File  string `json:"file"`
	Table string `json:"table"`
	Order int    `json:"order"`
	Value string `json:"value"`
}
