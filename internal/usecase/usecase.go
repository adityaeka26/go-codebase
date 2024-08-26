package usecase

import (
	"context"

	"github.com/adityaeka26/go-codebase/internal/dto"
)

type ExampleUsecase interface {
	Example(ctx context.Context, req dto.ExampleRequest) (*dto.ExampleResponse, error)
}
