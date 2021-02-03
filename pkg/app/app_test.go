package app

import (
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"

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
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}

	// GET /albums
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if len(albums) > 0 {
		t.Errorf("expected: %d, actual: %d", 0, len(albums))
	}

	// POST /albums
	testutils.DoAPIRequest(t, app, http.MethodPost, "/albums", `{
		"ean":    "4995879601242",
		"title":  "GREEN",
		"artist": "SEEDA"
	}`, &album)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(album, models.Album{
		EAN:    "4995879601242",
		Title:  "GREEN",
		Artist: "SEEDA",
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}

	// POST /albums
	testutils.DoAPIRequest(t, app, http.MethodPost, "/albums", `{
		"ean":    "4997184881425",
		"title":  "modal soul",
		"artist": "Nujabes"
	}`, &album)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(album, models.Album{
		EAN:    "4997184881425",
		Title:  "modal soul",
		Artist: "Nujabes",
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}

	// GET /albums
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(albums, []models.Album{
		{
			EAN:    "4995879601242",
			Title:  "GREEN",
			Artist: "SEEDA",
		}, {
			EAN:    "4997184881425",
			Title:  "modal soul",
			Artist: "Nujabes",
		},
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}

	// GET /albums/4995879601242
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums/4995879601242", "", &album)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(album, models.Album{
		EAN:    "4995879601242",
		Title:  "GREEN",
		Artist: "SEEDA",
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}

	// GET /albums/4997184881425
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums/4997184881425", "", &album)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(album, models.Album{
		EAN:    "4997184881425",
		Title:  "modal soul",
		Artist: "Nujabes",
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}

	// DELETE /albums/4995879601242
	statusCode = testutils.DoAPIRequest(t, app, http.MethodDelete, "/albums/4995879601242", "", nil)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}

	// GET /albums
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(albums, []models.Album{
		{
			EAN:    "4997184881425",
			Title:  "modal soul",
			Artist: "Nujabes",
		},
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}
