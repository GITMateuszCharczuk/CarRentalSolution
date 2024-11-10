package controllers

import (
	"encoding/json"
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/DeleteFile"
	commands "file-storage/Application/commands/delete_file"
	"net/http"
)

type DeleteFileController struct {
	commandHandler *commands.DeleteFileCommandHandler
}

func NewDeleteFileController(cmd *commands.DeleteFileCommandHandler) *DeleteFileController {
	return &DeleteFileController{commandHandler: cmd}
}

func (h *DeleteFileController) Handle(w http.ResponseWriter, r *http.Request) {
	var req contract.DeleteFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	command := mappers.MapToDeleteFileCommand(&req)
	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *DeleteFileController) Route() string {
	return "/files/delete"
}
