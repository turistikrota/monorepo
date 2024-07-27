package events

import (
	"fmt"

	"github.com/turistikrota/api/assets"
	"github.com/turistikrota/api/config"
	"github.com/turistikrota/api/internal/infra/mail"
)

type AuthRegistered struct {
	Name             string
	Email            string
	VerificationCode string
}

func OnAuthRegistered(e AuthRegistered) {
	go func() {
		mail.GetClient().SendWithTemplate(mail.SendWithTemplateConfig{
			SendConfig: mail.SendConfig{
				To:      []string{e.Email},
				Subject: "Turistikrota'ya Hoşgeldiniz",
			},
			Template: assets.Templates.AuthRegistered,
			Data: map[string]interface{}{
				"Name":            e.Name,
				"VerificationUrl": fmt.Sprintf("%s/auth/verify/%s", config.ReadValue().PublicHost, e.VerificationCode),
			},
		})
	}()
}
