package app

import (
	"github.com/turistikrota/api/internal/app/commands"
	"github.com/turistikrota/api/internal/app/queries"
	"github.com/turistikrota/api/internal/app/services"
)

type App struct {
	Commands commands.Handlers
	Queries  queries.Handlers
	Services services.Handlers
}
