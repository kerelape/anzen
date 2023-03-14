package signup

import (
	"net/http"
)

// Secret entry.
type Secret struct {
}

// NewSecret creates new Secret.
func NewSecret() *Secret {
	return &Secret{}
}

// Route this entry.
func (secret *Secret) Route() http.Handler {
	return secret
}

// ServeHTTP handles secret signup requests.
func (secret *Secret) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
