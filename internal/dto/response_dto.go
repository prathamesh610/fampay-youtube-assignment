package dto

type ResponseDTO struct {
	VideosCount    int64           `json:"total-videos"`
	PageNumber     int             `json:"page-number"`
	LastPageNumber int64           `json:"last-page-number"`
	Count          int             `json:"count"`
	Videos         []VideoResponse `json:"videos"`
}
