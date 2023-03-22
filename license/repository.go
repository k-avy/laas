package license

import (
	"context"
	"errors"
)

var (
	ErrCheck =errors.New("license already exists")
	ErrNotExist=errors.New("row doesnt exists")
	ErrUpdateFailed=errors.New("update failed")
	ErrDeleteFailed= errors.New("delete failed")
)

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context,license License) (*License,error)
	All(ctx context.Context) ([]License, error)
	GetByName(ctx context.Context,shortname string)(*License, error)
	Update(ctx context.Context, id int64, updated License) (*License,error)
	Delete(ctx context.Context, id int64) error
}