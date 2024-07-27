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

type PlaceFeatureAdminView struct {
	FeatureId uuid.UUID `json:"feature_id" params:"feature_id" validate:"required,uuid"`
}

type PlaceFeatureAdminViewHandler cqrs.HandlerFunc[PlaceFeatureAdminView, *entities.PlaceFeature]

func NewPlaceFeatureAdminViewHandler(t trace.Tracer, v validation.Service, placeFeatureRepo abstracts.PlaceFeatureRepo) PlaceFeatureAdminViewHandler {
	return func(ctx context.Context, query PlaceFeatureAdminView) (*entities.PlaceFeature, error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceFeatureAdminViewHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := placeFeatureRepo.FindById(ctx, query.FeatureId)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
