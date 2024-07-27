package queries

import (
	"context"
	"errors"

	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/state"
	"github.com/turistikrota/api/pkg/tracer"
	"go.opentelemetry.io/otel/trace"
)

type AuthCheck struct {
	VerifyToken string `json:"-"`
}

type AuthCheckHandler cqrs.HandlerFunc[AuthCheck, *cqrs.Empty]

func NewAuthCheckHandler(t trace.Tracer, verifyRepo abstracts.VerifyRepo) AuthCheckHandler {
	return func(ctx context.Context, query AuthCheck) (*cqrs.Empty, error) {
		ctx = tracer.Push(ctx, t, "queries.AuthCheckHandler")
		exists, err := verifyRepo.IsExists(ctx, query.VerifyToken, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, rescode.NotFound(errors.New("verify token not exists"))
		}
		return &cqrs.Empty{}, nil
	}
}
