package archive

import (
        "github.com/gorilla/mux"
        "fmt"
        "net/http"
        "time"
)

func AliveHandler(w http.ResponseWriter, r* http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"server\": \"mtg-server\", \"version\":\"0.0.0\"}");
    logger.Info("/alive call")
}

func V1AliveHandler(w http.ResponseWriter, r* http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"server\": \"mtg-server\", \"api-version\": \"1\", \"version\":\"0.0.0\"}");
    logger.Info("/mtg/v1/alive call")
}

func RunServer() {
    r := mux.NewRouter()
    r.HandleFunc("/alive", AliveHandler)

    mtgRouter := r.PathPrefix("/mtg/v1").Subrouter()
    mtgRouter.HandleFunc("/alive",V1AliveHandler)
    logger.Info("start server")

    srv := http.Server{
        Handler: r,
        Addr: ":8080",
        WriteTimeout: 30 * time.Second,
        ReadTimeout: 30 * time.Second,
    }
    srv.ListenAndServe()
}
