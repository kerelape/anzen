package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kerelape/anzen/internal/app/anzen/ui"
	"github.com/kerelape/anzen/internal/app/anzen/ui/api/v1/signup"
)

// APIv1 entry.
type APIv1 struct {
	login  ui.HTTPEntry
	signup ui.HTTPEntry
}

// NewAPIv1 create new apiv1.
func NewAPIv1() APIv1 {
	return APIv1{
		login:  nil,
		signup: signup.NewSignup(),
	}
}

// Route this entry.
func (api *APIv1) Route() http.Handler {
	router := chi.NewRouter()
	router.Mount("/login", api.login.Route())
	router.Mount("/signup", api.signup.Route())
	return router
}
