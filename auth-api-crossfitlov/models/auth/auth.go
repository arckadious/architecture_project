package auth

import (
	"auth-api-crossfitlov/parameters"
	"net/http"
)

// BasicAuth valide user et password
func BasicAuth(username string, password string, authOK bool) bool {

	config := parameters.Config

	if authOK == false || username != config.Auth.Username || password != config.Auth.Password {
		return false
	}

	return true
}

// checkAuth indique si l'auth est correct
func CheckAuth(w http.ResponseWriter, r *http.Request) bool {
	// Vérifier les droits
	username, password, authOK := r.BasicAuth()
	if !BasicAuth(username, password, authOK) {
		return false
	}

	return true

}
