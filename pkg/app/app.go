package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Context represents the context of the current HTTP request
type Context struct {
	echo.Context
}

// HandlerFunc is a function to serve HTTP requests.
type HandlerFunc func(env *Env, c *Context) error

// App is the main application.
type App struct {
	*echo.Echo
	starter func() error
}

// New creates an instance of App.
func New(conf *Config) *App {
	app := &App{
		Echo: echo.New(),
	}
	app.Logger.SetLevel(log.INFO)

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&Context{c})
		}
	})
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	env := &Env{
		Config: conf,
	}

	app.GET(env, "/", Poyo)

	app.starter = func() error {
		return app.Echo.Start(fmt.Sprintf(":%d", conf.Port))
	}

	return app
}

// Add registers a new route
func (app *App) Add(env *Env, method, path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Echo.Add(method, path, func(c echo.Context) error {
		return h(env, &Context{c})
	}, m...)
}

// GET registers a new GET route
func (app *App) GET(env *Env, path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return app.Add(env, http.MethodGet, path, h, m...)
}

// Start starts an HTTP server
func (app *App) Start() error {
	return app.starter()
}

// Env holds some application-level objects.
type Env struct {
	Config *Config
}

// Poyo is a sample HandlerFunc
func Poyo(env *Env, c *Context) error {
	return c.String(http.StatusOK, "poyopoyo")
}
