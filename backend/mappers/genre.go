package mappers

import (
	"backend/domain"
	"backend/models"
)

func GenreToDomain(m *models.Genre) *domain.Genre {
	if m == nil {
		return nil
	}

	return &domain.Genre{
		ID:   m.ID,
		Name: m.Name,
		Slug: m.Slug,
	}
}

func GenreToEntity(d *domain.Genre) *models.Genre {
	if d == nil {
		return nil
	}

	return &models.Genre{
		ID:   d.ID,
		Name: d.Name,
		Slug: d.Slug,
	}
}

func GenreToDomainList(list []models.Genre) []*domain.Genre {
	return MapValues(list, GenreToDomain)
}
