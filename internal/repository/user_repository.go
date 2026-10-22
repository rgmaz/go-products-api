package repostory

import (
	"context"

	"github.com/rgamz/go-products-api/internal/models"
)

// UserRepository defines methods for User data access
type UserRepository interface {
	Create(ctx context.Context, email, hashedPassword string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
}
