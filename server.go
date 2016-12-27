package archive

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/jarlaak/mtg-archive/server"
	"net/http"
	"time"
)

type Alive struct {
	Server      string `json:"server"`
	Version     string `json:"version"`
	Api_version int    `json:"api_version,omitempty"`
}

func AliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Alive{Server: "mtg-server", Version: "0.0.0"})
}

func V1AliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Alive{Server: "mtg-server", Version: "0.0.0", Api_version: 1})
}

func RunServer() {
	if logger == nil {
		logger = GetLogger()
	}
	server.UseLogger(logger)
	r := server.NewRouter()
	r.HandleFunc("/alive", AliveHandler)

	mtgRouter := r.PathPrefix("/mtg/v1").Subrouter()
	mtgRouter.HandleFunc("/alive", V1AliveHandler)
	logger.Info("start server")

	recoveryHandler := handlers.RecoveryHandler(handlers.RecoveryLogger(logger),
		handlers.PrintRecoveryStack(true))(r)
	srv := http.Server{
		Handler:      recoveryHandler,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	logger.Fatal(srv.ListenAndServe())
}
