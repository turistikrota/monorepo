package seeds

import (
	"time"

	"github.com/lib/pq"
	"github.com/turistikrota/api/config/roles"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/ptr"
	"gorm.io/gorm"
)

func runUserSeeds(db *gorm.DB) {
	var user entities.User
	if err := db.Model(&entities.User{}).Where("email = ?", "test@test.com").First(&user).Error; err != nil {
		db.Create(&entities.User{
			Email: "test@test.com",
			Name:  "Test",
			Roles: pq.StringArray{
				roles.Admin,
				roles.AdminSuper,
			},
			IsActive:   true,
			VerifiedAt: ptr.Time(time.Now()),
		})
	}
}
