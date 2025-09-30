package adapter

import (
	"my-app/internal/auth/domain"

	"github.com/jinzhu/copier"
)

// UserModel — GORM-модель, отображение User в БД
type UserModel struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

// ToDomain конвертирует UserModel → domain.User
func (m *UserModel) ToDomain() (*domain.User, error) {
	var user domain.User
	if err := copier.Copy(&user, m); err != nil {
		return nil, err
	}
	return &user, nil
}

// FromDomain конвертирует domain.User → UserModel
func FromDomain(u *domain.User) (*UserModel, error) {
	var model UserModel
	if err := copier.Copy(&model, u); err != nil {
		return nil, err
	}
	return &model, nil
}
