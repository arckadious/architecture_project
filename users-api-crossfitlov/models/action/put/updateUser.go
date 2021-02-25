package put

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"users-api-crossfitlov/models/auth"
	"users-api-crossfitlov/models/db/conn"
	"users-api-crossfitlov/models/db/dbUsers"
	"users-api-crossfitlov/models/structs/in"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//UpdateUser updates user information
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user in.UserInfos
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

	user.CrossfitlovID, err = strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userExist, err := dbUsers.UpdateUser(db, user)
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
