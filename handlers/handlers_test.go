package handlers

// Credits: https://speedscale.com/blog/testing-golang-with-httptest/

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
	strData := string(data)
	if string(data) != expected {
		t.Errorf("Expected [%v] but got [%v]", expected, strData)
	}
}

func TestCpu(t *testing.T) {
	expected := "832040"
	req := httptest.NewRequest(http.MethodGet, "/cpu?x=30", nil)
	w := httptest.NewRecorder()
	Cpu(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	strData := string(data)
	if string(data) != expected {
		t.Errorf("Expected [%v] but got [%v]", expected, strData)
	}
}

func TestTcp(t *testing.T) {
	expected := "TCP connection: OK\n"
	req := httptest.NewRequest(http.MethodGet, "/tcp?addr=google.com:443", nil)
	w := httptest.NewRecorder()
	Tcp(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	strData := string(data)
	if string(data) != expected {
		t.Errorf("Expected [%v] but got [%v]", expected, strData)
	}
}
