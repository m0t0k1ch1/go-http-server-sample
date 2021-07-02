package server

import (
	"database/sql"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/app"
)

// NewTestServer creates an instance of Server for test.
func NewTestServer(t *testing.T, db *sql.DB) *Server {
	s := &Server{
		Echo: echo.New(),
	}

	s.Use(middleware.Recover())

	s.env = &app.Env{
		DB: db,
	}
	s.initRoutes()

	return s
}
