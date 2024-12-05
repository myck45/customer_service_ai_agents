package data

import (
	"fmt"

	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserOrderRepositoryImpl struct {
	db *gorm.DB
}

// DeleteUserOrder implements UserOrderRepository.
func (u *UserOrderRepositoryImpl) DeleteUserOrder(orderCode string) error {
	result := u.db.Delete(&models.UserOrder{}, orderCode)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [DeleteUserOrder] Error deleting user order")
		return fmt.Errorf("error deleting user order with code %s", orderCode)
	}

	if result.RowsAffected == 0 {
		logrus.WithField("code", orderCode).Warn("*** [DeleteUserOrder] User order not found")
		return fmt.Errorf("user order with code %s not found", orderCode)
	}

	return nil
}

// GetUserOrderByCode implements UserOrderRepository.
func (u *UserOrderRepositoryImpl) GetUserOrderByCode(orderCode string) (*models.UserOrder, error) {

	var order models.UserOrder

	result := u.db.Where("order_code = ?", orderCode).First(&order)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetUserOrderByCode] Error fetching user order")
		return nil, fmt.Errorf("error fetching user order with code %s", orderCode)
	}

	return &order, nil
}

// GetUserOrdersByRestaurantID implements UserOrderRepository.
func (u *UserOrderRepositoryImpl) GetUserOrdersByRestaurantID(restaurantID uint) ([]models.UserOrder, error) {

	var orders []models.UserOrder

	result := u.db.Where("restaurant_id = ?", restaurantID).Find(&orders)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetUserOrdersByRestaurantID] Error fetching user orders")
		return nil, fmt.Errorf("error fetching user orders for restaurant with id %d", restaurantID)
	}

	return orders, nil
}

// SaveUserOrder implements UserOrderRepository.
func (u *UserOrderRepositoryImpl) SaveUserOrder(order *models.UserOrder) error {

	result := u.db.Create(order)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [SaveUserOrder] Error saving user order")
		return fmt.Errorf("error saving user order")
	}

	return nil
}

// UpdateUserOrder implements UserOrderRepository.
func (u *UserOrderRepositoryImpl) UpdateUserOrder(order *models.UserOrder) error {

	result := u.db.Save(order)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [UpdateUserOrder] Error updating user order")
		return fmt.Errorf("error updating user order")
	}

	return nil
}

// UpdateUserOrderStatus implements UserOrderRepository.
func (u *UserOrderRepositoryImpl) UpdateUserOrderStatus(orderCode string, status string) error {

	result := u.db.Model(&models.UserOrder{}).Where("order_code = ?", orderCode).Update("status", status)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [UpdateUserOrderStatus] Error updating user order status")
		return fmt.Errorf("error updating user order status")
	}

	if result.RowsAffected == 0 {
		logrus.WithField("code", orderCode).Warn("*** [UpdateUserOrderStatus] User order not found")
		return fmt.Errorf("user order with code %s not found", orderCode)
	}

	return nil
}

func NewUserOrderRepository(db *gorm.DB) UserOrderRepository {
	return &UserOrderRepositoryImpl{
		db: db,
	}
}
