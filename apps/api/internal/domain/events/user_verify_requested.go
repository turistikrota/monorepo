package events

import (
	"github.com/turistikrota/api/assets"
	"github.com/turistikrota/api/internal/infra/mail"
)

type UserVerifyRequested struct {
	Name             string
	Email            string
	VerificationCode string
}

func OnUserVerifyRequested(e UserVerifyRequested) {
	go func() {
		mail.GetClient().SendWithTemplate(mail.SendWithTemplateConfig{
			SendConfig: mail.SendConfig{
				To:      []string{e.Email},
				Subject: "Hesabınızı Yeniden Doğrulayın",
			},
			Template: assets.Templates.AuthVerifyRequested,
			Data: map[string]interface{}{
				"Name":             e.Name,
				"VerificationCode": e.VerificationCode,
			},
		})
	}()
}
