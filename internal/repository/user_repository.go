package repostory

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rgamz/go-products-api/internal/models"
)

// UserRepository defines methods for User data access
type UserRepository interface {
	Create(ctx context.Context, email, hashedPassword string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
}

type pgUserRepository struct {
	*pgxpool.Pool
}

func (r *pgUserRepository) Create(ctx context.Context, email, hashedPassword string) (*models.User, error) {
	query := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	RETURNING id, email, role, created_at
	`

	var user models.User
	err := r.QueryRow(ctx, query, email, hashedPassword).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("userRepo.Create: %w", err)
	}

	return &user, nil
}

func (r *pgUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
	SELECT id, email, password, role, created_at 
	FROM users 
	WHERE email = $1
	`

	var user models.User
	err := r.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		// Not found is not considered an error at this layer
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("userRepo.FindByEmail: %w", err)
	}

	return &user, nil
}

func (r *pgUserRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
	SELECT id, email, password, role, created_at 
	FROM users 
	WHERE id = $1
	`

	var user models.User
	err := r.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		// Not found is not considered an error at this layer
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("userRepo.FindByEmail: %w", err)
	}

	return &user, nil

}
