package route

import (
	"github.com/MicBun/go-job-list-api/controllers"
	"github.com/MicBun/go-job-list-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controllers.Login)

	jobRoutes := r.Group("/jobs")
	jobRoutes.Use(middleware.JwtAuthMiddleware())
	jobRoutes.GET("", controllers.GetJobListFromApi)
	jobRoutes.GET("/:id", controllers.GetJobDetailFromApi)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
