package ui

import "net/http"

// HTTPEntry that can be routed to a handler.
type HTTPEntry interface {

	// Route this entry to an http.Handler.
	Route() http.Handler
}
