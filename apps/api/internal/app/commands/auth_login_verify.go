package commands

import (
	"context"
	"errors"

	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/aggregates"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/state"
	"github.com/turistikrota/api/pkg/token"
	"github.com/turistikrota/api/pkg/tracer"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type AuthLoginVerify struct {
	VerifyToken string         `json:"-"`
	Code        string         `json:"code" validate:"required,numeric,len=4"`
	Device      *valobj.Device `json:"-"`
}

type AuthLoginVerifyRes struct {
	AccessToken  string `json:"-"`
	RefreshToken string `json:"-"`
}

type AuthLoginVerifyHandler cqrs.HandlerFunc[AuthLoginVerify, *AuthLoginVerifyRes]

func NewAuthLoginVerifyHandler(t trace.Tracer, v validation.Service, userRepo abstracts.UserRepo, verifyRepo abstracts.VerifyRepo, sessionRepo abstracts.SessionRepo) AuthLoginVerifyHandler {
	return func(ctx context.Context, cmd AuthLoginVerify) (*AuthLoginVerifyRes, error) {
		ctx = tracer.Push(ctx, t, "commands.AuthLoginVerifyHandler")
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		verify, err := verifyRepo.Find(ctx, cmd.VerifyToken, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if verify.IsExpired() {
			return nil, rescode.VerificationExpired(errors.New("verification expired"))
		}
		if verify.IsExceeded() {
			return nil, rescode.VerificationExceeded(errors.New("verification exceeded"))
		}
		if cmd.Code != verify.Code {
			verify.IncTryCount()
			err = verifyRepo.Save(ctx, cmd.VerifyToken, verify)
			if err != nil {
				return nil, err
			}
			return nil, rescode.VerificationInvalid(errors.New("verification invalid"))
		}
		err = verifyRepo.Delete(ctx, cmd.VerifyToken, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		user, err := userRepo.FindById(ctx, verify.UserId)
		if err != nil {
			return nil, err
		}
		accessToken, refreshToken, err := token.Client().Generate(token.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Roles: user.Roles,
		})
		if err != nil {
			return nil, rescode.Failed(err)
		}
		ses := aggregates.NewSession(*cmd.Device, state.GetDeviceId(ctx), accessToken, refreshToken)
		if err = sessionRepo.Save(ctx, user.Id, ses); err != nil {
			return nil, err
		}
		return &AuthLoginVerifyRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
}
