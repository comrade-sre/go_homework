package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1/?name=test", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := &Handler{}
	handler.ServeHTTP(reqRecorder, req)
	if status := reqRecorder.Code;status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Parsed param name with value: test`
	if reqRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqRecorder.Body.String(), expected)
	}
}
