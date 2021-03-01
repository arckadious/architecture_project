package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
	"users-api-crossfitlov/models/action/delete"
	"users-api-crossfitlov/models/action/get"
	"users-api-crossfitlov/models/action/post"
	"users-api-crossfitlov/models/action/put"
	"users-api-crossfitlov/parameters"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	// Timeout delay and graceful shutdown deadline
	Timeout = time.Second * 15
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

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "X-Accept-Charset,X-Accept,Content-Type,Authorization,Cache-Control,X-Http-Method-Override")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,GET,DELETE,OPTIONS")
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Expires", "-1")
		w.Header().Set("Keep-Alive", "timeout=5, max=99")
		w.Header().Set("Pragma", "no-cache")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func newRouter() http.Handler {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(CORS)

	router.Schemes(parameters.Config.Auth.Schemes)

	//CRUD interface
	router.HandleFunc("/v1/users", post.CreateUser).Methods("POST", "OPTIONS")          //Create a new user
	router.HandleFunc("/v1/users/get", post.GetUsers).Methods("POST", "OPTIONS")        //get users by criteria
	router.HandleFunc("/v1/users/{id}", get.GetUser).Methods("GET", "OPTIONS")          //get one user information
	router.HandleFunc("/v1/users/{id}", put.UpdateUser).Methods("PUT", "OPTIONS")       //update user information
	router.HandleFunc("/v1/users/{id}", delete.DeleteUser).Methods("DELETE", "OPTIONS") //delete user information

	router.HandleFunc("/v1/ping", get.Ping).Methods("GET")

	return router
}

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:" + parameters.Config.Auth.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: Timeout,
		ReadTimeout:  Timeout,
		IdleTimeout:  Timeout * 4,
		Handler:      newRouter(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Process signals channel
	sigChannel := make(chan os.Signal, 1)

	// Graceful shutdown via SIGINT
	signal.Notify(sigChannel, os.Interrupt)

	log.Info("Service running...")
	<-sigChannel // Block until SIGINT received

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	server.Shutdown(ctx)

	log.Info("Service shutdown")

}
