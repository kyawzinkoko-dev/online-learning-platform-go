package repository

import "github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/user/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
