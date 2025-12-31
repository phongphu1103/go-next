package resources

import "time"

type AricleResource struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Content     string           `json:"content"`
	Category    CategoryResource `json:"category"`
	Tags        []TagResource    `json:"tags"`
	CreatedAt   time.Time        `json:"created_at"`
}