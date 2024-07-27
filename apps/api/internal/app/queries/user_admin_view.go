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

type UserAdminView struct {
	Id uuid.UUID `json:"user_id" params:"user_id" validate:"required,uuid"`
}

type UserAdminViewHandler cqrs.HandlerFunc[UserAdminView, *entities.User]

func NewUserAdminViewHandler(t trace.Tracer, v validation.Service, userRepo abstracts.UserRepo) UserAdminViewHandler {
	return func(ctx context.Context, query UserAdminView) (*entities.User, error) {
		ctx = tracer.Push(ctx, t, "queries.UserAdminViewHandler")
		err := v.ValidateStruct(ctx, query)
		if err != nil {
			return nil, err
		}
		res, err := userRepo.FindById(ctx, query.Id)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
