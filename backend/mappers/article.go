package mappers

import (
    "backend/domain"
    "backend/models"
)

// model -> domain
func ArticleToDomain(m *models.Article) *domain.Article {
    if m == nil {
        return nil
    }

    article := &domain.Article{
        ID:          m.ID,
        Title:       m.Title,
        Slug:        m.Slug,
        Description: m.Description,
        Content:     m.Content,
        CategoryID:  m.CategoryID,
        CreatedAt:   m.CreatedAt,
    }

    // Category is nil when it was not preloaded
    if m.Category != nil {
        article.Category = CategoryToDomain(m.Category)
    }

    article.Tags = MapSlice(m.Tags, TagToDomain)

    return article
}

// domain -> model
func ArticleToEntity(d *domain.Article) *models.Article {
    if d == nil {
        return nil
    }

    model := &models.Article{
        Title:       d.Title,
        Slug:        d.Slug,
        Description: d.Description,
        Content:     d.Content,
        CategoryID:  d.CategoryID,
    }

    // Preserve ID for updates
    model.ID = d.ID

    // Optional category mapping
    if d.Category != nil {
        model.Category = CategoryToEntity(d.Category)
    }

    model.Tags = MapSlice(d.Tags, TagToEntity)

    return model
}

func ArticleToDomainList(list []*models.Article) []*domain.Article {
    return MapSlice(list, ArticleToDomain)
}