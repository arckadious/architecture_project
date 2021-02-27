package post

import (
	"auth-api-crossfitlov/models/auth"
	"auth-api-crossfitlov/models/client"
	"auth-api-crossfitlov/models/db/conn"
	"auth-api-crossfitlov/models/db/dbauth"
	"auth-api-crossfitlov/models/structs/in"
	"auth-api-crossfitlov/models/structs/out"
	"auth-api-crossfitlov/parameters"
	"encoding/json"
	"net/http"
	"time"

	// import the jwt-go library
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	ID string
	jwt.StandardClaims
}

// Signin Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//Vérifier les droits
	if !auth.CheckAuth(w, r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//connexion à la base de donnée du microservice
	db, err := conn.GetConn()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		logrus.Error(err)
		return
	}
	defer db.Close()

	var creds in.CredentialsData
	// Get the JSON body and decode into credentials
	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ID, encryptedExpectedPasswd, userExist, err := dbauth.GetPasswdAndID(creds.Login, db)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	//Check password correct and user exist
	if !userExist || !comparePasswords(encryptedExpectedPasswd, creds.Password) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(time.Duration(parameters.Config.JwtFields.ExpirationTime) * time.Second)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		ID: ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(parameters.Config.JwtFields.Key))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	var userInfos interface{}
	if creds.GetInfos {
		userInfos, err = client.GetUserInfos(ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}

	} else {
		userInfos = nil
	}

	err = json.NewEncoder(w).Encode(&out.TokenData{
		TokenInfos: out.TokenInfos{
			Name:      "token-CL",
			Value:     tokenString,
			ExpiresAt: expirationTime.Format("2006-01-02 15:04:05"),
		},
		UserInfos: userInfos,
	})
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

}

func comparePasswords(hashedPwd string, credPwd string) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(credPwd))
	if err != nil {
		return false
	}

	return true
}
