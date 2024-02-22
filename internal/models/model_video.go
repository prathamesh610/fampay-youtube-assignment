package models

import "time"

type Video struct {
	Id             int64     `gorm:"primaryKey" json:"id"`
	SearchQuery    string    `json:"searchQuery"`
	VideoTitle     string    `json:"videoTitle"`
	Description    string    `json:"description"`
	PublishingDate time.Time `json:"publishingDate"`
	VideoId        string    `json:"videoId"`
}
