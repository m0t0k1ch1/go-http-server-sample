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
func New(conf *common.Config) *App {
	app := &App{
		Echo: echo.New(),
	}
	app.Logger.SetLevel(log.INFO)

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.env = &common.Env{
		Config: conf,
	}

	app.GET("/ping", handlers.HandlePing)

	return app
}

// Add registers a new route.
func (app *App) Add(method, path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Echo.Add(method, path, func(c echo.Context) error {
		return h(app.env, &common.Context{
			Context: c,
		})
	}, m...)
}

// GET registers a new GET route.
func (app *App) GET(path string, h common.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(http.MethodGet, path, h, m...)
}

// Start starts an HTTP server.
func (app *App) Start() error {
	return app.Echo.Start(fmt.Sprintf(":%d", app.env.Config.Port))
}
