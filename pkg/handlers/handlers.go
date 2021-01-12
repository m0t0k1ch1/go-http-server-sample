package handlers

import (
	"net/http"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

// Ping is a sample HandlerFunc
func Ping(env *common.Env, c *common.Context) error {
	return c.String(http.StatusOK, "pong")
}
