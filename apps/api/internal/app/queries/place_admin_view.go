package queries

import (
	"context"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceAdminView struct {
	Id uuid.UUID `json:"place_id" params:"place_id" validate:"required,uuid"`
}

type PlaceAdminViewHandler cqrs.HandlerFunc[PlaceAdminView, *entities.Place]

func NewPlaceAdminViewHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceAdminViewHandler {
	return func(ctx context.Context, query PlaceAdminView) (*entities.Place, error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceAdminViewHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := placeRepo.FindById(ctx, query.Id)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
