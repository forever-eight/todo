package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/forever-eight/todo.git/internal/app/ds"
	"github.com/forever-eight/todo.git/internal/app/repository"
)

const (
	salt       = "kejeh3bjslsl0mdnls"
	TokenTTL   = 12 * time.Hour
	signingKey = "bwedjfvkfdsdndsnjds4mkejhdef2"
)

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

// Передаем на слой ниже - в репозиторий
func (s *AuthService) CreateUser(user ds.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims arent of right type")
	}
	return claims.UserId, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	tokenSigner := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(), // время на 12 часов больше, чем сейчас
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	token, err := tokenSigner.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("sign error: %w", err)
	}

	return token, nil
}
