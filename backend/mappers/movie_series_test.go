package mappers

import (
	"reflect"
	"testing"
	"time"

	"backend/models"
)

func ptr[T any](v T) *T { return &v }

func TestMovieRoundTrip(t *testing.T) {
	release := time.Date(2013, 6, 14, 0, 0, 0, 0, time.UTC)
	m := &models.Movie{
		ID:             1,
		Title:          "Man of Steel",
		Slug:           "man-of-steel",
		Synopsis:       ptr("Clark Kent learns he is from Krypton."),
		ReleaseDate:    &release,
		Duration:       ptr(uint(143)),
		BudgetMin:      ptr(int64(225000000)),
		BudgetMax:      ptr(int64(225000000)),
		BoxOffice:      ptr(int64(670145518)),
		UniverseID:     ptr(uint(1)),
		Universe:       &models.Universe{ID: 1, Name: "DC Extended Universe", ShortName: ptr("DCEU"), Slug: "dceu"},
		PrimaryGenreID: ptr(uint(1)),
		PrimaryGenre:   &models.Genre{ID: 1, Name: "Superhero", Slug: "superhero"},
		Genres:         []models.Genre{{ID: 1, Name: "Superhero", Slug: "superhero"}, {ID: 2, Name: "Action", Slug: "action"}},
		Franchises:     []models.Franchise{{ID: 1, Name: "Superman", Slug: "superman"}},
	}

	d := MovieToDomain(m)
	if len(d.Genres) != 2 || d.Genres[1].Slug != "action" || d.Universe.Slug != "dceu" || *d.BudgetMax != 225000000 {
		t.Fatalf("unexpected domain movie: %+v", d)
	}

	if back := MovieToEntity(d); !reflect.DeepEqual(back, m) {
		t.Fatalf("round trip mismatch:\n got  %+v\n want %+v", back, m)
	}
}

func TestMovieWithoutPreloads(t *testing.T) {
	d := MovieToDomain(&models.Movie{ID: 1, Title: "Man of Steel"})
	if d.Universe != nil || d.PrimaryGenre != nil || len(d.Genres) != 0 || len(d.Franchises) != 0 {
		t.Fatalf("relations should be empty when not preloaded: %+v", d)
	}
	if MovieToDomain(nil) != nil || MovieToEntity(nil) != nil {
		t.Fatal("nil input should map to nil")
	}
}

func TestSeriesRoundTrip(t *testing.T) {
	m := &models.Series{
		ID:     1,
		Title:  "Peacemaker",
		Slug:   "peacemaker",
		Status: models.SeriesStatusOngoing,
		Genres: []models.Genre{{ID: 1, Name: "Superhero", Slug: "superhero"}},
		Seasons: []models.Season{{
			ID: 1, SeriesID: 1, Number: 1,
			Episodes: []models.Episode{
				{ID: 1, SeasonID: 1, Number: 1, Title: "A Whole New Whirled", Duration: ptr(uint(47))},
				{ID: 2, SeasonID: 1, Number: 2, Title: "Best Friends, For Never"},
			},
		}},
	}

	d := SeriesToDomain(m)
	if d.Status != "ongoing" || len(d.Seasons) != 1 || len(d.Seasons[0].Episodes) != 2 {
		t.Fatalf("unexpected domain series: %+v", d)
	}

	if back := SeriesToEntity(d); !reflect.DeepEqual(back, m) {
		t.Fatalf("round trip mismatch:\n got  %+v\n want %+v", back, m)
	}
}
