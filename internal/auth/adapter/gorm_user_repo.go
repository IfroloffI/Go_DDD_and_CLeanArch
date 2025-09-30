package adapter

import (
	"my-app/internal/auth/domain"

	"gorm.io/gorm"
)

type GORMUserRepository struct {
	db *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{db: db}
}

func (r *GORMUserRepository) Save(user *domain.User) error {
	model, err := FromDomain(user)
	if err != nil {
		return err
	}
	if model.ID == "" {
		model.ID = generateID() // в продакшене — uuid
	}
	return r.db.Create(model).Error
}

func (r *GORMUserRepository) FindByEmail(email string) (*domain.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToDomain()
}

// Заглушка для ID (в реальности — github.com/google/uuid)
func generateID() string {
	return "temp-id"
}
