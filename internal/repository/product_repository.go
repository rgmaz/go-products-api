package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rgamz/go-products-api/internal/models"
)

// ProductRepository defines methods for Product data access
type ProductRepository interface {
	Create(ctx context.Context, userID int64, req *models.CreateProductRequest) (*models.Product, error)
	FindAll(ctx context.Context, page, limit int) ([]*models.Product, int64, error)
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	Update(ctx context.Context, id int64, req *models.UpdateProductRequest) (*models.Product, error)
	Delete(ctx context.Context, id int64) error
}

type pgProductRepository struct {
	*pgxpool.Pool
}

// NewProductReposiroty creates a new PostgreSQL-backed UserRepository
func NewProductReposiroty(pool *pgxpool.Pool) ProductRepository {
	return &pgProductRepository{pool}
}

func (r *pgProductRepository) Create(ctx context.Context, userID int64, req *models.CreateProductRequest) (*models.Product, error) {
	return nil, nil
}

func (r *pgProductRepository) FindAll(ctx context.Context, page, limit int) ([]*models.Product, int64, error) {
	return nil, 0, nil
}

func (r *pgProductRepository) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	return nil, nil
}

func (r *pgProductRepository) Update(ctx context.Context, id int64, req *models.UpdateProductRequest) (*models.Product, error) {
	return nil, nil
}

func (r *pgProductRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
