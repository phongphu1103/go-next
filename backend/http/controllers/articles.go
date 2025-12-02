package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"

	"backend/database"
	"backend/models"
)

// c *gin.Context is a pointer to a Gin request context
// gin wants to avoid copying big data structures
// allow the handler to modify the response and request context
func GetArticles(c *gin.Context) {
    category := c.Query("category")

    var cat models.Category
    database.DB.Where("slug = ?", category).First(&cat)

    var articles []models.Article
    database.DB.Where("category_id = ?", cat.ID).Find(&articles)

    c.JSON(http.StatusOK, gin.H{
        "category": cat.Name,
        "articles": articles,
    })
}

func GetArticle(c *gin.Context) {
    slug := c.Param("slug")

    var article models.Article
    result := database.DB.Preload("Category").Where("slug = ?", slug).First(&article)

    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
        return
    }

    c.JSON(http.StatusOK, article)
}


func GetRelated(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "articleID": id,
        "related": []string{"Related 1", "Related 2"},
    })
}
