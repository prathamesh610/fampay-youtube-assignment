package dto

import "time"

type VideoResponse struct {
	VideoTitle string
	//ChannelName  string
	Description  string
	ThumbnailUrl string
	PublishedAt  time.Time
}
