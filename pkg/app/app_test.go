package app

import (
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"

	"github.com/m0t0k1ch1/go-http-server-sample/internal/testutils"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/rdb/models"
)

func TestMain(m *testing.M) {
	testutils.Run(m)
}

func TestApp(t *testing.T) {
	rdb, truncate := testutils.SetUpRDB()
	defer truncate()

	app := NewTestApp(t, rdb)

	var statusCode int

	// GET /ping
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/ping", "", nil)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}

	// GET /albums
	var albums []models.Album
	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/albums", "", &albums)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, acttual: %d", http.StatusOK, statusCode)
	}
	if diff := cmp.Diff(albums, []models.Album{
		{EAN: "4988002758807", Title: "Juice", Artist: "iri"},
		{EAN: "4988005553027", Title: "This Is The One", Artist: "Utada"},
		{EAN: "4988008803235", Title: "MADRUGADA / TIGER EYES", Artist: "Jazztronik"},
		{EAN: "4995879601242", Title: "GREEN", Artist: "SEEDA"},
		{EAN: "4997184881425", Title: "modal soul", Artist: "Nujabes"},
	}); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}
