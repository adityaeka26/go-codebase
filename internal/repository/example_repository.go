package repository

import (
	"context"

	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/model"
)

type exampleRepository struct {
	config *config.EnvConfig
}

func NewExampleRepository(config *config.EnvConfig) (ExampleRepository, error) {
	return &exampleRepository{
		config: config,
	}, nil
}

func (r *exampleRepository) Example(ctx context.Context, id int) (*model.ExampleModel, error) {
	// do something here like query to database, call external API, etc

	return &model.ExampleModel{
		Id:   id,
		Name: "test",
	}, nil
}
