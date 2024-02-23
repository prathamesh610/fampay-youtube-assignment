package dto

import (
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"time"
)

type VideoResponse struct {
	VideoTitle   string            `json:"video-title"`
	Description  string            `json:"description"`
	ThumbnailUrl models.Thumbnails `json:"thumbnail-url"`
	PublishedAt  time.Time         `json:"published-at"`
}
