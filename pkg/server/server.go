package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/app"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/handlers"
)

// Server is the main application server.
type Server struct {
	*echo.Echo
	env *app.Env
}

// New creates an instance of Server.
func New(conf app.Config) (*Server, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	s := &Server{
		Echo: e,
	}

	s.Logger.SetLevel(log.INFO)

	s.Use(middleware.Logger())
	s.Use(middleware.Recover())

	env, err := app.NewEnv(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create an instance of app.Env: %w", err)
	}

	s.env = env
	s.initRoutes()

	return s, nil
}

func (s *Server) initRoutes() {
	s.GET("/ping", handlers.HandlePing)
	s.POST("/albums", handlers.HandlePostAlbum)
	s.GET("/albums", handlers.HandleGetAlbums)
	s.GET("/albums/:ean", handlers.HandleGetAlbum)
	s.PATCH("/albums/:ean", handlers.HandlePatchAlbum)
	s.DELETE("/albums/:ean", handlers.HandleDeleteAlbum)
}

// Add registers a new route.
func (s *Server) Add(method, path string, h app.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.Echo.Add(method, path, func(c echo.Context) error {
		return h(s.env, &app.Context{
			Context: c,
		})
	}, m...)
}

// POST registers a new POST route.
func (s *Server) POST(path string, h app.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.Add(http.MethodPost, path, h, m...)
}

// GET registers a new GET route.
func (s *Server) GET(path string, h app.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.Add(http.MethodGet, path, h, m...)
}

// PATCH registers a new PATCH route.
func (s *Server) PATCH(path string, h app.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.Add(http.MethodPatch, path, h, m...)
}

// DELETE registers a new DELETE route.
func (s *Server) DELETE(path string, h app.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.Add(http.MethodDelete, path, h, m...)
}

// Start starts an HTTP server.
func (s *Server) Start() error {
	return s.Echo.Start(fmt.Sprintf(":%d", s.env.Config.Port))
}
