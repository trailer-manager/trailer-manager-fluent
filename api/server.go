package api

import (
	"SiverPineValley/trailer-manager/config"
	db "SiverPineValley/trailer-manager/db/rdb"
	router "SiverPineValley/trailer-manager/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	config config.Config
	Router *echo.Echo
}

func NewServer(store db.Store) (*Server, error) {
	conf := config.GetConfig()
	server := &Server{store: store, config: conf}

	corsConfig := middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	}

	if len(conf.AllowOrigins) > 0 {
		corsConfig.AllowOrigins = conf.AllowOrigins
	}

	e := echo.New()
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = methodNotAllowdHandler
	e.Use(transactionIdHandler)
	e.Use(httpLogHandler)
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(corsConfig))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentSecurityPolicy: "default-src 'self'",
	}))
	e.HTTPErrorHandler = errHandler

	// setup Router
	router.InitRouter(e)

	server.Router = e
	return server, nil
}
