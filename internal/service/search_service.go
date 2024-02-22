package service

import (
	"context"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dto"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
)

func SearchUsingQueryInDB(ctx context.Context, client database.DatabaseClient, searchQuery string) (*dto.ResponseDTO, error) {

	videos, err := client.SearchVideo(ctx, searchQuery)
	if err != nil {
		return nil, err
	}

	videoIds := getVideoIds(*videos)

	thumbnails, err := client.SearchThumbnails(ctx, videoIds)

	videoResponse := assignThumbnailsToVideos(*videos, *thumbnails)

	responseDto := &dto.ResponseDTO{
		SearchQuery: searchQuery,
		Videos:      videoResponse,
	}
	return responseDto, nil
}

func assignThumbnailsToVideos(videos []models.Video, thumbnails []models.ThumbnailUrl) []dto.VideoResponse {
	thumbnailMap := make(map[string]string)
	for _, thumbnail := range thumbnails {
		thumbnailMap[thumbnail.VideoId] = thumbnail.ThumbnailUrl
	}

	videoSlice := make([]dto.VideoResponse, 0)

	for _, video := range videos {
		thumbnailUrl := thumbnailMap[video.VideoId]
		videoModel := &dto.VideoResponse{
			VideoTitle:   video.VideoTitle,
			Description:  video.Description,
			ThumbnailUrl: thumbnailUrl,
			PublishedAt:  video.PublishingDate,
		}
		videoSlice = append(videoSlice, *videoModel)
	}

	return videoSlice
}

func getVideoIds(videos []models.Video) []string {
	result := make([]string, 0)
	for _, video := range videos {
		result = append(result, video.VideoId)
	}
	return result
}
