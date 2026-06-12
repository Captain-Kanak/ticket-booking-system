package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	jwtSecret            = "$2a$10$ZHHEQEz7gd97TL7nasjbAnsL4cx5BJX/kzdo2Q2xjutw10ub6"
	defaultTokenDuration = 24 * time.Hour
)

type JWTClaims struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(id uuid.UUID, name, email string) (string, error)
	ValidateToken(token string) error
}

type jwtService struct {
	secret        string
	tokenDuration time.Duration
}

func NewJWTService(secret string) JWTService {
	if secret == "" {
		secret = jwtSecret
	}

	return &jwtService{
		secret:        secret,
		tokenDuration: defaultTokenDuration,
	}
}

func (j *jwtService) GenerateToken(id uuid.UUID, name, email string) (string, error) {
	claims := JWTClaims{
		ID:    id,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ticket-booking-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtService) ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secret), nil
	})
	return err
}
