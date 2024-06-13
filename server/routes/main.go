package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vyas-git/demyst-reports/docs"
)

func Init() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Add logic to allow specific origins if needed
			return true
		},
		MaxAge: 12 * time.Hour,
	}
	r.Use(cors.New(config))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	initAccountsRoutes(api)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Demyst APIs"
	docs.SwaggerInfo.Description = "This shows Demyst backend apis."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "Demyst.com"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
