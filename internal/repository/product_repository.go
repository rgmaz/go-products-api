package repository

import (
	"context"

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
