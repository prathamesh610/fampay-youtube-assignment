package thirdparty

import (
	"encoding/json"
	"fmt"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dberrors"
	"io"
	"net/http"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/constants"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models/response"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/utils"
)

var keys map[string]bool
var initialized bool

func initializeMap() {
	keys = make(map[string]bool)
}
func addKeys(key string) {
	keys[key] = false
}

func InitializeAndAddKeys(key string) {
	if !initialized {
		initialized = true
		initializeMap()
	}
	addKeys(key)
}

func exhaustedKey(key string) {
	keys[key] = true
}

func GetYoutubeResultAndPopulateDB(client database.DatabaseClient, searchQuery string) error {

	nextPageToken := ""

	result := make([]response.YoutubeSearchResponse, 0)

	for i := 0; i < 5; i++ {
		res, _ := getSearchResult(searchQuery, nextPageToken)
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
func getSearchResult(searchQuery string, nextPageToken string) (*response.YoutubeSearchResponse, error) {

	url := constants.SEARCH + "?q=" + searchQuery + "&kind=youtube%23searchListResponse&publishedAfter=2024-01-01T00:00:00Z"

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

	url := constants.VIDEOS + "?part=snippet&id=" + videoId

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
	var apiKey string

	for key, val := range keys {
		if !val {
			apiKey = key
			break
		}
	}

	if apiKey == "" {
		return result, &dberrors.KeysExhausted{}
	}
	url += "&key=" + apiKey

	res, err := http.Get(url)
	if res.StatusCode == http.StatusForbidden {
		exhaustedKey(apiKey)
		parse, err := getResultAndParse(url, result)
		if err != nil {
			return parse, err
		}
		return parse, nil
	}

	if err != nil {
		fmt.Printf("Unable to get results for url %s with error %v\n", url, err)
		return result, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Unable to read body in [thirdparty.getResultAndParse] with error %v\n", err)
		return result, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Can not unmarshal JSON with error %v\n", err)
		return result, err
	}
	return result, nil

}
