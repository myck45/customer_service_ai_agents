package router

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
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

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "alive",
		})
	})

	baseRoute := router.Group("/api/v1")
	{
		botCRUDRoute := baseRoute.Group("/bot")
		{
			botCRUDRoute.POST("", r.botCRUDController.CreateBot)
			botCRUDRoute.DELETE("/:id", r.botCRUDController.DeleteBotByID)
			botCRUDRoute.GET("/all", r.botCRUDController.GetAllBots)
			botCRUDRoute.GET("/:id", r.botCRUDController.GetBotByID)
			botCRUDRoute.GET("/restaurant/:id", r.botCRUDController.GetBotByRestaurantID)
			botCRUDRoute.GET("/whatsapp/:whatsapp", r.botCRUDController.GetBotByWspNumber)
			botCRUDRoute.POST("/update/:id", r.botCRUDController.UpdateBot)
		}

		botRoute := baseRoute.Group("/bot-response")
		{
			botRoute.POST("/twilio/webhook", r.botController.BotResponse)
		}
	}

	r.ginLambda = ginadapter.New(router)

	return r
}

func (r *Router) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.ginLambda.ProxyWithContext(ctx, req)
}
