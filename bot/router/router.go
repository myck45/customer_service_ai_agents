package router

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/bot/controller"
)

type Router struct {
	botCRUDController controller.BotCRUDController
	botController     controller.BotController
	ginLambda         *ginadapter.GinLambda
}

func NewRouter(botCRUDController controller.BotCRUDController, botController controller.BotController) *Router {
	return &Router{
		botCRUDController: botCRUDController,
		botController:     botController,
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
		ctx.JSON(200, gin.H{
			"message": "alive",
		})
	})

	baseRoute := router.Group("/api/v1")
	{
		botRoute := baseRoute.Group("/bot")
		{
			botRoute.POST("", r.botCRUDController.CreateBot)
			botRoute.DELETE("/:id", r.botCRUDController.DeleteBotByID)
			botRoute.GET("", r.botCRUDController.GetAllBots)
			botRoute.GET("/:id", r.botCRUDController.GetBotByID)
			botRoute.GET("/restaurant/:id", r.botCRUDController.GetBotByRestaurantID)
			botRoute.GET("/whatsapp/:whatsapp", r.botCRUDController.GetBotByWspNumber)
			botRoute.PUT("/:id", r.botCRUDController.UpdateBot)
			botRoute.POST("/twilio", r.botController.BotResponse)
		}
	}

	r.ginLambda = ginadapter.New(router)

	return r
}

func (r *Router) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.ginLambda.ProxyWithContext(ctx, req)
}
