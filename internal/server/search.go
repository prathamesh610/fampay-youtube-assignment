package server

import (
	"github.com/labstack/echo/v4"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/service"
	"net/http"
	"strconv"
)

func (s *EchoServer) Search(ctx echo.Context) error {
	query := ctx.Param("query")
	pageString := ctx.QueryParam("page")
	pageNumber := 1
	if pageString != "" {
		page, err := strconv.Atoi(pageString)
		if err != nil || page < 1 {
			return ctx.JSON(http.StatusBadRequest, "Please provide valid pageNumber")
		}
		pageNumber = page
	}
	videos, err := service.SearchUsingQueryInDB(s.DB, query, pageNumber)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, videos)
}

func (s *EchoServer) GetAllVideos(ctx echo.Context) error {
	pageString := ctx.QueryParam("page")
	pageNumber := 1
	if pageString != "" {
		page, err := strconv.Atoi(pageString)
		if err != nil || page < 1 {
			return ctx.JSON(http.StatusBadRequest, "Please provide valid pageNumber")
		}
		pageNumber = page
	}

	res, err := service.GetAllVideos(s.DB, pageNumber)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
