package data

import (
	"fmt"

	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
	db *gorm.DB
}

// CreateMenu implements MenuRepository.
func (m *MenuRepositoryImpl) CreateMenu(menu *models.Menu) error {
	result := m.db.Create(menu)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error creating menu")
		return fmt.Errorf("error creating menu")
	}

	return nil
}

// DeleteMenu implements MenuRepository.
func (m *MenuRepositoryImpl) DeleteMenu(id uint) error {
	result := m.db.Delete(&models.Menu{}, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error deleting menu")
		return fmt.Errorf("error deleting menu with id %d", id)
	}

	if result.RowsAffected == 0 {
		logrus.WithField("id", id).Warn("Menu not found")
		return fmt.Errorf("menu with id %d not found", id)
	}

	return nil
}

// GetAllMenus implements MenuRepository.
func (m *MenuRepositoryImpl) GetAllMenus() ([]models.Menu, error) {
	var menus []models.Menu

	result := m.db.Find(&menus)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error fetching menus")
		return nil, fmt.Errorf("error fetching menus")
	}

	return menus, nil
}

// GetMenuByID implements MenuRepository.
func (m *MenuRepositoryImpl) GetMenuByID(id uint) (*models.Menu, error) {
	var menu models.Menu

	result := m.db.First(&menu, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error fetching menu")
		return nil, fmt.Errorf("error fetching menu")
	}

	return &menu, nil
}

// SemanticSearchMenu implements MenuRepository.
func (m *MenuRepositoryImpl) SemanticSearchMenu(queryEmbedding []float32, similarityThreshold float32, matchCount int, restaurantID uint) ([]dto.MenuSearchResponse, error) {
	var results []dto.MenuSearchResponse

	// result := m.db.Model(&models.Menu{}).
	// 	Select(`
	// 		id,
	// 		item_name,
	// 		price,
	// 		description,
	// 		likes,
	// 		embedding <#> ? AS similarity
	// 	`, queryEmbedding).
	// 	Where("restaurant_id = ?", restaurantID).
	// 	Where("embedding <#> ? < ?", queryEmbedding, similarityThreshold).
	// 	Order("similarity").
	// 	Limit(matchCount).
	// 	Scan(&results)

	result := m.db.Raw(`
		SELECT
			id,
			item_name,
			price,
			description,
			likes,
			embedding <#> ? AS similarity
		FROM menus
		WHERE restaurant_id = ? AND embedding <#> ? < ?
		ORDER BY similarity
		LIMIT ?
	`, queryEmbedding, restaurantID, queryEmbedding, similarityThreshold, matchCount).Scan(&results)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error performing semantic search")
		return nil, fmt.Errorf("error performing semantic search")
	}

	return results, nil
}

// UpdateMenu implements MenuRepository.
func (m *MenuRepositoryImpl) UpdateMenu(menu *models.Menu) error {
	result := m.db.Save(menu)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error updating menu")
		return fmt.Errorf("error updating menu with id %d", menu.ID)
	}

	return nil
}

func NewMenuRepositoryImpl(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		db: db,
	}
}
