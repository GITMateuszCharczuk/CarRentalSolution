package jwt_token_service

import (
	"fmt"
	"time"

	"identity-api/Domain/constants"
	models "identity-api/Domain/models/token"
	repository_interfaces "identity-api/Domain/repository_interfaces/refresh_token_repository"
	user_repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenServiceImpl struct {
	commandRepo     repository_interfaces.RefreshTokenCommandRepository
	queryRepo       repository_interfaces.RefreshTokenQueryRepository
	userQueryRepo   user_repository_interfaces.UserQueryRepository
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SecretKey       []byte
}

func NewJWTTokenService(accessTokenTTL,
	refreshTokenTTL time.Duration,
	secretKey string,
	commandRepo repository_interfaces.RefreshTokenCommandRepository,
	queryRepo repository_interfaces.RefreshTokenQueryRepository,
	userQueryRepo user_repository_interfaces.UserQueryRepository) *JWTTokenServiceImpl {
	return &JWTTokenServiceImpl{
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
		commandRepo:     commandRepo,
		queryRepo:       queryRepo,
		SecretKey:       []byte(secretKey),
		userQueryRepo:   userQueryRepo,
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

	err = s.commandRepo.SaveRefreshToken(userID, string(refreshToken), int(s.RefreshTokenTTL.Hours()))
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
	refreshToken := fmt.Sprintf("%s:%d", userID, time.Now().UnixNano())
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
	userID, err := s.queryRepo.GetRefreshToken(string(refreshToken.RefreshToken))
	if err != nil {
		return models.NewJwtToken(""), fmt.Errorf("invalid refresh token: %v", err)
	}

	// Retrieve user claims (roles) from a user repository or service
	user, err := s.userQueryRepo.GetUserByID(userID)
	if err != nil {
		return models.NewJwtToken(""), fmt.Errorf("could not retrieve user roles: %v", err)
	}

	newAccessToken, _, err := s.GenerateTokens(userID, user.Roles)
	if err != nil {
		return models.NewJwtToken(""), fmt.Errorf("could not generate new access token: %v", err)
	}

	return newAccessToken, nil
}

func (s *JWTTokenServiceImpl) RevokeToken(token models.JwtRefreshToken) error {
	err := s.commandRepo.RevokeRefreshToken(string(token.RefreshToken))
	if err != nil {
		return fmt.Errorf("could not revoke refresh token: %v", err)
	}
	return nil
}
