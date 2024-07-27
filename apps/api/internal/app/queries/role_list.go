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

type RoleList struct {
	Pagi    list.PagiRequest
	Filters valobj.BaseFilters
}

type RoleListHandler cqrs.HandlerFunc[RoleList, *list.PagiResponse[*dtos.RoleList]]

func NewRoleListHandler(t trace.Tracer, v validation.Service, roleRepo abstracts.RoleRepo) RoleListHandler {
	return func(ctx context.Context, query RoleList) (*list.PagiResponse[*dtos.RoleList], error) {
		ctx = tracer.Push(ctx, t, "queries.RoleListHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := roleRepo.Filter(ctx, &query.Pagi, &query.Filters)
		if err != nil {
			return nil, err
		}
		return dtos.NewRoleList(res), nil
	}
}
