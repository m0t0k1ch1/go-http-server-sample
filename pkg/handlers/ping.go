package handlers

import (
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

// HandlePing is a sample HandlerFunc.
func HandlePing(env *common.Env, c *common.Context) error {
	if err := env.DB.Ping(); err != nil {
		return c.InternalServerError(err)
	}

	return c.SuccessWithEmpty()
}
