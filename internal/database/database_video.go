package database

import (
	"context"
	"errors"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/dberrors"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c Client) GetVideo(videoId string) (*models.Video, error) {
	video1 := &models.Video{}
	filter := bson.M{"_id": videoId}
	err := c.DB.Database("fampay").Collection("videos").FindOne(context.TODO(), filter).Decode(&video1)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &dberrors.NotFoundError{
				Entity: "video",
				Query:  videoId,
			}
		}
		return nil, err
	}

	return video1, nil

}

func (c Client) SaveVideoInDB(video *models.Video) error {
	_, err := c.DB.Database("fampay").Collection("videos").InsertOne(context.TODO(), video)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) UpdateVideoInDB(video *models.Video) error {
	update := bson.M{
		"$set": video,
	}
	_, err := c.DB.Database("fampay").Collection("videos").UpdateByID(context.TODO(), video.VideoId, update)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) SearchVideo(searchQuery string, pageNumber int) (*[]models.Video, error) {
	videos := &[]models.Video{}
	regex := bson.M{"$regex": searchQuery, "$options": "i"}
	// define the filter to search multiple fields using fuzzy logic
	filter := bson.M{
		"$or": []bson.M{
			{"title": regex},
			{"description": regex},
		},
	}

	sort := bson.M{"publishingDate": -1}

	res, err := c.DB.Database("fampay").Collection("videos").Find(context.TODO(), filter, options.Find().SetSort(sort).SetSkip(int64(pageNumber*5)).SetLimit(5))
	if err != nil {
		return nil, err
	}

	if err = res.All(context.TODO(), videos); err != nil {
		return nil, err
	}

	if len(*videos) == 0 {
		return nil, &dberrors.NotFoundError{Entity: "Video", Query: searchQuery}
	}

	return videos, nil
}

func (c Client) GetVideosCount() (int64, error) {
	res, err := c.DB.Database("fampay").Collection("videos").EstimatedDocumentCount(context.TODO())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, &dberrors.NotFoundError{
				Entity: "video",
				Query:  "count",
			}
		}
		return 0, err
	}
	return res, nil
}

func (c Client) GetVideosPaginated(pageNumber int) (*[]models.Video, error) {
	videos := &[]models.Video{}
	sort := bson.M{"publishingDate": -1}
	res, err := c.DB.Database("fampay").Collection("videos").Find(context.TODO(), bson.D{}, options.Find().SetSort(sort).SetSkip(int64(pageNumber*5)).SetLimit(5))
	if err != nil {
		return nil, err
	}

	if err = res.All(context.TODO(), videos); err != nil {
		return nil, err
	}

	if len(*videos) == 0 {
		return nil, &dberrors.NotFoundError{Entity: "Video", Query: "videos-paginated"}
	}

	return videos, nil
}
