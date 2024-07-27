package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/api/internal/app"
	"github.com/turistikrota/api/internal/app/queries"
	"github.com/turistikrota/api/pkg/claguard"
	"github.com/turistikrota/api/pkg/rescode"
)

type ClaimGuardConfig struct {
	Claims []string
}

func NewClaimGuard(app app.App, claims []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := AccessMustParse(ctx)
		if len(u.Roles) == 0 {
			return rescode.PermissionDenied(errors.New("permission denied"))
		}
		roles, err := app.Queries.RoleListIds(ctx.UserContext(), queries.RoleListIds{
			Ids: u.Roles,
		})
		if err != nil {
			return err
		}
		allClaims := make([]string, 0)
		for _, role := range roles {
			allClaims = append(allClaims, role.Claims...)
		}
		if claguard.Check(allClaims, claims) {
			return ctx.Next()
		}
		return rescode.PermissionDenied(errors.New("permission denied"))
	}
}
