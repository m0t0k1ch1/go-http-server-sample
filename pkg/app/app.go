package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/handlers"
)

// App is the main application.
type App struct {
	*echo.Echo
	env *common.Env
}

// New creates an instance of App.
func New(conf common.Config) (*App, error) {
	app := &App{
		Echo: echo.New(),
	}
	app.Logger.SetLevel(log.INFO)

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	env, err := common.NewEnv(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create an instance of common.Env: %w", err)
	}

	app.env = env
	app.initRoutes()

	return app, nil
}

func (app *App) initRoutes() {
	app.GET("/ping", handlers.HandlePing)
	app.POST("/albums", handlers.HandlePostAlbum)
	app.GET("/albums", handlers.HandleGetAlbums)
	app.GET("/albums/:ean", handlers.HandleGetAlbum)
	app.PATCH("/albums/:ean", handlers.HandlePatchAlbum)
	app.DELETE("/albums/:ean", handlers.HandleDeleteAlbum)
}

// Add registers a new route.
func (app *App) Add(method, path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Echo.Add(method, path, func(c echo.Context) error {
		return h(app.env, &common.Context{
			Context: c,
		})
	}, m...)
}

// POST registers a new POST route.
func (app *App) POST(path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(http.MethodPost, path, h, m...)
}

// GET registers a new GET route.
func (app *App) GET(path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(http.MethodGet, path, h, m...)
}

// PATCH registers a new PATCH route.
func (app *App) PATCH(path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(http.MethodPatch, path, h, m...)
}

// DELETE registers a new DELETE route.
func (app *App) DELETE(path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(http.MethodDelete, path, h, m...)
}

// Start starts an HTTP server.
func (app *App) Start() error {
	return app.Echo.Start(fmt.Sprintf(":%d", app.env.Config.Port))
}
