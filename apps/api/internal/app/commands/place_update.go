package commands

import (
	"context"

	"github.com/9ssi7/txn"
	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceUpdate struct {
	UserId       uuid.UUID       `json:"user_id" validate:"-"`
	Id           uuid.UUID       `json:"place_id" params:"place_id" validate:"required,uuid"`
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

type PlaceUpdateHandler cqrs.HandlerFunc[PlaceUpdate, *cqrs.Empty]

func NewPlaceUpdateHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceUpdateHandler {
	return func(ctx context.Context, cmd PlaceUpdate) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.PlaceUpdateHandler")
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
		ids := make([]uuid.UUID, len(cmd.FeatureIds))
		for i, id := range cmd.FeatureIds {
			ids[i] = uuid.MustParse(id)
		}
		place.Update(cmd.UserId, ids, valobj.PlaceKind(cmd.Kind), cmd.Title, cmd.Description, *cmd.Seo, cmd.Latitude, cmd.Longitude, cmd.Images, *cmd.MinTimeSpent, *cmd.MaxTimeSpent, *cmd.IsPayed)
		if err := placeRepo.Save(ctx, place); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
