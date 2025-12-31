package repositories

import (
    "context"
    "gorm.io/gorm"

    "backend/domain"
    "backend/models"
    "backend/mappers"
)

type IArticleRepository interface {
    Create(ctx context.Context, article *models.Article) error
    FindByID(ctx context.Context, id uint) (*domain.Article, error)
    FindBySlug(ctx context.Context, slug string) (*domain.Article, error)
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
    return mappers.ArticleToDomain(&article), nil
}

func (r *ArticleRepository) FindBySlug(ctx context.Context, slug string) (*domain.Article, error) {
    var article models.Article
    if err := r.exec(r.Query(ctx).Where("slug = ?", slug).First(&article)); err != nil {
        return nil, err
    }
    return mappers.ArticleToDomain(&article), nil
}

func (r *ArticleRepository) Create(ctx context.Context, article *domain.Article) (*domain.Article, error) {
    article := mappers.ArticleToEntity(article)
    if err := r.BaseRepository.Create(ctx, article); err != nil {
        return nil, err
    }
    return mappers.ArticleToDomain(&article), nil
}

func (r *ArticleRepository) List(ctx context.Context) ([]domain.Article, error) {
    var articles []models.Article
    query : = r.QueryWithPreload(ctx, "Category", "Tags")
    if err := r.BaseRepository.FindAllWithPreload(ctx, &articles, "Category", "Tags"); err != nil {
        return nil, err
    }
    return mappers.ArticleToDomainList(articles), nil
}

