package migrations

import (
	"github.com/turistikrota/api/internal/domain/entities"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}
