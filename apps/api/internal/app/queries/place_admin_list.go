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

type PlaceAdminList struct {
	Pagi    list.PagiRequest
	Filters valobj.PlaceFilters
}

type PlaceAdminListHandler cqrs.HandlerFunc[PlaceAdminList, *list.PagiResponse[*dtos.PlaceAdminList]]

func NewPlaceAdminListHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.PlaceRepo) PlaceAdminListHandler {
	return func(ctx context.Context, query PlaceAdminList) (*list.PagiResponse[*dtos.PlaceAdminList], error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceAdminListHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := placeRepo.Filter(ctx, &query.Pagi, &query.Filters)
		if err != nil {
			return nil, err
		}
		return dtos.NewPlaceAdminList(res), nil
	}
}
