package main

import (
    "log"
    "net/http"

    "file-storage/API/routes"
    "file-storage/Infrastructure/config"
    "file-storage/Infrastructure/db"
    "file-storage/Infrastructure/queue"
)

func main() {
    config.LoadConfig()

    db.ConnectMongo()

    queue.InitializeJetStream()

    routes.RegisterRoutes()

    log.Println("Starting server on :8080")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}
