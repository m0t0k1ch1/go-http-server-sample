package app

import (
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/m0t0k1ch1/go-http-server-sample/internal/testutils"
)

func TestMain(m *testing.M) {
	testutils.Run(m)
}

func TestApp(t *testing.T) {
	rdb, truncate := testutils.SetUpRDB()
	defer truncate()

	app := NewTestApp(t, rdb)

	var statusCode int

	statusCode = testutils.DoAPIRequest(t, app, http.MethodGet, "/ping", "", nil)
	if statusCode != http.StatusOK {
		t.Errorf("expected: %d, actual: %d", http.StatusOK, statusCode)
	}
}
