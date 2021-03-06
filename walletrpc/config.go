package walletrpc

import (
	"net/http"
)

// Config holds the configuration of a monero rpc client.
type Config struct {
	Scheme        string
	Host          string
	Port          string
	CustomHeaders map[string]string
	Transport     http.RoundTripper
}
