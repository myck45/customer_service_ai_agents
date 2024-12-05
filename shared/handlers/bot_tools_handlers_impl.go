package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"github.com/proyectos01-a/shared/constants"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/dto/req"
	"github.com/proyectos01-a/shared/dto/res"
	"github.com/proyectos01-a/shared/models"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sirupsen/logrus"
)

type BotToolsHandlerImpl struct {
	menuRepo      data.MenuRepository
	userOrderRepo data.UserOrderRepository
	botUtils      utils.BotUtils
}

// HandleGetMenuItemsFromImage implements BotToolsHandler.
func (b *BotToolsHandlerImpl) HandleGetMenuItemsFromImage(data string, restaurantID uint) error {
	// Parse the data into a slice of ExtractedMenuItemResponse
	var extractedItems []res.ExtractedMenuItemResponse
	if err := json.Unmarshal([]byte(data), &extractedItems); err != nil {
		logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to unmarshal data")
		return fmt.Errorf("failed to unmarshal data: %v", err)
	}

	// Iterate over the extracted items and create a menu for each
	for _, item := range extractedItems {
		// Marshal the item into a string to generate an embedding
		itemStr, err := json.Marshal(item)
		if err != nil {
			logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to marshal item")
			return fmt.Errorf("failed to marshal item: %v", err)
		}
		embedding, err := b.botUtils.GenerateEmbedding(string(itemStr))
		if err != nil {
			logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to generate embedding")
			return fmt.Errorf("failed to generate embedding: %v", err)
		}

		// Create a new menu by each item
		menu := &models.Menu{
			ItemName:     item.ItemName,
			Description:  item.Description,
			Price:        item.Price,
			Likes:        0,
			Embedding:    pgvector.NewVector(embedding),
			RestaurantID: restaurantID,
		}
		if err := b.menuRepo.CreateMenu(menu); err != nil {
			logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to create menu")
			return fmt.Errorf("failed to create menu: %v", err)
		}
	}

	return nil
}

// HandleGetUserOrder implements BotToolsHandler.
func (b *BotToolsHandlerImpl) HandleGetUserOrder(data string, chatInfo dto.ChatInfoRequest) (*models.UserOrder, error) {

	// Parse the data into a UserOrderRequest
	var orderRequest req.UserOrderRequest
	if err := json.Unmarshal([]byte(data), &orderRequest); err != nil {
		logrus.WithError(err).Error("[HandleGetUserOrder] failed to unmarshal data")
		return nil, fmt.Errorf("failed to unmarshal data: %v", err)
	}

	// Create a new user order
	order := &models.UserOrder{
		OrderMenuItems:  make([]models.OrderMenuItem, 0),
		OrderCode:       uuid.New(),
		DeliveryAddress: orderRequest.DeliveryAddress,
		UserName:        orderRequest.UserName,
		PhoneNumber:     orderRequest.PhoneNumber,
		PaymentMethod:   orderRequest.PaymentMethod,
		TotalPrice:      0,
		BotWspNumber:    chatInfo.BotWspNumber,
		SenderWspNumber: chatInfo.SenderWspNumber,
		RestaurantID:    chatInfo.RestaurantID,
		Status:          constants.OrderStatusPending,
	}

	// Iterate over the menu items and create an order menu item for each
	for _, item := range orderRequest.MenuItems {
		orderItem := models.OrderMenuItem{
			ItemName: item.ItemName,
			Quantity: item.Quantity,
			Price:    item.Price,
			Subtotal: item.Price * item.Quantity,
		}
		order.TotalPrice += item.Price * item.Quantity
		order.OrderMenuItems = append(order.OrderMenuItems, orderItem)
	}

	// Save the user order
	if err := b.userOrderRepo.SaveUserOrder(order); err != nil {
		logrus.WithError(err).Error("[HandleGetUserOrder] failed to save user order")
		return nil, fmt.Errorf("failed to save user order: %v", err)
	}

	return order, nil
}

func NewBotToolsHandler(menuRepo data.MenuRepository, botUtils utils.BotUtils, userOrderRepo data.UserOrderRepository) BotToolsHandler {
	return &BotToolsHandlerImpl{
		menuRepo:      menuRepo,
		botUtils:      botUtils,
		userOrderRepo: userOrderRepo,
	}
}
