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
	}
}
