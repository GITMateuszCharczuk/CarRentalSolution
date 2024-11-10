// handlers/save_file_handler.go

package controllers

import (
	"encoding/json"
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/SaveFile"
	command "file-storage/Application/commands/save_file"
	"io"
	"net/http"
)

type SaveFileController struct {
	commandHandler *command.SaveFileCommandHandler
}

func NewSaveFileController(cmd *command.SaveFileCommandHandler) *SaveFileController {
	return &SaveFileController{commandHandler: cmd}
}

func (h *SaveFileController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req contract.SaveFileRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	command := mappers.MapToSaveFileCommand(&req)
	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *SaveFileController) Route() string {
	return "/files"
}
