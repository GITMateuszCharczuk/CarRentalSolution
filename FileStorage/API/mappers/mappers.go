package mappers

import (
	delete_contract "file-storage/Application.contract/DeleteFile"
	get_contract "file-storage/Application.contract/GetFile"
	save_contract "file-storage/Application.contract/SaveFile"
	delete_command "file-storage/Application/commands/delete_file"
	save_command "file-storage/Application/commands/save_file"
	get_queries "file-storage/Application/queries/get_file"
)

func MapToDeleteFileCommand(req *delete_contract.DeleteFileRequest) delete_command.DeleteFileCommand {
	return delete_command.DeleteFileCommand{
		FileID:  req.FileID,
		OwnerID: req.OwnerID,
	}
}

func MapToSaveFileCommand(req *save_contract.SaveFileRequest) save_command.SaveFileCommand {
	return save_command.SaveFileCommand{
		FileID:   req.FileID,
		OwnerID:  req.OwnerID,
		FileName: req.FileName,
		Content:  req.Content,
	}
}

func MapToGetFileQuery(req *get_contract.GetFileRequest) get_queries.GetFileQuery {
	return get_queries.GetFileQuery{
		FileID:  req.FileID,
		OwnerID: req.OwnerID,
	}
}
