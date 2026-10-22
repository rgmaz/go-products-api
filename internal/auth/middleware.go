package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/rgamz/go-products-api/internal/models"
	"github.com/rgamz/go-products-api/internal/utils"
)

type contextKey string

const claimsContextKey contextKey = "jwt_claims"

// Middleware returns an http.Hanlder middleware that validates the JWT on every request.
func (m *Manager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"authorization header is required",
			)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"authorization heder format must be: Bearer {token}",
			)
			return
		}

		claims, err := m.Validate(parts[1])
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Inject valid claims into the request context
		ctx := context.WithValue(r.Context(), claimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClaimsFromContext retrieves the validated JWT claims from the request context.
//
// Returns nil if no present (i.e., called on an unprotected route).
func ClaimsFromContext(ctx context.Context) *models.JWTClaims {
	claims, _ := ctx.Value(claimsContextKey).(*models.JWTClaims)
	return claims
}

// RequireRole returns a middleware that enforces a specific user role.
func (m *Manager) RequireRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := ClaimsFromContext(r.Context())
		if claims == nil || claims.Role != role {
			utils.WriteError(
				w,
				http.StatusForbidden,
				"you do not have permission to perform this action",
			)

			return
		}

		next.ServeHTTP(w, r)
	})
}
