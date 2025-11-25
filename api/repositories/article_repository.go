package repository

import "api/models"

type ArticleRepository {
	Create(article *models.Article) error
	Find() ([]models.Article, error)
}
