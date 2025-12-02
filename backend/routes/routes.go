package routes

import (
    "github.com/gin-gonic/gin"
    "backend/http/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/articles", controllers.GetArticles)
        api.GET("/articles/:slug", controllers.GetArticle)
        api.GET("/articles/:id/related", controllers.GetRelated)
        api.GET("/categories", controllers.GetCategories)
    }
}
