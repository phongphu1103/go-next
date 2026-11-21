package mappers

import (
	"backend/domain"
	"backend/models"
)

func UniverseToDomain(m *models.Universe) *domain.Universe {
	if m == nil {
		return nil
	}

	return &domain.Universe{
		ID:          m.ID,
		Name:        m.Name,
		ShortName:   m.ShortName,
		Slug:        m.Slug,
		Description: m.Description,
	}
}

func UniverseToEntity(d *domain.Universe) *models.Universe {
	if d == nil {
		return nil
	}

	return &models.Universe{
		ID:          d.ID,
		Name:        d.Name,
		ShortName:   d.ShortName,
		Slug:        d.Slug,
		Description: d.Description,
	}
}

func UniverseToDomainList(list []models.Universe) []*domain.Universe {
	return MapValues(list, UniverseToDomain)
}
