package services

import (
	"context"

	"backend/domain"
	"backend/repositories"
)

type IArticleService interface {
	Create(ctx context.Context, input ArticleCreateInput) (*domain.Article, error)
    FindByID(ctx context.Context, id uint) (*domain.Article, error)
    List(ctx context.Context) ([]domain.Article, error)
}

type ArticleService struct {
	repo repositories.IArticleRepository
}

type ArticleCreateInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
    Content     string `json:"content" binding:"required"`
}

func NewArticleService(repo repositories.IArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(ctx context.Context, input ArticleCreateInput) (*models.Article, error) {
	article := &domain.Article{
		Title:       input.Title,
		Description: input.Descripion,
		Content:     input.Content,
	}
	return s.repo.Create(ctx, article)
}

func (s *ArticleService) FindByID(ctx context.Context, id uint) (*domain.Article, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ArticleService) List(ctx context.Context) ([]domain.Article, error) {
	return s.repo.List(ctx)
}