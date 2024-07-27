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

type UserAdminDisable struct {
	CurrentUserId uuid.UUID `json:"-" validate:"-"`
	UserId        uuid.UUID `json:"user_id" validate:"required,uuid"`
}

type UserAdminDisableHandler cqrs.HandlerFunc[UserAdminDisable, *cqrs.Empty]

func NewUserAdminDisableHandler(t trace.Tracer, v validation.Service, userRepo abstracts.UserRepo) UserAdminDisableHandler {
	return func(ctx context.Context, cmd UserAdminDisable) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.UserAdminDisableHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		tx := txn.New()
		tx.Register(userRepo.GetTxnAdapter())
		onError := func(ctx context.Context, err error) error {
			tx.Rollback(ctx)
			return err
		}
		user, err := userRepo.FindById(ctx, cmd.UserId)
		if err != nil {
			return nil, onError(ctx, err)
		}
		if !user.IsActive {
			return nil, onError(ctx, rescode.UserAlreadyDisabled(errors.New("user is not active")))
		}
		user.Disable(cmd.CurrentUserId)
		if err := userRepo.Save(ctx, user); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
