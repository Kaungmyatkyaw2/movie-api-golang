package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	app := newTestApplicatoin(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/v1/healthcheck")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	var m map[string]any
	err := json.NewDecoder(strings.NewReader(string(body))).Decode(&m)

	if err != nil {
		t.Fatal("invalid json response")
	}

	if m["status"] != "available" {
		t.Errorf("want status: %v; got %v", "available", m["status"])
	}

}
