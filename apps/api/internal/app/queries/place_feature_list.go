package queries

import (
	"context"

	"github.com/turistikrota/api/internal/app/dtos"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/list"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceFeatureList struct {
	Pagi    list.PagiRequest
	Filters valobj.BaseFilters
}

type PlaceFeatureListHandler cqrs.HandlerFunc[PlaceFeatureList, *list.PagiResponse[*dtos.PlaceFeatureList]]

func NewPlaceFeatureListHandler(t trace.Tracer, v validation.Service, placeFeatureRepo abstracts.PlaceFeatureRepo) PlaceFeatureListHandler {
	return func(ctx context.Context, query PlaceFeatureList) (*list.PagiResponse[*dtos.PlaceFeatureList], error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceFeatureListHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		query.Filters.IsActive = "1" // Only active features for public view
		res, err := placeFeatureRepo.Filter(ctx, &query.Pagi, &query.Filters)
		if err != nil {
			return nil, err
		}
		return dtos.NewPlaceFeatureList(res), nil
	}
}
