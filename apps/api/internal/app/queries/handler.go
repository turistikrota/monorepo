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

	PlaceFeatureAdminList PlaceFeatureAdminListHandler
	PlaceFeatureAdminView PlaceFeatureAdminViewHandler
	PlaceFeatureList      PlaceFeatureListHandler
}

func NewHandler(tracer trace.Tracer, r abstracts.Repositories, v validation.Service) Handlers {
	return Handlers{
		AuthCheck:         NewAuthCheckHandler(tracer, r.VerifyRepo),
		AuthVerifyAccess:  NewAuthVerifyAccessHandler(tracer, r.SessionRepo),
		AuthVerifyRefresh: NewAuthVerifyRefreshHandler(tracer, r.SessionRepo),

		PlaceFeatureAdminList: NewPlaceFeatureAdminListHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureAdminView: NewPlaceFeatureAdminViewHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureList:      NewPlaceFeatureListHandler(tracer, v, r.PlaceFeatureRepo),
	}
}
