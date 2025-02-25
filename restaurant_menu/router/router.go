package router

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/restaurantMenu/controller"
)

type Router struct {
	restaurantController controller.RestaurantController
	menuController       controller.MenuController
	menuFileController   controller.MenuFileController
	ginLambda            *ginadapter.GinLambda
}

func NewRouter(
	restaurantController controller.RestaurantController,
	menuController controller.MenuController,
	menuFileController controller.MenuFileController,
) *Router {
	return &Router{
		restaurantController: restaurantController,
		menuController:       menuController,
		menuFileController:   menuFileController,
	}
}

func (r *Router) InitRoutes() *Router {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Limit upload file size for multipart forms
	router.MaxMultipartMemory = 16 << 20 // 16 MB

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
		restaurantRoute := baseRoute.Group("/restaurant")
		{
			restaurantRoute.POST("", r.restaurantController.CreateRestaurant)
			restaurantRoute.DELETE("/:id", r.restaurantController.DeleteRestaurant)
			restaurantRoute.GET("", r.restaurantController.GetAllRestaurants)
			restaurantRoute.GET("/:id", r.restaurantController.GetRestaurantByID)
			restaurantRoute.PUT("/:id", r.restaurantController.UpdateRestaurant)
		}

		menuRoute := baseRoute.Group("/menu")
		{
			menuRoute.POST("", r.menuController.CreateMenu)
			menuRoute.DELETE("/:id", r.menuController.DeleteMenu)
			menuRoute.GET("", r.menuController.GetAllMenus)
			menuRoute.GET("/search", r.menuController.SemanticSearchMenu)
			menuRoute.GET("/:id", r.menuController.GetMenuByID)
			menuRoute.PUT("/:id", r.menuController.UpdateMenu)
		}

		menuFileRoute := baseRoute.Group("/menu-files")
		{
			menuFileRoute.POST("", r.menuFileController.CreateMenuFile)
			menuFileRoute.DELETE("/:id", r.menuFileController.DeleteMenuFile)
			menuFileRoute.GET("/:id", r.menuFileController.GetMenuFileByID)
			menuFileRoute.GET("/restaurant/:restaurant_id", r.menuFileController.GetMenuFileByRestaurantID)
			menuFileRoute.PUT("/:id", r.menuFileController.UpdateMenuFile)
		}
	}

	r.ginLambda = ginadapter.New(router)

	return r
}

func (r *Router) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.ginLambda.ProxyWithContext(ctx, req)
}
