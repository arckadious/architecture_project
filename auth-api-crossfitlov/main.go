package main

import (
	"auth-api-crossfitlov/models/action/get"
	"auth-api-crossfitlov/models/action/post"
	"auth-api-crossfitlov/models/action/put"
	"auth-api-crossfitlov/parameters"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

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
	router.Use(mux.CORSMethodMiddleware(router)) //add a access-control-allow-methods if an OPTIONS method is available on the router.HandleFunc().Methods()

	router.Schemes(parameters.Config.Auth.Schemes)

	router.HandleFunc("/v1/signup", put.Signup).Methods("PUT", "OPTIONS")

	router.HandleFunc("/v1/signin", post.Signin).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/check", post.Check).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/refresh", post.Refresh).Methods("POST", "OPTIONS")

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
