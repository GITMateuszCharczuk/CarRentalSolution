// package handlers

// import (
// 	"file-storage/Application/commands"
// 	"file-storage/Application/queries"
// 	"io"
// 	"net/http"
// )

// func SaveFileHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		var cmd commands.SaveFileCommand
// 		fileID := r.FormValue("file_id")
// 		ownerID := r.FormValue("owner_id")
// 		fileName := r.FormValue("owner_id")

// 		fileContent, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "Could not read file", http.StatusBadRequest)
// 			return
// 		}

// 		cmd = commands.SaveFileCommand{
// 			FileID:   fileID,
// 			OwnerID:  ownerID,
// 			FileName: fileName,
// 			Content:  fileContent,
// 		}

// 		if err := cmd.Execute(); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusCreated)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func GetFileHandler(w http.ResponseWriter, r *http.Request) {
// 	fileID := r.URL.Query().Get("fileID")
// 	ownerID := r.URL.Query().Get("ownerID")
// 	if fileID == "" || ownerID == "" {
// 		http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
// 		return
// 	}
// 	query := queries.GetFileQuery{FileID: fileID, OwnerID: ownerID}
// 	filePath, err := query.Execute()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write([]byte(filePath.Content))
// }

// func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
// 	fileID := r.URL.Query().Get("fileID")
// 	ownerID := r.URL.Query().Get("ownerID")
// 	if fileID == "" || ownerID == "" {
// 		http.Error(w, "fileID and ownerID are required", http.StatusBadRequest)
// 		return
// 	}
// 	command := commands.DeleteFileCommand{FileID: fileID, OwnerID: ownerID}
// 	if err := command.Execute(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
