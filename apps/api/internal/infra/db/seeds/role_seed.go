package seeds

import (
	"github.com/turistikrota/api/config/claims"
	"github.com/turistikrota/api/internal/domain/entities"
	"gorm.io/gorm"
)

func runRoleSeeds(db *gorm.DB) {
	var role entities.Role
	if err := db.Model(&entities.Role{}).Where("name = ?", "Admin").First(&role).Error; err == gorm.ErrRecordNotFound {
		db.Create(&entities.Role{
			Name:        "Admin",
			Description: "Turistikrota admin role",
			IsActive:    true,
			IsLocked:    true,
			Claims: []string{
				claims.AdminSuper,
				claims.Admin,
			},
		})
	}
}
