package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler ...
type Handler struct {
	authService AuthInterface
}

// NewHandler ...
func NewHandler(authService AuthInterface) *Handler {
	return &Handler{authService: authService}
}

type loginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	credentials := &loginCredentials{}
	err = json.Unmarshal(body, credentials)
	if err != nil {
		log.Print(err)
		return
	}
	loggedIn := h.authService.LoginUser(credentials.Email, credentials.Password)
	log.Print(loggedIn)
	w.Write([]byte("authed"))
}
