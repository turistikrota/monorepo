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

type AuthVerifyRefresh struct {
	AccessToken  string
	RefreshToken string
	IpAddr       string
}

type AuthVerifyRefreshRes struct {
	User *token.UserClaim
}

type AuthVerifyRefreshHandler cqrs.HandlerFunc[AuthVerifyRefresh, *AuthVerifyRefreshRes]

func NewAuthVerifyRefreshHandler(t trace.Tracer, sessionRepo abstracts.SessionRepo) AuthVerifyRefreshHandler {
	return func(ctx context.Context, query AuthVerifyRefresh) (*AuthVerifyRefreshRes, error) {
		ctx = tracer.Push(ctx, t, "queries.AuthVerifyRefreshHandler")
		claims, err := token.Client().Parse(query.RefreshToken)
		if err != nil {
			return nil, rescode.Failed(err)
		}
		isValid, err := token.Client().Verify(query.RefreshToken)
		if err != nil {
			return nil, rescode.Failed(err)
		}
		if !isValid {
			return nil, rescode.InvalidOrExpiredToken(errors.New("invalid or expired refresh token"))
		}
		session, notFound, err := sessionRepo.FindByIds(ctx, claims.Id, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if notFound {
			return nil, rescode.InvalidRefreshToken(errors.New("invalid refresh with access token and ip"))
		}
		if !session.IsRefreshValid(query.AccessToken, query.RefreshToken, query.IpAddr) {
			return nil, rescode.InvalidRefreshToken(errors.New("invalid refresh with access token and ip"))
		}
		return &AuthVerifyRefreshRes{
			User: claims,
		}, nil
	}
}
