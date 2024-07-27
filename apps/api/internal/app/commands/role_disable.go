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

type RoleDisable struct {
	UserId uuid.UUID `json:"user_id" validate:"-"`
	Id     uuid.UUID `json:"role_id" params:"role_id" validate:"required,uuid"`
}

type RoleDisableHandler cqrs.HandlerFunc[RoleDisable, *cqrs.Empty]

func NewRoleDisableHandler(t trace.Tracer, v validation.Service, roleRepo abstracts.RoleRepo) RoleDisableHandler {
	return func(ctx context.Context, cmd RoleDisable) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.RoleDisableHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
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
		if role.IsLocked {
			return nil, onError(ctx, rescode.RoleIsLocked(errors.New("role is locked")))
		}
		role.Disable(cmd.UserId)
		if err := roleRepo.Save(ctx, role); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
