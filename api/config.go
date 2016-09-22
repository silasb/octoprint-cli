package api

import (
	"net/http"
)

type Config struct {
	// Logger is the logger used by client library.
	//Logger Logger

	Endpoint string
	Key      string

	Transport http.RoundTripper
}
