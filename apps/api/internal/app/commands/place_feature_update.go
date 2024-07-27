package commands

import (
	"context"

	"github.com/9ssi7/txn"
	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceFeatureUpdate struct {
	UserId      uuid.UUID `json:"user_id" validate:"-"`
	FeatureId   uuid.UUID `json:"feature_id" params:"feature_id" validate:"required,uuid"`
	Title       string    `json:"title" validate:"required,min=3,max=255"`
	Description string    `json:"description" validate:"required,min=3,max=255"`
	Icon        string    `json:"icon" validate:"required,min=3,max=255"`
}

type PlaceFeatureUpdateHandler cqrs.HandlerFunc[PlaceFeatureUpdate, *cqrs.Empty]

func NewPlaceFeatureUpdateHandler(t trace.Tracer, v validation.Service, placeFeatureRepo abstracts.PlaceFeatureRepo) PlaceFeatureUpdateHandler {
	return func(ctx context.Context, cmd PlaceFeatureUpdate) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.PlaceFeatureUpdateHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		tx := txn.New()
		tx.Register(placeFeatureRepo.GetTxnAdapter())
		onError := func(ctx context.Context, err error) error {
			tx.Rollback(ctx)
			return err
		}
		feature, err := placeFeatureRepo.FindById(ctx, cmd.FeatureId)
		if err != nil {
			return nil, onError(ctx, err)
		}
		feature.Update(cmd.UserId, cmd.Title, cmd.Description, cmd.Icon)
		if err := placeFeatureRepo.Save(ctx, feature); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
