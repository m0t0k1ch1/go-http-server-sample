package app

import (
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/m0t0k1ch1/go-http-server-sample/internal/testutils"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/models"
)

func TestMain(m *testing.M) {
	testutils.Run(m)
}

func TestApp(t *testing.T) {
	db, truncate := testutils.SetUpDB()
	defer truncate()

	app := NewTestApp(t, db)

	var statusCode int

	var album models.Album
	var albums []models.Album

	// GET /ping
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/ping", "", nil)
	testutils.Equal(t, statusCode, http.StatusOK)

	// GET /albums
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, albums, []models.Album{})

	// POST /albums
	testutils.DoAPIRequest(t, app, http.MethodPost, "/albums", `{
		"ean":    "4995879601242",
		"title":  "GREEN",
		"artist": "YOSEEDA"
	}`, &album)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, album, models.Album{
		EAN:    "4995879601242",
		Title:  "GREEN",
		Artist: "YOSEEDA",
	})

	// POST /albums
	testutils.DoAPIRequest(t, app, http.MethodPost, "/albums", `{
		"ean":    "4997184881425",
		"title":  "modal soul",
		"artist": "Nujabes"
	}`, &album)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, album, models.Album{
		EAN:    "4997184881425",
		Title:  "modal soul",
		Artist: "Nujabes",
	})

	// GET /albums
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, albums, []models.Album{
		{
			EAN:    "4995879601242",
			Title:  "GREEN",
			Artist: "YOSEEDA",
		}, {
			EAN:    "4997184881425",
			Title:  "modal soul",
			Artist: "Nujabes",
		},
	})

	// GET /albums/4995879601242
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums/4995879601242", "", &album)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, album, models.Album{
		EAN:    "4995879601242",
		Title:  "GREEN",
		Artist: "YOSEEDA",
	})

	// GET /albums/4997184881425
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums/4997184881425", "", &album)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, album, models.Album{
		EAN:    "4997184881425",
		Title:  "modal soul",
		Artist: "Nujabes",
	})

	// PATCH /albums/4995879601242
	statusCode = testutils.DoAPIRequest(t, app, http.MethodPatch, "/albums/4995879601242", `{
		"artist": "SEEDA"
	}`, nil)
	testutils.Equal(t, statusCode, http.StatusOK)

	// GET /albums/4995879601242
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums/4995879601242", "", &album)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, album, models.Album{
		EAN:    "4995879601242",
		Title:  "GREEN",
		Artist: "SEEDA",
	})

	// DELETE /albums/4995879601242
	statusCode = testutils.DoAPIRequest(t, app, http.MethodDelete, "/albums/4995879601242", "", nil)
	testutils.Equal(t, statusCode, http.StatusOK)

	// GET /albums
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	testutils.Equal(t, statusCode, http.StatusOK)
	testutils.Equal(t, albums, []models.Album{
		{
			EAN:    "4997184881425",
			Title:  "modal soul",
			Artist: "Nujabes",
		},
	})
}
