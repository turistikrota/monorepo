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

type UserAdminList struct {
	Pagi    list.PagiRequest
	Filters valobj.BaseFilters
}

type UserAdminListHandler cqrs.HandlerFunc[UserAdminList, *list.PagiResponse[*dtos.UserAdminList]]

func NewUserAdminListHandler(t trace.Tracer, v validation.Service, userRepo abstracts.UserRepo) UserAdminListHandler {
	return func(ctx context.Context, query UserAdminList) (*list.PagiResponse[*dtos.UserAdminList], error) {
		ctx = tracer.Push(ctx, t, "queries.UserAdminListHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := userRepo.Filter(ctx, &query.Pagi, &query.Filters)
		if err != nil {
			return nil, err
		}
		return dtos.NewUserAdminList(res), nil
	}
}
