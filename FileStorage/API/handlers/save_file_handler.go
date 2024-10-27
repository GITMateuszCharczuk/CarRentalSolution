// handlers/save_file_handler.go

package handlers

import (
	"file-storage/Application/commands"
	"io"
	"net/http"
)

type SaveFileHandler struct {
	commands *commands.Commands
}

func NewSaveFileHandler(cmds *commands.Commands) *SaveFileHandler {
	return &SaveFileHandler{commands: cmds}
}

func (h *SaveFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fileID := r.FormValue("file_id")
	ownerID := r.FormValue("owner_id")
	fileName := r.FormValue("file_name")

	fileContent, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusBadRequest)
		return
	}

	cmd := h.commands.SaveFile
	cmd.FileID = fileID
	cmd.OwnerID = ownerID
	cmd.FileName = fileName
	cmd.Content = fileContent

	if err := cmd.Execute(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
