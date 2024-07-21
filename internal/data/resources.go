package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/theakhandpatel/NerdStore/internal/validator"
)

type Resource struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title,omitempty"`
	Link      string    `json:"link"`
	Tags      []string  `json:"tags,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateResource(v *validator.Validator, resource *Resource) {
	v.Check(resource.Title != "", "title", "must be provided")
	v.Check(len(resource.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(resource.Link != "", "link", "must be provided")
	//TODO:
	// v.Check(validator.IsUrl(resource.Link), "link", "must be a valid URL")

	v.Check(len(resource.Tags) <= 10, "tags", "must not contain more than 10 tags")
	v.Check(validator.Unique(resource.Tags), "tags", "must not contain duplicate values")
}

type ResourceModel struct {
	DB *sql.DB
}

func (r ResourceModel) Insert(resource *Resource) error {
	query := `
		INSERT INTO resources (title, link, tags)
		VALUES($1, $2, $3)
		RETURNING id, created_at, version`

	args := []any{resource.Title, resource.Link, pq.Array(resource.Tags)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return r.DB.QueryRowContext(ctx, query, args...).Scan(&resource.ID, &resource.CreatedAt, &resource.Version)
}

func (r ResourceModel) Get(id int64) (*Resource, error) {
	query := `
		SELECT id, created_at, title, link, tags, version
		FROM resources
		WHERE id = $1`

	var resource Resource

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&resource.ID,
		&resource.CreatedAt,
		&resource.Title,
		&resource.Link,
		pq.Array(&resource.Tags),
		&resource.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &resource, nil
}

func (r ResourceModel) GetAll(title string, tags []string, filters Filters) ([]*Resource, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, title, link, tags, version
		FROM resources
		WHERE 
		(to_tsvector('simple',title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (tags @> $2 OR $2= '{}')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.Sort, filters.Order)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{title, pq.Array(tags), filters.limit(), filters.offset()}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	resources := []*Resource{}
	totalRecords := 0
	for rows.Next() {
		var resource Resource

		err := rows.Scan(
			&totalRecords,
			&resource.ID,
			&resource.CreatedAt,
			&resource.Title,
			&resource.Link,
			pq.Array(&resource.Tags),
			&resource.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		resources = append(resources, &resource)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return resources, metadata, nil
}

func (r ResourceModel) Update(resource *Resource) error {
	query := `
		UPDATE resources
		SET title = $1, link = $2, tags=$3, version = version + 1
		WHERE id = $4 AND version = $5
		RETURNING version`

	args := []any{
		resource.Title,
		resource.Link,
		pq.Array(resource.Tags),
		resource.ID,
		resource.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&resource.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil

}

func (r ResourceModel) Delete(id int64) error {
	query := `
		DELETE FROM resources
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
