package microservice_connector

import (
	models "email-service/Domain/models/external"
	interfaces "email-service/Domain/service_interfaces"
	mappers "email-service/Infrastructure/microservice_connector/mappers"
	responses "email-service/Infrastructure/microservice_connector/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type MicroserviceConnectorImpl struct {
	IdentityApiBaseUrl string
	httpClient         *http.Client
}

func NewMicroserviceConnectorImpl(IdentityApiBaseUrl string) interfaces.MicroserviceConnector {
	return &MicroserviceConnectorImpl{
		IdentityApiBaseUrl: IdentityApiBaseUrl,
		httpClient:         &http.Client{},
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
