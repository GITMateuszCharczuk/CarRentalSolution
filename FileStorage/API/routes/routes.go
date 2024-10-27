// routes/router.go

package routes

import (
	"file-storage/API/handlers"
	"net/http"
)

type Router struct {
	fileHandler *handlers.FileHandler
}

func NewRouter(fileHandler *handlers.FileHandler) *Router {
	return &Router{fileHandler: fileHandler}
}

func (r *Router) RegisterRoutes() {
	http.HandleFunc("/files", r.fileHandler.SaveFileHandler)
	http.HandleFunc("/files/delete", r.fileHandler.DeleteFileHandler)
	http.HandleFunc("/files/get", r.fileHandler.GetFileHandler)
}
