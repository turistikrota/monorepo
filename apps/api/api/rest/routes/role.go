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

func Roles(router fiber.Router, srv restsrv.Srv, app app.App) {
	group := router.Group("/roles")
	group.Post("/", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.Role.Create), srv.Timeout(roleCreate(app)))
	group.Put("/:role_id", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.Role.Update), srv.Timeout(roleUpdate(app)))
	group.Patch("/:role_id/disable", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.Role.Disable), srv.Timeout(roleDisable(app)))
	group.Patch("/:role_id/enable", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.Role.Enable), srv.Timeout(roleEnable(app)))
	group.Get("/", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.Role.List), srv.Timeout(roleList(app)))
	group.Get("/:role_id", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Role.Super, claims.Role.View), srv.Timeout(roleView(app)))
}

func roleCreate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.RoleCreate
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.RoleCreate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

func roleUpdate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.RoleUpdate
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.RoleUpdate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func roleEnable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.RoleEnable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.RoleEnable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func roleDisable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.RoleDisable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.RoleDisable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func roleList(app app.App) fiber.Handler {
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
		query := queries.RoleList{
			Pagi:    pagi,
			Filters: filters,
		}
		res, err := app.Queries.RoleList(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func roleView(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query queries.RoleView
		if err := c.ParamsParser(&query); err != nil {
			return err
		}
		res, err := app.Queries.RoleView(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
