package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

type APIResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *APIClient) AddTokenToURL(baseURL, token string) string {
	if token == "" {
		return baseURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return baseURL
	}

	q := u.Query()
	q.Set("token", token)
	u.RawQuery = q.Encode()

	return u.String()
}

func (c *APIClient) handleResponse(resp *http.Response, operation string) (*APIResponse, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		log.Printf("Authorization failed for %s. Status: %d, Body: %s", operation, resp.StatusCode, string(body))
		return nil, fmt.Errorf("unauthorized: %s failed with status 401", operation)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Printf("Error unmarshaling response: %v, Body: %s", err, string(body))
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("%s failed with status: %d", operation, resp.StatusCode)
		}
		// If we can't unmarshal but status is OK, return empty response
		return &APIResponse{Success: true}, nil
	}

	if resp.StatusCode >= 400 {
		log.Printf("Request failed for %s. Status: %d, Body: %s", operation, resp.StatusCode, string(body))
		return &apiResp, fmt.Errorf("%s failed with status: %d", operation, resp.StatusCode)
	}

	return &apiResp, nil
}

func (c *APIClient) Post(endpoint string, data interface{}, token string) (*APIResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	url = c.AddTokenToURL(url, token)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp, "POST request")
}

func (c *APIClient) UploadFile(endpoint, filename, mimeType string, content []byte, token string) (*APIResponse, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	fileField, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("error creating form file: %w", err)
	}

	if _, err := io.Copy(fileField, bytes.NewReader(content)); err != nil {
		return nil, fmt.Errorf("error writing file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("error closing multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	url = c.AddTokenToURL(url, token)

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp, "file upload")
}
