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
	Ids []string `json:"role_ids" validate:"required,gt=0,dive,uuid"`
}

type RoleListIdsHandler cqrs.HandlerFunc[RoleListIds, []*entities.Role]

func NewRoleListIdsHandler(t trace.Tracer, v validation.Service, roleRepo abstracts.RoleRepo) RoleListIdsHandler {
	return func(ctx context.Context, query RoleListIds) ([]*entities.Role, error) {
		ctx = tracer.Push(ctx, t, "queries.RoleListIdsHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		ids := make([]uuid.UUID, 0, len(query.Ids))
		for _, id := range query.Ids {
			uid, err := uuid.Parse(id)
			if err != nil {
				return nil, err
			}
			ids = append(ids, uid)
		}
		res, err := roleRepo.FindByIds(ctx, ids)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
