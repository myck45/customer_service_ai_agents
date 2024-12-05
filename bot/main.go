package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/proyectos01-a/bot/controller"
	"github.com/proyectos01-a/bot/router"
	"github.com/proyectos01-a/bot/service"
	"github.com/proyectos01-a/shared/config"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/proyectos01-a/shared/providers"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sirupsen/logrus"
)

var r *router.Router

func init() {
	logrus.Info("Initializing bot service")

	// Instance Database
	db := config.DatabaseConnection()

	// Instance Supabase client
	supabaseDB, err := config.NewSupabaseClient()
	if err != nil {
		logrus.WithError(err).Fatal("Error initializing Supabase client")
	}

	// Instance openai client
	openaiClient := providers.NewOpenAIClient()

	// Instance Twilio client
	twilioClient := providers.NewTwilioClient()

	// Instance Twilio Utils
	twilioUtils := utils.NewTwilioUtilsImpl(twilioClient)

	// Instance Bot repository
	botRepo := data.NewBotRepositoryImpl(db)

	// Instance chat history repository
	chatHistoryRepo := data.NewChatHistoryRepositoryImpl(db)

	// Instance Menu repository
	menuRepo := data.NewMenuRepositoryImpl(db, supabaseDB)

	// Instance User Order repository
	userOrderRepo := data.NewUserOrderRepository(db)

	// Instance Bot CRUD Service
	botCRUDService := service.NewBotCRUDServiceImpl(botRepo)

	// Instance Bot Tools
	botTools := utils.NewBotTools()

	// Instance Bot Utils
	botUtils := utils.NewBotUtilsImpl(openaiClient, menuRepo, botTools)

	// Instance Bot Tools Handler
	botToolHandler := handlers.NewBotToolsHandler(menuRepo, botUtils, userOrderRepo)

	// Instance Bot Service
	botService := service.NewBotServiceImpl(openaiClient, twilioUtils, botUtils, chatHistoryRepo, botRepo, menuRepo, botTools, botToolHandler)

	// Instance Response Handler
	responseHandler := handlers.NewResponseHandlersImpl()

	// Instance Bot CRUD Controller
	botCRUDController := controller.NewBotCRUDControllerImpl(botCRUDService, responseHandler)

	// Instance Bot Controller
	botController := controller.NewBotControllerImpl(botService, responseHandler)

	// Instance Router
	r = router.NewRouter(botCRUDController, botController)
	r.InitRoutes()

	logrus.Info("Bot service initialized")
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
