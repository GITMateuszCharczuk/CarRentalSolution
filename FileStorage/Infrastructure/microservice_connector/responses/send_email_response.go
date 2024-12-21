package microservice_connector

type SendEmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
