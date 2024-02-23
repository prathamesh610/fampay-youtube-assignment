package service

import (
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dto"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
)

func SearchUsingQueryInDB(client database.DatabaseClient, searchQuery string, pageNumber int) (*dto.ResponseDTO, error) {

	pageNumber = pageNumber - 1
	videos, err := client.SearchVideo(searchQuery, pageNumber)
	if err != nil {
		return nil, err
	}

	videoResponse := assignThumbnailsToVideos(*videos)

	responseDto := &dto.ResponseDTO{
		SearchQuery: searchQuery,
		Videos:      videoResponse,
	}
	return responseDto, nil
}

func assignThumbnailsToVideos(videos []models.Video) []dto.VideoResponse {

	videoSlice := make([]dto.VideoResponse, 0)

	for _, video := range videos {
		videoModel := &dto.VideoResponse{
			VideoTitle:   video.VideoTitle,
			Description:  video.Description,
			ThumbnailUrl: video.Thumbnails,
			PublishedAt:  video.PublishingDate,
		}
		videoSlice = append(videoSlice, *videoModel)
	}

	return videoSlice
}
