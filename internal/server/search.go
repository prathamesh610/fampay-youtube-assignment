package server

import (
	"github.com/labstack/echo/v4"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/service"
	"net/http"
)

func (s *EchoServer) Search(ctx echo.Context) error {
	query := ctx.Param("query")
	videos, err := service.SearchUsingQueryInDB(ctx.Request().Context(), s.DB, query)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, videos)
}
