package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", strings.NewReader("foo"))
	w := httptest.NewRecorder()
	Handler(w, req)

	if got := w.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	body, _ := io.ReadAll(w.Body)
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	if w.Result().StatusCode != http.StatusCreated {
		t.FailNow()
	}
}
