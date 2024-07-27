package commands

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/turistikrota/api/config/claims"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type RoleCreate struct {
	UserId      uuid.UUID `json:"user_id" validate:"-"`
	Name        string    `json:"name" validate:"required,min=3,max=255"`
	Description string    `json:"description" validate:"required,min=3,max=255"`
	Claims      []string  `json:"claims" validate:"required,gte=1,dive,min=3,max=255"`
}

type RoleCreateHandler cqrs.HandlerFunc[RoleCreate, *uuid.UUID]

func NewRoleCreateHandler(t trace.Tracer, v validation.Service, roleRepo abstracts.RoleRepo) RoleCreateHandler {
	return func(ctx context.Context, cmd RoleCreate) (*uuid.UUID, error) {
		ctx = tracer.Push(ctx, t, "commands.RoleCreateHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		for _, claim := range cmd.Claims {
			if !claims.IsReal(claim) {
				return nil, rescode.ClaimIsNotReal(errors.New("claim is not real"))
			}
		}
		exists, err := roleRepo.IsExistsByName(ctx, cmd.Name)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, rescode.RoleNameAlreadyExists(errors.New("role already exists"))
		}
		role := entities.NewRole(cmd.UserId, cmd.Name, cmd.Description, cmd.Claims)
		if err := roleRepo.Save(ctx, role); err != nil {
			return nil, err
		}
		return &role.Id, nil
	}
}
