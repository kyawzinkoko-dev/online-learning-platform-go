package service

import (
	"github.com/kyawzinkoko-dev/online-learning-platform/configs"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/dto"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/token"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/user/domain"
	userRepository "github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/user/repository"
	"github.com/kyawzinkoko-dev/online-learning-platform/pkg/hash"
)

type AuthService struct {
	userRepo userRepository.UserRepository
	config   *configs.Config
}

func NewAuthService(userRepo userRepository.UserRepository, config *configs.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) error {
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     domain.RoleStudent,
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	isValid := hash.CheckPassword(req.Password, user.Password)
	if !isValid {
		return "", err
	}

	token, err := token.GenerateJwt(user.ID.String(), s.config.JWTSecret)

	if err != nil {
		return "", err
	}
	return token, nil
}
