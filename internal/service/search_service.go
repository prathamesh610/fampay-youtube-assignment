package service

import (
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dberrors"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dto"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
)

func SearchUsingQueryInDB(client database.DatabaseClient, searchQuery string, pageNumber int) (*dto.SearchResponseDTO, error) {

	pageNumber = pageNumber - 1
	videos, err := client.SearchVideo(searchQuery, pageNumber)
	if err != nil {
		return nil, err
	}

	videoResponse := convertVideoSliceToResponse(*videos)

	responseDto := &dto.SearchResponseDTO{
		SearchQuery: searchQuery,
		Videos:      videoResponse,
	}
	return responseDto, nil
}

func convertVideoSliceToResponse(videos []models.Video) []dto.VideoResponse {

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

func GetAllVideos(client database.DatabaseClient, pageNumber int) (*dto.ResponseDTO, error) {
	pageNumber = pageNumber - 1

	count, err := client.GetVideosCount()
	if err != nil {
		return nil, err
	}

	totalPages := count / 5
	if count == 5 {
		totalPages = 0
	}

	if int64(pageNumber) > totalPages {
		return nil, &dberrors.InvalidPageNumber{
			PageNumber: 0,
			ValidTill:  0,
		}
	}

	videos, err := client.GetVideosPaginated(pageNumber)
	if err != nil {
		return nil, err
	}
	videoResponse := convertVideoSliceToResponse(*videos)

	responseDto := &dto.ResponseDTO{
		VideosCount:    count,
		PageNumber:     pageNumber + 1,
		Videos:         videoResponse,
		LastPageNumber: totalPages + 1,
		Count:          len(*videos),
	}
	return responseDto, nil
}
