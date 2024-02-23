package models

import "time"

type Video struct {
	VideoId        string     `bson:"_id,omitempty" json:"id"`
	VideoTitle     string     `bson:"videoTitle" json:"videoTitle"`
	Description    string     `bson:"description" json:"description"`
	PublishingDate time.Time  `bson:"publishingDate" json:"publishingDate"`
	Thumbnails     Thumbnails `bson:"thumbnails" json:"thumbnails"`
}
