// handlers/delete_file_handler.go

package handlers

import (
	"file-storage/Application/commands"
	"net/http"
)

type DeleteFileHandler struct {
	commands *commands.Commands
}

func NewDeleteFileHandler(cmds *commands.Commands) *DeleteFileHandler {
	return &DeleteFileHandler{commands: cmds}
}

func (h *DeleteFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("fileID")
	ownerID := r.URL.Query().Get("ownerID")
	if fileID == "" || ownerID == "" {
		http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
		return
	}

	cmd := h.commands.DeleteFile
	cmd.FileID = fileID
	cmd.OwnerID = ownerID

	if err := cmd.Execute(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
