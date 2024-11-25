package router

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/user/controller"
)

type Router struct {
	userController  controller.UserController
	loginController controller.AuthController
	ginLambda       *ginadapter.GinLambda
}

func NewRouter(userController controller.UserController, loginController controller.AuthController) *Router {
	return &Router{
		userController:  userController,
		loginController: loginController,
	}
}

func (r *Router) InitRoutes() *Router {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	//Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "alive",
		})
	})

	baseRoute := router.Group("/api/v1")
	{
		userRoute := baseRoute.Group("/user")
		{
			userRoute.POST("", r.userController.CreateUser)
			userRoute.DELETE("/:id", r.userController.DeleteUser)
			userRoute.GET("", r.userController.GetAllUsers)
			userRoute.GET("/:id", r.userController.GetUserByID)
			userRoute.GET("/email/:email", r.userController.GetUserByEmail)
			userRoute.PUT("/:id", r.userController.UpdateUser)
			userRoute.POST("/login", r.loginController.Login)
		}
	}

	r.ginLambda = ginadapter.New(router)
	return r
}

func (r *Router) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.ginLambda.ProxyWithContext(ctx, req)
}
