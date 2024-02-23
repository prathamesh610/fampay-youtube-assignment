package thirdparty

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"net/http"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/constants"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models/response"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/utils"
)

// type YoutubeService interface {
// 	GetYoutubeResultAndPopulateDB(ctx context.Context, client database.DatabaseClient, searchQuery string) error
// }

func GetYoutubeResultAndPopulateDB(ctx context.Context, client database.DatabaseClient, searchQuery string) error {

	nextPageToken := ""

	result := make([]response.YoutubeSearchResponse, 0)

	for i := 0; i < 5; i++ {
		res, _ := getSerachResult(searchQuery, nextPageToken)
		nextPageToken = res.NextPageToken
		result = append(result, *res)
		if nextPageToken == "" {
			break
		}
	}

	videoResult := make([]response.YoutubeVideoResponse, 0)

	for _, val := range result {
		for _, val1 := range val.Items {
			res, _ := getVideoDetails(val1.ID.VideoID)
			videoResult = append(videoResult, *res)
		}
	}

	videoModels := utils.ConvertConvertYoutubeVideoResToDbModelSlice(videoResult)

	for _, model := range *videoModels {
		video, err := client.GetVideo(model.VideoId)
		if err != nil {
			err := client.SaveVideoInDB(&model)
			if err != nil {
				return err
			}

		}
		if video != nil {
			err := client.UpdateVideoInDB(&model)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// This method fetches search response for a query
func getSerachResult(searchQuery string, nextPageToken string) (*response.YoutubeSearchResponse, error) {

	url := constants.SEARCH + "?key=" + "AIzaSyCB5H390Q04K9SawzKYBTHVPc9mE4tU200" + "&q=" + searchQuery + "&kind=youtube%23searchListResponse&publishedAfter=2024-01-01T00:00:00Z"

	if nextPageToken != "" {
		url = url + "&pageToken=" + nextPageToken
	}

	result := &response.YoutubeSearchResponse{}

	result, err := getResultAndParse(url, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// This method for every video id fetches video details from youtube
func getVideoDetails(videoId string) (*response.YoutubeVideoResponse, error) {

	url := constants.VIDEOS + "?key=" + "AIzaSyCB5H390Q04K9SawzKYBTHVPc9mE4tU200" + "&part=snippet&id=" + videoId

	result := &response.YoutubeVideoResponse{}

	result, err := getResultAndParse(url, result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// Generic function to get and parse response from youtube returning
// the parsed object.
func getResultAndParse[T any](url string, result T) (T, error) {

	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("Unable to get results for url %s with error %v\n", url, err)
		return result, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Unable to read body in [thirdparty.getResultAndParse] with error %v\n", err)
		return result, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Can not unmarshal JSON with error %v\n", err)
		return result, err
	}
	return result, nil

}
