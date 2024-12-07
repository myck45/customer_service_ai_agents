package main

import (
	"context"
	"net/http"
	"os"

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

	// Instance Supabase client
	supabaseDB, err := config.NewSupabaseClient()
	if err != nil {
		logrus.WithError(err).Fatal("Error initializing Supabase client")
	}

	s3Client := providers.NewAWSS3Client("sa-east-1")

	// Instance openai client
	openaiClient := providers.NewOpenAIClient()

	// Instance repository
	restaurantRepo := data.NewRestaurantRepositoryImpl(db)

	// Instance menu repository
	menuRepo := data.NewMenuRepositoryImpl(db, supabaseDB)

	// S3 Bucket name
	s3BucketName := os.Getenv("S3_BUCKET_NAME")

	// Instance menu file repository
	s3FileRepo := data.NewS3FileRepositoryImpl(s3Client, s3BucketName, "menu-files")

	// Instance menu file repository
	menuFileRepo := data.NewMenuFileRepositoryImpl(db)

	// Instance User Order Repository
	userOrderRepo := data.NewUserOrderRepository(db)

	// Instance Restaurant Service
	restaurantService := service.NewRestaurantServiceImpl(restaurantRepo)

	// Instance Bot Tools
	botTools := utils.NewBotTools()

	// Instance Bot Utils
	botUtils := utils.NewBotUtilsImpl(openaiClient, menuRepo, botTools)

	// Instance utils
	utils := utils.NewUtilsImpl()

	// Instance Bot Tools Handler
	botToolsHandler := handlers.NewBotToolsHandler(menuRepo, botUtils, userOrderRepo, utils)

	// Instance Menu Service
	menuService := service.NewMenuServiceImpl(menuRepo, botUtils)

	// Instance Menu File Service
	menuFileService := service.NewMenuFileService(menuFileRepo, s3FileRepo, botUtils, botToolsHandler)

	// Instance Response Handler
	responseHandler := handlers.NewResponseHandlersImpl()

	// Instance Restaurant Controller
	restaurantController := controller.NewRestaurantControllerImpl(restaurantService, responseHandler)

	// Instance Menu Controller
	menuController := controller.NewMenuControllerImpl(menuService, responseHandler)

	// Instance Menu File Controller
	menuFileController := controller.NewMenuFileControllerImpl(menuFileService, responseHandler)

	// Instance Router
	r = router.NewRouter(restaurantController, menuController, menuFileController)
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

	return res, nil
}

func main() {
	logrus.Info("Starting restaurant_menu service")
	lambda.Start(Handler)
}
