package repositories

import (
    "context"
    "gorm.io/gorm"

    "backend/models"
)

type IArticleRepository interface {
    Create(ctx context.Context, article *models.Article) error
    FindByID(ctx context.Context, id uint) (*domain.Article, error)
    List(ctx context.Context) ([]models.Article, error)
}

type ArticleRepository struct {
    *BaseRepository[models.Article]
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
    return &ArticleRepository{
        BaseRepository: NewBaseRepository[models.Article](db),
    }
}

func (r *ArticleRepository) FindByID(ctx context.Context, id uint) (*domain.Article, error) {
    var article models.Article
    if err := r.BaseRepository.FindByID(ctx, &article, id); err != nil {
        return nil, err
    }
    return &article, nil
}

func (r *ArticleRepository) Create(ctx context.Context, article *models.Article) error {
    return r.BaseRepository.Create(ctx, article)
}

func (r *ArticleRepository) List(ctx context.Context) ([]models.Article, error) {
    var articles []models.Article
    if err := r.BaseRepository.FindAll(ctx, &articles); err != nil {
        return nil, err
    }
    return articles, nil
}

