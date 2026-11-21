package mappers

import (
	"backend/domain"
	"backend/models"
)

func FranchiseToDomain(m *models.Franchise) *domain.Franchise {
	if m == nil {
		return nil
	}

	return &domain.Franchise{
		ID:          m.ID,
		Name:        m.Name,
		Slug:        m.Slug,
		Description: m.Description,
	}
}

func FranchiseToEntity(d *domain.Franchise) *models.Franchise {
	if d == nil {
		return nil
	}

	return &models.Franchise{
		ID:          d.ID,
		Name:        d.Name,
		Slug:        d.Slug,
		Description: d.Description,
	}
}

func FranchiseToDomainList(list []models.Franchise) []*domain.Franchise {
	return MapValues(list, FranchiseToDomain)
}
