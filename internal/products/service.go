package products

import (
	"context"
	repo "ecom-local/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	ListProductById(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) ListProductById(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.ListProductById(ctx, id)
	if err != nil {
		return repo.Product{}, err
	}
	// exemple de règle métier supplémentaire : On applique la réduction de 50%
	// (On divise le prix en centimes par 2)
	// product.PriceInCents = product.PriceInCents / 2

	return product, nil
}
