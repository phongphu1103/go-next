package mappers

import (
    "backend/domain"
    "backend/models"
)

func CategoryToDomain(m *models.Category) *domain.Category {
    if m == nil {
        return nil
    }

    return &domain.Category{
        ID:   m.ID,
        Name: m.Name,
        Slug: m.Slug,
    }
}

func CategoryToEntity(d *domain.Category) *models.Category {
    if d == nil {
        return nil
    }

    return &models.Category{
        ID:   d.ID,
        Name: d.Name,
        Slug: d.Slug,
    }
}
