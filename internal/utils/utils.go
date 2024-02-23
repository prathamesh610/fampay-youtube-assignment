package utils

import (
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models/response"
)

func ConvertYoutubeVideoResToDbModel(videoRes response.YoutubeVideoResponse) *[]models.Video {
	videoResult := make([]models.Video, 0)

	for _, val := range videoRes.Items {
		thumbnails := &models.Thumbnails{
			Default: models.Thumbnail{
				URL:    val.Snippet.Thumbnails.Default.URL,
				Width:  val.Snippet.Thumbnails.Default.Width,
				Height: val.Snippet.Thumbnails.Default.Height,
			},
			Medium: models.Thumbnail{
				URL:    val.Snippet.Thumbnails.Medium.URL,
				Width:  val.Snippet.Thumbnails.Default.Width,
				Height: val.Snippet.Thumbnails.Default.Height,
			},
			High: models.Thumbnail{
				URL:    val.Snippet.Thumbnails.High.URL,
				Width:  val.Snippet.Thumbnails.High.Width,
				Height: val.Snippet.Thumbnails.High.Height,
			},
			Standard: models.Thumbnail{
				URL:    val.Snippet.Thumbnails.Standard.URL,
				Width:  val.Snippet.Thumbnails.Standard.Width,
				Height: val.Snippet.Thumbnails.Standard.Height,
			},
			Maxres: models.Thumbnail{
				URL:    val.Snippet.Thumbnails.Maxres.URL,
				Width:  val.Snippet.Thumbnails.Maxres.Width,
				Height: val.Snippet.Thumbnails.Maxres.Height},
		}

		videoModel := &models.Video{
			Thumbnails:     *thumbnails,
			VideoTitle:     val.Snippet.Title,
			Description:    val.Snippet.Description,
			PublishingDate: val.Snippet.PublishedAt,
			VideoId:        val.ID,
		}

		videoResult = append(videoResult, *videoModel)

	}

	return &videoResult
}

func ConvertConvertYoutubeVideoResToDbModelSlice(videoRes []response.YoutubeVideoResponse) *[]models.Video {
	videoResult := make([]models.Video, 0)
	for _, val := range videoRes {
		resVideo := ConvertYoutubeVideoResToDbModel(val)

		videoResult = append(videoResult, *resVideo...)
	}

	return &videoResult
}
