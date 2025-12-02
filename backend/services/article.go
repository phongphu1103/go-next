package services

import (
	"context"

	"backend/models"
	"backend/repositories"
)

type IArticleService interface {
	Create(ctx context.Context, input ArticleCreateInput) (*models.Article, error)
    FindByID(ctx context.Context, id uint) (*models.Article, error)
    List(ctx context.Context) ([]models.Article, error)
}

type ArticleService struct {
	repo repositories.IArticleRepository
}

func NewArticleService(repo repositories.IArticleRepository) *ArticleService {
    return &ArticleService{repo: repo}
}

type ArticleCreateInput struct {
    Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
    Content     string `json:"content" binding:"required"`
}