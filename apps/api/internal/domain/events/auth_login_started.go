package events

import (
	"github.com/turistikrota/api/assets"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/internal/infra/mail"
)

type AuthLoginStarted struct {
	Email  string
	Code   string
	Device valobj.Device
}

func OnAuthLoginStarted(e AuthLoginStarted) {
	go func() {
		mail.GetClient().SendWithTemplate(mail.SendWithTemplateConfig{
			SendConfig: mail.SendConfig{
				To:      []string{e.Email},
				Subject: "Doğrulama Kodunuz",
				Message: e.Code,
			},
			Template: assets.Templates.AuthVerify,
			Data: map[string]interface{}{
				"Code":    e.Code,
				"IP":      mail.GetField(e.Device.IP),
				"Browser": mail.GetField(e.Device.Name),
				"OS":      mail.GetField(e.Device.OS),
			},
		})

	}()
}
