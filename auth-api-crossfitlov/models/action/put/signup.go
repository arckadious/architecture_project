package put

import (
	"auth-api-crossfitlov/models/auth"
	"auth-api-crossfitlov/models/client"
	"auth-api-crossfitlov/models/db/conn"
	"auth-api-crossfitlov/models/db/dbauth"
	"auth-api-crossfitlov/models/structs/out"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//Signup add a new crossfitlov user
func Signup(w http.ResponseWriter, r *http.Request) {

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

	var userRegisterInfos out.UserRegisterInfos
	// Get the JSON body and decode into user infos
	err = json.NewDecoder(r.Body).Decode(&userRegisterInfos)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//GENERER CROSSFITLOVID en faisant un auto-incrément dans db auth, puis voir si on peut récupérer le numéro de ligne inséré
	lastInsertedID, err := dbauth.GetLastInsertID(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	crossfitLovID := lastInsertedID + 1

	hashpasswd, err := bcrypt.GenerateFromPassword([]byte(userRegisterInfos.Credentials.Password), bcrypt.MinCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	//insert en db des credentials
	err = dbauth.CreatePasswdAndID(db, strconv.Itoa(crossfitLovID), userRegisterInfos.Credentials.Login, string(hashpasswd))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	//SI OK create user
	userRegisterInfos.UserInfos.CrossfitlovID = crossfitLovID
	err = client.CreateUserInfos(userRegisterInfos.UserInfos)
	if err != nil {
		dbauth.DeletePasswdAndID(userRegisterInfos.Credentials.Login, db)
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
