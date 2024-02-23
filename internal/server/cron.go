package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/service/thirdparty"
	"net/http"
	"time"
)

func (s *EchoServer) StartCron(context echo.Context) error {
	searchQuery := context.Param("searchQuery")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Task executed at", time.Now())
			err := thirdparty.GetYoutubeResultAndPopulateDB(s.DB, searchQuery)
			if err != nil {
				return context.JSON(http.StatusInternalServerError, err)
			}
		}
	}
}
