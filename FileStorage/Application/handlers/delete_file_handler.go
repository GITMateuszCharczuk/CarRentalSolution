// handlers/delete_file_handler.go

package handlers

import (
	"file-storage/Application/commands"
	"net/http"
)

type DeleteFileHandler struct {
	command *commands.DeleteFileCommand
}

func NewDeleteFileHandler(cmd *commands.DeleteFileCommand) *DeleteFileHandler {
	return &DeleteFileHandler{command: cmd}
}

func (h *DeleteFileHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("fileID")
	ownerID := r.URL.Query().Get("ownerID")
	if fileID == "" || ownerID == "" {
		http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
		return
	}

	cmd := h.command
	cmd.FileID = fileID
	cmd.OwnerID = ownerID

	if err := cmd.Execute(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DeleteFileHandler) Route() string {
	return "/files/delete"
}
