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
	"go.opentelemetry.io/otel/trace"
)

type UserDisable struct {
	UserId uuid.UUID `json:"-" validate:"-"`
}

type UserDisableHandler cqrs.HandlerFunc[UserDisable, *cqrs.Empty]

func NewUserDisableHandler(t trace.Tracer, userRepo abstracts.UserRepo) UserDisableHandler {
	return func(ctx context.Context, cmd UserDisable) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.UserDisableHandler")
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
		user.Disable(cmd.UserId)
		if err := userRepo.Save(ctx, user); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		return &cqrs.Empty{}, nil
	}
}
