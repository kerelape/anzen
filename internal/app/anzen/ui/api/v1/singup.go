package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kerelape/anzen/internal/app/anzen/ui"
	"github.com/kerelape/anzen/internal/app/anzen/ui/api/v1/signup"
)

// Signup entry.
type Signup struct {
	secret ui.HTTPEntry
}

// NewSignup create new Signup.
func NewSignup() *Signup {
	return &Signup{
		secret: signup.NewSecret(),
	}
}

// Router this entry.
func (signup *Signup) Route() http.Handler {
	router := chi.NewRouter()
	router.Mount("/secret", signup.secret.Route())
	return router
}
