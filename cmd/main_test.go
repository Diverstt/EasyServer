package main

import (
	"net/http"
	"testing"
	"time"
)

func TestMiddleware(t *testing.T) {
	go main()
	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}
