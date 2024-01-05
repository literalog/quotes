package quote

import "context"

type Service interface {
	Create(ctx context.Context, q *Quote) error
	Update(ctx context.Context, q *Quote) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]Quote, error)
	GetById(ctx context.Context, id string) (*Quote, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(ctx context.Context, q *Quote) error {
	return s.repository.Create(ctx, q)
}

func (s *service) Update(ctx context.Context, q *Quote) error {
	return s.repository.Update(ctx, q)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]Quote, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetById(ctx context.Context, id string) (*Quote, error) {
	return s.repository.GetById(ctx, id)
}
