package data

import (
	"fmt"

	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

// DeleteUser implements UserRepository.
func (u *UserRepositoryImpl) DeleteUser(id uint) error {
	result := u.db.Delete(&models.User{}, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [DeleteUser] Error deleting user")
		return fmt.Errorf("error deleting user with id %d", id)
	}

	if result.RowsAffected == 0 {
		logrus.WithField("id", id).Warn("*** [DeleteUser] User not found")
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}

// GetAllUsers implements UserRepository.
func (u *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := u.db.Find(&users)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetAllUsers] Error fetching users")
		return nil, fmt.Errorf("error fetching users")
	}

	return users, nil
}

// GetUserByEmail implements UserRepository.
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	result := u.db.Where("user_email = ?", email).First(&user)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetUserByEmail] Error fetching user")
		return nil, fmt.Errorf("error fetching user with email %s", email)
	}

	return &user, nil
}

// GetUserByID implements UserRepository.
func (u *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {

	var user models.User

	result := u.db.First(&user, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetUserByID] Error fetching user")
		return nil, fmt.Errorf("error fetching user with id %d", id)
	}

	return &user, nil
}

// SaveUser implements UserRepository.
func (u *UserRepositoryImpl) SaveUser(user *models.User) (*models.User, error) {

	result := u.db.Create(user)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [SaveUser] Error creating user")
		return nil, fmt.Errorf("error creating user")
	}

	return user, nil
}

// UpdateUser implements UserRepository.
func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	result := u.db.Save(user)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [UpdateUser] Error updating user")
		return fmt.Errorf("error updating user")
	}

	return nil
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
