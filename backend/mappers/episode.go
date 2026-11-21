package mappers

import (
	"backend/domain"
	"backend/models"
)

func EpisodeToDomain(m *models.Episode) *domain.Episode {
	if m == nil {
		return nil
	}

	return &domain.Episode{
		ID:       m.ID,
		SeasonID: m.SeasonID,
		Number:   m.Number,
		Title:    m.Title,
		Synopsis: m.Synopsis,
		AirDate:  m.AirDate,
		Duration: m.Duration,
	}
}

func EpisodeToEntity(d *domain.Episode) *models.Episode {
	if d == nil {
		return nil
	}

	return &models.Episode{
		ID:       d.ID,
		SeasonID: d.SeasonID,
		Number:   d.Number,
		Title:    d.Title,
		Synopsis: d.Synopsis,
		AirDate:  d.AirDate,
		Duration: d.Duration,
	}
}

func EpisodeToDomainList(list []models.Episode) []*domain.Episode {
	return MapValues(list, EpisodeToDomain)
}
