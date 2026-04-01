package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	apiError "github.com/gonotes/api/error"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

type Claims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

type Service struct {
	userRepo  UserRepository
	tokenRepo RefreshTokenRepository
	jwtSecret []byte
}

func NewService(userRepo UserRepository, tokenRepo RefreshTokenRepository, jwtSecret []byte) *Service {
	return &Service{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) generateAccessToken(userID string) (string, time.Time, error) {
	expiresAt := time.Now().Add(AccessTokenExpiry)

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expiresAt, nil
}

func (s *Service) generateRefreshToken(userID string) (*RefreshToken, error) {
	tokenStr := uuid.New().String()
	refreshToken := NewRefreshToken(userID, tokenStr, time.Now().Add(RefreshTokenExpiry))

	err := s.tokenRepo.Save(*refreshToken)
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func (s *Service) generateAuthResponse(userID string) (*AuthResponseDTO, *apiError.ErrorDTO) {
	accessToken, expiresAt, err := s.generateAccessToken(userID)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to generate access token.")
	}

	refreshToken, err := s.generateRefreshToken(userID)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to generate refresh token.")
	}

	return &AuthResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Token,
		ExpiresAt:    expiresAt,
	}, nil
}

func (s *Service) Register(registerRequestDTO RegisterRequestDTO) (*UserResponseDTO, *apiError.ErrorDTO) {
	existingUser, err := s.userRepo.GetByEmail(registerRequestDTO.Email)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to register user.")
	}

	if existingUser != nil {
		return nil, apiError.NewBadRequestError("Email already exists.")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequestDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to register user.")
	}

	newUser := NewUser(registerRequestDTO.Email, string(passwordHash))

	err = s.userRepo.Save(*newUser)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to register user.")
	}

	return &UserResponseDTO{
		ID:        newUser.ID,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (s *Service) Login(loginRequestDTO LoginRequestDTO) (*AuthResponseDTO, *apiError.ErrorDTO) {
	user, err := s.userRepo.GetByEmail(loginRequestDTO.Email)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to login.")
	}

	if user == nil {
		return nil, apiError.NewUnauthorizedError("Invalid email or password.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequestDTO.Password))
	if err != nil {
		return nil, apiError.NewUnauthorizedError("Invalid email or password.")
	}

	return s.generateAuthResponse(user.ID)
}

func (s *Service) Refresh(refreshTokenRequestDTO RefreshTokenRequestDTO) (*AuthResponseDTO, *apiError.ErrorDTO) {
	storedToken, err := s.tokenRepo.GetByToken(refreshTokenRequestDTO.RefreshToken)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to refresh token.")
	}

	if storedToken == nil {
		return nil, apiError.NewUnauthorizedError("Invalid or expired refresh token.")
	}

	err = s.tokenRepo.DeleteByToken(refreshTokenRequestDTO.RefreshToken)
	if err != nil {
		return nil, apiError.NewInternalServerError("Failed to refresh token.")
	}

	if storedToken.ExpiresAt.Before(time.Now()) {
		return nil, apiError.NewUnauthorizedError("Invalid or expired refresh token.")
	}

	return s.generateAuthResponse(storedToken.UserID)
}

func (s *Service) Logout(userID string) *apiError.ErrorDTO {
	err := s.tokenRepo.DeleteAllTokensByUserID(userID)
	if err != nil {
		return apiError.NewInternalServerError("Failed to logout.")
	}
	return nil
}
