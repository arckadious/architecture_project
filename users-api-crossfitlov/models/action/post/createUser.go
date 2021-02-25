package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"users-api-crossfitlov/models/auth"
	"users-api-crossfitlov/models/db/conn"
	"users-api-crossfitlov/models/db/dbUsers"
	"users-api-crossfitlov/models/structs/in"

	"github.com/sirupsen/logrus"
)

//CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user in.UserInfos

	w.Header().Set("Content-Type", "application/json")

	//Vérifier les droits
	if !auth.CheckAuth(w, r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db, err := conn.GetConn()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		logrus.Error(err)
		return
	}

	defer db.Close()

	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dbUsers.InsertUser(db, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
