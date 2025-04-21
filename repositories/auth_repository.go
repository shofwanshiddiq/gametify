package repositories

import (
	"gametify/models"

	"gorm.io/gorm"
)

// Rename the interface to avoid conflict with the struct
type AuthRepositoryInterface interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

// Rename the struct to avoid conflict with the interface
type AuthRepository struct {
	db *gorm.DB
}

// Constructor function now returns a pointer to the struct
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// Implement the methods for the struct
func (r *AuthRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
