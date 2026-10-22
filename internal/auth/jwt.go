package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rgamz/go-products-api/internal/models"
)

// jwtCustomClaims embeds the standard RegistereClaims, annding custom fields
type jwtCustomClaims struct {
	models.JWTClaims
	jwt.RegisteredClaims
}

// Manager handles all JWT operations
type Manager struct {
	secretKey     []byte
	expiryMinutes int
}

// Generate creates a signed JWT string for a given user
func (m *Manager) Generate(userID int64, email, role string) (string, error) {
	claims := jwtCustomClaims{
		JWTClaims: models.JWTClaims{
			UserID: userID,
			Email:  email,
			Role:   role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(m.expiryMinutes * int(time.Minute)))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-products-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", fmt.Errorf("auth: failed to sign token: %w", err)
	}

	return signedToken, nil
}

// Validate parses and validates a JWT string, returning the embedded claims.
func (m *Manager) Validate(tokenstr string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenstr,
		&jwtCustomClaims{},
		func(t *jwt.Token) (any, error) {
			// Verify signing method to prevent algorithm substitution attacks
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"auth: unexpected signing method %v",
					t.Header["alg"],
				)
			}

			return m.secretKey, nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("auth: token has expired")
		}

		return nil, fmt.Errorf("auth: invalid token: %w", err)
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("auth: token claims are invalid")
	}

	return &claims.JWTClaims, nil
}

// NewManager creates a new JWT Manager.
func NewManager(secret string, expiryMinutes int) *Manager {
	return &Manager{
		secretKey:     []byte(secret),
		expiryMinutes: expiryMinutes,
	}
}
