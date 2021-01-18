package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

// HandlePing is a sample HandlerFunc.
func HandlePing(env *common.Env, c *common.Context) error {
	if err := env.RDB.Ping(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "pong")
}
