package handlers

import (
	"net/http"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

// HandlePing is a sample HandlerFunc.
func HandlePing(env *common.Env, c *common.Context) error {
	return c.String(http.StatusOK, "pong")
}
