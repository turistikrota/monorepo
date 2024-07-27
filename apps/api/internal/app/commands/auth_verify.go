package commands

import (
	"context"

	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/tracer"
	"go.opentelemetry.io/otel/trace"
)

type AuthVerify struct {
	Token string `json:"token" validate:"required,uuid"`
}

type AuthVerifyHandler cqrs.HandlerFunc[AuthVerify, *cqrs.Empty]

func NewAuthVerifyHandler(t trace.Tracer, userRepo abstracts.UserRepo) AuthVerifyHandler {
	return func(ctx context.Context, cmd AuthVerify) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.AuthVerifyHandler")
		u, err := userRepo.FindByToken(ctx, cmd.Token)
		if err != nil {
			return nil, err
		}
		u.Verify()
		err = userRepo.Save(ctx, u)
		if err != nil {
			return nil, err
		}
		return &cqrs.Empty{}, nil
	}
}
