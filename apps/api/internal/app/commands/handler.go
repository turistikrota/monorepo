package commands

import (
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/pkg/validation"
	"go.opentelemetry.io/otel/trace"
)

type Handlers struct {
	AuthLoginVerify AuthLoginVerifyHandler
	AuthRegister    AuthRegisterHandler
	AuthLoginStart  AuthLoginStartHandler
	AuthRefresh     AuthRefreshHandler
	AuthLogout      AuthLogoutHandler
	AuthVerify      AuthVerifyHandler

	PlaceFeatureCreate  PlaceFeatureCreateHandler
	PlaceFeatureUpdate  PlaceFeatureUpdateHandler
	PlaceFeatureDisable PlaceFeatureDisableHandler
	PlaceFeatureEnable  PlaceFeatureEnableHandler

	PlaceCreate  PlaceCreateHandler
	PlaceUpdate  PlaceUpdateHandler
	PlaceDisable PlaceDisableHandler
	PlaceEnable  PlaceEnableHandler

	RoleCreate  RoleCreateHandler
	RoleUpdate  RoleUpdateHandler
	RoleDisable RoleDisableHandler
	RoleEnable  RoleEnableHandler

	UserRoleAdd      UserRoleAddHandler
	UserRoleRemove   UserRoleRemoveHandler
	UserDisable      UserDisableHandler
	UserEnable       UserEnableHandler
	UserAdminDisable UserAdminDisableHandler
	UserAdminEnable  UserAdminEnableHandler
}

func NewHandler(tracer trace.Tracer, r abstracts.Repositories, v validation.Service) Handlers {
	return Handlers{
		AuthLoginVerify: NewAuthLoginVerifyHandler(tracer, v, r.UserRepo, r.VerifyRepo, r.SessionRepo),
		AuthLoginStart:  NewAuthLoginStartHandler(tracer, v, r.VerifyRepo, r.UserRepo),
		AuthLogout:      NewAuthLogoutHandler(tracer, r.SessionRepo),
		AuthRefresh:     NewAuthRefreshHandler(tracer, r.SessionRepo, r.UserRepo),
		AuthRegister:    NewAuthRegisterHandler(tracer, v, r.UserRepo),
		AuthVerify:      NewAuthVerifyHandler(tracer, r.UserRepo),

		PlaceFeatureCreate:  NewPlaceFeatureCreateHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureUpdate:  NewPlaceFeatureUpdateHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureDisable: NewPlaceFeatureDisableHandler(tracer, v, r.PlaceFeatureRepo),
		PlaceFeatureEnable:  NewPlaceFeatureEnableHandler(tracer, v, r.PlaceFeatureRepo),

		PlaceCreate:  NewPlaceCreateHandler(tracer, v, r.PlaceRepo),
		PlaceUpdate:  NewPlaceUpdateHandler(tracer, v, r.PlaceRepo),
		PlaceDisable: NewPlaceDisableHandler(tracer, v, r.PlaceRepo),
		PlaceEnable:  NewPlaceEnableHandler(tracer, v, r.PlaceRepo),

		RoleCreate:  NewRoleCreateHandler(tracer, v, r.RoleRepo),
		RoleUpdate:  NewRoleUpdateHandler(tracer, v, r.RoleRepo),
		RoleDisable: NewRoleDisableHandler(tracer, v, r.RoleRepo),
		RoleEnable:  NewRoleEnableHandler(tracer, v, r.RoleRepo),

		UserRoleAdd:      NewUserRoleAddHandler(tracer, v, r.UserRepo, r.RoleRepo),
		UserRoleRemove:   NewUserRoleRemoveHandler(tracer, v, r.UserRepo, r.RoleRepo),
		UserDisable:      NewUserDisableHandler(tracer, r.UserRepo),
		UserEnable:       NewUserEnableHandler(tracer, v, r.UserRepo),
		UserAdminDisable: NewUserAdminDisableHandler(tracer, v, r.UserRepo),
		UserAdminEnable:  NewUserAdminEnableHandler(tracer, v, r.UserRepo),
	}
}
