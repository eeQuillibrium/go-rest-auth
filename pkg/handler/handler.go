package handler

import (
	"net/http"

	auth "github.com/eeQuillibrium/go-rest-auth"
	"github.com/eeQuillibrium/go-rest-auth/pkg/service"
)

type Handler struct {
	htmpl    *auth.TemplateWrapper
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	tmpl := new(auth.TemplateWrapper)
	tmpl.Initialize("assets/*.html") //relative to the go-rest-auth/template.go
	return &Handler{services: services, htmpl: tmpl}
}

func (h *Handler) InitRoutes(fileserver http.Handler) *http.ServeMux {

	mux := http.NewServeMux()
	mux.Handle("/auth", http.HandlerFunc(h.authHandler))
	mux.Handle("/auth/signUp", http.HandlerFunc(h.signUpHandler))
	mux.Handle("/auth/signIn", http.HandlerFunc(h.signInHandler))
	mux.Handle("/auth/finishAuth", http.HandlerFunc(h.finishAuthHandler))

	mux.Handle("/assets/", fileserver)
	return mux
}
