package microservice_connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	models "rental-api/Domain/models/external"
	interfaces "rental-api/Domain/service_interfaces"
	mappers "rental-api/Infrastructure/microservice_connector/mappers"
	responses "rental-api/Infrastructure/microservice_connector/responses"
)

type MicroserviceConnectorImpl struct {
	IdentityApiBaseUrl  string
	EmailServiceBaseUrl string
	httpClient          *http.Client
}

func NewMicroserviceConnectorImpl(IdentityApiBaseUrl string, EmailServiceBaseUrl string) interfaces.MicroserviceConnector {
	return &MicroserviceConnectorImpl{
		IdentityApiBaseUrl:  IdentityApiBaseUrl,
		EmailServiceBaseUrl: EmailServiceBaseUrl,
		httpClient:          &http.Client{},
	}
}

func (df *MicroserviceConnectorImpl) ValidateToken(token models.JwtToken) (*models.TokenInfo, error) {
	endpoint := fmt.Sprintf("%s/identity-api/api/token/validate", df.IdentityApiBaseUrl)

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Add("token", token.Token)
	u.RawQuery = q.Encode()

	resp, err := df.httpClient.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var result responses.ValidateTokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		if result.Success {
			tokenInfo := mappers.MapToTokenInfo(result)
			return &tokenInfo, nil
		}
		return nil, fmt.Errorf("token not valid: %d", resp.StatusCode)
	default:
		return nil, fmt.Errorf("token not valid: %d", resp.StatusCode)
	}
}

func (df *MicroserviceConnectorImpl) GetUserInternalInfo(token models.JwtToken) (*models.UserInfo, error) {
	endpoint := fmt.Sprintf("%s/identity-api/api/user/internal", df.IdentityApiBaseUrl)

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Add("token", token.Token)
	u.RawQuery = q.Encode()

	resp, err := df.httpClient.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var result responses.GetUserInternalResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		if result.Success {
			userInfo := mappers.MapToUserInfo(result)
			return &userInfo, nil
		}
		return nil, fmt.Errorf("user not found: %d", resp.StatusCode)
	default:
		return nil, fmt.Errorf("failed to fetch user info: %d", resp.StatusCode)
	}
}

func (df *MicroserviceConnectorImpl) SendEmail(email models.InternalEmail) error {
	endpoint := fmt.Sprintf("%s/email-service/api/send-internal-email", df.EmailServiceBaseUrl)

	jsonBody, err := json.Marshal(email)
	if err != nil {
		return fmt.Errorf("failed to marshal email request: %w", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := df.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse responses.SendEmailResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("email service returned error status: %d", resp.StatusCode)
		}
		if !errorResponse.Success {
			return fmt.Errorf("email service error: %s", errorResponse.Message)
		}
	}
	return nil
}
