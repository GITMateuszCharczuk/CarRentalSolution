package mappers

import (
	"errors"
	"file-storage/Domain/models"
	"fmt"
)

var (
	ErrInvalidDataType = errors.New("invalid data type")
)

func MapToFile(data interface{}) (models.File, error) {
	fileData, ok := data.(map[string]interface{})
	if !ok {
		return models.File{}, fmt.Errorf("%w: expected map[string]interface{}", ErrInvalidDataType)
	}

	file := models.File{
		ID:       fileData["ID"].(string),
		OwnerID:  fileData["OwnerID"].(string),
		FileName: fileData["FileName"].(string),
		Content:  fileData["Content"].([]byte),
	}

	return file, nil
}

func MapToFileID(data interface{}) (string, error) {
	fileID, ok := data.(string)
	if !ok {
		return "", fmt.Errorf("%w: expected string", ErrInvalidDataType)
	}

	return fileID, nil
}
