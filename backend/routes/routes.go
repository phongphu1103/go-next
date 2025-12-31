package routes

import (
    "github.com/gin-gonic/gin"
    "backend/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        // api.GET("/articles", controllers.GetArticles)
        api.GET("/articles/:slug", handlers.GetArticle)
        // api.GET("/articles/:id/related", handlers.GetRelated)
        // api.GET("/categories", controllers.GetCategories)
    }
}
