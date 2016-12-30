package archive

import (
	"github.com/gorilla/handlers"
	ctrl "github.com/jarlaak/mtg-archive/controllers"
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
	server.SendJSON(w, Alive{Server: "mtg-server", Version: "0.0.0"})
}

func V1AliveHandler(w http.ResponseWriter, r *http.Request) {
	server.SendJSON(w, Alive{Server: "mtg-server", Version: "0.0.0", Api_version: 1})
}

func InitServer() {
	if logger == nil {
		logger = GetLogger()
	}
	server.UseLogger(logger)

	InitializeDatabase()
	db.SetLogger(logger)

	ctrl.Init(db, logger)
	logger.Info("server initialized")
}

func RunServer() {
	InitServer()
	defer db.Close()

	r := server.NewRouter()
	r.HandleFunc("/alive", AliveHandler)

	mtgRouter := r.PathPrefix("/mtg/v1").Subrouter()
	mtgRouter.HandleFunc("/alive", V1AliveHandler)

	recoveryHandler := handlers.RecoveryHandler(handlers.RecoveryLogger(logger),
		handlers.PrintRecoveryStack(true))(r)
	srv := http.Server{
		Handler:        recoveryHandler,
		Addr:           ":8080",
		WriteTimeout:   30 * time.Second,
		ReadTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Fatal(srv.ListenAndServe())
}
