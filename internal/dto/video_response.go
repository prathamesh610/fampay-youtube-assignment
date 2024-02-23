package dto

import (
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"time"
)

type VideoResponse struct {
	VideoTitle   string
	Description  string
	ThumbnailUrl models.Thumbnails
	PublishedAt  time.Time
}
