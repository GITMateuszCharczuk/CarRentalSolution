package mappers

import (
	"encoding/base64"
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

	file := models.File{}

	if id, ok := fileData["ID"].(string); ok {
		file.ID = id
	} else {
		return models.File{}, fmt.Errorf("invalid type for ID")
	}

	if ownerID, ok := fileData["OwnerID"].(string); ok {
		file.OwnerID = ownerID
	} else {
		return models.File{}, fmt.Errorf("invalid type for OwnerID")
	}

	if fileName, ok := fileData["FileName"].(string); ok {
		file.FileName = fileName
	} else {
		return models.File{}, fmt.Errorf("invalid type for FileName")
	}

	if contentStr, ok := fileData["Content"].(string); ok {
		contentBytes, err := base64.StdEncoding.DecodeString(contentStr)
		if err != nil {
			return models.File{}, fmt.Errorf("failed to decode Content: %w", err)
		}
		file.Content = contentBytes
	} else {
		return models.File{}, fmt.Errorf("invalid type for Content")
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
