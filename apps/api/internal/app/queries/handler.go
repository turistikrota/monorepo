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

	PlaceAdminList PlaceAdminListHandler
	PlaceAdminView PlaceAdminViewHandler
	PlaceList      PlaceListHandler
	PlaceView      PlaceViewHandler

	RoleList    RoleListHandler
	RoleListIds RoleListIdsHandler
	RoleView    RoleViewHandler

	UserAdminList UserAdminListHandler
	UserAdminView UserAdminViewHandler
}

func NewHandler(tracer trace.Tracer, r abstracts.Repositories, v validation.Service) Handlers {
	return Handlers{
		AuthCheck:         NewAuthCheckHandler(tracer, r.VerifyRepo),
		AuthVerifyAccess:  NewAuthVerifyAccessHandler(tracer, r.SessionRepo),
		AuthVerifyRefresh: NewAuthVerifyRefreshHandler(tracer, r.SessionRepo),

		PlaceFeatureAdminList: NewPlaceFeatureAdminListHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureAdminView: NewPlaceFeatureAdminViewHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureList:      NewPlaceFeatureListHandler(tracer, v, r.PlaceFeatureRepo),

		PlaceAdminList: NewPlaceAdminListHandler(tracer, v, r.PlaceRepo),
		PlaceAdminView: NewPlaceAdminViewHandler(tracer, v, r.PlaceRepo),
		PlaceList:      NewPlaceListHandler(tracer, v, r.PlaceRepo),
		PlaceView:      NewPlaceViewHandler(tracer, v, r.PlaceRepo),

		RoleList:    NewRoleListHandler(tracer, v, r.RoleRepo),
		RoleListIds: NewRoleListIdsHandler(tracer, v, r.RoleRepo),
		RoleView:    NewRoleViewHandler(tracer, v, r.RoleRepo),

		UserAdminList: NewUserAdminListHandler(tracer, v, r.UserRepo),
		UserAdminView: NewUserAdminViewHandler(tracer, v, r.UserRepo),
	}
}
