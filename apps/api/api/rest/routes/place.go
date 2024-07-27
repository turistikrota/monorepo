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

func Places(router fiber.Router, srv restsrv.Srv, app app.App) {
	group := router.Group("/places")
	group.Post("/", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Place.Super, claims.Place.Create), srv.Timeout(placeCreate(app)))
	group.Put("/:place_id", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Place.Super, claims.Place.Update), srv.Timeout(placeUpdate(app)))
	group.Patch("/:place_id/disable", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Place.Super, claims.Place.Disable), srv.Timeout(placeDisable(app)))
	group.Patch("/:place_id/enable", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Place.Super, claims.Place.Enable), srv.Timeout(placeEnable(app)))
	group.Get("/", srv.Timeout(placeList(app)))
	group.Get("/admin", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Place.Super, claims.Place.List), srv.Timeout(placeAdminList(app)))
	group.Get("/admin/:place_id", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(claims.Place.Super, claims.Place.View), srv.Timeout(placeAdminView(app)))
	group.Get("/:slug", srv.Timeout(placeView(app)))
}

func placeCreate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceCreate
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceCreate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

func placeUpdate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceUpdate
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceUpdate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeDisable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceDisable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceDisable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeEnable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceEnable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceEnable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeList(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pagi list.PagiRequest
		if err := c.QueryParser(&pagi); err != nil {
			return err
		}
		var filters valobj.PlaceFilters
		if err := c.QueryParser(&filters); err != nil {
			return err
		}
		pagi.Default()
		query := queries.PlaceList{
			Pagi:    pagi,
			Filters: filters,
		}
		res, err := app.Queries.PlaceList(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeAdminList(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pagi list.PagiRequest
		if err := c.QueryParser(&pagi); err != nil {
			return err
		}
		var filters valobj.PlaceFilters
		if err := c.QueryParser(&filters); err != nil {
			return err
		}
		pagi.Default()
		query := queries.PlaceAdminList{
			Pagi:    pagi,
			Filters: filters,
		}
		res, err := app.Queries.PlaceAdminList(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeAdminView(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query queries.PlaceAdminView
		if err := c.ParamsParser(&query); err != nil {
			return err
		}
		res, err := app.Queries.PlaceAdminView(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeView(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query queries.PlaceView
		if err := c.ParamsParser(&query); err != nil {
			return err
		}
		res, err := app.Queries.PlaceView(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
