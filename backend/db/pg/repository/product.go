package repository

import (
	"context"
	"fmt"
	"spring-assessment-backend/db/pg/model"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type ProductRepository interface {
	ListProducts(ctx context.Context, cursor uuid.UUID, count int) ([]model.Product, error)
	CreateProducts(ctx context.Context, count int) error
	CreateProductsWithBody(ctx context.Context, products []model.Product) error
	SearchProducts(ctx context.Context, query string) ([]model.Product, error)
}

type productRepository struct {
	db *pg.DB
}

func NewProductRepository(db *pg.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) ListProducts(ctx context.Context, cursor uuid.UUID, count int) ([]model.Product, error) {
	products := make([]model.Product, 0)

	if err := r.db.ModelContext(ctx, &products).Where("id > ?", cursor).Limit(count).Select(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) CreateProducts(ctx context.Context, count int) error {
	products := make([]model.Product, 0)

	for i := 0; i < count; i++ {
		products = append(products, model.Product{
			Name:          "test name",
			Description:   "test description",
			Category:      "test category",
			Brand:         "test brand",
			StockQuantity: 10,
			SKU:           "TEST-SKU",
		})
	}

	_, err := r.db.ModelContext(ctx, &products).Insert()

	return err
}

func (r *productRepository) CreateProductsWithBody(ctx context.Context, products []model.Product) error {
	_, err := r.db.ModelContext(ctx, &products).Insert()

	return err
}

func (r *productRepository) SearchProducts(ctx context.Context, query string) ([]model.Product, error) {
	products := make([]model.Product, 0)

	if err := r.db.ModelContext(ctx, &products).Where("name_search like ?", fmt.Sprintf("%%%s%%", query)).Select(); err != nil {
		return nil, err
	}

	return products, nil
}
