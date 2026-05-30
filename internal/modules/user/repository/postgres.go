package repository

import (
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/user/domain"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) UserRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error

}

func (r *postgresRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil

}
