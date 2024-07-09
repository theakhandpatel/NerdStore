package data

import (
	"time"

	"github.com/theakhandpatel/NerdStore/internal/validator"
)

type Resource struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title,omitempty"`
	Link      string    `json:"link"`
	Tags      []string  `json:"tags,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateResource(v *validator.Validator, resource *Resource) {
	v.Check(resource.Title != "", "title", "must be provided")
	v.Check(len(resource.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(resource.Link != "", "link", "must be provided")
	v.Check(validator.IsUrl(resource.Link), "link", "must be a valid URL")

	v.Check(len(resource.Tags) <= 10, "tags", "must not contain more than 10 tags")
	v.Check(validator.Unique(resource.Tags), "tags", "must not contain duplicate values")
}
