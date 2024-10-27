// handlers/get_file_handler.go

package handlers

import (
	"file-storage/Application/queries"
	"net/http"
)

type GetFileHandler struct {
	queries *queries.Queries
}

func NewGetFileHandler(qrs *queries.Queries) *GetFileHandler {
	return &GetFileHandler{queries: qrs}
}

func (h *GetFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("fileID")
	ownerID := r.URL.Query().Get("ownerID")
	if fileID == "" || ownerID == "" {
		http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
		return
	}

	query := h.queries.GetFile
	query.FileID = fileID
	query.OwnerID = ownerID

	filePath, err := query.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(filePath.Content))
}
