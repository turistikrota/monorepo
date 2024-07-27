package commands

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/aggregates"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/internal/domain/events"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/state"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type AuthLoginStart struct {
	Phone  string         `json:"phone" validate:"required_without=Email,omitempty,phone"`
	Email  string         `json:"email" validate:"required_without=Phone,omitempty,email"`
	Device *valobj.Device `json:"-"`
}

type AuthLoginStartRes struct {
	VerifyToken string `json:"-"`
}

type AuthLoginStartHandler cqrs.HandlerFunc[AuthLoginStart, *AuthLoginStartRes]

func NewAuthLoginStartHandler(t trace.Tracer, v validation.Service, verifyRepo abstracts.VerifyRepo, userRepo abstracts.UserRepo) AuthLoginStartHandler {
	return func(ctx context.Context, cmd AuthLoginStart) (*AuthLoginStartRes, error) {
		ctx = tracer.Push(ctx, t, "commands.AuthLoginStartHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		var user *entities.User
		if cmd.Phone != "" {
			user, err = userRepo.FindByPhone(ctx, cmd.Phone)
			if err != nil {
				return nil, err
			}
		} else {
			user, err = userRepo.FindByEmail(ctx, cmd.Email)
			if err != nil {
				return nil, err
			}
		}
		if user == nil {
			return nil, rescode.NotFound(errors.New("user not found"))
		}
		if !user.IsActive {
			return nil, rescode.UserDisabled(errors.New("user disabled"))
		}
		if user.TempToken != nil && *user.TempToken != "" {
			return nil, rescode.UserVerifyRequired(errors.New("user verify required"))
		}
		verifyToken := uuid.New().String()
		verify := aggregates.NewVerify(user.Id, state.GetDeviceId(ctx), state.GetLocale(ctx))
		err = verifyRepo.Save(ctx, verifyToken, verify)
		if err != nil {
			return nil, err
		}
		events.OnAuthLoginStarted(events.AuthLoginStarted{
			Email:  user.Email,
			Code:   verify.Code,
			Device: *cmd.Device,
		})
		return &AuthLoginStartRes{
			VerifyToken: verifyToken,
		}, nil
	}
}
