package queries

import (
	"context"

	"github.com/turistikrota/api/internal/app/dtos"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceView struct {
	Slug string `json:"slug" params:"slug" validate:"required,slug"`
}

type PlaceViewHandler cqrs.HandlerFunc[PlaceView, *dtos.PlaceView]

func NewPlaceViewHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceViewHandler {
	return func(ctx context.Context, query PlaceView) (*dtos.PlaceView, error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceViewHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := placeRepo.FindBySlug(ctx, query.Slug)
		if err != nil {
			return nil, err
		}
		return dtos.NewPlaceView(res), nil
	}
}
