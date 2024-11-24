package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/proyectos01-a/restaurantMenu/controller"
	"github.com/proyectos01-a/restaurantMenu/router"
	"github.com/proyectos01-a/restaurantMenu/service"
	"github.com/proyectos01-a/shared/config"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/proyectos01-a/shared/providers"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sirupsen/logrus"
)

var r *router.Router

func init() {

	logrus.Info("Initializing restaurant_menu service")

	// Instance Database
	db := config.DatabaseConnection()

	// Instance openai client
	openaiClient := providers.NewOpenAIClient()

	// Instance Bot Utils
	botUtils := utils.NewBotUtilsImpl(openaiClient)

	// Instance repository
	restaurantRepo := data.NewRestaurantRepositoryImpl(db)

	// Instance menu repository
	menuRepo := data.NewMenuRepositoryImpl(db)

	// Instance Restaurant Service
	restaurantService := service.NewRestaurantServiceImpl(restaurantRepo)

	// Instance Menu Service
	menuService := service.NewMenuServiceImpl(menuRepo, botUtils)

	// Instance Response Handler
	responseHandler := handlers.NewResponseHandlersImpl()

	// Instance Restaurant Controller
	restaurantController := controller.NewRestaurantControllerImpl(restaurantService, responseHandler)

	// Instance Menu Controller
	menuController := controller.NewMenuControllerImpl(menuService, responseHandler)

	// Instance Router
	r = router.NewRouter(restaurantController, menuController)
	r.InitRoutes()

	logrus.Info("Restaurant_menu service initialized")
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logrus.Info("Handling request", req.RequestContext.RequestID)
	res, err := r.Handler(ctx, req)
	if err != nil {
		logrus.WithError(err).Error("Error handling request")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Gateway error, check logs for more information",
		}, err

	}

	logrus.Info("Request handled successfully")

	res.Headers["Access-Control-Allow-Origin"] = "*"
	res.Headers["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	res.Headers["Access-Control-Allow-Methods"] = "OPTIONS,POST,GET,PUT,DELETE"
	return res, nil
}

func main() {
	logrus.Info("Starting restaurant_menu service")
	lambda.Start(Handler)
}
