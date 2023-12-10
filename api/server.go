package api

import (
	"github.com/labstack/echo/v4"
	"github.com/trailer-manager/trailer-manager-common/config"
	"github.com/trailer-manager/trailer-manager-common/server"
	db "github.com/trailer-manager/trailer-manager-fluent/db/rdb"
	router "github.com/trailer-manager/trailer-manager-fluent/router"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	config config.Config
	Router *echo.Echo
}

func NewFluentServer(store db.Store) (*Server, error) {
	conf := config.GetConfig()
	svc := &Server{store: store, config: conf}
	e := server.NewServer()
	router.InitRouter(e)
	svc.Router = e
	return svc, nil
}
