package service

import (
	"errors"
	"my-app/internal/auth/domain"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(email, password string) (string, error) {
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return "", errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &domain.User{
		Email:    email,
		Password: string(hashed),
	}
	if err := user.Validate(); err != nil {
		return "", err
	}

	if err := s.userRepo.Save(user); err != nil {
		return "", err
	}

	return user.ID, nil
}

func (s *AuthService) Login(email, password string) error {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
