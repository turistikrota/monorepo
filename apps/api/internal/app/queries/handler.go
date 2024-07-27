package queries

import (
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type Handlers struct {
	AuthCheck         AuthCheckHandler
	AuthVerifyAccess  AuthVerifyAccessHandler
	AuthVerifyRefresh AuthVerifyRefreshHandler
}

func NewHandler(tracer trace.Tracer, r abstracts.Repositories, v validation.Service) Handlers {
	return Handlers{
		AuthCheck:         NewAuthCheckHandler(tracer, r.VerifyRepo),
		AuthVerifyAccess:  NewAuthVerifyAccessHandler(tracer, r.SessionRepo),
		AuthVerifyRefresh: NewAuthVerifyRefreshHandler(tracer, r.SessionRepo),
	}
}
