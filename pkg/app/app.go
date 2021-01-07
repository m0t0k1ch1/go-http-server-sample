package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// App is the main application.
type App struct {
	*echo.Echo
}

// New creates an instance of App.
func New() *App {
	app := &App{echo.New()}
	app.Logger.SetLevel(log.INFO)

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&Context{c})
		}
	})
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c *Context) error {
		return c.String(http.StatusOK, "poyopoyo")
	})

	return app
}

// HandlerFunc is a function to serve HTTP requests.
type HandlerFunc func(c *Context) error

// Add registers a new route
func (app *App) Add(method, path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Echo.Add(method, path, func(c echo.Context) error {
		cc := c.(*Context)
		return h(cc)
	}, m...)
}

// GET registers a new GET route
func (app *App) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(http.MethodGet, path, h, m...)
}
