package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/proyectos01-a/shared/config"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/proyectos01-a/shared/utils"
	"github.com/proyectos01-a/user/auth"
	"github.com/proyectos01-a/user/controller"
	"github.com/proyectos01-a/user/router"
	"github.com/proyectos01-a/user/service"
	"github.com/sirupsen/logrus"
)

var r *router.Router

func init() {
	logrus.Info("Initializing user service")

	// Instance Database
	db := config.DatabaseConnection()

	// Instance repository
	userRepo := data.NewUserRepositoryImpl(db)

	// Instance bcrypt
	bcryptUtil := auth.NewBcryptImpl()

	// Instance utils
	utils := utils.NewUtilsImpl()

	// Instance auth
	auth := auth.NewAuth()

	// Instance user service
	userService := service.NewUserServiceImpl(userRepo, bcryptUtil, utils)

	// Instance auth service
	authService := service.NewAuthServiceImpl(auth, bcryptUtil, userRepo)

	// Instance response handler
	responseHandler := handlers.NewResponseHandlersImpl()

	// Instance controller
	userController := controller.NewUserControllerImpl(userService, responseHandler)

	// Instance auth controller
	authController := controller.NewAuthControllerImpl(authService, responseHandler)

	// Instance router
	r = router.NewRouter(userController, authController)
	r.InitRoutes()

	logrus.Info("User service initialized")
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logrus.Info("Handling request", req.RequestContext.RequestID)
	res, err := r.Handler(ctx, req)
	if err != nil {
		logrus.WithError(err).Error("Error handling request")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Gateway error, check logs for more information",
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
				"Access-Control-Allow-Methods": "OPTIONS,POST,GET,PUT,DELETE",
			},
		}, err

	}

	logrus.Info("Request handled successfully")

	res.Headers["Access-Control-Allow-Origin"] = "*"
	res.Headers["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	res.Headers["Access-Control-Allow-Methods"] = "OPTIONS,POST,GET,PUT,DELETE"
	return res, nil
}

func main() {
	logrus.Info("Starting user service")
	lambda.Start(Handler)
}
