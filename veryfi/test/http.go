package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HTTPServer describes a mock http/https server.
type HTTPServer struct {
	// URL has the `host:port` format.
	URL string

	// server is the mock http server.
	server *httptest.Server

	// mux is the http request multiplexer.
	mux *http.ServeMux
}

// NewHTTPServer returns an instance of a mock http server.
func NewHTTPServer() HTTPServer {
	m := http.NewServeMux()
	s := httptest.NewTLSServer(m)

	return HTTPServer{
		URL:    s.URL[8:],
		server: s,
		mux:    m,
	}
}

// ServeVersioned serves a versioned endpoint.
func (s HTTPServer) Serve(t *testing.T, uri string, statusCode int, response interface{}) {
	serve(s.mux, t, uri, statusCode, response)
}

// Close closes the server connection.
func (s HTTPServer) Close() {
	s.server.Close()
}

// serve registers a path handler and writes to its responses.
func serve(m *http.ServeMux, t *testing.T, uri string, statusCode int, response interface{}) {
	m.HandleFunc(
		uri,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			fprint(t, w, response)
		},
	)
}

// fprint calls fmt.Fprint and validates its returned error.
func fprint(t *testing.T, w io.Writer, a interface{}) {
	_, err := fmt.Fprint(w, a)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}
