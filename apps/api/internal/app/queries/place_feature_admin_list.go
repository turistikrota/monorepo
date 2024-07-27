package queries

import (
	"context"

	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/list"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type PlaceFeatureAdminList struct {
	Pagi    list.PagiRequest
	Filters valobj.BaseFilters
}

type PlaceFeatureAdminListHandler cqrs.HandlerFunc[PlaceFeatureAdminList, *list.PagiResponse[*entities.PlaceFeature]]

func NewPlaceFeatureAdminListHandler(t trace.Tracer, v validation.Service, placeFeatureRepo abstracts.PlaceFeatureRepo) PlaceFeatureAdminListHandler {
	return func(ctx context.Context, query PlaceFeatureAdminList) (*list.PagiResponse[*entities.PlaceFeature], error) {
		ctx = tracer.Push(ctx, t, "queries.PlaceFeatureAdminListHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := placeFeatureRepo.Filter(ctx, &query.Pagi, &query.Filters)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
