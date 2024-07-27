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

type RoleView struct {
	Id uuid.UUID `json:"role_id" params:"role_id" validate:"required,uuid"`
}

type RoleViewHandler cqrs.HandlerFunc[RoleView, *entities.Role]

func NewRoleViewHandler(t trace.Tracer, v validation.Service, placeRepo abstracts.RoleRepo) RoleViewHandler {
	return func(ctx context.Context, query RoleView) (*entities.Role, error) {
		ctx = tracer.Push(ctx, t, "queries.RoleViewHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := placeRepo.FindById(ctx, query.Id)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
