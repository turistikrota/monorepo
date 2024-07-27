package queries

import (
	"context"
	"errors"

	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/state"
	"github.com/turistikrota/api/pkg/token"
	"github.com/turistikrota/api/pkg/tracer"
	"go.opentelemetry.io/otel/trace"
)

type AuthVerifyAccess struct {
	AccessToken  string
	IpAddr       string
	IsUnverified bool
}

type AuthVerifyAccessRes struct {
	User *token.UserClaim
}

type AuthVerifyAccessHandler cqrs.HandlerFunc[AuthVerifyAccess, *AuthVerifyAccessRes]

func NewAuthVerifyAccessHandler(t trace.Tracer, sessionRepo abstracts.SessionRepo) AuthVerifyAccessHandler {
	return func(ctx context.Context, query AuthVerifyAccess) (*AuthVerifyAccessRes, error) {
		ctx = tracer.Push(ctx, t, "queries.AuthVerifyAccessHandler")
		var claims *token.UserClaim
		var err error
		if query.IsUnverified {
			claims, err = token.Client().Parse(query.AccessToken)
		} else {
			claims, err = token.Client().VerifyAndParse(query.AccessToken)
		}
		if err != nil {
			return nil, rescode.Failed(err)
		}
		session, notExists, err := sessionRepo.FindByIds(ctx, claims.Id, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if notExists {
			return nil, rescode.InvalidAccess(errors.New("invalid access with token and ip"))
		}
		if !session.IsAccessValid(query.AccessToken, query.IpAddr) {
			return nil, rescode.InvalidAccess(errors.New("invalid access with token and ip"))
		}
		return &AuthVerifyAccessRes{
			User: claims,
		}, nil
	}
}
