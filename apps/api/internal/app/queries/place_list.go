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

type PlaceList struct {
	Pagi    list.PagiRequest
	Filters valobj.PlaceFilters
}

type PlaceListHandler cqrs.HandlerFunc[PlaceList, *list.PagiResponse[*dtos.PlaceList]]

func NewPlaceListHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceListHandler {
	return func(ctx context.Context, query PlaceList) (*list.PagiResponse[*dtos.PlaceList], error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceListHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		query.Filters.IsActive = "1"
		res, err := placeRepo.Filter(ctx, &query.Pagi, &query.Filters)
		if err != nil {
			return nil, err
		}
		return dtos.NewPlaceList(res), nil
	}
}
