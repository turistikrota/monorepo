package seeds

import (
	"time"

	"github.com/lib/pq"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/ptr"
	"gorm.io/gorm"
)

func runUserSeeds(db *gorm.DB) {
	var user entities.User
	var role entities.Role
	db.Model(&entities.Role{}).Select("id").Where("name = ?", "Admin").First(&role)
	if err := db.Model(&entities.User{}).Where("email = ?", "test@test.com").First(&user).Error; err == gorm.ErrRecordNotFound {
		db.Create(&entities.User{
			Email: "test@test.com",
			Name:  "Test",
			Roles: pq.StringArray{
				role.Id.String(),
			},
			IsActive:   true,
			VerifiedAt: ptr.Time(time.Now()),
		})
	}
}
