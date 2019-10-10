package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(viewHRTFHandler))
	defer server.Close()
	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}
	//router := NewRouter()

	if string(data) !=


	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	//router.ServeHTTP(rec, req)

	//assert.Equal(t, http.StatusOK, rec.Code)
	//assert.Equal(t, helloMessage, rec.Body.String())
}
