package data

import (
	"fmt"

	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MenuFileRepositoryImpl struct {
	db *gorm.DB
}

// DeleteMenuFile implements MenuFileRepository.
func (m *MenuFileRepositoryImpl) DeleteMenuFile(menuFileID uint) error {
	result := m.db.Delete(&models.MenuFile{}, menuFileID)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[MenuFileRepositoryImpl] failed to delete menu file")
		return fmt.Errorf("failed to delete menu file: %v", result.Error)
	}

	return nil
}

// GetMenuFileByID implements MenuFileRepository.
func (m *MenuFileRepositoryImpl) GetMenuFileByID(menuFileID uint) (*models.MenuFile, error) {
	var menuFile models.MenuFile

	result := m.db.First(&menuFile, menuFileID)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[MenuFileRepositoryImpl] failed to get menu file")
		return nil, fmt.Errorf("failed to get menu file: %v", result.Error)
	}

	return &menuFile, nil
}

// GetMenuFileByRestaurantID implements MenuFileRepository.
func (m *MenuFileRepositoryImpl) GetMenuFileByRestaurantID(restaurantID uint) ([]models.MenuFile, error) {
	var menuFiles []models.MenuFile

	result := m.db.Where("restaurant_id = ?", restaurantID).Find(&menuFiles)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[MenuFileRepositoryImpl] failed to get menu files")
		return nil, fmt.Errorf("failed to get menu files: %v", result.Error)
	}

	return menuFiles, nil
}

// SaveMenuFile implements MenuFileRepository.
func (m *MenuFileRepositoryImpl) SaveMenuFile(menuFile *models.MenuFile) (*models.MenuFile, error) {
	result := m.db.Create(menuFile)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[MenuFileRepositoryImpl] failed to save menu file")
		return nil, fmt.Errorf("failed to save menu file: %v", result.Error)
	}

	return menuFile, nil
}

// UpdateMenuFile implements MenuFileRepository.
func (m *MenuFileRepositoryImpl) UpdateMenuFile(menuFile *models.MenuFile) (*models.MenuFile, error) {
	result := m.db.Save(menuFile)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("[MenuFileRepositoryImpl] failed to update menu file")
		return nil, fmt.Errorf("failed to update menu file: %v", result.Error)
	}

	return menuFile, nil
}

func NewMenuFileRepositoryImpl(db *gorm.DB) MenuFileRepository {
	return &MenuFileRepositoryImpl{db: db}
}
