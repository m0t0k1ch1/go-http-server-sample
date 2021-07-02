package handlers

import (
	"context"
	"database/sql"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/app"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/db"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/handlers/requests"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/models"
	v "github.com/m0t0k1ch1/go-http-server-sample/pkg/validation"
)

// HandlePostAlbum is an HandlerFunc to create a new album.
func HandlePostAlbum(env *app.Env, c *app.Context) error {
	var req models.Album
	if err := c.Bind(&req); err != nil {
		return c.BadRequest("invalid json format")
	}

	if err := req.Validate(); err != nil {
		return c.BadRequest(err.Error())
	}

	ctx := context.Background()

	albumDup, err := db.FetchAlbum(ctx, env.DB, req.EAN)
	if err != nil {
		return c.InternalServerError(err)
	}
	if albumDup != nil {
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
func HandleGetAlbums(env *app.Env, c *app.Context) error {
	albums, err := db.FetchAlbums(context.Background(), env.DB)
	if err != nil {
		return c.InternalServerError(err)
	}

	return c.Success(albums)
}

// HandleGetAlbum is an HandlerFunc to fetch an album by specifying EAN.
func HandleGetAlbum(env *app.Env, c *app.Context) error {
	ean := c.Param("ean")
	if err := v.ValidateEAN(ean); err != nil {
		return c.BadRequest("invalid ean")
	}

	album, err := db.FetchAlbum(context.Background(), env.DB, ean)
	if err != nil {
		return c.InternalServerError(err)
	}
	if album == nil {
		return c.NotFound("album not found")
	}

	return c.Success(album)
}

// HandlePatchAlbum is an HandlerFunc to patch an album by specifying EAN.
func HandlePatchAlbum(env *app.Env, c *app.Context) error {
	ean := c.Param("ean")
	if err := v.ValidateEAN(ean); err != nil {
		return c.BadRequest("invalid ean")
	}

	var req requests.PatchAlbumRequest
	if err := c.Bind(&req); err != nil {
		return c.BadRequest("invalid json format")
	}
	if req.IsEmpty() {
		return c.BadRequest("no parameter specified")
	}

	params := db.QueryParams{}

	if req.HasTitle() {
		title, err := req.GetTitleWithValidation()
		if err != nil {
			return c.BadRequest(err.Error())
		}
		params["title"] = title
	}
	if req.HasArtist() {
		artist, err := req.GetArtistWithValidation()
		if err != nil {
			return c.BadRequest(err.Error())
		}
		params["artist"] = artist
	}

	ctx := context.Background()

	album, err := db.FetchAlbum(ctx, env.DB, ean)
	if err != nil {
		return c.InternalServerError(err)
	}
	if album == nil {
		return c.NotFound("album not found")
	}

	if err := db.Transact(ctx, env.DB, func(txCtx context.Context, tx *sql.Tx) error {
		var txErr error

		if album, txErr = db.FetchAlbumForUpdate(txCtx, tx, ean); txErr != nil {
			return txErr
		}

		if txErr = db.UpdateAlbum(txCtx, tx, ean, params); txErr != nil {
			return txErr
		}

		return nil
	}); err != nil {
		return c.InternalServerError(err)
	}

	return c.SuccessWithEmpty()
}

// HandleDeleteAlbum is an HandlerFunc to delete an album by specifying EAN.
func HandleDeleteAlbum(env *app.Env, c *app.Context) error {
	ean := c.Param("ean")
	if err := v.ValidateEAN(ean); err != nil {
		return c.BadRequest("invalid ean")
	}

	ctx := context.Background()

	album, err := db.FetchAlbum(ctx, env.DB, ean)
	if err != nil {
		return c.InternalServerError(err)
	}
	if album == nil {
		return c.NotFound("album not found")
	}

	if err := db.Transact(ctx, env.DB, func(txCtx context.Context, tx *sql.Tx) error {
		var txErr error

		if album, txErr = db.FetchAlbumForUpdate(txCtx, tx, ean); txErr != nil {
			return txErr
		}

		if txErr = db.DeleteAlbum(txCtx, tx, ean); txErr != nil {
			return txErr
		}

		return nil
	}); err != nil {
		return c.InternalServerError(err)
	}

	return c.SuccessWithEmpty()
}
