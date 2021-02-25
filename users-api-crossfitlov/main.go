package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"users-api-crossfitlov/models/action/delete"
	"users-api-crossfitlov/models/action/get"
	"users-api-crossfitlov/models/action/post"
	"users-api-crossfitlov/models/action/put"
	"users-api-crossfitlov/parameters"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// init chargement config
func init() {

	p, _ := os.Getwd()

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: binary -config=[config]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.String("config", "", "JSON config file")
	dbg := flag.String("debug", "", "Trace, Debug, Info, Warn, Error, Fatal, Panic")

	flag.Parse()

	configDir := flag.Lookup("config")
	if configDir.Value.String() == configDir.DefValue {
		flag.Set("config", filepath.Join(p, "parameters", "parameters.json"))
	}

	strLevel := *dbg
	if strLevel == "" {
		strLevel = "Warning"
	}
	level, _ := log.ParseLevel(strLevel)

	//Initialiser les paramètres
	parameters.LoadConfiguration(flag.Lookup("config").Value.String())

	log.SetLevel(level)
	log.SetReportCaller(true)

}

func newServer() http.Handler {

	router := mux.NewRouter().StrictSlash(true)

	for _, host := range parameters.Config.Auth.Hosts {

		router.Host(host).Subrouter()

		router.Schemes(parameters.Config.Auth.Schemes)

		//CRUD interface
		router.HandleFunc("/v1/users", post.CreateUser).Methods("POST") //Create a new user
		//router.HandleFunc("/v1/users", post.GetUsers).Methods("POST")            //get users by criteria
		router.HandleFunc("/v1/users/{id}", get.GetUser).Methods("GET")          //get one user information
		router.HandleFunc("/v1/users/{id}", put.UpdateUser).Methods("PUT")       //update user information
		router.HandleFunc("/v1/users/{id}", delete.DeleteUser).Methods("DELETE") //delete user information

		router.HandleFunc("/v1/ping", get.Ping).Methods("GET")

	}

	return router
}

func main() {

	log.Info("Service running...")

	log.Fatal(http.ListenAndServe(":"+parameters.Config.Auth.Port, newServer()))

}
