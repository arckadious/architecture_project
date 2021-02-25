package get

import (
	"encoding/json"
	"net/http"
	"users-api-crossfitlov/models/auth"
	"users-api-crossfitlov/models/db/conn"
	"users-api-crossfitlov/models/db/dbUsers"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//GetUser gets a user by crossfitlovID
func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

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

	usersInfos, userExist, err := dbUsers.SelectUser(db, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	//si pas d'utilisateur, on retourne 204 pour indiquer qu'on a pas trouvé d'utilisateur avec cet ID
	if !userExist {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(usersInfos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
