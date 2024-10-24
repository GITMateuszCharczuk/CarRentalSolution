package routes

import (
    "file-storage/API/handlers"
    "net/http"
)

// RegisterRoutes registers the API routes and their handlers
func RegisterRoutes() {
    http.HandleFunc("/files", handlers.FileHandler)
    http.HandleFunc("/files/delete", handlers.DeleteFileHandler)
    http.HandleFunc("/files/get", handlers.GetFileHandler)
}
