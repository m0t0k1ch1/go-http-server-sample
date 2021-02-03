package handlers

import (
	"context"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/db"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/models"
)

// HandlePostAlbum is an HandlerFunc to create a new album.
func HandlePostAlbum(env *common.Env, c *common.Context) error {
	var req models.Album
	if err := c.Bind(&req); err != nil {
		return c.BadRequest("invalid json format")
	}

	if err := req.Validate(); err != nil {
		return c.BadRequest(err.Error())
	}

	ctx := context.Background()

	dup, err := db.FetchAlbum(ctx, env.DB, req.EAN)
	if err != nil {
		return c.InternalServerError(err)
	}
	if dup != nil {
		return c.BadRequest("album already exists")
	}

	if err := db.CreateAlbum(ctx, env.DB, &req); err != nil {
		return c.InternalServerError(err)
	}

	album, err := db.FetchAlbum(ctx, env.DB, req.EAN)
	if err != nil {
		return c.InternalServerError(err)
	}

	return c.Success(album)
}

// HandleGetAlbums is an HandlerFunc to fetch all albums.
func HandleGetAlbums(env *common.Env, c *common.Context) error {
	albums, err := db.FetchAlbums(context.Background(), env.DB)
	if err != nil {
		return c.InternalServerError(err)
	}

	return c.Success(albums)
}

// HandleGetAlbum is an HandlerFunc to fetch an album by specifying EAN.
func HandleGetAlbum(env *common.Env, c *common.Context) error {
	ean := c.Param("ean")

	// TODO: validation

	album, err := db.FetchAlbum(context.Background(), env.DB, ean)
	if err != nil {
		return c.InternalServerError(err)
	}
	if album == nil {
		return c.NotFound("album not found")
	}

	return c.Success(album)
}

// HandleDeleteAlbum is an HandlerFunc to delete an album by specifying EAN.
func HandleDeleteAlbum(env *common.Env, c *common.Context) error {
	return nil
}
