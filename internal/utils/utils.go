package utils

import (
	"context"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models/response"
)

func ConvertYoutubeVideoResToDbModel(ctx context.Context, client database.DatabaseClient, searchQuery string, videoRes response.YoutubeVideoResponse) (*[]models.Video, *[]models.ThumbnailUrl) {
	videoResult := make([]models.Video, 0)
	thumbanilsResult := make([]models.ThumbnailUrl, 0)

	for _, val := range videoRes.Items {
		videoModel := &models.Video{
			SearchQuery:    searchQuery,
			VideoTitle:     val.Snippet.Title,
			Description:    val.Snippet.Description,
			PublishingDate: val.Snippet.PublishedAt,
			VideoId:        val.ID,
		}

		client.GetVideo(ctx, *videoModel)
		thumbnailModel := &models.ThumbnailUrl{
			VideoId:      val.ID,
			ThumbnailUrl: val.Snippet.Thumbnails.Default.URL,
		}

		client.GetThumbnail(ctx, *thumbnailModel)
		videoResult = append(videoResult, *videoModel)
		thumbanilsResult = append(thumbanilsResult, *thumbnailModel)

	}

	return &videoResult, &thumbanilsResult
}

func ConvertConvertYoutubeVideoResToDbModelSlice(ctx context.Context, client database.DatabaseClient, searchQuery string, videoRes []response.YoutubeVideoResponse) (*[]models.Video, *[]models.ThumbnailUrl) {
	videoResult := make([]models.Video, 0)
	thumbanilsResult := make([]models.ThumbnailUrl, 0)
	for _, val := range videoRes {
		resVideo, resThumbnail := ConvertYoutubeVideoResToDbModel(ctx, client, searchQuery, val)

		videoResult = append(videoResult, *resVideo...)
		thumbanilsResult = append(thumbanilsResult, *resThumbnail...)
	}

	return &videoResult, &thumbanilsResult
}
