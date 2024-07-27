package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceCreate struct {
	UserId       uuid.UUID       `json:"user_id" validate:"-"`
	FeatureIds   []string        `json:"feature_ids" validate:"required,dive,uuid"`
	Title        string          `json:"title" validate:"required,min=3,max=255"`
	Description  string          `json:"description" validate:"required,min=3,max=255"`
	Seo          *valobj.Seo     `json:"seo" validate:"required,dive"`
	Images       []*valobj.Image `json:"images" validate:"required,dive"`
	Latitude     float64         `json:"latitude" validate:"required,latitude"`
	Longitude    float64         `json:"longitude" validate:"required,longitude"`
	MinTimeSpent *int16          `json:"min_time_spent" validate:"required,min=0,max=1440"`
	MaxTimeSpent *int16          `json:"max_time_spent" validate:"required,min=0,max=1440"`
	IsPayed      *bool           `json:"is_payed" validate:"required"`
	Kind         string          `json:"kind" validate:"required,place_kind"`
}

type PlaceCreateHandler cqrs.HandlerFunc[PlaceCreate, *uuid.UUID]

func NewPlaceCreateHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceCreateHandler {
	return func(ctx context.Context, cmd PlaceCreate) (*uuid.UUID, error) {
		ctx = tracer.Push(ctx, t, "commands.PlaceCreateHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		ids := make([]uuid.UUID, len(cmd.FeatureIds))
		for i, id := range cmd.FeatureIds {
			ids[i] = uuid.MustParse(id)
		}
		place := entities.NewPlace(cmd.UserId, ids, valobj.PlaceKind(cmd.Kind), cmd.Title, cmd.Description, *cmd.Seo, cmd.Latitude, cmd.Longitude, cmd.Images, *cmd.MinTimeSpent, *cmd.MaxTimeSpent, *cmd.IsPayed)
		if err := placeRepo.Save(ctx, place); err != nil {
			return nil, err
		}
		return &place.Id, nil
	}
}
