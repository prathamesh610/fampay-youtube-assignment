package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dberrors"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"gorm.io/gorm"
)

func (c Client) SaveThumbnailsToDB(ctx context.Context, thumbnails *[]models.ThumbnailUrl) (*[]models.ThumbnailUrl, error) {
	result := c.DB.WithContext(ctx).Create(&thumbnails)

	fmt.Printf("Saved THUMBNAIL to db with: %v", thumbnails)
	if result.Error != nil {
		return nil, result.Error
	}
	return thumbnails, nil
}

func (c Client) GetThumbnail(ctx context.Context, thumbnail models.ThumbnailUrl) (*models.ThumbnailUrl, error) {
	thumbnail1 := &models.ThumbnailUrl{}
	result := c.DB.WithContext(ctx).Where(&models.ThumbnailUrl{
		VideoId: thumbnail.VideoId,
	}).First(&thumbnail1)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &thumbnail, nil
		}
		return nil, result.Error
	}

	return thumbnail1, nil
}

func (c Client) SearchThumbnails(ctx context.Context, videoIds []string) (*[]models.ThumbnailUrl, error) {
	thumbnails := make([]models.ThumbnailUrl, 0)
	result := c.DB.WithContext(ctx).Where("video_id IN ? ", videoIds).Find(&thumbnails)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "Thumbnail", Query: videoIds[0]}
		}
		return nil, result.Error
	}

	return &thumbnails, nil
}
