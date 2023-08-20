package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// https://speedscale.com/blog/testing-golang-with-httptest/

func TestOk(t *testing.T) {
	expected := "OK"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	Ok(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected Hello john but got %v", string(data))
	}
}
