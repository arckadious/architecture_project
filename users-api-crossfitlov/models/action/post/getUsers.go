package post

import (
	"encoding/json"
	"net/http"
	"users-api-crossfitlov/models/auth"
	"users-api-crossfitlov/models/db/conn"
	"users-api-crossfitlov/models/db/dbUsers"

	"github.com/sirupsen/logrus"
)

//GetUsers gets users by criteria
func GetUsers(w http.ResponseWriter, r *http.Request) {

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

	usersInfosList, err := dbUsers.SelectUsers(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	json.NewEncoder(w).Encode(usersInfosList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
