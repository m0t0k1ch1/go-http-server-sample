package handlers

import (
	"context"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/db/models"
)

// HandleGetAlbums is an HandlerFunc to get all albums.
func HandleGetAlbums(env *common.Env, c *common.Context) error {
	albums, err := models.FetchAlbums(context.Background(), env.RDB)
	if err != nil {
		return c.InternalServerError(err)
	}

	return c.Success(albums)
}
