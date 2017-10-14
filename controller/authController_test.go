package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func TestIsAdmin(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		err := IsAdmin(w, r)
		if err == nil{
			t.Fatal("UPS!! I am admin, but i not have admin cookies!!1")
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	http.Get(server.URL)
}

func TestLogout(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Logout(w, r)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, nil)
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	if resp.Header["Location"][0] != "/login"{
		t.Fatal("UPS!! No redirect after press logout button")
	}
}