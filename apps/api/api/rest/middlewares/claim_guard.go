package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/api/pkg/claguard"
	"github.com/turistikrota/api/pkg/rescode"
)

type ClaimGuardConfig struct {
	Claims []string
}

func NewClaimGuard(claims []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := AccessMustParse(ctx)
		if claguard.Check(u.Roles, claims) {
			return ctx.Next()
		}
		return rescode.PermissionDenied(errors.New("permission denied"))
	}
}
