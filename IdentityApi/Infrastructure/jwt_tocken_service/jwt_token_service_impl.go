package jwt_token_service

import (
	"fmt"
	"time"

	"identity-api/Domain/constants"
	"identity-api/Domain/models"
	"identity-api/Domain/repository_interfaces"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenServiceImpl struct {
	refreshTokenRepo repository_interfaces.RefreshTokenRepository
	AccessTokenTTL   time.Duration
	RefreshTokenTTL  time.Duration
	SecretKey        []byte
}

func NewJWTTokenService(accessTokenTTL, refreshTokenTTL time.Duration, secretKey string, refreshTokenRepo repository_interfaces.RefreshTokenRepository) *JWTTokenServiceImpl {
	return &JWTTokenServiceImpl{
		AccessTokenTTL:   accessTokenTTL,
		RefreshTokenTTL:  refreshTokenTTL,
		refreshTokenRepo: refreshTokenRepo,
		SecretKey:        []byte(secretKey),
	}
}

func (s *JWTTokenServiceImpl) GenerateTokens(userID string, roles []constants.JWTRole) (models.JwtToken, models.JwtRefreshToken, error) {
	accessToken, err := s.createAccessToken(userID, roles)
	if err != nil {
		return models.NewJwtToken(""), models.NewRefreshToken(""), fmt.Errorf("could not generate access token: %v", err)
	}

	refreshToken, err := s.createRefreshToken(userID)
	if err != nil {
		return models.NewJwtToken(""), models.NewRefreshToken(""), fmt.Errorf("could not generate refresh token: %v", err)
	}

	err = s.refreshTokenRepo.SaveRefreshToken(userID, string(refreshToken), int(s.RefreshTokenTTL.Hours()))
	if err != nil {
		return models.NewJwtToken(""), models.NewRefreshToken(""), fmt.Errorf("could not save refresh token: %v", err)
	}

	return models.NewJwtToken(accessToken), models.NewRefreshToken(refreshToken), nil
}

func (s *JWTTokenServiceImpl) createAccessToken(userID string, roles []constants.JWTRole) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"roles": roles,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(s.AccessTokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *JWTTokenServiceImpl) createRefreshToken(userID string) (string, error) {
	refreshToken := fmt.Sprintf("%s:%d", userID, time.Now().UnixNano()) //random number
	return refreshToken, nil
}

func (s *JWTTokenServiceImpl) ValidateToken(token models.JwtToken) (string, []constants.JWTRole, error) {
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(string(token.Token), claims, func(token *jwt.Token) (interface{}, error) {
		return s.SecretKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return "", nil, fmt.Errorf("invalid token: %v", err)
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", nil, fmt.Errorf("invalid token claims: missing user ID")
	}

	roles, ok := claims["roles"].([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid token claims: missing roles")
	}

	var jwtRoles []constants.JWTRole
	for _, role := range roles {
		if r, ok := role.(string); ok {
			jwtRoles = append(jwtRoles, constants.JWTRole(r))
		}
	}

	return userID, jwtRoles, nil
}

func (s *JWTTokenServiceImpl) RefreshToken(refreshToken models.JwtRefreshToken) (models.JwtToken, error) {
	userID, err := s.refreshTokenRepo.GetRefreshToken(string(refreshToken.RefreshToken))
	if err != nil {
		return models.NewJwtToken(""), fmt.Errorf("invalid refresh token: %v", err)
	}

	// // Retrieve user claims (roles) from a user repository or service
	// roles, err := s.getUserRoles(userID) // Implement this method to fetch roles
	// if err != nil {
	// 	return models.NewJwtToken(""), fmt.Errorf("could not retrieve user roles: %v", err)
	// }

	newAccessToken, _, err := s.GenerateTokens(userID, []constants.JWTRole{constants.User})
	if err != nil {
		return models.NewJwtToken(""), fmt.Errorf("could not generate new access token: %v", err)
	}

	return newAccessToken, nil
}

func (s *JWTTokenServiceImpl) RevokeToken(token models.JwtRefreshToken) error {
	err := s.refreshTokenRepo.RevokeRefreshToken(string(token.RefreshToken))
	if err != nil {
		return fmt.Errorf("could not revoke refresh token: %v", err)
	}
	return nil
}
