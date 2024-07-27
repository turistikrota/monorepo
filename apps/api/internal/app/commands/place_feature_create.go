package commands

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

type PlaceFeatureCreate struct {
	UserId      uuid.UUID `json:"user_id" validate:"-"`
	Title       string    `json:"title" validate:"required,min=3,max=255"`
	Description string    `json:"description" validate:"required,min=3,max=255"`
	Icon        string    `json:"icon" validate:"required,min=3,max=255"`
}

type PlaceFeatureCreateHandler cqrs.HandlerFunc[PlaceFeatureCreate, *uuid.UUID]

func NewPlaceFeatureCreateHandler(t trace.Tracer, v validation.Service, placeFeatureRepo abstracts.PlaceFeatureRepo) PlaceFeatureCreateHandler {
	return func(ctx context.Context, cmd PlaceFeatureCreate) (*uuid.UUID, error) {
		ctx = tracer.Push(ctx, t, "commands.PlaceFeatureCreateHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		placeFeature := entities.NewPlaceFeature(cmd.UserId, cmd.Title, cmd.Description, cmd.Icon)
		if err := placeFeatureRepo.Save(ctx, placeFeature); err != nil {
			return nil, err
		}
		return &placeFeature.Id, nil
	}
}
