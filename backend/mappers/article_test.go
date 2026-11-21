package mappers

import (
	"testing"

	"backend/models"
)

func TestArticleToDomainWithoutCategory(t *testing.T) {
	// Category not preloaded: used to panic with a nil pointer dereference.
	d := ArticleToDomain(&models.Article{ID: 1, Title: "Hello", CategoryID: 3})
	if d.Category != nil || d.CategoryID != 3 {
		t.Fatalf("unexpected: %+v", d)
	}
}

func TestArticleToDomainWithCategory(t *testing.T) {
	d := ArticleToDomain(&models.Article{
		ID:         1,
		CategoryID: 3,
		Category:   &models.Category{ID: 3, Name: "News", Slug: "news"},
		Tags:       []*models.Tag{{ID: 1, Name: "go", Slug: "go"}},
	})
	if d.Category == nil || d.Category.Slug != "news" || len(d.Tags) != 1 {
		t.Fatalf("unexpected: %+v", d)
	}
}
