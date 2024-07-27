package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/api/api/rest/middlewares"
	restsrv "github.com/turistikrota/api/api/rest/srv"
	"github.com/turistikrota/api/config/claims"
	"github.com/turistikrota/api/internal/app"
	"github.com/turistikrota/api/internal/app/commands"
	"github.com/turistikrota/api/internal/app/queries"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/list"
)

func Users(router fiber.Router, srv restsrv.Srv, app app.App) {
	group := router.Group("/users")
	group.Patch("/disable", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(userDisable(app)))
	group.Patch("/enable", srv.Turnstile(), srv.Timeout(userEnable(app)))
	group.Get("/admin", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.User.Super, claims.User.List), srv.Timeout(userAdminList(app)))
	group.Get("/admin/:user_id", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.User.Super, claims.User.View), srv.Timeout(userAdminView(app)))
	group.Patch("/:user_id/role-add", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.User.Super, claims.User.RoleAdd), srv.Timeout(userRoleAdd(app)))
	group.Patch("/:user_id/role-remove", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.User.Super, claims.User.RoleRemove), srv.Timeout(userRoleRemove(app)))
	group.Patch("/:user_id/disable", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.User.Super, claims.User.Disable), srv.Timeout(userAdminDisable(app)))
	group.Patch("/:user_id/enable", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.User.Enable), srv.Timeout(userAdminEnable(app)))
}

func userAdminDisable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.UserAdminDisable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.CurrentUserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.UserAdminDisable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userAdminEnable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.UserAdminEnable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.CurrentUserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.UserAdminEnable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userDisable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cmd := commands.UserDisable{
			UserId: middlewares.AccessParse(c).Id,
		}
		res, err := app.Commands.UserDisable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userEnable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.UserEnable
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		res, err := app.Commands.UserEnable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userRoleAdd(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.UserRoleAdd
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.CurrentUserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.UserRoleAdd(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userRoleRemove(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.UserRoleRemove
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.CurrentUserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.UserRoleRemove(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userAdminList(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pagi list.PagiRequest
		if err := c.QueryParser(&pagi); err != nil {
			return err
		}
		var filters valobj.BaseFilters
		if err := c.QueryParser(&filters); err != nil {
			return err
		}
		pagi.Default()
		query := queries.UserAdminList{
			Pagi:    pagi,
			Filters: filters,
		}
		res, err := app.Queries.UserAdminList(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func userAdminView(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query queries.UserAdminView
		if err := c.ParamsParser(&query); err != nil {
			return err
		}
		res, err := app.Queries.UserAdminView(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
