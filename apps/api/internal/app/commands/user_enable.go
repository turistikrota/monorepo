package commands

import (
	"context"
	"errors"

	"github.com/9ssi7/txn"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/events"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type UserEnable struct {
	Email string `json:"email" validate:"required,email"`
}

type UserEnableHandler cqrs.HandlerFunc[UserEnable, *cqrs.Empty]

func NewUserEnableHandler(t trace.Tracer, v validation.Service, userRepo abstracts.UserRepo) UserEnableHandler {
	return func(ctx context.Context, cmd UserEnable) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.UserEnableHandler")
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
		user, err := userRepo.FindByEmail(ctx, cmd.Email)
		if err != nil {
			return nil, onError(ctx, err)
		}
		if user.IsActive {
			return nil, onError(ctx, rescode.UserAlreadyEnabled(errors.New("user is active")))
		}
		if user.UpdatedBy != nil && user.UpdatedBy.String() != user.Id.String() {
			return nil, onError(ctx, rescode.OnlyAdminCanEnableUser(errors.New("only admin can enable user")))
		}
		user.UnVerify()
		if err := userRepo.Save(ctx, user); err != nil {
			return nil, onError(ctx, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, onError(ctx, err)
		}
		events.OnUserVerifyRequested(events.UserVerifyRequested{
			Name:             user.Name,
			Email:            user.Email,
			VerificationCode: *user.TempToken,
		})
		return &cqrs.Empty{}, nil
	}
}
