package microservice_connector

type ValidateTokenResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Valid   bool     `json:"valid"`
	Roles   []string `json:"roles"`
}
