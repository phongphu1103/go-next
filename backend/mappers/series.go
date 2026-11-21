package mappers

import (
	"backend/domain"
	"backend/models"
)

// model -> domain
// Relations are only filled when they were preloaded
// (e.g. Preload("Seasons.Episodes")); otherwise they stay nil / empty.
func SeriesToDomain(m *models.Series) *domain.Series {
	if m == nil {
		return nil
	}

	return &domain.Series{
		ID:             m.ID,
		Title:          m.Title,
		Slug:           m.Slug,
		Synopsis:       m.Synopsis,
		FirstAirDate:   m.FirstAirDate,
		LastAirDate:    m.LastAirDate,
		Status:         domain.SeriesStatus(m.Status),
		UniverseID:     m.UniverseID,
		Universe:       UniverseToDomain(m.Universe),
		PrimaryGenreID: m.PrimaryGenreID,
		PrimaryGenre:   GenreToDomain(m.PrimaryGenre),
		Genres:         MapValues(m.Genres, GenreToDomain),
		Franchises:     MapValues(m.Franchises, FranchiseToDomain),
		Seasons:        MapValues(m.Seasons, SeasonToDomain),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

// domain -> model
func SeriesToEntity(d *domain.Series) *models.Series {
	if d == nil {
		return nil
	}

	return &models.Series{
		ID:             d.ID, // preserve ID for updates
		Title:          d.Title,
		Slug:           d.Slug,
		Synopsis:       d.Synopsis,
		FirstAirDate:   d.FirstAirDate,
		LastAirDate:    d.LastAirDate,
		Status:         string(d.Status),
		UniverseID:     d.UniverseID,
		Universe:       UniverseToEntity(d.Universe),
		PrimaryGenreID: d.PrimaryGenreID,
		PrimaryGenre:   GenreToEntity(d.PrimaryGenre),
		Genres:         MapToValues(d.Genres, GenreToEntity),
		Franchises:     MapToValues(d.Franchises, FranchiseToEntity),
		Seasons:        MapToValues(d.Seasons, SeasonToEntity),
	}
}

func SeriesToDomainList(list []models.Series) []*domain.Series {
	return MapValues(list, SeriesToDomain)
}

func SeasonToDomain(m *models.Season) *domain.Season {
	if m == nil {
		return nil
	}

	return &domain.Season{
		ID:       m.ID,
		SeriesID: m.SeriesID,
		Number:   m.Number,
		Title:    m.Title,
		AirDate:  m.AirDate,
		Episodes: MapValues(m.Episodes, EpisodeToDomain),
	}
}

func SeasonToEntity(d *domain.Season) *models.Season {
	if d == nil {
		return nil
	}

	return &models.Season{
		ID:       d.ID,
		SeriesID: d.SeriesID,
		Number:   d.Number,
		Title:    d.Title,
		AirDate:  d.AirDate,
		Episodes: MapToValues(d.Episodes, EpisodeToEntity),
	}
}

func SeasonToDomainList(list []models.Season) []*domain.Season {
	return MapValues(list, SeasonToDomain)
}
