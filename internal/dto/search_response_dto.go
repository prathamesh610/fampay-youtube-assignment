package dto

type SearchResponseDTO struct {
	SearchQuery string          `json:"search-query"`
	Videos      []VideoResponse `json:"videos"`
}
