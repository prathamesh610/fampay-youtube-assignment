package database

import (
	"context"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"

	//"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool

	GetVideo(videoId string) (*models.Video, error)
	GetVideosCount() (int64, error)
	GetVideosPaginated(pageNumber int) (*[]models.Video, error)
	SaveVideoInDB(video *models.Video) error
	UpdateVideoInDB(video *models.Video) error
	SearchVideo(searchQuery string, pageNumber int) (*[]models.Video, error)
}

type Client struct {
	DB *mongo.Client
}

func NewDatabaseClient() (DatabaseClient, error) {
	// MongoDB connection string
	mongoURI := "mongodb://mongodb:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Client{DB: client}, nil
}

func (c *Client) Ready() bool {
	err := c.DB.Ping(context.Background(), nil)
	if err != nil {
		return false
	}
	return true
}
