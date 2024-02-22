package database

import (
	"context"
	"fmt"
	"time"

	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool

	GetVideo(ctx context.Context, video models.Video) (*models.Video, error)
	SaveVideosToDB(ctx context.Context, videos *[]models.Video) (*[]models.Video, error)
	SearchVideo(ctx context.Context, searchQuery string) (*[]models.Video, error)

	GetThumbnail(ctx context.Context, video models.ThumbnailUrl) (*models.ThumbnailUrl, error)
	SaveThumbnailsToDB(ctx context.Context, thumbnails *[]models.ThumbnailUrl) (*[]models.ThumbnailUrl, error)
	SearchThumbnails(ctx context.Context, videoId []string) (*[]models.ThumbnailUrl, error)
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		5432,
		"disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "fampay.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}

	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)

	if tx.Error != nil {
		return false
	}

	if ready == "1" {
		return true
	}

	return false
}
