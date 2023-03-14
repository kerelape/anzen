package signup

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kerelape/anzen/internal/app/anzen/ui"
)

// Signup entry.
type Signup struct {
	secret ui.HTTPEntry
}

// NewSignup create new Signup.
func NewSignup() *Signup {
	return &Signup{
		secret: nil,
	}
}

// Router this entry.
func (signup *Signup) Route() http.Handler {
	router := chi.NewRouter()
	router.Mount("/secret", signup.secret.Route())
	return router
}
