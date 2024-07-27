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

type PlaceDisable struct {
	UserId uuid.UUID `json:"user_id" validate:"-"`
	Id     uuid.UUID `json:"place_id" params:"place_id" validate:"required,uuid"`
}

type PlaceDisableHandler cqrs.HandlerFunc[PlaceDisable, *cqrs.Empty]

func NewPlaceDisableHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceDisableHandler {
	return func(ctx context.Context, cmd PlaceDisable) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.PlaceDisableHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		tx := txn.New()
		tx.Register(placeRepo.GetTxnAdapter())
		onError := func(ctx context.Context, err error) error {
			tx.Rollback(ctx)
			return err
		}
		place, err := placeRepo.FindById(ctx, cmd.Id)
		if err != nil {
			return nil, onError(ctx, err)
		}
		place.Disable(cmd.UserId)
		if err := placeRepo.Save(ctx, place); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
