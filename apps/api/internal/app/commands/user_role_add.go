package commands

import (
	"context"
	"errors"

	"github.com/9ssi7/txn"
	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type UserRoleAdd struct {
	CurrentUserId uuid.UUID `json:"-" validate:"-"`
	UserId        uuid.UUID `json:"user_id" validate:"required,uuid"`
	RoleId        uuid.UUID `json:"role_id" validate:"required,uuid"`
}

type UserRoleAddHandler cqrs.HandlerFunc[UserRoleAdd, *cqrs.Empty]

func NewUserRoleAddHandler(t trace.Tracer, v validation.Service, userRepo abstracts.UserRepo, roleRepo abstracts.RoleRepo) UserRoleAddHandler {
	return func(ctx context.Context, cmd UserRoleAdd) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.UserRoleAddHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		tx := txn.New()
		tx.Register(userRepo.GetTxnAdapter())
		tx.Register(roleRepo.GetTxnAdapter())
		onError := func(ctx context.Context, err error) error {
			tx.Rollback(ctx)
			return err
		}
		user, err := userRepo.FindById(ctx, cmd.UserId)
		if err != nil {
			return nil, onError(ctx, err)
		}
		role, err := roleRepo.FindById(ctx, cmd.RoleId)
		if err != nil {
			return nil, onError(ctx, err)
		}
		if !role.IsActive {
			return nil, onError(ctx, rescode.RoleIsNotActive(errors.New("role is not active")))
		}
		if user.CheckRole(cmd.RoleId.String()) {
			return nil, onError(ctx, rescode.RoleAlreadyAssigned(errors.New("role already assigned")))
		}
		user.AddRole(cmd.CurrentUserId, cmd.RoleId.String())
		if err := userRepo.Save(ctx, user); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
