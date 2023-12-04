package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"

	auth "github.com/eeQuillibrium/go-rest-auth"
)

func (h *Handler) signInHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		h.htmpl.Tmpl.ExecuteTemplate(w, "signIn.html", nil)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Fatal().
			Err(err)
	}

	attemptedUser := auth.User{Login: r.FormValue("login"), Password: r.FormValue("pass")}

	token, err := h.services.Authorization.CheckUser(attemptedUser)
	if err != nil {
		log.Fatal().AnErr("Check user err", err)
	}

	log.Debug().Msg("SignIn executed token:" + token)

	c := &http.Cookie{
		Name:  "token",
		Value: token,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/auth/finishAuth", http.StatusPermanentRedirect)
}
func (h *Handler) signUpHandler(w http.ResponseWriter, r *http.Request) {
	h.htmpl.Tmpl.ExecuteTemplate(w, "signUp.html", nil)

	if r.Method != "POST" {
		return
	}
	log.Debug().Msg("signUp executed")

	if err := r.ParseForm(); err != nil {
		log.Fatal().
			Err(err)
	}

	inputUser := auth.User{Login: r.FormValue("login"), Password: r.FormValue("pass")}
	h.services.Authorization.CreateUser(inputUser)

	http.Redirect(w, r, "/finishAuth", http.StatusPermanentRedirect)
}
func (h *Handler) finishAuthHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		log.Fatal().AnErr("Unauthorized", err)
		return
	}
	log.Debug().Msg("JWT token: " + c.Value)
	h.htmpl.Tmpl.ExecuteTemplate(w, "authFinish.html", nil)
}
func (h *Handler) authHandler(w http.ResponseWriter, r *http.Request) {
	h.htmpl.Tmpl.ExecuteTemplate(w, "auth.html", nil)
}
