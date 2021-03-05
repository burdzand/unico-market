package api

import (
	"unico/docs"
	"unico/internal/unico/database"

	logger "unico/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const versionPath = "/v1"

var log *logger.Logger

func init() {
	log = logger.NewLogger()
}

func ConfigureRoutes(r *gin.Engine, db database.Database) {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Unico Market API"
	docs.SwaggerInfo.Description = "Api Unico to Market"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	handler := controlerHandler{db, validator.New()}

	routes := r.Group(docs.SwaggerInfo.BasePath)
	{
		// Set Swagger
		routes.POST("import-data", handler.ImportData)
		market := routes.Group("/market")
		{
			market.POST("/", handler.CreateMarket)
			market.GET("/", handler.GetSearchMarket)
			market.GET(":record", handler.GetMarket)
			market.PUT(":record", handler.UpdateMarket)
			market.DELETE(":record", handler.RemoveMarket)
		}

		district := routes.Group("/district")
		{
			district.GET(":subprefecture", handler.GetAllDistrict)
		}

		subprefecture := routes.Group("/subprefecture")
		{
			subprefecture.GET("/", handler.GetAllSubprefecture)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
