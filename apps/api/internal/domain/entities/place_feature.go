package entities

import (
	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/valobj"
)

type PlaceFeature struct {
	Base
	valobj.Audit
	Icon        string `json:"icon" gorm:"type:varchar(255);not null"`
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	IsActive    bool   `json:"is_active" gorm:"type:boolean;not null;default:true"`
}

func (p *PlaceFeature) Enable() {
	p.IsActive = true
}

func (p *PlaceFeature) Disable() {
	p.IsActive = false
}

func (p *PlaceFeature) Update(adminId uuid.UUID, title string, description string, icon string) {
	p.Title = title
	p.Description = description
	p.Icon = icon
	p.Audit.UpdatedBy = &adminId
}

func NewPlaceFeature(adminId uuid.UUID, title string, description string, icon string) *PlaceFeature {
	return &PlaceFeature{
		Audit: valobj.Audit{
			MakedBy: &adminId,
		},
		Title:       title,
		Description: description,
		Icon:        icon,
		IsActive:    true,
	}
}
