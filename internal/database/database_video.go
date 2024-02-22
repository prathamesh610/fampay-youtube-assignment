package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dberrors"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"gorm.io/gorm"
)

func (c Client) SaveVideosToDB(ctx context.Context, videos *[]models.Video) (*[]models.Video, error) {
	result := c.DB.WithContext(ctx).Save(&videos)

	fmt.Printf("Saved VIDEO to db with: %v", videos)

	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func (c Client) GetVideo(ctx context.Context, video models.Video) (*models.Video, error) {
	video1 := &models.Video{}
	result := c.DB.WithContext(ctx).Where(&models.Video{
		SearchQuery: video.SearchQuery,
	}).First(&video1)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &video, nil
		}
		return nil, result.Error
	}

	return video1, nil

}

func (c Client) SearchVideo(ctx context.Context, searchQuery string) (*[]models.Video, error) {
	videos := &[]models.Video{}
	// TODO: sort by published at
	result := c.DB.WithContext(ctx).Where(&models.Video{SearchQuery: searchQuery}).Order("publishing_date desc").Find(&videos)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "Video", Query: searchQuery}
		}
		return nil, result.Error
	}

	return videos, nil
}
