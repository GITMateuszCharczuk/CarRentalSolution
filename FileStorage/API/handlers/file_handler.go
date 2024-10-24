package handlers

import (
    "encoding/json"
    "file-storage/Application/commands"
    "file-storage/Application/queries"
    "net/http"
)

// FileHandler handles file uploads
func FileHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        var cmd commands.SaveFileCommand
        if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }
        if err := commands.SaveFile(cmd); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// GetFileHandler handles file retrieval by ID
func GetFileHandler(w http.ResponseWriter, r *http.Request) {
    fileID := r.URL.Query().Get("fileID")
    ownerID := r.URL.Query().Get("ownerID")
    if fileID == "" || ownerID == "" {
        http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
        return
    }
    query := queries.GetFileQuery{FileID: fileID, OwnerID: ownerID}
    filePath, err := queries.GetFile(query)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write([]byte(filePath))
}

// DeleteFileHandler handles file deletion by ID
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
    fileID := r.URL.Query().Get("fileID")
    ownerID := r.URL.Query().Get("ownerID")
    if fileID == "" || ownerID == "" {
        http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
        return
    }
    query := commands.DeleteFileCommand{FileID: fileID, OwnerID: ownerID}
    if err := commands.DeleteFile(query); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
