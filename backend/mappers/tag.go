package mappers

import (
    "backend/domain"
    "backend/models"
)

func TagToDomain(m *models.Tag) *domain.Tag {
    if m == nil {
        return nil
    }
    
    return &domain.Tag{
        ID:   m.ID,
        Name: m.Name,
        Slug: m.Slug,
    }
}

func TagToEntity(d *domain.Tag) *models.Tag {
    if d == nil {
        return nil
    }
    
    return &models.Tag{
        ID:   d.ID,
        Name: d.Name,
        Slug: d.Slug,
    }
}
