package app

import (
	"database/sql"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

// NewTestApp creates an instance of App for test.
func NewTestApp(t *testing.T, rdb *sql.DB) *App {
	app := &App{
		Echo: echo.New(),
	}

	app.Use(middleware.Recover())

	app.env = &common.Env{
		RDB: rdb,
	}
	app.initRoutes()

	return app
}
