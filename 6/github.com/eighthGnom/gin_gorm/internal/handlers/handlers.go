package handlers

import (
	"github.com/eighthGnom/gin_gorm/internal/halpers"
	"github.com/eighthGnom/gin_gorm/models"
	"github.com/gin-gonic/gin"
)

func GetAllArticles(ctx *gin.Context) {
	var articles []models.Article
	err := models.GetAllArticles(&articles)
	if err != nil {
		halpers.RespondJSON(ctx, 404, articles)
		return
	}
	halpers.RespondJSON(ctx, 200, articles)
}
func PostNewArticle(ctx *gin.Context) {
	var article models.Article
	ctx.BindJSON(&article)
	err := models.AddNewArticle(&article)
	if err != nil {
		halpers.RespondJSON(ctx, 400, article)
		return
	}
	halpers.RespondJSON(ctx, 201, article)
}
func GetArticleByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var article models.Article
	err := models.GetArticleByID(&article, id)
	if err != nil {
		halpers.RespondJSON(ctx, 404, article)
		return
	}
	halpers.RespondJSON(ctx, 200, article)

}
func UpdateArticleByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var article models.Article
	err := models.GetArticleByID(&article, id)
	if err != nil {
		halpers.RespondJSON(ctx, 404, article)
		return
	}
	ctx.BindJSON(&article)
	err = models.UpdateArticleByID(&article, id)
	if err != nil {
		halpers.RespondJSON(ctx, 404, article)
		return
	}
	halpers.RespondJSON(ctx, 202, article)

}
func DeleteArticleByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var article models.Article
	err := models.DeleteArticleByID(&article, id)
	if err != nil {
		halpers.RespondJSON(ctx, 404, article)
		return
	}
	halpers.RespondJSON(ctx, 202, article)
}
