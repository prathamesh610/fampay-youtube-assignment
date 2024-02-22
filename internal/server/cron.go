package server

import (
	"fmt"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/service/thirdparty"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) StartCron(context echo.Context) error {
	searchQuery := context.Param("searchQuery")
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Task executed at", time.Now())
			err := thirdparty.GetYoutubeResultAndPopulateDB(context.Request().Context(), s.DB, searchQuery)
			if err != nil {
				return context.JSON(http.StatusInternalServerError, err)
			}
		}
	}
}
