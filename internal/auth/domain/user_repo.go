package domain

// UserRepository — порт из домена (не зависит от GORM!)
type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Save(user *User) error
}
