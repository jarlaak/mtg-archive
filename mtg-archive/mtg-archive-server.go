package main

import (
        "github.com/gorilla/mux"
        "fmt"
        "net/http"
        "time"
        "github.com/jarlaak/mtg-archive"
)

var logger = archive.GetLogger()

func AliveHandler(w http.ResponseWriter, r* http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"server\": \"mtg-server\", \"version\":\"0.0.0\"}");
    logger.Info("/alive call")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/alive", AliveHandler)

    logger.Info("start server")

    srv := http.Server{
        Handler: r,
        Addr: ":8080",
        WriteTimeout: 30 * time.Second,
        ReadTimeout: 30 * time.Second,
    }
    srv.ListenAndServe()
}
