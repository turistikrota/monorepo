package commands

import (
	"context"
	"errors"

	"github.com/9ssi7/txn"
	"github.com/google/uuid"
	"github.com/turistikrota/api/config/claims"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type RoleUpdate struct {
	UserId      uuid.UUID `json:"user_id" validate:"-"`
	Id          uuid.UUID `json:"role_id" params:"role_id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,min=3,max=255"`
	Description string    `json:"description" validate:"required,min=3,max=255"`
	Claims      []string  `json:"claims" validate:"required,gte=1,dive,min=3,max=255"`
}

type RoleUpdateHandler cqrs.HandlerFunc[RoleUpdate, *cqrs.Empty]

func NewRoleUpdateHandler(t trace.Tracer, v validation.Service, roleRepo abstracts.RoleRepo) RoleUpdateHandler {
	return func(ctx context.Context, cmd RoleUpdate) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.RoleUpdateHandler")
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
		tx := txn.New()
		tx.Register(roleRepo.GetTxnAdapter())
		onError := func(ctx context.Context, err error) error {
			tx.Rollback(ctx)
			return err
		}
		role, err := roleRepo.FindById(ctx, cmd.Id)
		if err != nil {
			return nil, onError(ctx, err)
		}
		role.Update(cmd.UserId, cmd.Name, cmd.Description, cmd.Claims)
		if err := roleRepo.Save(ctx, role); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
