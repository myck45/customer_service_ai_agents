package data

import (
	"github.com/proyectos01-a/shared/models"
)

type UserOrderRepository interface {
	// SaveUserOrder saves a user order to the database.
	SaveUserOrder(order *models.UserOrder) error

	// GetUserOrderByCode retrieves a user order by its code.
	GetUserOrderByCode(orderCode string) (*models.UserOrder, error)

	// UpdateUserOrder updates a user order in the database.
	UpdateUserOrder(order *models.UserOrder) error

	// UpdateUserOrderStatus updates the status of a user order in the database.
	UpdateUserOrderStatus(orderCode string, status string) error

	// UpdateUserOrderByCode updates a user order by its code.
	UpdateUserOrderByCode(orderCode string, order *models.UserOrder) error

	// DeleteUserOrder deletes a user order from the database.
	DeleteUserOrder(orderCode string) error

	// GetUserOrdersByRestaurantID retrieves all user orders for a restaurant.
	GetUserOrdersByRestaurantID(restaurantID uint) ([]models.UserOrder, error)
}
