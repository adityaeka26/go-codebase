package repository

import (
	"context"

	"github.com/adityaeka26/go-codebase/internal/model"
)

type ExampleRepository interface {
	Example(ctx context.Context, id int) (*model.ExampleModel, error)
}
