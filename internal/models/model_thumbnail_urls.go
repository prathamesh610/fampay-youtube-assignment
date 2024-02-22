package models

type ThumbnailUrl struct {
	Id           int64 `gorm:"primaryKey"`
	VideoId      string
	ThumbnailUrl string
}
