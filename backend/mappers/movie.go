package mappers

import (
	"backend/domain"
	"backend/models"
)

// model -> domain
// Relations (Universe, PrimaryGenre, Genres, Franchises) are only filled
// when they were preloaded; otherwise they stay nil / empty.
func MovieToDomain(m *models.Movie) *domain.Movie {
	if m == nil {
		return nil
	}

	return &domain.Movie{
		ID:             m.ID,
		Title:          m.Title,
		Slug:           m.Slug,
		Synopsis:       m.Synopsis,
		ReleaseDate:    m.ReleaseDate,
		Duration:       m.Duration,
		BudgetMin:      m.BudgetMin,
		BudgetMax:      m.BudgetMax,
		BoxOffice:      m.BoxOffice,
		UniverseID:     m.UniverseID,
		Universe:       UniverseToDomain(m.Universe),
		PrimaryGenreID: m.PrimaryGenreID,
		PrimaryGenre:   GenreToDomain(m.PrimaryGenre),
		Genres:         MapValues(m.Genres, GenreToDomain),
		Franchises:     MapValues(m.Franchises, FranchiseToDomain),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

// domain -> model
// Genres / Franchises are passed through so GORM writes the join tables on
// Create. Set only UniverseID / PrimaryGenreID (not the objects) when linking
// to existing rows, so GORM does not try to upsert them.
func MovieToEntity(d *domain.Movie) *models.Movie {
	if d == nil {
		return nil
	}

	return &models.Movie{
		ID:             d.ID, // preserve ID for updates
		Title:          d.Title,
		Slug:           d.Slug,
		Synopsis:       d.Synopsis,
		ReleaseDate:    d.ReleaseDate,
		Duration:       d.Duration,
		BudgetMin:      d.BudgetMin,
		BudgetMax:      d.BudgetMax,
		BoxOffice:      d.BoxOffice,
		UniverseID:     d.UniverseID,
		Universe:       UniverseToEntity(d.Universe),
		PrimaryGenreID: d.PrimaryGenreID,
		PrimaryGenre:   GenreToEntity(d.PrimaryGenre),
		Genres:         MapToValues(d.Genres, GenreToEntity),
		Franchises:     MapToValues(d.Franchises, FranchiseToEntity),
	}
}

func MovieToDomainList(list []models.Movie) []*domain.Movie {
	return MapValues(list, MovieToDomain)
}
