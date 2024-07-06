package data

import "time"

type Resource struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title,omitempty"`
	Link      string    `json:"link"`
	Tags      []string  `json:"tags,omitempty"`
	Version   int32     `json:"version"`
}
