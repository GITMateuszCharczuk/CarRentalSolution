package controllers

import (
	"encoding/json"
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/GetFile"
	queries "file-storage/Application/queries/get_file"
	"net/http"
)

type GetFileController struct {
	queryHandler *queries.GetFileQueryHandler
}

func NewGetFileController(qrs *queries.GetFileQueryHandler) *GetFileController {
	return &GetFileController{queryHandler: qrs}
}

func (h *GetFileController) Handle(w http.ResponseWriter, r *http.Request) {
	var req contract.GetFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	query := mappers.MapToGetFileQuery(&req)
	resp, err := h.queryHandler.Execute(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp.File))
}

func (h *GetFileController) Route() string {
	return "/files/get"
}
