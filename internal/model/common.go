package model

type MetadataRequest struct {
	Limit uint8 `json:"limit"`
	Page  uint8 `json:"page"`
}

type MetadataResponse struct {
	Limit uint8  `json:"limit"`
	Page  uint8  `json:"page"`
	Count uint32 `json:"count"`
}
