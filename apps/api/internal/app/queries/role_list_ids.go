package queries

import (
	"context"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type RoleListIds struct {
	Ids []uuid.UUID `json:"role_ids" validate:"required,gt=0,dive,uuid"`
}

type RoleListIdsHandler cqrs.HandlerFunc[RoleListIds, []*entities.Role]

func NewRoleListIdsHandler(t trace.Tracer, v validation.Service, roleRepo abstracts.RoleRepo) RoleListIdsHandler {
	return func(ctx context.Context, query RoleListIds) ([]*entities.Role, error) {
		ctx = tracer.Push(ctx, t, "queries.RoleListIdsHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := roleRepo.FindByIds(ctx, query.Ids)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
