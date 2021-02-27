package post

import (
	"auth-api-crossfitlov/models/auth"
	"auth-api-crossfitlov/models/structs/in"
	"auth-api-crossfitlov/models/structs/out"
	"auth-api-crossfitlov/parameters"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

//Refresh renew the token if the remaining time before expiration is less than 30s.
func Refresh(w http.ResponseWriter, r *http.Request) {

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
	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(time.Duration(parameters.Config.JwtFields.ExpirationTime))
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(parameters.Config.JwtFields.Key))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	err = json.NewEncoder(w).Encode(&out.TokenData{
		TokenInfos: out.TokenInfos{
			Name:      "token-CL",
			Value:     tokenString,
			ExpiresAt: expirationTime.Format("2006-01-02 15:04:05"),
		},
		UserInfos: nil,
	})
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}
}
