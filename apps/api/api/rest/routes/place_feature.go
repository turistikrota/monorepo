package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/api/api/rest/middlewares"
	restsrv "github.com/turistikrota/api/api/rest/srv"
	"github.com/turistikrota/api/internal/app"
	"github.com/turistikrota/api/internal/app/commands"
	"github.com/turistikrota/api/internal/app/queries"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/list"
)

func PlaceFeature(router fiber.Router, srv restsrv.Srv, app app.App) {
	group := router.Group("/place-features")
	group.Post("/", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(placeFeatureCreate(app)))
	group.Put("/:feature_id", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(placeFeatureUpdate(app)))
	group.Patch("/:feature_id/disable", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(placeFeatureDisable(app)))
	group.Patch("/:feature_id/enable", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(placeFeatureEnable(app)))
	group.Get("/", srv.Timeout(placeFeatureList(app)))
	group.Get("/admin", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(placeFeatureAdminList(app)))
	group.Get("/admin/:feature_id", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(placeFeatureAdminView(app)))
}

func placeFeatureCreate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceFeatureCreate
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceFeatureCreate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

func placeFeatureUpdate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceFeatureUpdate
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceFeatureUpdate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeFeatureDisable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceFeatureDisable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceFeatureDisable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeFeatureEnable(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.PlaceFeatureEnable
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		cmd.UserId = middlewares.AccessParse(c).Id
		res, err := app.Commands.PlaceFeatureEnable(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeFeatureList(app app.App) fiber.Handler {
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
		query := queries.PlaceFeatureList{
			Pagi:    pagi,
			Filters: filters,
		}
		res, err := app.Queries.PlaceFeatureList(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeFeatureAdminList(app app.App) fiber.Handler {
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
		query := queries.PlaceFeatureAdminList{
			Pagi:    pagi,
			Filters: filters,
		}
		res, err := app.Queries.PlaceFeatureAdminList(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func placeFeatureAdminView(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query queries.PlaceFeatureAdminView
		if err := c.ParamsParser(&query); err != nil {
			return err
		}
		res, err := app.Queries.PlaceFeatureAdminView(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
