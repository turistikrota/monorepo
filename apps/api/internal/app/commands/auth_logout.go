package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/state"
	"github.com/turistikrota/api/pkg/tracer"
	"go.opentelemetry.io/otel/trace"
)

type AuthLogout struct {
	UserId uuid.UUID `json:"-"`
}

type AuthLogoutHandler cqrs.HandlerFunc[AuthLogout, *cqrs.Empty]

func NewAuthLogoutHandler(t trace.Tracer, sessionRepo abstracts.SessionRepo) AuthLogoutHandler {
	return func(ctx context.Context, cmd AuthLogout) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "commands.AuthLogoutHandler")
		err := sessionRepo.Destroy(ctx, cmd.UserId, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		return &cqrs.Empty{}, nil
	}
}
