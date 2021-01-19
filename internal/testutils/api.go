package testutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// DoAPIRequest sends an HTTP request to the test server and returns an HTTP status code.
func DoAPIRequest(t *testing.T, h http.Handler, method, path, body string, res interface{}) int {
	t.Helper()

	srv := httptest.NewServer(h)
	defer srv.Close()

	req, err := http.NewRequest(method, srv.URL+path, strings.NewReader(body))
	if err != nil {
		t.Errorf("failed to create an instance of http.Request: %s", err)
		return 0
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to execute the request: %s", err)
		return 0
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read the response body: %s", err)
		return resp.StatusCode
	}

	if res != nil {
		if err := json.Unmarshal(b, res); err != nil {
			t.Errorf("failed to unmarshal the response body: %s", err)
		}
	}

	return resp.StatusCode
}
