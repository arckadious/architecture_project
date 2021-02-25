package delete

import (
	"net/http"
	"users-api-crossfitlov/models/auth"
	"users-api-crossfitlov/models/db/conn"
	"users-api-crossfitlov/models/db/dbUsers"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//DeleteUser deletes user informations
func DeleteUser(w http.ResponseWriter, r *http.Request) {

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

	userExist, err := dbUsers.DeleteUser(db, vars["id"])
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

	w.WriteHeader(http.StatusOK)
}
