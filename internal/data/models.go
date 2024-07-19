package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Resources ResourceModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Resources: ResourceModel{DB: db},
	}
}
