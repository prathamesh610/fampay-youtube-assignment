package server

import (
	"github.com/labstack/echo/v4"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/service/thirdparty"
)

func (s *EchoServer) AddNewKey(ctx echo.Context) error {
	key := ctx.Param("key")
	thirdparty.InitializeAndAddKeys(key)
	return nil
}
