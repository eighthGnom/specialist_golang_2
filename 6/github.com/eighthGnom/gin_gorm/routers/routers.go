package routers

import (
	"github.com/eighthGnom/gin_gorm/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1Group := router.Group("/api/v1")
	apiV1Group.GET("article", handlers.GetAllArticles)
	apiV1Group.POST("article", handlers.PostNewArticle)
	apiV1Group.GET("article/:id", handlers.GetArticleByID)
	apiV1Group.PUT("article/:id", handlers.UpdateArticleByID)
	apiV1Group.DELETE("article/:id", handlers.DeleteArticleByID)
	return router
}
