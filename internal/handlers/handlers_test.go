package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDummyLoginHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dummyLogin?user_type=client", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DummyLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var expected string = `{"token":"token_client"}`
	fmt.Println(rr.Body.String())
	fmt.Println(expected)
	fmt.Println(rr.Body.String())
	if strings.Split(rr.Body.String(), "\n")[0] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
