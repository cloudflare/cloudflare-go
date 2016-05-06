package cloudflare

import (
	"net/http"
	"net/http/httptest"
)

var (
	// mux is the HTTP request mulitplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested
	client *API

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// CloudFlare client configured to use test server
	client = New("cloudflare@example.org", "deadbeef")
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}
