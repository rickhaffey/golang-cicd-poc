package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	home(w, req)
	res := w.Result()

	defer func() {
		_ = res.Body.Close()
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	got := string(data)
	want := "Hello from simple API"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
