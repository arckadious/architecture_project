package post

import (
	"auth-api-crossfitlov/models/auth"
	"auth-api-crossfitlov/models/structs/in"
	"auth-api-crossfitlov/parameters"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//Check check is the token is still valid
func Check(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//Vérifier les droits
	if !auth.CheckAuth(w, r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var tokenStruct in.TokenStruct
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&tokenStruct)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := tokenStruct.Token
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(parameters.Config.JwtFields.Key), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
