package commands

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/cqrs"
	"github.com/turistikrota/api/pkg/rescode"
	"github.com/turistikrota/api/pkg/state"
	"github.com/turistikrota/api/pkg/token"
	"github.com/turistikrota/api/pkg/tracer"
	"go.opentelemetry.io/otel/trace"
)

type AuthRefresh struct {
	AccessToken  string
	RefreshToken string
	IpAddress    string
	UserId       uuid.UUID
}

type AuthRefreshRes struct {
	AccessToken string
}

type AuthRefreshHandler cqrs.HandlerFunc[AuthRefresh, *AuthRefreshRes]

func NewAuthRefreshHandler(t trace.Tracer, sessionRepo abstracts.SessionRepo, userRepo abstracts.UserRepo) AuthRefreshHandler {
	return func(ctx context.Context, cmd AuthRefresh) (*AuthRefreshRes, error) {
		ctx = tracer.Push(ctx, t, "commands.AuthRefreshHandler")
		session, notFound, err := sessionRepo.FindByIds(ctx, cmd.UserId, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if notFound {
			return nil, rescode.InvalidRefreshOrAccessTokens(errors.New("invalid refresh with access token and ip"))
		}
		if !session.IsRefreshValid(cmd.AccessToken, cmd.RefreshToken, cmd.IpAddress) {
			return nil, rescode.InvalidRefreshOrAccessTokens(errors.New("invalid refresh with access token and ip"))
		}
		user, err := userRepo.FindById(ctx, cmd.UserId)
		if err != nil {
			return nil, err
		}
		accessToken, err := token.Client().GenerateAccessToken(token.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Roles: user.Roles,
		})
		if err != nil {
			return nil, err
		}
		session.Refresh(accessToken)
		if err := sessionRepo.Save(ctx, user.Id, session); err != nil {
			return nil, err
		}
		return &AuthRefreshRes{
			AccessToken: accessToken,
		}, nil
	}
}
