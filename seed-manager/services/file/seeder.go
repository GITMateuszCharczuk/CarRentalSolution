package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	c "seeder-manager/api"
	"seeder-manager/config"
	"seeder-manager/reference_store"
	"seeder-manager/services/file/factories"
)

type FileSeeder struct {
	fileFactory *factories.FileFactory
	refStore    *reference_store.InMemoryStore
	apiBaseUrl  string
	token       string
	BlogImages  int
	CarImages   int
}

type SaveFileResponse struct {
	FileID string `json:"id"`
}

func NewFileSeeder(refStore *reference_store.InMemoryStore, cfg *config.Config) *FileSeeder {
	return &FileSeeder{
		fileFactory: factories.NewFileFactory(),
		refStore:    refStore,
		apiBaseUrl:  cfg.FileServiceURL,
		token:       cfg.JWTToken,
		BlogImages:  cfg.SeedCount.BlogImages,
		CarImages:   cfg.SeedCount.CarImages,
	}
}

func (s *FileSeeder) uploadFile(fileHeader *multipart.FileHeader) (string, error) {
	// Create a buffer to store the form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form file field
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	// Get the content from the header
	content := fileHeader.Header.Get("X-Content")
	if content == "" {
		return "", fmt.Errorf("no content found in file header")
	}

	// Write the content to the form field
	if _, err = io.Copy(part, bytes.NewReader([]byte(content))); err != nil {
		return "", fmt.Errorf("failed to copy file content: %v", err)
	}

	// Close the multipart writer
	if err = writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Create the request
	url := s.apiBaseUrl + "/file-storage/api/files"
	clientInst := c.NewAPIClient("")
	url = clientInst.AddTokenToURL(url, s.token)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse the response
	var response SaveFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return response.FileID, nil
}

func (s *FileSeeder) SeedCarImages(count int) error {
	log.Println("Seeding car images")
	for i := 0; i < count; i++ {
		// Generate a car image
		fileHeader, err := s.fileFactory.CreateCarImage()
		if err != nil {
			return fmt.Errorf("failed to create car image: %v", err)
		}

		// Upload the image
		fileID, err := s.uploadFile(fileHeader)
		if err != nil {
			return fmt.Errorf("failed to upload car image: %v", err)
		}

		// Store the file ID
		s.refStore.AddCarFileID(fileID)
		log.Printf("Created car image %d/%d with ID: %s", i+1, count, fileID)
	}
	log.Println("Seeding car images done")
	return nil
}

func (s *FileSeeder) SeedBlogImages(count int) error {
	log.Println("Seeding blog images")
	for i := 0; i < count; i++ {
		// Generate a blog image
		fileHeader, err := s.fileFactory.CreateBlogImage()
		if err != nil {
			return fmt.Errorf("failed to create blog image: %v", err)
		}

		// Upload the image
		fileID, err := s.uploadFile(fileHeader)
		if err != nil {
			return fmt.Errorf("failed to upload blog image: %v", err)
		}

		// Store the file ID
		s.refStore.AddBlogPostImageID(fileID)
		log.Printf("Created blog image %d/%d with ID: %s", i+1, count, fileID)
	}
	log.Println("Seeding blog images done")
	return nil
}

func (s *FileSeeder) Seed(refStore *reference_store.InMemoryStore, token string) error {
	log.Println("Seeding file service")
	err := s.SeedCarImages(s.CarImages)
	if err != nil {
		return fmt.Errorf("failed to seed car images: %v", err)
	}
	err = s.SeedBlogImages(s.BlogImages)
	if err != nil {
		return fmt.Errorf("failed to seed blog images: %v", err)
	}
	return nil
}

func (s *FileSeeder) Cleanup() error {
	return nil
}
