package datafetcher

import (
	"encoding/json"
	"fmt"
	models "identity-api/Domain/models/external"
	interfaces "identity-api/Infrastructure/data_fetcher/interfaces"
	mappers "identity-api/Infrastructure/data_fetcher/mappers"
	responses "identity-api/Infrastructure/data_fetcher/responses"
	"net/http"
	"net/url"
)

type DataFetcherImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDataFetcherImpl(baseURL string) interfaces.DataFetcher {
	return &DataFetcherImpl{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (df *DataFetcherImpl) ValidateToken(token models.JwtToken) (*models.TokenInfo, error) {
	endpoint := fmt.Sprintf("%s/identity-api/api/token/validate", df.baseURL)

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

func (df *DataFetcherImpl) GetUserInternalInfo(token models.JwtToken) (*models.UserInfo, error) {
	endpoint := fmt.Sprintf("%s/identity-api/api/user/internal", df.baseURL)

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
