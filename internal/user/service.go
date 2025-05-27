package user

import "context"

type Service interface {
	GetUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, u User) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) CreateUser(ctx context.Context, u User) error {
	return s.repo.Create(ctx, u)
}
