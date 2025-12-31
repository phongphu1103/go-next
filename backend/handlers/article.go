package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "backend/http/controllers"
    "backend/services"
)

type ArticleHandler struct {
    controller *controllers.ArticleController
}

func NewArticleHandler(controller *controllers.ArticleController) *ArticleHandler {
    return &ArticleHandler{
        controller: controller,
    }
}

func (h *ArticleHandler) Create(ctx *gin.Context) {
    var input services.ArticleCreateInput

    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    article, err := h.controller.Create(ctx, input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"data": article})
}

func (h *ArticleHandler) FindByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    article, err := h.controller.FindByID(ctx, uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": article})
}

// GET /articles/:slug
func (h *ArticleHandler) FindBySlug(ctx *gin.Context) {
    slug := ctx.Param("slug")

    article, err := h.controller.FindBySlug(ctx, slug)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": article})
}

// GET /articles/:slug/related
func (h *ArticleHandler) FindBySlugWithRelated(ctx *gin.Context) {
    slug := ctx.Param("slug")

    article, related, err := h.controller.FindBySlugWithRelated(ctx, slug)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": gin.H{
            "article":  article,
            "related":  related,
        },
    })
}

// GET /articles
func (h *ArticleHandler) List(ctx *gin.Context) {
    page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

    var title *string
	if t := ctx.Query("title"); t != "" {
		title = &t
	}
    
	var category *string
	if cat := ctx.Query("category"); cat != "" {
        category = &cat
	}
    
    pagination := domain.Pagination{
		Page:     page,
		Limit:    limit,
		Title:    title,
		Category: category,
	}
    
    articles, err := h.controller.List(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": articles})
}


