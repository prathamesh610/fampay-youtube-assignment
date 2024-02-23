package server

import (
	"fmt"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	StartCron(context echo.Context) error

	Search(ctx echo.Context) error
	GetAllVideos(ctx echo.Context) error

	AddNewKey(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {

	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server shut down unexpectedly with error %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/cron")
	cg.GET("/:searchQuery", s.StartCron)

	sg := s.echo.Group("/search")
	sg.GET("", s.GetAllVideos)
	sg.GET("/:query", s.Search)

	yg := s.echo.Group("/keys")
	yg.POST("/:key", s.AddNewKey)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}

	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
